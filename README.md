# go18exp 
> A test implementation of some go generics concepts


## Example usage of some functional concepts
```go 
import (
	"fmt"
	"github.com/modfin/go18exp/result"
	"github.com/modfin/go18exp/slicez"
	"github.com/modfin/go18exp/slicez/pipe"
	"strconv"
)

func main() {

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}

	squares := pipe.Of(slice).
		Map(func(a int) int { // Squaring numbers
			return a * a
		}).
		Filter(func(a int) bool { // Filtering even
			return a%2 == 0
		}).
		Slice()
	fmt.Println(squares)
	// [4 16 36 64]

	sum := slicez.Fold(squares, func(acc int, val int) int {
		return acc + val
	}, 0)
	// 120

	avg := sum / len(squares)
	// 30

	small, big := slicez.Partition(squares, func(i int) bool {
		return i < avg
	})
	fmt.Println(small, big)
	// [4 16] [36 64]

	zipped := slicez.Zip(small, big, func(s int, b int) string {
		return fmt.Sprintf("(%d, %d)", s, b)
	})
	fmt.Println(zipped)
	// [(4, 36) (16, 64)]

	// This is all well and good, but does not cover error handling.
	// Introducing error handling in map function could give the following api
	//
	//  mapped, err := Map(slice, func(i int) (string, error){
	//	  i, err := strconv.ParseInt(str, 10, 64)
	//	  return int(i), err
	// })
	//
	// However, this would make the slicez api harder to work with in the usecases 
	// where errors is not needed to be returned.
	// Using something like wrapping, boxing or optional might be a better solution
	// eg.

	strslice := []string{"1", "2", "NaN", "4", "inf"}
	maybeInts := slicez.Map(strslice, func(str string) result.Result[int] {
		i, err := strconv.ParseInt(str, 10, 64)
		return result.From(int(i), err)
	})
	fmt.Println(maybeInts)
	// [{1} {2} {strconv.ParseInt: parsing "NaN": invalid syntax} {4} {strconv.ParseInt: parsing "inf": invalid syntax}]

	fmt.Println(result.SliceOk(maybeInts))
	// false

	fmt.Println(result.SliceError(maybeInts))
	// strconv.ParseInt: parsing "NaN": invalid syntax

	fmt.Println(result.SliceErrors(maybeInts))
	// [strconv.ParseInt: parsing "NaN": invalid syntax strconv.ParseInt: parsing "inf": invalid syntax]

	resultInts := slicez.Filter(maybeInts, result.ValueFilter[int])
	// or
	// resultInts := slicez.Filter(maybeInts, func(a result.Result[int]) bool {
	// 	 return a.Ok()
	// })
	fmt.Println(resultInts)
	// [{1} {2} {4}]

	ints := result.SliceValues(resultInts)
	// or
	// int := slicez.Map(resultInts, result.ValueMapper[int])
	// or
	// ints := slicez.Map(resultInts, func(a result.Result[int]) int {
	// 	 return a.Value()
	// })
	fmt.Println(ints)
	// [1 2 4]

}
```



## Example usage a wrapped version of container/heap as a priority queue
```go 
import (
	"fmt"
	"github.com/modfin/go18exp/containers/heap"
	"math/rand"
	"sync"
)

type Work struct {
	id       int
	priority int
}

func (w Work) Do() {
	fmt.Printf("- Consuming work with prio on %d and id %d\n", w.priority, w.id)
}

func main() {
	wg := sync.WaitGroup{}

	priorityQueue := heap.New[Work](func(a, b Work) bool {
		return a.priority < b.priority
	})

	producer := func() {
		for id := range make([]int, 10) {
			prio := rand.Intn(10)
			fmt.Printf("+ Producing work with prio on %d and id %d\n", prio, id)
			priorityQueue.Push(Work{
				id:       id,
				priority: prio,
			})

		}
		wg.Done()
	}

	consumer := func() {
		for range make([]int, 10) {
			// Spinnlock waiting for reads
			for priorityQueue.Len() == 0 {
			}
			work := priorityQueue.Pop()
			work.Do()

		}
		wg.Done()
	}

	wg.Add(1)
	go producer()
	wg.Add(1)
	go consumer()

	wg.Wait()
}
```


## Example usage a wrapped version of sort
```go 
import (
	"fmt"
	"github.com/modfin/go18exp/compare"
	"github.com/modfin/go18exp/slicez"
)

func main() {
	in := []int{2, 3, 5, 1, 12, 3, 6, 7, 34, 123, 65, 4631, 1, 1323}

	in = slicez.Sort(in)
	fmt.Println("Sorted in Ascending order")
	fmt.Println(in)
	fmt.Println()

	in = slicez.SortFunc(in, compare.Reverse(compare.Less[int]))
	fmt.Println("Sorted in Descending order")
	fmt.Println(in)
	fmt.Println()

	i, e := slicez.Search(slicez.Sort(in), func(e int) bool {
		return e >= 40
	})
	fmt.Println("Searching slice")
	fmt.Printf("index %d, element %d\n", i, e)
}
```

## Example usage of generics for common number calculations
```go 
import (
	"fmt"
	"github.com/modfin/go18exp/numberz"
	"github.com/modfin/go18exp/slicez"
)

func main() {

	ints := []int{1, 2, 3, 4, 5, 23, 12, 5, 231, 12, 12}
	fmt.Println("Integers")
	fmt.Println("Mean:", numberz.Mean(ints...))
	fmt.Println("Median:", numberz.Median(ints...))
	fmt.Println("Mode:", numberz.Mode(ints...))
	fmt.Println("Min:", numberz.Min(ints...))
	fmt.Println("Max:", numberz.Max(ints...))
	fmt.Println("Sum:", numberz.Sum(ints...))
	fmt.Println("x^2:", numberz.VPow(ints, 2))
	fmt.Println("Variance:", numberz.Var(ints...))  // Sample
	fmt.Println("StdDev:", numberz.StdDev(ints...)) // Sample
	fmt.Println()

	floats := slicez.Map(ints, numberz.MapFloat64[int])
	floats = slicez.Map(floats, func(a float64) float64 {
		return a + 0.5
	})
	fmt.Println("Floats")
	fmt.Println("Mean:", numberz.Mean(floats...))
	fmt.Println("Median:", numberz.Median(floats...))
	fmt.Println("Mode:", numberz.Mode(floats...))
	fmt.Println("Min:", numberz.Min(floats...))
	fmt.Println("Max:", numberz.Max(floats...))
	fmt.Println("Sum:", numberz.Sum(floats...))
	fmt.Println("x^2:", numberz.VPow(floats, 2))
	fmt.Println("Variance:", numberz.Var(floats...))
	fmt.Println("StdDev:", numberz.StdDev(floats...))

}
```