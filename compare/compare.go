package compare

import "constraints"

type IsLess[E any] func(a, b E) bool
type IsEqual[E any] func(a, b E) bool

func Less[E constraints.Ordered](a, b E) bool {
	return a < b
}
func Greater[E constraints.Ordered](a, b E) bool {
	return a > b
}

func Reverse[E any](less IsLess[E]) IsLess[E] {
	return func(a, b E) bool {
		return less(b, a)
	}
}

func Equal[N comparable](a, b N) bool {
	return a == b
}

func NotEqual[N comparable](a, b N) bool {
	return a != b
}

func EqualBy[N comparable](n N) N {
	return n
}
