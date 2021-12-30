# Henry 
> A test implementation of some functional concepts with go generics


## Example usage of some functional concepts
```go 
package main

import (
	"fmt"
	"github.com/crholm/henry/henry"
	"github.com/crholm/henry/henry/pipe"
)

func main() {

	even := func(index int, i int) bool {
		return i%2 == 0
	}
	negate := func(index int, i int) int {
		return -i
	}

	var pslice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var nslice = pipe.Of(pslice).Map(negate).Reverse().Slice()

	positive, negative := pipe.Of(pslice).
		Concat(nslice).                       // Reversing original slice, negating numbers and concatenating the res
		Filter(even).                         // Filtering and keeping even numbers
		DropLeft(1).                          // Dropping one number on the left
		DropRight(1).                         // Dropping one number of the right
		Reverse().                            // Revers the slice
		Partition(func(_ int, val int) bool { // Partitioning into slice into positive and negative numbers
			return val > 0
		})

	// Mapping number to string
	toStr := func(index int, v int) string {
		return fmt.Sprintf("%d", v)
	}
	var pStrSlice = henry.Map(positive, toStr)
	var nStrSlice = henry.Map(negative, toStr)

	// Joining []string to string
	joiner := func(index int, accumulator string, val string) string {
		return fmt.Sprintf("%s, %s", accumulator, val)
	}

	var pStr = henry.FoldLeft(henry.Tail(pStrSlice), joiner, henry.Head(pStrSlice).Value())
	var nStr = henry.FoldLeft(henry.Tail(nStrSlice), joiner, henry.Head(nStrSlice).Value())

	fmt.Printf("(%s), (%s)\n", pStr, nStr)
	// (8, 6, 4), (-4, -6, -8)
}
```



## Example usage a wrapped version of container/heap
```go 
import (
	"fmt"
	"github.com/crholm/henry/heap"
)

func main() {
	// Min heap
	h := heap.New[int](heap.Ordered[int], 1235, 543, 2)
	h.Push(4)
	h.Push(982130, 41, 15, 62, 1, 4, 11, 5, 64, 45, 4, 48, 85, 23, 12)

	fmt.Println("Min Heap popping")
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	fmt.Println()

	// Max heap
	h = heap.New[int](heap.Reverse(heap.Ordered[int]), 1235, 543, 2)
	h.Push(4)
	h.Push(982130, 41, 15, 62, 1, 4, 11, 5, 64, 45, 4, 48, 85, 23, 12)

	fmt.Println("Max Heap popping")
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	fmt.Println()
}
```


## Example usage a wrapped version of sort
```go 
import (
	"fmt"
	"github.com/crholm/henry/sort"
)

func main() {
	in := []int{2, 3, 5, 1, 12, 3, 6, 7, 34, 123, 65, 4631, 1, 1323}

	sort.Slice(in, sort.Reverse(sort.Ordered[int]))
	fmt.Println("Sorted in Descending order")
	fmt.Println(in)
	fmt.Println()

	sort.Slice(in, sort.Ordered[int])
	fmt.Println("Sorted in Ascending order")
	fmt.Println(in)
	fmt.Println()

	i, e := sort.Search(in, func(e int) bool {
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
	"github.com/crholm/henry"
	"github.com/crholm/henry/numbers"
)

func main() {

	ints := []int{1, 2, 3, 4, 5, 23, 12, 5, 231, 12, 12}
	fmt.Println("Integers")
	fmt.Println("Mean:", numbers.Mean(ints...))
	fmt.Println("Median:", numbers.Median(ints...))
	fmt.Println("Mode:", numbers.Mode(ints...))
	fmt.Println("Min:", numbers.Min(ints...))
	fmt.Println("Max:", numbers.Max(ints...))
	fmt.Println("Sum:", numbers.Sum(ints...))
	fmt.Println("x^2:", numbers.Pow(ints, 2))
	fmt.Println("Variance:", numbers.Variance(ints...)) // Sample
	fmt.Println("StdDev:", numbers.StdDev(ints...))     // Sample
	fmt.Println()

	floats := henry.Map(ints, numbers.MapFloat64[int])
	floats = henry.Map(floats, func(_ int, a float64) float64 {
		return a + 0.5
	})
	fmt.Println("Floats")
	fmt.Println("Mean:", numbers.Mean(floats...))
	fmt.Println("Median:", numbers.Median(floats...))
	fmt.Println("Mode:", numbers.Mode(floats...))
	fmt.Println("Min:", numbers.Min(floats...))
	fmt.Println("Max:", numbers.Max(floats...))
	fmt.Println("Sum:", numbers.Sum(floats...))
	fmt.Println("x^2:", numbers.Pow(floats, 2))
	fmt.Println("Variance:", numbers.Variance(floats...))
	fmt.Println("StdDev:", numbers.StdDev(floats...))

}
```