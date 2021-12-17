package henry

import (
	"github.com/crholm/henry/maybe"
	"math/rand"
	"time"
)

func Each[A any](slice []A, apply func(i int, a A)) {
	for i, a := range slice {
		apply(i, a)
	}
}

func Concat[A any](slices ...[]A) []A {
	var ret []A
	for _, slice := range slices {
		ret = append(ret, slice...)
	}
	return ret
}

func Reverse[A any](slice []A) []A {
	l := len(slice)
	res := make([]A, l)
	for i, val := range slice {
		res[l-i-1] = val
	}
	return res
}

func Head[A any](slice []A) maybe.Value[A] {
	if len(slice) > 0 {
		maybe.Some(slice[0])
	}
	var zero A
	return maybe.None(zero)
}
func Tail[A any](slice []A) []A {
	return DropLeft(slice, 1)
}

func Last[A any](slice []A) maybe.Value[A] {
	if len(slice) > 0 {
		return maybe.Some(slice[len(slice)-1])
	}
	var zero A
	return maybe.None(zero)
}

func Nth[A any](slice []A, i int) maybe.Value[A] {
	if i < 0 {
		i = len(slice) + i
	}

	if i < len(slice) {
		return maybe.Some(slice[i])
	}
	var zero A
	return maybe.None(zero)
}

func TakeLeftWhile[A any](slice []A, take func(i int, a A) bool) []A {
	var res []A
	for i, val := range slice {
		if !take(i, val) {
			break
		}
		res = append(res, val)
	}
	return res
}
func TakeRightWhile[A any](slice []A, take func(i int, a A) bool) []A {
	var l = len(slice)
	var res []A
	for i := range slice {
		i = l - i - 1
		val := slice[i]
		if !take(i, val) {
			break
		}
		res = append([]A{val}, res...)
	}
	return res
}

func TakeLeft[A any](slice []A, i int) []A {
	return TakeLeftWhile(slice, func(j int, _ A) bool {
		return j < i
	})
}
func TakeRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	return TakeRightWhile(slice, func(j int, _ A) bool {
		return j > i
	})
}

func DropLeftWhile[A any](slice []A, drop func(i int, a A) bool) []A {
	if len(slice) == 0 {
		return nil
	}

	var index int = -1
	for i, val := range slice {
		if !drop(i, val) {
			break
		}
		index = i
	}

	var a []A

	if index == -1 {
		a = make([]A, len(slice))
		copy(a, slice)
		return a
	}

	if index+1 < len(slice) {
		a = make([]A, len(slice)-index-1)
		copy(a, slice[index+1:])
		return a
	}

	return a
}

func DropRightWhile[A any](slice []A, drop func(i int, a A) bool) []A {

	if len(slice) == 0 {
		return nil
	}

	var index int = -1
	var l = len(slice)
	for i := range slice {
		i = l - i - 1
		val := slice[i]
		if !drop(i, val) {
			break
		}
		index = i
	}
	var a []A
	if index == -1 {
		a = make([]A, len(slice))
		copy(a, slice)
		return a
	}

	if 0 < index && index < len(slice) {
		a = make([]A, index)
		copy(a, slice[:index])
		return a
	}
	return a

}

func DropLeft[A any](slice []A, i int) []A {
	return DropLeftWhile(slice, func(j int, _ A) bool {
		return j < i
	})
}
func DropRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	return DropRightWhile(slice, func(j int, _ A) bool {
		return j > i
	})
}

func Filter[A any](slice []A, include func(i int, a A) bool) []A {
	var res []A
	for i, val := range slice {
		if include(i, val) {
			res = append(res, val)
		}
	}
	return res
}

func Reject[A any](slice []A, exclude func(i int, a A) bool) []A {
	return Filter(slice, func(i int, a A) bool {
		return !exclude(i, a)
	})
}

func Every[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if !predicate(val) {
			return false
		}
	}
	return true
}
func Some[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if predicate(val) {
			return true
		}
	}
	return false
}

func Has[A any](slice []A, target A, predicate func(a A, b A) bool) bool {
	return Some(slice, func(a A) bool {
		return predicate(target, a)
	})
}

func Partition[A any](slice []A, predicate func(i int, a A) bool) (satisfied, notSatisfied []A) {
	for i, a := range slice {
		if predicate(i, a) {
			satisfied = append(satisfied, a)
			continue
		}
		notSatisfied = append(notSatisfied, a)
	}
	return satisfied, notSatisfied
}

func None[A any](slice []A, predicate func(A) bool) bool {
	return !Some(slice, predicate)
}

func Shuffle[A any](slice []A) []A {
	rand.Seed(time.Now().UnixNano())
	var ret []A
	for _, a := range slice {
		ret = append(ret, a)
	}
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

func Sample[A any](slice []A, n int) []A {
	var ret []A

	if n > len(slice) {
		n = len(slice)
	}

	if n > len(slice)/3 { // sqare root?
		ret = Shuffle(slice)
		return ret[:n]
	}

	idxs := map[int]struct{}{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		var idx int
		for {
			idx = rand.Intn(len(slice))
			_, found := idxs[idx]
			if found {
				continue
			}
			idxs[idx] = struct{}{}
			break
		}

		ret = append(ret, slice[idx])
	}
	return ret
}
