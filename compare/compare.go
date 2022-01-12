package compare

import "constraints"

func Compare[N constraints.Ordered](e1 N, e2 N) int {
	switch {
	case e1 < e2:
		return -1
	case e1 > e2:
		return +1
	default:
		return 0
	}
}

func Identity[N comparable](n N) N {
	return n
}

func Equal[N comparable](a, b N) bool {
	return a == b
}
func Less[E constraints.Ordered](a, b E) bool {
	return a < b
}

func Negate[A any](f func(a, b A) bool) func(A, A) bool {
	return func(a, b A) bool {
		return !f(a, b)
	}
}

func EqualOf[N comparable](needle N) func(b N) bool {
	return func(b N) bool {
		return needle == b
	}
}

func NegateOf[A any](f func(A) bool) func(A) bool {
	return func(a A) bool {
		return !f(a)
	}
}
