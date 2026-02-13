package slicez

import (
	"errors"
	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/slicez/sort"
	"math/rand"
)

// Equal takes two slices of that is of the interface comparable. It returns true if they are of equal length and each
// element in a[x] == b[x] for every element
func Equal[A comparable](s1, s2 []A) bool {
	return EqualBy(s1, s2, compare.Equal[A])
}

// EqualBy takes two slices and an equality check function. It returns true if they are of equal length and each
// element in eq(a[x], b[x]) == true for every element
func EqualBy[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
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
	return IndexBy(s, func(e E) bool {
		return needle == e
	})
}

// IndexBy finds the first index of an element where the passed in function returns true. It returns -1 if it is not present
func IndexBy[E any](s []E, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

// LastIndex finds the last index of an element in an array. It returns -1 if it is not present
func LastIndex[E comparable](s []E, needle E) int {
	return LastIndexBy(s, func(e E) bool {
		return e == needle
	})
}

// LastIndexBy finds the last index of an element where the passed in function returns true. It returns -1 if it is not present
func LastIndexBy[E any](s []E, f func(E) bool) int {
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
	return CutBy(s, func(e E) bool {
		return e == needle
	})
}

// CutBy will cut a slice into a left and a right part at the first instance where the on function returns true.
// The element that makes the "on" function return true will not be included.
func CutBy[E any](s []E, on func(E) bool) (left, right []E, found bool) {
	i := IndexBy(s, on)
	if i == -1 {
		return s, nil, false
	}
	return s[:i], s[i+1:], true
}

// Replace replaces occurrences needle in haystack n times. It Replaces all for n < 0
func Replace[E comparable](haystack []E, needle E, replacement E, n int) []E {
	return Map(haystack, func(e E) E {
		if n != 0 && e == needle {
			n--
			return replacement
		}
		return e
	})
}

// ReplaceFirst replaces first occurrences needle in haystack.
func ReplaceFirst[E comparable](haystack []E, needle E, replacement E) []E {
	return Replace(haystack, needle, replacement, 1)
}

// ReplaceAll replaces all occurrences needle in haystack.
func ReplaceAll[E comparable](haystack []E, needle E, replacement E) []E {
	return Replace(haystack, needle, replacement, -1)
}

// Find will find the first instance of an element in a slice where the equal func returns true
func Find[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := IndexBy(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

// FindLast will find the last instance of an element in a slice where the equal func returns true
func FindLast[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := LastIndexBy(s, equal)
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

// ContainsBy returns true if the passed in func returns true on any of the element in the slice
func ContainsBy[E any](s []E, f func(e E) bool) bool {
	return IndexBy(s, f) >= 0
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
	return CompareBy(s1, s2, compare.Compare[E])
}

// CompareBy will compare two slices using a compare function
func CompareBy[E1, E2 any](s1 []E1, s2 []E2, cmp func(E1, E2) int) int {
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

// RepeatBy creates a slice of i length and assign each element with result of the by function
func RepeatBy[A any](i int, by func(i int) A) []A {
	res := make([]A, 0, i)

	for n := 0; n < i; n++ {
		res = append(res, by(n))
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

// Initial gets all but the last element of the slice
func Initial[A any](slice []A) []A {
	return DropRight(slice, 1)
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

// ForEach will apply the "apply" func on each element of the slice
func ForEach[A any](slice []A, apply func(a A)) {
	for _, a := range slice {
		apply(a)
	}
}

// ForEachRight will apply the "apply" func on each element of the slice
func ForEachRight[A any](slice []A, apply func(a A)) {
	length := len(slice)
	for i := 0; i < length; i++ {
		apply(slice[length-1-i])
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
	// Pre-allocate with full capacity, then truncate to actual size
	// This avoids reallocation in the common case where filter keeps many elements
	res := make([]A, 0, len(slice))
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

// Without creates a new slice excluding all given values.
func Without[A comparable](slice []A, exclude ...A) []A {
	set := Set(exclude)
	return Reject(slice, func(a A) bool {
		return set[a]
	})
}

// Every returns true if every element in the slice is equal to the needle
func Every[A comparable](slice []A, needle A) bool {
	return EveryBy(slice, compare.EqualOf[A](needle))

}

// EveryBy returns true if the predicate function returns true for every element in the slice
func EveryBy[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if !predicate(val) {
			return false
		}
	}
	return true
}

// Some returns true there exist an element in the slice that is equal to the needle, an alias for Contains
func Some[A comparable](slice []A, needle A) bool {
	return SomeBy(slice, compare.EqualOf[A](needle))
}

// SomeBy returns true if there is an element in the slice for which the predicate function returns true
func SomeBy[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if predicate(val) {
			return true
		}
	}
	return false
}

// None returns true if there is no element in the slice that matches the needle
func None[A comparable](slice []A, needle A) bool {
	return !SomeBy(slice, compare.EqualOf[A](needle))
}

// NoneBy returns true if there are no element in the slice for which the predicate function returns true
func NoneBy[A any](slice []A, predicate func(A) bool) bool {
	return !SomeBy(slice, predicate)
}

// Partition will partition a slice into to two slices. One where every element for which the predicate function returns true
// and where it returns false
func Partition[A any](slice []A, predicate func(a A) bool) (satisfied, notSatisfied []A) {
	// Pre-allocate both slices with estimated capacity
	// In worst case, one slice could have all elements, so we allocate len(slice) for each
	// but the actual memory won't be used until elements are appended
	satisfied = make([]A, 0, len(slice))
	notSatisfied = make([]A, 0, len(slice))

	for _, a := range slice {
		if predicate(a) {
			satisfied = append(satisfied, a)
			continue
		}
		notSatisfied = append(notSatisfied, a)
	}
	return satisfied, notSatisfied
}

// PartitionBy will partition a slice into to a slice of slices.
// Returns an array of elements split into groups.
// The order of grouped values is determined by the order they occur in collection.
// The grouping is generated from the results of running each element of collection through iteratee.
func PartitionBy[A any, B comparable](slice []A, by func(a A) B) [][]A {
	if len(slice) == 0 {
		return nil
	}

	m := make(map[B][]A, len(slice)/2) // Pre-allocate with estimated capacity
	var order []B
	var seen = make(map[B]struct{}, len(slice)/2)

	for _, v := range slice {
		k := by(v)
		m[k] = append(m[k], v)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			order = append(order, k)
		}
	}

	// Build result without calling Uniq (which creates another map)
	result := make([][]A, len(order))
	for i, k := range order {
		result[i] = m[k]
	}
	return result
}

// Chunk will make de-flatten a slice into chunks
// eg slicez.Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
// // [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}
func Chunk[A any](slice []A, n int) [][]A {
	var i int
	var c = n
	return PartitionBy(slice, func(a A) int {
		r := i
		c--
		if c == 0 {
			c = n
			i++
		}
		return r
	})
}

// Interleave Round-robin alternating input slices and sequentially appending value at index into result.
// interleaved := Interleave([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func Interleave[A any](slices ...[]A) []A {
	var total int
	var max int
	for _, s := range slices {
		total += len(s)
		max = Max(max, len(s))
	}
	var res = make([]A, 0, total)
	for i := 0; i < max; i++ {
		for _, s := range slices {
			if i < len(s) {
				res = append(res, s[i])
			}
		}
	}
	return res
}

// Shuffle will return a new slice where the elements from the original slice is shuffled
func Shuffle[A any](slice []A) []A {
	var ret = append([]A{}, slice...)
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

// Sample will return a slice containing "n" random elements from the original slice
func Sample[A any](slice []A, n int) []A {
	if n > len(slice) {
		n = len(slice)
	}

	if n <= 0 {
		return []A{}
	}

	// For large samples (>50% of slice), use Fisher-Yates shuffle approach
	// This is O(n) and avoids the birthday paradox problem
	if n > len(slice)/2 {
		// Create a copy to avoid mutating original
		shuffled := make([]A, len(slice))
		copy(shuffled, slice)
		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})
		return shuffled[:n]
	}

	// For smaller samples, use swap-to-end algorithm (O(n) time, O(n) space for copy)
	// This is more efficient than retry-loop with map when slice is large
	if len(slice) <= 10000 {
		// Create mutable copy and swap selected elements to the front
		mut := make([]A, len(slice))
		copy(mut, slice)

		ret := make([]A, n)
		for i := 0; i < n; i++ {
			// Pick random index from remaining elements
			j := i + rand.Intn(len(mut)-i)
			ret[i] = mut[j]
			// Swap to keep selected elements at the front
			mut[i], mut[j] = mut[j], mut[i]
		}
		return ret
	}

	// For very large slices with small samples, use set-based approach
	// to avoid copying the entire slice
	ret := make([]A, 0, n)
	idxs := make(map[int]struct{}, n)
	for len(idxs) < n {
		idx := rand.Intn(len(slice))
		if _, found := idxs[idx]; !found {
			idxs[idx] = struct{}{}
			ret = append(ret, slice[idx])
		}
	}
	return ret
}

// Sort will return a new slice that is sorted in the natural order
func Sort[A compare.Ordered](slice []A) []A {
	return SortBy(slice, compare.Less[A])
}

// SortBy will return a new slice that is sorted using the supplied less function for natural ordering
func SortBy[A any](slice []A, less func(a, b A) bool) []A {
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
	return CompactBy(slice, compare.Equal[A])
}

// CompactBy will remove any duplicate elements following each other determined by the equal func.
// eg removing duplicate whitespaces from a string might look like
//
//	CompactBy([]rune("a    b"), func(a, b rune) {
//	 	return a == ' ' && a == b
//	})
//
// resulting in "a b"
func CompactBy[A any](slice []A, equal func(a, b A) bool) []A {
	if len(slice) == 0 {
		return slice
	}
	// Pre-allocate result with same length as input
	// In worst case (no duplicates), we'll use full capacity
	result := make([]A, 1, len(slice))
	result[0] = slice[0]
	last := slice[0]

	for i := 1; i < len(slice); i++ {
		current := slice[i]
		if !equal(last, current) {
			result = append(result, current)
			last = current
		}
	}
	return result
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

// Min returns the smallest element of the slice
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

// FlatMap will map entries in one slice to entries in another slice and then flatten the map
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

// Associate will iterate over a slice turning each object into a key/value pair in a map. Alias SliceToMap
func Associate[E any, K comparable, V any](slice []E, mapper func(e E) (key K, value V)) map[K]V {
	acc := make(map[K]V, len(slice))
	for _, e := range slice {
		k, v := mapper(e)
		acc[k] = v
	}
	return acc
}

// Set will create a Set in the form of map[E]bool,
// This can be used to lookup if a item was present in the slice or not
func Set[E comparable](slice []E) map[E]bool {
	return Associate(slice, func(a E) (key E, value bool) {
		return a, true
	})
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
	capacity := Min(len(aSlice), len(bSlice))
	if capacity == 0 {
		return []C{}
	}
	cSlice := make([]C, capacity)
	for i := 0; i < capacity; i++ {
		cSlice[i] = zipper(aSlice[i], bSlice[i])
	}
	return cSlice
}

// Unzip will unzip a slice slices, c, into two slices, a and b, using the supplied unziper function
func Unzip[A any, B any, C any](cSlice []C, unzipper func(c C) (a A, b B)) ([]A, []B) {
	if len(cSlice) == 0 {
		return []A{}, []B{}
	}
	aSlice := make([]A, len(cSlice))
	bSlice := make([]B, len(cSlice))
	for i, c := range cSlice {
		aSlice[i], bSlice[i] = unzipper(c)
	}
	return aSlice, bSlice
}

// Zip2 will zip three slices, a, b and c, into one slice, d, using the zip function to combined elements
func Zip2[A any, B any, C any, D any](aSlice []A, bSlice []B, cSlice []C, zipper func(a A, b B, c C) D) []D {
	capacity := Min(len(aSlice), len(bSlice), len(cSlice))
	if capacity == 0 {
		return []D{}
	}
	dSlice := make([]D, capacity)
	for i := 0; i < capacity; i++ {
		dSlice[i] = zipper(aSlice[i], bSlice[i], cSlice[i])
	}
	return dSlice
}

// Unzip2 will unzip a slice slices, d, into three slices, a, b and c, using the supplied unziper function
func Unzip2[A any, B any, C any, D any](dSlice []D, unzipper func(d D) (a A, b B, c C)) ([]A, []B, []C) {
	if len(dSlice) == 0 {
		return []A{}, []B{}, []C{}
	}
	aSlice := make([]A, len(dSlice))
	bSlice := make([]B, len(dSlice))
	cSlice := make([]C, len(dSlice))
	for i, d := range dSlice {
		aSlice[i], bSlice[i], cSlice[i] = unzipper(d)
	}
	return aSlice, bSlice, cSlice
}

// Zip3 will zip three slices, a, b and c, into one slice, d, using the zip function to combined elements
func Zip3[A any, B any, C any, D any, E any](aSlice []A, bSlice []B, cSlice []C, dSlice []D, zipper func(a A, b B, c C, d D) E) []E {
	capacity := Min(len(aSlice), len(bSlice), len(cSlice), len(dSlice))
	if capacity == 0 {
		return []E{}
	}
	eSlice := make([]E, capacity)
	for i := 0; i < capacity; i++ {
		eSlice[i] = zipper(aSlice[i], bSlice[i], cSlice[i], dSlice[i])
	}
	return eSlice
}

// Unzip3 will unzip a slice slices, d, into three slices, a, b and c, using the supplied unziper function
func Unzip3[A any, B any, C any, D any, E any](eSlice []E, unzipper func(e E) (a A, b B, c C, d D)) ([]A, []B, []C, []D) {
	if len(eSlice) == 0 {
		return []A{}, []B{}, []C{}, []D{}
	}
	aSlice := make([]A, len(eSlice))
	bSlice := make([]B, len(eSlice))
	cSlice := make([]C, len(eSlice))
	dSlice := make([]D, len(eSlice))
	for i, e := range eSlice {
		aSlice[i], bSlice[i], cSlice[i], dSlice[i] = unzipper(e)
	}
	return aSlice, bSlice, cSlice, dSlice
}

func XOR[A comparable](slices ...[]A) []A {
	return XORBy(compare.Identity[A], slices...)
}

func XORBy[A any, B comparable](by func(A) B, slices ...[]A) []A {
	seen := map[B]int{}
	var res []A
	for _, slice := range slices {
		for _, e := range slice {
			k := by(e)
			seen[k] = seen[k] + 1
		}
	}
	for _, slice := range slices {
		for _, e := range slice {
			k := by(e)
			if seen[k] > 1 {
				continue
			}
			res = append(res, e)
		}
	}

	return res
}

// ScanLeft returns all intermediate results of folding from left to right.
// Like Fold but returns all accumulator values including the initial value.
//
// Example:
//
//	ScanLeft([]int{1, 2, 3}, func(acc, val int) int { return acc + val }, 0)
//	// Returns []int{0, 1, 3, 6} (running sums)
func ScanLeft[I any, A any](slice []I, combine func(accumulator A, val I) A, init A) []A {
	result := make([]A, len(slice)+1)
	result[0] = init
	for i, val := range slice {
		init = combine(init, val)
		result[i+1] = init
	}
	return result
}

// ScanRight returns all intermediate results of folding from right to left.
// Like ScanLeft but operates from right to left.
//
// Example:
//
//	ScanRight([]int{1, 2, 3}, func(acc, val int) int { return acc + val }, 0)
//	// Returns []int{0, 3, 5, 6} (running sums from right)
func ScanRight[I any, A any](slice []I, combine func(accumulator A, val I) A, init A) []A {
	result := make([]A, len(slice)+1)
	result[len(slice)] = init
	for i := len(slice) - 1; i >= 0; i-- {
		init = combine(init, slice[i])
		result[i] = init
	}
	return result
}

// Scan is an alias for ScanLeft.
func Scan[I any, A any](slice []I, combine func(accumulator A, val I) A, init A) []A {
	return ScanLeft(slice, combine, init)
}

// SlidingWindow creates sliding windows of size n from the slice.
// Returns a slice of slices where each inner slice has n consecutive elements.
//
// Example:
//
//	SlidingWindow([]int{1, 2, 3, 4, 5}, 3)
//	// Returns [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
func SlidingWindow[A any](slice []A, n int) [][]A {
	if n <= 0 || len(slice) < n {
		return nil
	}
	result := make([][]A, 0, len(slice)-n+1)
	for i := 0; i <= len(slice)-n; i++ {
		window := make([]A, n)
		copy(window, slice[i:i+n])
		result = append(result, window)
	}
	return result
}

// Transpose transposes a matrix (slice of slices), swapping rows and columns.
// Assumes all rows have the same length. Returns nil for empty input.
//
// Example:
//
//	Transpose([][]int{{1, 2, 3}, {4, 5, 6}})
//	// Returns [][]int{{1, 4}, {2, 5}, {3, 6}}
func Transpose[A any](matrix [][]A) [][]A {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	rows := len(matrix)
	cols := len(matrix[0])

	// Verify all rows have same length
	for _, row := range matrix {
		if len(row) != cols {
			return nil // or panic, or handle differently
		}
	}

	result := make([][]A, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]A, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

// Intersperse inserts element between each element of the slice.
// Returns a new slice with the element inserted between consecutive elements.
//
// Example:
//
//	Intersperse([]int{1, 2, 3}, 0)
//	// Returns []int{1, 0, 2, 0, 3}
func Intersperse[A any](slice []A, element A) []A {
	if len(slice) <= 1 {
		return Clone(slice)
	}
	result := make([]A, 0, len(slice)*2-1)
	for i, v := range slice {
		result = append(result, v)
		if i < len(slice)-1 {
			result = append(result, element)
		}
	}
	return result
}

// SplitAt splits the slice at the given index.
// Returns two slices: elements before index and elements from index onward.
//
// Example:
//
//	SplitAt([]int{1, 2, 3, 4, 5}, 2)
//	// Returns ([]int{1, 2}, []int{3, 4, 5})
func SplitAt[A any](slice []A, index int) (before, after []A) {
	if index <= 0 {
		return []A{}, Clone(slice)
	}
	if index >= len(slice) {
		return Clone(slice), []A{}
	}
	return Clone(slice[:index]), Clone(slice[index:])
}

// Span splits the slice at the first element that does not satisfy the predicate.
// Returns two slices: elements satisfying the predicate, and remaining elements.
// More efficient than calling TakeWhile + DropWhile separately.
//
// Example:
//
//	Span([]int{1, 2, 3, 4, 5}, func(n int) bool { return n < 4 })
//	// Returns ([]int{1, 2, 3}, []int{4, 5})
func Span[A any](slice []A, predicate func(a A) bool) (init, rest []A) {
	for i, a := range slice {
		if !predicate(a) {
			return Clone(slice[:i]), Clone(slice[i:])
		}
	}
	return Clone(slice), []A{}
}

// MapIdx maps a function over the slice with index awareness.
// The mapper function receives both the index and the element.
//
// Example:
//
//	MapIdx([]string{"a", "b", "c"}, func(i int, s string) string {
//	    return fmt.Sprintf("%d:%s", i, s)
//	})
//	// Returns []string{"0:a", "1:b", "2:c"}
func MapIdx[A any, B any](slice []A, mapper func(index int, a A) B) []B {
	result := make([]B, len(slice))
	for i, a := range slice {
		result[i] = mapper(i, a)
	}
	return result
}

// FilterIdx filters elements with index awareness.
// The predicate receives both the index and the element.
//
// Example:
//
//	FilterIdx([]int{10, 20, 30, 40}, func(i int, n int) bool {
//	    return i%2 == 0 && n > 15
//	})
//	// Returns []int{30} (index 2 is even and value > 15)
func FilterIdx[A any](slice []A, include func(index int, a A) bool) []A {
	result := make([]A, 0, len(slice))
	for i, a := range slice {
		if include(i, a) {
			result = append(result, a)
		}
	}
	return result
}

// RejectIdx is the complement of FilterIdx, filtering out elements that satisfy the predicate.
func RejectIdx[A any](slice []A, exclude func(index int, a A) bool) []A {
	return FilterIdx(slice, func(i int, a A) bool {
		return !exclude(i, a)
	})
}

// IsAllUnique returns true if all elements in the slice are unique (no duplicates).
// Uses a map for O(n) time complexity.
//
// Example:
//
//	IsAllUnique([]int{1, 2, 3})    // Returns true
//	IsAllUnique([]int{1, 2, 2, 3}) // Returns false
func IsAllUnique[A comparable](slice []A) bool {
	seen := make(map[A]struct{}, len(slice))
	for _, a := range slice {
		if _, exists := seen[a]; exists {
			return false
		}
		seen[a] = struct{}{}
	}
	return true
}

// IsSorted returns true if the slice is sorted in ascending order.
// Uses the natural ordering of the elements.
//
// Example:
//
//	IsSorted([]int{1, 2, 3})    // Returns true
//	IsSorted([]int{3, 2, 1})    // Returns false
//	IsSorted([]int{1, 2, 2, 3}) // Returns true (allows duplicates)
func IsSorted[A compare.Ordered](slice []A) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// IsSortedBy returns true if the slice is sorted according to the provided comparison function.
// The less function should return true if a should come before b.
//
// Example:
//
//	IsSortedBy([]string{"a", "bb", "ccc"}, func(a, b string) bool {
//	    return len(a) < len(b)
//	}) // Returns true (sorted by length)
func IsSortedBy[A any](slice []A, less func(a, b A) bool) bool {
	for i := 1; i < len(slice); i++ {
		if less(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}

// GroupByEntry represents a single entry in a grouped result with preserved order.
type GroupByEntry[K comparable, V any] struct {
	Key    K
	Values []V
}

// GroupByOrdered groups slice elements by a key function while preserving insertion order.
// Unlike GroupBy which returns a map (with random iteration), this returns a slice
// where the order of groups matches the order keys first appear in the input.
//
// Example:
//
//	words := []string{"apple", "banana", "avocado", "blueberry", "cherry"}
//	GroupByOrdered(words, func(s string) string { return string(s[0]) })
//	// Returns []GroupByEntry{{"a", ["apple", "avocado"]}, {"b", ["banana", "blueberry"]}, {"c", ["cherry"]}}
func GroupByOrdered[A any, K comparable](slice []A, by func(A) K) []GroupByEntry[K, A] {
	if len(slice) == 0 {
		return nil
	}

	groups := make(map[K][]A, len(slice))
	order := make([]K, 0, len(slice))
	seen := make(map[K]struct{}, len(slice))

	for _, v := range slice {
		k := by(v)
		groups[k] = append(groups[k], v)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			order = append(order, k)
		}
	}

	result := make([]GroupByEntry[K, A], len(order))
	for i, k := range order {
		result[i] = GroupByEntry[K, A]{Key: k, Values: groups[k]}
	}
	return result
}

// ChunkBy groups consecutive elements that satisfy the predicate.
// The predicate receives consecutive elements (a, b) and should return true
// if they should be grouped together.
//
// Example:
//
//	ChunkBy([]int{1, 1, 1, 2, 2, 3, 3, 3}, func(a, b int) bool { return a == b })
//	// Returns [][]int{{1, 1, 1}, {2, 2}, {3, 3, 3}}
//
//	ChunkBy([]int{1, 2, 3, 2, 2, 1}, func(a, b int) bool { return a <= b })
//	// Returns [][]int{{1, 2, 3}, {2, 2}, {1}}
func ChunkBy[A any](slice []A, predicate func(a, b A) bool) [][]A {
	if len(slice) == 0 {
		return nil
	}

	var result [][]A
	current := []A{slice[0]}

	for i := 1; i < len(slice); i++ {
		if predicate(slice[i-1], slice[i]) {
			current = append(current, slice[i])
		} else {
			result = append(result, current)
			current = []A{slice[i]}
		}
	}
	result = append(result, current)
	return result
}

// Deduplicate removes consecutive duplicate elements from a slice.
// Unlike Uniq which removes all duplicates, this only removes consecutive duplicates.
//
// Example:
//
//	Deduplicate([]int{1, 1, 2, 2, 2, 3, 3}) // Returns []int{1, 2, 3}
//	Deduplicate([]int{1, 2, 1, 2, 1})      // Returns []int{1, 2, 1, 2, 1}
func Deduplicate[A comparable](slice []A) []A {
	if len(slice) == 0 {
		return []A{}
	}

	result := make([]A, 1, len(slice))
	result[0] = slice[0]

	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}
	return result
}

// Cycle returns a slice that cycles through the input slice infinitely.
// The returned slice can be indexed modulo the length of the input.
// To get the nth cycled element: result[n % len(slice)]
//
// Example:
//
//	cycle := Cycle([]int{1, 2, 3})
//	cycle[0] // 1, cycle[3] // 1, cycle[6] // 1
//	cycle[1] // 2, cycle[4] // 2, cycle[7] // 2
func Cycle[A any](slice []A) []A {
	if len(slice) == 0 {
		return nil
	}
	// Return a copy to prevent mutation of original
	return Clone(slice)
}

// CycleValue returns the element at position index in a cycled slice.
// Useful for getting values from an infinitely cycling sequence.
//
// Example:
//
//	CycleValue([]int{1, 2, 3}, 0)  // Returns 1
//	CycleValue([]int{1, 2, 3}, 3)  // Returns 1 (cycles back)
//	CycleValue([]int{1, 2, 3}, 7)  // Returns 2 (7 % 3 = 1)
func CycleValue[A any](slice []A, index int) A {
	if len(slice) == 0 {
		var zero A
		return zero
	}
	return slice[index%len(slice)]
}

// Fill creates a slice of length n filled with the given value.
//
// Example:
//
//	Fill(5, 42)    // Returns []int{42, 42, 42, 42, 42}
//	Fill(3, "x")   // Returns []string{"x", "x", "x"}
func Fill[A any](n int, value A) []A {
	if n <= 0 {
		return []A{}
	}
	result := make([]A, n)
	for i := range result {
		result[i] = value
	}
	return result
}

// Range creates a slice of integers from start to end (inclusive).
//
// Example:
//
//	Range(1, 5)   // Returns []int{1, 2, 3, 4, 5}
//	Range(5, 1)   // Returns []int{} (empty when start > end)
//	Range(3, 3)   // Returns []int{3}
func Range(start, end int) []int {
	if start > end {
		return []int{}
	}
	n := end - start + 1
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = start + i
	}
	return result
}

// RangeFrom creates a slice of n integers starting from start.
//
// Example:
//
//	RangeFrom(0, 5)   // Returns []int{0, 1, 2, 3, 4}
//	RangeFrom(10, 3)  // Returns []int{10, 11, 12}
//	RangeFrom(5, 0)   // Returns []int{} (empty when n <= 0)
func RangeFrom(start, n int) []int {
	if n <= 0 {
		return []int{}
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = start + i
	}
	return result
}

// RangeStep creates a slice from start to end with a step value.
//
// Example:
//
//	RangeStep(0, 10, 2)   // Returns []int{0, 2, 4, 6, 8, 10}
//	RangeStep(10, 0, -2)  // Returns []int{10, 8, 6, 4, 2, 0}
//	RangeStep(0, 10, 3)   // Returns []int{0, 3, 6, 9}
func RangeStep(start, end, step int) []int {
	if step == 0 {
		return []int{}
	}
	if (step > 0 && start > end) || (step < 0 && start < end) {
		return []int{}
	}

	// Calculate number of steps
	// For positive step: include values while value <= end
	// For negative step: include values while value >= end
	n := 0
	if step > 0 {
		if start <= end {
			n = (end-start)/step + 1
			// Check if last value exceeds end
			if start+(n-1)*step > end {
				n--
			}
		}
	} else {
		if start >= end {
			n = (start-end)/(-step) + 1
			// Check if last value is below end
			if start+(n-1)*step < end {
				n--
			}
		}
	}

	if n <= 0 {
		return []int{}
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = start + i*step
	}
	return result
}

// Repeat repeats the slice n times and returns the concatenated result.
//
// Example:
//
//	Repeat([]int{1, 2}, 3)    // Returns []int{1, 2, 1, 2, 1, 2}
//	Repeat([]string{"a"}, 5)   // Returns []string{"a", "a", "a", "a", "a"}
//	Repeat([]int{1, 2}, 0)    // Returns []int{} (empty when n <= 0)
func Repeat[A any](slice []A, n int) []A {
	if n <= 0 || len(slice) == 0 {
		return []A{}
	}
	result := make([]A, 0, len(slice)*n)
	for i := 0; i < n; i++ {
		result = append(result, slice...)
	}
	return result
}
