package slicez

func Flatten[A any](slice [][]A) []A {
	var res []A
	for _, val := range slice {
		res = append(res, val...)
	}
	return res
}

func Map[A any, B any](slice []A, f func(a A) B) []B {
	res := make([]B, 0, len(slice))
	for _, a := range slice {
		res = append(res, f(a))
	}
	return res
}

func FlatMap[A any, B any](slice []A, f func(a A) []B) []B {
	return Flatten(Map(slice, f))
}

func Fold[I any, A any](slice []I, combined func(accumulator A, val I) A, accumulator A) A {
	for _, val := range slice {
		accumulator = combined(accumulator, val)
	}
	return accumulator
}

func FoldRight[I any, A any](slice []I, combined func(accumulator A, val I) A, accumulator A) A {
	l := len(slice)
	for i := range slice {
		i := l - i - 1
		accumulator = combined(accumulator, slice[i])
	}
	return accumulator
}

func KeyBy[A any, B comparable](slice []A, key func(a A) B) map[B]A {

	m := make(map[B]A)

	for _, v := range slice {
		k := key(v)
		m[k] = v
	}
	return m
}

func GroupBy[A any, B comparable](slice []A, key func(a A) B) map[B][]A {

	m := make(map[B][]A)

	for _, v := range slice {
		k := key(v)
		m[k] = append(m[k], v)
	}
	return m
}

func Uniq[A comparable](slice []A) []A {
	return UniqBy(func(a A) A {
		return a
	}, slice)
}

func UniqBy[A any, B comparable](by func(a A) B, slice []A) []A {
	var res []A
	var set = map[B]struct{}{}
	for _, e := range slice {
		key := by(e)
		_, ok := set[key]
		if ok {
			continue
		}
		set[key] = struct{}{}
		res = append(res, e)
	}
	return res
}

func Union[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var res []A
	var set = map[B]struct{}{}
	for _, slice := range slices {
		for _, e := range slice {
			key := by(e)
			_, ok := set[key]
			if ok {
				continue
			}
			set[key] = struct{}{}
			res = append(res, e)
		}
	}
	return res
}

func Intersection[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var res = UniqBy(by, slices[0])
	for _, slice := range slices[1:] {
		var set = map[B]bool{}
		for _, e := range slice {
			set[by(e)] = true
		}
		res = Filter(res, func(a A) bool {
			return set[by(a)]
		})
	}
	return res
}

func Difference[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var exclude = map[B]bool{}
	for _, v := range Intersection(by, slices...) {
		exclude[by(v)] = true
	}

	var res []A
	for _, slice := range slices {
		for _, e := range slice {
			key := by(e)
			if exclude[key] {
				continue
			}
			exclude[key] = true
			res = append(res, e)
		}
	}
	return res
}

func Complement[A any, B comparable](by func(a A) B, a, b []A) []A {
	if len(a) == 0 {
		return b
	}

	var exclude = map[B]bool{}
	for _, e := range a {
		exclude[by(e)] = true
	}

	var res []A
	for _, e := range b {
		key := by(e)
		if exclude[key] {
			continue
		}
		exclude[key] = true
		res = append(res, e)
	}

	return res
}
