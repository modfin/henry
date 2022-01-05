package main

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
