package main

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
