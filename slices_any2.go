package henry

func Flatten[A any](slice [][]A) []A {
	var res []A
	for _, val := range slice {
		res = append(res, val...)
	}
	return res
}

func Map[A any, B any](slice []A, f func(i int, a A) B) []B {
	res := make([]B, 0, len(slice))
	for i, a := range slice {
		res = append(res, f(i, a))
	}
	return res
}

func FlatMap[A any, B any](slice []A, f func(i int, a A) []B) []B {
	return Flatten(Map(slice, f))
}

func FoldLeft[I any, A any](slice []I, combined func(i int, accumulator A, val I) A, accumulator A) A {
	for i, val := range slice {
		accumulator = combined(i, accumulator, val)
	}
	return accumulator
}

func FoldRight[I any, A any](slice []I, combined func(i int, accumulator A, val I) A, accumulator A) A {
	l := len(slice)
	for i := range slice {
		i := l - i - 1
		accumulator = combined(i, accumulator, slice[i])
	}
	return accumulator
}

func KeyBy[A any, B comparable](slice []A, key func(i int, a A) B) map[B]A {

	m := make(map[B]A)

	for i, v := range slice {
		k := key(i, v)
		m[k] = v
	}
	return m
}

func GroupBy[A any, B comparable](slice []A, key func(i int, a A) B) map[B][]A {

	m := make(map[B][]A)

	for i, v := range slice {
		k := key(i, v)
		m[k] = append(m[k], v)
	}
	return m
}
