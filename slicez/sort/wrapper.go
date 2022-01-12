package sort

import (
	sort2 "sort"
)

type sortable[E any] struct {
	data []E
	less func(a, b E) bool
}

func (s sortable[E]) Len() int {
	return len(s.data)
}
func (s sortable[E]) Less(i, j int) bool {
	return s.less(s.data[i], s.data[j])
}
func (s sortable[E]) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func Slice[E any](data []E, less func(a, b E) bool) {
	sort2.Sort(sortable[E]{
		data: data,
		less: less,
	})

}

func StableSlice[E any](data []E, less func(a, b E) bool) {
	sort2.Stable(sortable[E]{
		data: data,
		less: less,
	})
}

func IsSorted[E any](data []E, less func(a, b E) bool) bool {
	return sort2.IsSorted(sortable[E]{
		data: data,
		less: less,
	})
}

// Search given a slice data sorted in ascending order,
// the call Search[int](data, func(e int) bool { return e >= 23 })
// returns the smallest index i and element e such that e >= 23.
func Search[E any](data []E, f func(e E) bool) (int, E) {
	i := sort2.Search(len(data)-1, func(i int) bool {
		return f(data[i])
	})
	return i, data[i]
}
