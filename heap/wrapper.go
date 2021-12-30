package heap

import (
	heap2 "container/heap"
	"github.com/crholm/henry/compare"
	"sync"
)

type Heap[E any] interface {
	Pop() E
	Push(e ...E)
	Slice() []E
	Range(apply func(e E))
	Len() int
}

func New[E any](less compare.IsLess[E], init ...E) Heap[E] {
	h := &heap[E]{data: init, less: less}

	heap2.Init(h)

	return &wrapper[E]{heap: h}
}

type wrapper[E any] struct {
	mu   sync.Mutex
	heap *heap[E]
}

func (w *wrapper[E]) pop() E {
	e := heap2.Pop(w.heap)
	return e.(E)
}

func (w *wrapper[E]) Pop() E {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.pop()
}
func (w *wrapper[E]) push(e ...E) {
	if len(e) == 1 {
		heap2.Push(w.heap, e[0])
		return
	}
	// Probably not a great idea...
	w.heap.data = append(w.heap.data, e...)
	heap2.Init(w.heap)

}
func (w *wrapper[E]) Push(e ...E) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.push(e...)
}
func (w *wrapper[E]) len() int {
	return w.heap.Len()
}
func (w *wrapper[E]) Len() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.len()
}

func (w *wrapper[E]) Range(apply func(e E)) {
	w.mu.Lock()
	defer w.mu.Unlock()
	for w.len() > 0 {
		apply(w.pop())
	}
}

func (w *wrapper[E]) Slice() []E {
	w.mu.Lock()
	defer w.mu.Unlock()
	var s []E
	for w.len() > 0 {
		s = append(s, w.pop())
	}
	return s
}

type heap[E any] struct {
	less compare.IsLess[E]
	data []E
}

func (h heap[E]) Len() int           { return len(h.data) }
func (h heap[E]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h heap[E]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *heap[E]) Push(x interface{}) {
	h.data = append(h.data, x.(E))
}

func (h *heap[E]) Pop() interface{} {
	n := len(h.data)
	e := h.data[n-1]
	h.data = h.data[0 : n-1]
	return e
}
