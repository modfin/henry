package heap

import "constraints"

type Heap[E any] interface {
	Pop() E
	Push(e E)
	Slice() []E
	Range(apply func(e E))
	Len() int
}

type heap[E any] struct {
	eless func(a, b E) bool
	data  []E
}

func Ordered[E constraints.Ordered](a, b E) bool {
	return a < b
}

func Min[E any](less func(a, b E) bool, init ...E) Heap[E] {
	h := &heap[E]{
		eless: less,
	}

	for _, v := range init {
		h.data = append(h.data, v)
	}
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}

	return h
}

func Max[E any](less func(a, b E) bool, init ...E) Heap[E] {
	return Min(func(a, b E) bool {
		return less(b, a)
	}, init...)
}

func (h *heap[E]) less(i, j int) bool {
	return h.eless(h.data[i], h.data[j])
}

func (h *heap[E]) Len() int {
	return len(h.data)
}
func (h *heap[E]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *heap[E]) Slice() []E {
	var s []E
	for h.Len() > 0 {
		s = append(s, h.Pop())
	}
	return s
}
func (h *heap[E]) Range(apply func(e E)) {
	for h.Len() > 0 {
		apply(h.Pop())
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *heap[E]) Push(e E) {
	h.data = append(h.data, e)
	h.up(h.Len() - 1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *heap[E]) Pop() E {
	n := h.Len() - 1
	h.swap(0, n)
	h.down(0, n)
	return h.pop()
}
func (h *heap[E]) pop() (e E) {
	e, h.data = h.data[h.Len()-1], h.data[:h.Len()-1]
	return e
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *heap[E]) Remove(i int) any {
	n := h.Len() - 1
	if n != i {
		h.swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *heap[E]) Fix(i int) {
	if !h.down(i, len(h.data)) {
		h.up(i)
	}
}

func (h *heap[E]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *heap[E]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}
