package heap

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	h := Min[int](Ordered[int], 3, 2, 1, 67, 83, 21, 3, 1, 23)
	h.Push(3)
	h.Push(4)
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}
	h = Min[int](Ordered[int], 3, 2, 1, 67, 83, 21, 3, 1, 23)
	h.Push(3)
	h.Push(4)
	fmt.Println(h.Slice())

	h = Min[int](Ordered[int], 3, 2, 1, 67, 83, 21, 3, 1, 23)
	h.Push(3)
	h.Push(4)
	h.Range(func(e int) {
		fmt.Println(e)
	})

}
