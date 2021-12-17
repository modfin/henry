package main

import (
	"fmt"
	"github.com/crholm/henry"
)

func main() {

	even := func(index int, i int) bool {
		return i%2 == 0
	}
	negate := func(index int, i int) int {
		return -i
	}

	var pslice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var nslice = henry.PipeOf(pslice).Map(negate).Reverse().Slice()

	var all = henry.PipeOf(pslice).
		Concat(nslice). // Reversing original slice, negating numbers and concatenating the res
		Filter(even).   // Filtering and keeping even numbers
		DropLeft(1).    // Dropping one number on the left
		DropRight(1).   // Dropping one number of the right
		Reverse().      // Revers the slice
		Slice()         // Retrive the resulting slice from pipe

	// Partitioning into slice into positive and negative numbers
	positive, negative := henry.Partition(all, func(_ int, val int) bool {
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
