# go18exp 
> A test implementation of some go generics concepts


## Example usage of some functional concepts
```go 

import (
	"fmt"
	"github.com/modfin/go18exp/slicez"
	"github.com/modfin/go18exp/slicez/pipe"
)

func main() {

	even := func(i int) bool {
		return i%2 == 0
	}
	negate := func(i int) int {
		return -i
	}

	var pslice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var nslice = pipe.Of(pslice).Map(negate).Reverse().Slice()

	positive, negative := pipe.Of(pslice).
		Concat(nslice).                // Reversing original slice, negating numbers and concatenating the res
		Filter(even).                  // Filtering and keeping even numbers
		Drop(1).                       // Dropping one number on the left
		DropRight(1).                  // Dropping one number of the right
		Reverse().                     // Revers the slice
		Partition(func(val int) bool { // Partitioning into slice into positive and negative numbers
			return val > 0
		})

	// Mapping number to string
	toStr := func(v int) string {
		return fmt.Sprintf("%d", v)
	}
	var pStrSlice = slicez.Map(positive, toStr)
	var nStrSlice = slicez.Map(negative, toStr)

	// Joining []string to string
	joiner := func(accumulator string, val string) string {
		return fmt.Sprintf("%s, %s", accumulator, val)
	}

	h1, _ := slicez.Head(pStrSlice)
	var pStr = slicez.Fold(slicez.Tail(pStrSlice), joiner, h1)
	h2, _ := slicez.Head(nStrSlice)
	var nStr = slicez.Fold(slicez.Tail(nStrSlice), joiner, h2)

	fmt.Printf("(%s), (%s)\n", pStr, nStr)
	// (8, 6, 4), (-4, -6, -8)
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