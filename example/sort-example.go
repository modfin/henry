package main

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
