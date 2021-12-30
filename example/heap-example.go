package main

import (
	"fmt"
	"github.com/crholm/henry/compare"
	"github.com/crholm/henry/heap"
)

func main() {
	// Min heap
	h := heap.New[int](compare.Less[int], 1235, 543, 2)
	h.Push(4)
	h.Push(982130, 41, 15, 62, 1, 4, 11, 5, 64, 45, 4, 48, 85, 23, 12)

	fmt.Println("Min Heap popping")
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	fmt.Println()

	// Max heap
	h = heap.New[int](compare.Reverse(compare.Less[int]), 1235, 543, 2)
	h.Push(4)
	h.Push(982130, 41, 15, 62, 1, 4, 11, 5, 64, 45, 4, 48, 85, 23, 12)

	fmt.Println("Max Heap popping")
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	fmt.Println()
}
