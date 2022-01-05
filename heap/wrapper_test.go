package heap

import (
	"github.com/crholm/go18exp/compare"
	"testing"
)

func TestMax(t *testing.T) {
	h := New[int](compare.Less[int])
	h.Push(4)
	h.Push(982130, 41, 15, 62, 1, 4, 11, 5, 64, 45, 4, 48, 85, 23, 12)

	i := 0
	for h.Len() > 0 {
		p := h.Pop()
		if p < i {
			t.Log("Expected p to be greater or equal to i", p, ">=", i)
			t.Fail()
		}
		i = p
	}

	h = New[int](compare.Less[int], 3, 2, 1, 67, 83, 21, 3, 1, 23)
	h.Push(3)
	h.Push(4)
	i = 0
	for _, p := range h.Slice() {
		if p < i {
			t.Log("Expected p to be greater or equal to i", p, ">=", i)
			t.Fail()
		}
		i = p
	}

	h = New[int](compare.Less[int], 3, 2, 1, 67, 83, 21, 3, 1, 23)
	h.Push(3)
	h.Push(4)

	i = 0
	h.Range(func(p int) {
		if p < i {
			t.Log("Expected p to be greater or equal to i", p, ">=", i)
			t.Fail()
		}
		i = p
	})

}
