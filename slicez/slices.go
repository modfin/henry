package slicez

import (
	"errors"
	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/slicez/sort"
	"math/rand"
	"time"
)

// Equal takes two slices of that is of the interface comparable. It returns true if they are of equal length and each
// element in a[x] == b[x] for every element
func Equal[A comparable](s1, s2 []A) bool {
	return EqualFunc(s1, s2, compare.Equal[A])
}

// EqualFunc takes two slices and an equality check function. It returns true if they are of equal length and each
// element in eq(a[x], b[x]) == true for every element
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

// Index finds the first index of an element in an array. It returns -1 if it is not present
func Index[E comparable](s []E, needle E) int {
	return IndexFunc(s, func(e E) bool {
		return needle == e
	})
}

// IndexFunc finds the first index of an element where the passed in function returns true. It returns -1 if it is not present
func IndexFunc[E any](s []E, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

// LastIndex finds the last index of an element in an array. It returns -1 if it is not present
func LastIndex[E comparable](s []E, needle E) int {
	return LastIndexFunc(s, func(e E) bool {
		return e == needle
	})
}

// LastIndexFunc finds the last index of an element where the passed in function returns true. It returns -1 if it is not present
func LastIndexFunc[E any](s []E, f func(E) bool) int {
	n := len(s)

	for i := 0; i < n; i++ {
		if f(s[n-i-1]) {
			return n - i - 1
		}
	}
	return -1
}

// Cut will cut a slice into a left and a right part at the first instance where the needle is found. The needle is not included
func Cut[E comparable](s []E, needle E) (left, right []E, found bool) {
	return CutFunc(s, func(e E) bool {
		return e == needle
	})
}

// CutFunc will cut a slice into a left and a right part at the first instance where the on function returns true.
// The element that makes the "on" function return true will not be included.
func CutFunc[E any](s []E, on func(E) bool) (left, right []E, found bool) {
	i := IndexFunc(s, on)
	if i == -1 {
		return s, nil, false
	}
	return s[:i], s[i+1:], true
}

// Find will find the first instance of an element in a slice where the equal func returns true
func Find[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := IndexFunc(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

// FindLast will find the last instance of an element in a slice where the equal func returns true
func FindLast[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := LastIndexFunc(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

// Join will join a two-dimensional slice into a one dimensional slice with the glue slice between them.
// Similar to strings.Join or bytes.Join
func Join[E any](slices [][]E, glue []E) []E {
	if len(slices) == 0 {
		return []E{}
	}
	if len(slices) == 1 {
		return append([]E(nil), slices[0]...)
	}
	n := len(glue) * (len(slices) - 1)
	for _, v := range slices {
		n += len(v)
	}

	b := make([]E, n)
	bp := copy(b, slices[0])
	for _, v := range slices[1:] {
		bp += copy(b[bp:], glue)
		bp += copy(b[bp:], v)
	}
	return b
}

// Contains returns true if the needle is present in the slice
func Contains[E comparable](s []E, needle E) bool {
	return Index(s, needle) >= 0
}

// ContainsFunc returns true if the passed in func returns true on any of the element in the slice
func ContainsFunc[E any](s []E, f func(e E) bool) bool {
	return IndexFunc(s, f) >= 0
}

// Clone will create a copy of the slice
func Clone[E any](s []E) []E {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append([]E{}, s...)
}

// Compare will compare two slices
func Compare[E compare.Ordered](s1, s2 []E) int {
	return CompareFunc(s1, s2, compare.Compare[E])
}

// CompareFunc will compare two slices using a compare function
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

// Concat will concatenate supplied slices in the given order into a new slice
func Concat[A any](slices ...[]A) []A {
	var capacity int
	for _, s := range slices {
		capacity += len(s)
	}
	var ret = make([]A, 0, capacity)
	for _, slice := range slices {
		ret = append(ret, slice...)
	}
	return ret
}

// Reverse will return a news slice, but reversed of the original one
func Reverse[A any](slice []A) []A {
	l := len(slice)
	res := make([]A, l)
	for i, val := range slice {
		res[l-i-1] = val
	}
	return res
}

// Head will return the first element of the slice, or an error if the length of the slice is 0
func Head[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[0], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

// Tail will return a new slice with all but the first element
func Tail[A any](slice []A) []A {
	return Drop(slice, 1)
}

// Last will return the last element of the slice, or an error if the length of the slice is 0
func Last[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[len(slice)-1], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

// Nth will return the nth element in the slice. It returns the zero value if len(slice) == 0.
// Nth looks as the slice of a modul group and will wrap around from both ends. Eg Nth(-1) will return the last element
// and Nth(10) where len(slice) == 10 will return the first element
func Nth[A any](slice []A, i int) A {
	var zero A
	n := len(slice)
	if n == 0 {
		return zero
	}
	if n == 1 {
		return slice[0]
	}

	i = i % n

	if i < 0 {
		i = len(slice) + i
	}
	return slice[i]
}

// Each will apply the "apply" func on each element of the slice
func Each[A any](slice []A, apply func(a A)) {
	for _, a := range slice {
		apply(a)
	}
}

// TakeWhile will produce a new slice containing all elements from the left until the "take" func returns false
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

// TakeRightWhile will produce a new slice containing all elements from the right until the "take" func returns false
func TakeRightWhile[A any](slice []A, take func(a A) bool) []A {
	idx := len(slice) - 1
	for ; 0 <= idx; idx-- {
		if !take(slice[idx]) {
			break
		}
	}
	res := make([]A, len(slice)-1-idx)
	copy(res, slice[idx+1:])
	return res
}

// Take will produce a new slice containing the "i" first element of the passed in slice
func Take[A any](slice []A, i int) []A {
	var j int
	return TakeWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}

// TakeRight will produce a new slice containing the "i" last element of the passed in slice
func TakeRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return TakeRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

// DropWhile will produce a new slice, where the left most elements are dropped until the first instance the
// "drop" function returns false
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

// DropRightWhile will produce a new slice, where the right most elements are dropped until the first instance the
// "drop" function returns false
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

// Drop will produce a new slice where the "i" first element of the passed in slice are removed
func Drop[A any](slice []A, i int) []A {
	var j int
	return DropWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}

// DropRight will produce a new slice where the "i" last element of the passed in slice are removed
func DropRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return DropRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

// Filter will produce a new slice only containing elements where the "include" function returns true
func Filter[A any](slice []A, include func(a A) bool) []A {
	res := make([]A, 0, len(slice)/2)
	for _, val := range slice {
		if include(val) {
			res = append(res, val)
		}
	}
	return res
}

// Reject is the complement of Filter and will produce a new slice only containing elements where the "exclude" function returns false
func Reject[A any](slice []A, exclude func(a A) bool) []A {
	return Filter(slice, func(a A) bool {
		return !exclude(a)
	})
}

// Every returns true if every element in the slice is equal to the needle
func Every[A comparable](slice []A, needle A) bool {
	return EveryFunc(slice, compare.EqualOf[A](needle))

}

// EveryFunc returns true if the predicate function returns true for every element in the slice
func EveryFunc[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if !predicate(val) {
			return false
		}
	}
	return true
}

// Some returns true there exist an element in the slice that is equal to the needle, an alias for Contains
func Some[A comparable](slice []A, needle A) bool {
	return SomeFunc(slice, compare.EqualOf[A](needle))
}

// SomeFunc returns true if there is an element in the slice for which the predicate function returns true
func SomeFunc[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if predicate(val) {
			return true
		}
	}
	return false
}

// None returns true if there is no element in the slice that matches the needle
func None[A comparable](slice []A, needle A) bool {
	return !SomeFunc(slice, compare.EqualOf[A](needle))
}

// NoneFunc returns true if there are no element in the slice for which the predicate function returns true
func NoneFunc[A any](slice []A, predicate func(A) bool) bool {
	return !SomeFunc(slice, predicate)
}

// Partition will partition a slice into to slices. One where every element for which the predicate function returns true
// and where it returns false
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

// Shuffle will return a new slice where the elements from the original slice is shuffled
func Shuffle[A any](slice []A) []A {
	var ret = append([]A{}, slice...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

// Sample will return a slice containing "n" random elements from the original slice
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

// Sort will return a new slice that is sorted in the natural order
func Sort[A compare.Ordered](slice []A) []A {
	return SortFunc(slice, compare.Less[A])
}

// SortFunc will return a new slice that is sorted using the supplied less function for natural ordering
func SortFunc[A any](slice []A, less func(a, b A) bool) []A {
	var res = append([]A{}, slice...)
	sort.Slice(res, less)
	return res
}

// Search given a slice data sorted in ascending order,
// the call
//
//	Search[int](data, func(e int) bool { return e >= 23 })
//
// returns the smallest index i and element e such that e >= 23.
func Search[A any](slice []A, f func(e A) bool) (index int, e A) {
	return sort.Search(slice, f)
}

// Compact will remove any duplicate elements following each other in a slice, eg
//
//	{1,1,2,1,2,2,2} => {1,2,1,2}
func Compact[A comparable](slice []A) []A {
	return CompactFunc(slice, compare.Equal[A])
}

// CompactFunc will remove any duplicate elements following each other determined by the equal func.
// eg removing duplicate whitespaces from a string might look like
//
//	CompactFunc([]rune("a    b"), func(a, b rune) {
//	 	return a == ' ' && a == b
//	})
//
// resulting in "a b"
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

// Max returns the largest element of the slice
func Max[E compare.Ordered](slice ...E) E {
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

// Min returns the smalest element of the slice
func Min[E compare.Ordered](slice ...E) E {
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

// Flatten will flatten a 2d slice into a 1d slice
func Flatten[A any](slice [][]A) []A {
	var capacity int
	for _, s := range slice {
		capacity += len(s)
	}
	var res = make([]A, 0, capacity)
	for _, val := range slice {
		res = append(res, val...)
	}
	return res
}

// Map will map entries in one slice to entries in another slice
func Map[A any, B any](slice []A, f func(a A) B) []B {
	res := make([]B, 0, len(slice))
	for _, a := range slice {
		res = append(res, f(a))
	}
	return res
}

// FlatMap will map entries in one slice to enteris in another slice and then flatten the map
func FlatMap[A any, B any](slice []A, f func(a A) []B) []B {
	return Flatten(Map(slice, f))
}

// Fold will iterate through the slice, from the left, and execute the combine function on each element accumulating the result into a value
func Fold[I any, A any](slice []I, combined func(accumulator A, val I) A, init A) A {
	for _, val := range slice {
		init = combined(init, val)
	}
	return init
}

// FoldRight will iterate through the slice, from the right, and execute the combine function on each element accumulating the result into a value
func FoldRight[I any, A any](slice []I, combined func(accumulator A, val I) A, init A) A {
	l := len(slice)
	for i := range slice {
		i := l - i - 1
		init = combined(init, slice[i])
	}
	return init
}

// SliceToMap will iterate over a slice turning each object into a key/value pair in a map
func SliceToMap[E any, K comparable, V any](slice []E, mapper func(a E) (key K, value V)) map[K]V {
	return Associate(slice, mapper)
}

// Associate will iterate over a slice turning each object into a key/value pair in a map
func Associate[E any, K comparable, V any](slice []E, mapper func(a E) (key K, value V)) map[K]V {
	return Fold(slice, func(acc map[K]V, e E) map[K]V {
		k, v := mapper(e)
		acc[k] = v
		return acc
	}, map[K]V{})
}

// KeyBy will iterate through the slice and create a map where the key function generates the key value pair.
// If multiple values generate the same key, it is the first value that is stored in the map
func KeyBy[A any, B comparable](slice []A, key func(a A) B) map[B]A {

	m := make(map[B]A)

	for _, v := range slice {
		k := key(v)
		_, exist := m[k]
		if exist {
			continue
		}
		m[k] = v
	}
	return m
}

// GroupBy will iterate through the slice and create a map where entries are grouped into slices using the key function generates the key .
func GroupBy[A any, B comparable](slice []A, key func(a A) B) map[B][]A {

	m := make(map[B][]A)

	for _, v := range slice {
		k := key(v)
		m[k] = append(m[k], v)
	}
	return m
}

// Uniq returns a slice with no duplicate entries
func Uniq[A comparable](slice []A) []A {
	return UniqBy(slice, compare.Identity[A])
}

// UniqBy returns a slice with no duplicate entries using the by function to determine the key
func UniqBy[A any, B comparable](slice []A, by func(a A) B) []A {
	set := make(map[B]struct{}, len(slice))
	res := make([]A, 0, len(slice))
	for _, e := range slice {
		key := by(e)
		_, exist := set[key]
		if exist {
			continue
		}
		set[key] = struct{}{}
		res = append(res, e)
	}
	return res
}

// Union will return the union of an arbitrary number of slices. This is equivalent to Uniq(Concat(sliceA, sliceB))
func Union[A comparable](slices ...[]A) []A {
	return UnionBy(compare.Identity[A], slices...)
}

// UnionBy will return the union of an arbitrary number of slices where the by function is used to determine the key. This is equivalent to UniqBy(Concat(sliceA, sliceB), by)
func UnionBy[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var maxCapacity = 0
	for _, slice := range slices {
		if len(slice) > maxCapacity {
			maxCapacity = len(slice)
		}
	}
	var res = make([]A, 0, maxCapacity)
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

// Intersection returns a slice containing the intersection between passed in slices
func Intersection[A comparable](slices ...[]A) []A {
	return IntersectionBy(compare.Identity[A], slices...)
}

// IntersectionBy returns a slice containing the intersection between passed in slices determined by the "by" function
func IntersectionBy[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var res = UniqBy(slices[0], by)
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

// Difference returns a slice containing the difference between passed in slices
func Difference[A comparable](slices ...[]A) []A {
	return DifferenceBy(compare.Identity[A], slices...)
}

// DifferenceBy returns a slice containing the difference between passed in slices determined by the "by" function
func DifferenceBy[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var exclude = map[B]bool{}
	for _, v := range IntersectionBy(by, slices...) {
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

// Complement returns a slice containing all elements in "b" that is not present in "a"
func Complement[A comparable](a, b []A) []A {
	return ComplementBy(compare.Identity[A], a, b)
}

// ComplementBy returns a slice containing all elements in "b" that is not present in "a" determined using the "by" function
func ComplementBy[A any, B comparable](by func(a A) B, a, b []A) []A {
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

// Zip will zip two slices, a and b, into one slice, c, using the zip function to combined elements
func Zip[A any, B any, C any](aSlice []A, bSlice []B, zipper func(a A, b B) C) []C {
	var capacity = Min(len(aSlice), len(bSlice))
	var cSlice = make([]C, 0, capacity)
	for i := 0; i < capacity; i++ {
		cSlice = append(cSlice, zipper(aSlice[i], bSlice[i]))
	}
	return cSlice
}

// Unzip will unzip a slice slices, c, into two slices, a and b, using the supplied unziper function
func Unzip[A any, B any, C any](cSlice []C, unzipper func(c C) (a A, b B)) ([]A, []B) {
	var aSlice = make([]A, 0, len(cSlice))
	var bSlice = make([]B, 0, len(cSlice))
	for _, c := range cSlice {
		a, b := unzipper(c)
		aSlice = append(aSlice, a)
		bSlice = append(bSlice, b)
	}
	return aSlice, bSlice

}

// Zip2 will zip three slices, a, b and c, into one slice, d, using the zip function to combined elements
func Zip2[A any, B any, C any, D any](aSlice []A, bSlice []B, cSlice []C, zipper func(a A, b B, c C) D) []D {
	var capacity = Min(len(aSlice), len(bSlice), len(cSlice))
	var dSlice = make([]D, 0, capacity)
	for i := 0; i < capacity; i++ {
		dSlice = append(dSlice, zipper(aSlice[i], bSlice[i], cSlice[i]))
	}
	return dSlice
}

// Unzip2 will unzip a slice slices, d, into three slices, a, b and c, using the supplied unziper function
func Unzip2[A any, B any, C any, D any](dSlice []D, unzipper func(d D) (a A, b B, c C)) ([]A, []B, []C) {
	var aSlice = make([]A, 0, len(dSlice))
	var bSlice = make([]B, 0, len(dSlice))
	var cSlice = make([]C, 0, len(dSlice))
	for _, d := range dSlice {
		a, b, c := unzipper(d)
		aSlice = append(aSlice, a)
		bSlice = append(bSlice, b)
		cSlice = append(cSlice, c)
	}
	return aSlice, bSlice, cSlice
}
