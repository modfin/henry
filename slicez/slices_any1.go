package slicez

import (
	"constraints"
	"errors"
	"fmt"
	"github.com/modfin/go18exp/compare"
	"github.com/modfin/go18exp/slicez/sort"
	"math/rand"
	"time"
)

func Equal[A comparable](s1, s2 []A) bool {
	return EqualFunc(s1, s2, compare.Equal[A])
}
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

func Index[E comparable](s []E, needle E) int {
	return IndexFunc(s, func(e E) bool {
		return needle == e
	})
}

func IndexFunc[E any](s []E, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

func Contains[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

func Compare[E constraints.Ordered](s1, s2 []E) int {
	return CompareFunc(s1, s2, compare.Compare[E])
}

func Clone[E any](s []E) []E {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append([]E{}, s...)
}

func CompareFunc[E1, E2 any](s1 []E1, s2 []E2, cmp func(E1, E2) int) int {
	s2len := len(s2)
	for i, v1 := range s1 {
		if i >= s2len {
			return +1
		}
		v2 := s2[i]
		if c := cmp(v1, v2); c != 0 {
			return c
		}
	}
	if len(s1) < s2len {
		return -1
	}
	return 0
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

func Head[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[0], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

func Last[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[len(slice)-1], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

func Nth[A any](slice []A, i int) (A, error) {
	if i < 0 {
		i = len(slice) + i
	}

	if i < len(slice) {
		return slice[i], nil
	}
	var zero A
	return zero, fmt.Errorf("slice of len %d does not contain element %d", len(slice), i)
}

func Tail[A any](slice []A) []A {
	return Drop(slice, 1)
}

func Each[A any](slice []A, apply func(a A)) {
	for _, a := range slice {
		apply(a)
	}
}

func TakeWhile[A any](slice []A, take func(a A) bool) []A {
	var res []A
	for _, val := range slice {
		if !take(val) {
			break
		}
		res = append(res, val)
	}
	return res
}
func TakeRightWhile[A any](slice []A, take func(a A) bool) []A {
	var l = len(slice)
	var res []A
	for i := range slice {
		i = l - i - 1
		val := slice[i]
		if !take(val) {
			break
		}
		res = append([]A{val}, res...)
	}
	return res
}

func Take[A any](slice []A, i int) []A {
	var j int
	return TakeWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}
func TakeRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return TakeRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

func DropWhile[A any](slice []A, drop func(a A) bool) []A {
	if len(slice) == 0 {
		return nil
	}

	var index int = -1
	for i, val := range slice {
		if !drop(val) {
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

func DropRightWhile[A any](slice []A, drop func(a A) bool) []A {

	if len(slice) == 0 {
		return nil
	}

	var index int = -1
	var l = len(slice)
	for i := range slice {
		i = l - i - 1
		val := slice[i]
		if !drop(val) {
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

func Drop[A any](slice []A, i int) []A {
	var j int
	return DropWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}
func DropRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return DropRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

func Filter[A any](slice []A, include func(a A) bool) []A {
	var res []A
	for _, val := range slice {
		if include(val) {
			res = append(res, val)
		}
	}
	return res
}

func Reject[A any](slice []A, exclude func(a A) bool) []A {
	return Filter(slice, func(a A) bool {
		return !exclude(a)
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

func Has[A comparable](slice []A, needle A) bool {
	return Some(slice, func(a A) bool {
		return needle == a
	})
}

func Partition[A any](slice []A, predicate func(a A) bool) (satisfied, notSatisfied []A) {
	for _, a := range slice {
		if predicate(a) {
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

func Sort[A constraints.Ordered](slice []A) []A {
	return SortFunc(slice, compare.Less[A])
}
func SortFunc[A any](slice []A, less func(a, b A) bool) []A {
	var res = make([]A, len(slice))
	copy(res, slice)
	sort.Slice(res, less)
	return res
}

func Search[A any](slice []A, f func(e A) bool) (index int, e A) {
	return sort.Search(slice, f)
}

func Compact[A comparable](slice []A) []A {
	return CompactFunc(slice, compare.Equal[A])
}

func CompactFunc[A any](slice []A, equal func(a, b A) bool) []A {
	if len(slice) == 0 {
		return slice
	}
	head := slice[0]
	last := head
	tail := Fold(slice[1:], func(accumulator []A, current A) []A {
		if equal(last, current) {
			return accumulator
		}
		last = current
		return append(accumulator, current)
	}, []A{})
	return append([]A{head}, tail...)
}

func Max[E constraints.Ordered](slice ...E) E {
	var zero E
	if slice == nil || len(slice) == 0 {
		return zero
	}
	cur := slice[0]
	for _, c := range slice {
		if cur < c {
			cur = c
		}
	}
	return cur
}

func Min[E constraints.Ordered](slice ...E) E {
	var zero E
	if slice == nil || len(slice) == 0 {
		return zero
	}
	cur := slice[0]
	for _, c := range slice {
		if cur > c {
			cur = c
		}
	}
	return cur
}
