package slicez

import (
	"errors"
	"math/rand"

	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/slicez/sort"
)

// Equal checks if two slices are equal.
// Two slices are considered equal if they have the same length and
// each element at the same index is equal (== comparison).
//
// Example:
//
//	slicez.Equal([]int{1, 2, 3}, []int{1, 2, 3})        // returns true
//	slicez.Equal([]int{1, 2, 3}, []int{1, 2, 3, 4})    // returns false (different lengths)
//	slicez.Equal([]string{"a", "b"}, []string{"a", "b"}) // returns true
//	slicez.Equal([]int{1, 2}, []int{2, 1})             // returns false (different order)
func Equal[A comparable](s1, s2 []A) bool {
	return EqualBy(s1, s2, compare.Equal[A])
}

// EqualBy checks if two slices are equal using a custom equality function.
// It returns true if both slices have the same length and the equality function
// returns true for every pair of elements at the same index.
//
// This is useful when comparing slices of different types, or when you need
// custom comparison logic (e.g., case-insensitive string comparison).
//
// Example:
//
//	// Compare strings ignoring case
//	s1 := []string{"hello", "world"}
//	s2 := []string{"HELLO", "WORLD"}
//	slicez.EqualBy(s1, s2, func(a, b string) bool {
//	    return strings.EqualFold(a, b)
//	}) // returns true
//
//	// Compare different types
//	ints := []int{1, 2, 3}
//	floats := []float64{1.0, 2.0, 3.0}
//	slicez.EqualBy(ints, floats, func(i int, f float64) bool {
//	    return float64(i) == f
//	}) // returns true
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

// Index returns the index of the first occurrence of needle in s.
// It returns -1 if needle is not found.
// Uses == comparison for equality.
//
// Example:
//
//	slicez.Index([]int{1, 2, 3, 2, 1}, 2)        // returns 1 (first occurrence)
//	slicez.Index([]string{"a", "b", "c"}, "b")    // returns 1
//	slicez.Index([]int{1, 2, 3}, 4)             // returns -1 (not found)
func Index[E comparable](s []E, needle E) int {
	return IndexBy(s, func(e E) bool {
		return needle == e
	})
}

// IndexBy returns the index of the first element where f returns true.
// It returns -1 if no element satisfies the condition.
//
// Example:
//
//	// Find first even number
//	slicez.IndexBy([]int{1, 3, 5, 6, 7}, func(n int) bool {
//	    return n%2 == 0
//	}) // returns 3
//
//	// Find first string longer than 3 characters
//	slicez.IndexBy([]string{"a", "bb", "ccc", "dddd"}, func(s string) bool {
//	    return len(s) > 3
//	}) // returns 3
func IndexBy[E any](s []E, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last occurrence of needle in s.
// It returns -1 if needle is not found.
// Uses == comparison for equality.
//
// Example:
//
//	slicez.LastIndex([]int{1, 2, 3, 2, 1}, 2)     // returns 3 (last occurrence)
//	slicez.LastIndex([]string{"a", "b", "c", "b"}, "b") // returns 3
//	slicez.LastIndex([]int{1, 2, 3}, 4)          // returns -1 (not found)
func LastIndex[E comparable](s []E, needle E) int {
	return LastIndexBy(s, func(e E) bool {
		return e == needle
	})
}

// LastIndexBy returns the index of the last element where f returns true.
// It returns -1 if no element satisfies the condition.
//
// Example:
//
//	// Find last even number
//	slicez.LastIndexBy([]int{2, 4, 6, 7, 8}, func(n int) bool {
//	    return n%2 == 0
//	}) // returns 4
//
//	// Find last negative number
//	slicez.LastIndexBy([]int{1, -2, 3, -4, 5}, func(n int) bool {
//	    return n < 0
//	}) // returns 3
func LastIndexBy[E any](s []E, f func(E) bool) int {
	n := len(s)

	for i := 0; i < n; i++ {
		if f(s[n-i-1]) {
			return n - i - 1
		}
	}
	return -1
}

// Cut splits the slice at the first occurrence of needle.
// It returns left (elements before needle), right (elements after needle), and a boolean
// indicating if the needle was found. The needle itself is not included in either part.
//
// Example:
//
//	left, right, found := slicez.Cut([]int{1, 2, 3, 4, 5}, 3)
//	// left = []int{1, 2}, right = []int{4, 5}, found = true
//
//	left, right, found := slicez.Cut([]string{"a", "b", "c"}, "d")
//	// left = []string{"a", "b", "c"}, right = nil, found = false
func Cut[E comparable](s []E, needle E) (left, right []E, found bool) {
	return CutBy(s, func(e E) bool {
		return e == needle
	})
}

// CutBy splits the slice at the first element where on returns true.
// It returns left (elements before the match), right (elements after the match), and a boolean
// indicating if a match was found. The matching element is not included in either part.
//
// Example:
//
//	// Cut at first number > 5
//	left, right, found := slicez.CutBy([]int{1, 2, 6, 7, 8}, func(n int) bool {
//	    return n > 5
//	})
//	// left = []int{1, 2}, right = []int{7, 8}, found = true
//
//	// Cut at first string starting with "c"
//	left, right, found := slicez.CutBy([]string{"apple", "banana", "cherry"}, func(s string) bool {
//	    return strings.HasPrefix(s, "c")
//	})
//	// left = []string{"apple", "banana"}, right = []string{}, found = true
func CutBy[E any](s []E, on func(E) bool) (left, right []E, found bool) {
	i := IndexBy(s, on)
	if i == -1 {
		return s, nil, false
	}
	return s[:i], s[i+1:], true
}

// Replace replaces up to n occurrences of needle with replacement in haystack.
// If n < 0, all occurrences are replaced.
// Returns a new slice; the original is not modified.
//
// Example:
//
//	// Replace first 2 occurrences
//	slicez.Replace([]int{1, 2, 1, 2, 1}, 1, 99, 2)
//	// returns []int{99, 2, 99, 2, 1}
//
//	// Replace all occurrences (n < 0)
//	slicez.Replace([]int{1, 2, 1, 2, 1}, 1, 99, -1)
//	// returns []int{99, 2, 99, 2, 99}
//
//	// Replace first occurrence only
//	slicez.Replace([]string{"a", "b", "a"}, "a", "z", 1)
//	// returns []string{"z", "b", "a"}
func Replace[E comparable](haystack []E, needle E, replacement E, n int) []E {
	return Map(haystack, func(e E) E {
		if n != 0 && e == needle {
			n--
			return replacement
		}
		return e
	})
}

// ReplaceFirst replaces the first occurrence of needle with replacement in haystack.
// Returns a new slice; the original is not modified.
//
// Example:
//
//	slicez.ReplaceFirst([]int{1, 2, 1, 3}, 1, 99)
//	// returns []int{99, 2, 1, 3}
func ReplaceFirst[E comparable](haystack []E, needle E, replacement E) []E {
	return Replace(haystack, needle, replacement, 1)
}

// ReplaceAll replaces all occurrences of needle with replacement in haystack.
// Returns a new slice; the original is not modified.
//
// Example:
//
//	slicez.ReplaceAll([]int{1, 2, 1, 3, 1}, 1, 99)
//	// returns []int{99, 2, 99, 3, 99}
func ReplaceAll[E comparable](haystack []E, needle E, replacement E) []E {
	return Replace(haystack, needle, replacement, -1)
}

// Find returns the first element that satisfies the predicate function.
// Returns the element and true if found, or the zero value and false if not found.
//
// Example:
//
//	// Find first number > 10
//	val, found := slicez.Find([]int{1, 5, 12, 7, 15}, func(n int) bool {
//	    return n > 10
//	})
//	// val = 12, found = true
//
//	// Find first string containing "e"
//	val, found := slicez.Find([]string{"cat", "dog", "elephant"}, func(s string) bool {
//	    return strings.Contains(s, "e")
//	})
//	// val = "elephant", found = true
func Find[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := IndexBy(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

// FindLast returns the last element that satisfies the predicate function.
// Returns the element and true if found, or the zero value and false if not found.
//
// Example:
//
//	// Find last even number
//	val, found := slicez.FindLast([]int{1, 2, 3, 4, 5}, func(n int) bool {
//	    return n%2 == 0
//	})
//	// val = 4, found = true
//
//	// Find last string starting with "a"
//	val, found := slicez.FindLast([]string{"apple", "banana", "apricot", "cherry"}, func(s string) bool {
//	    return strings.HasPrefix(s, "a")
//	})
//	// val = "apricot", found = true
func FindLast[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := LastIndexBy(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

// Join concatenates multiple slices with glue inserted between each pair.
// Similar to strings.Join but works with any slice type.
// Returns an empty slice if slices is empty or nil.
// If only one slice is provided, returns a copy of that slice.
//
// Example:
//
//	// Join with separator
//	slicez.Join([][]int{{1, 2}, {3, 4}, {5, 6}}, []int{0, 0})
//	// returns []int{1, 2, 0, 0, 3, 4, 0, 0, 5, 6}
//
//	// Join strings (like strings.Join but returns []string)
//	slicez.Join([][]string{{"hello"}, {"world"}}, []string{" "})
//	// returns []string{"hello", " ", "world"}
//
//	// Single slice
//	slicez.Join([][]int{{1, 2, 3}}, []int{0})
//	// returns []int{1, 2, 3}
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

// Contains checks if needle exists in slice.
// Uses == comparison for equality.
// Returns false if slice is empty.
//
// Example:
//
//	slicez.Contains([]int{1, 2, 3}, 2)          // returns true
//	slicez.Contains([]string{"a", "b"}, "c")     // returns false
//	slicez.Contains([]int{}, 1)                  // returns false
func Contains[E comparable](s []E, needle E) bool {
	return Index(s, needle) >= 0
}

// ContainsBy checks if any element satisfies the predicate function.
// Returns false if slice is empty or no element satisfies the condition.
//
// Example:
//
//	// Check if any number is even
//	slicez.ContainsBy([]int{1, 3, 5, 6, 7}, func(n int) bool {
//	    return n%2 == 0
//	}) // returns true
//
//	// Check if any string is longer than 5 chars
//	slicez.ContainsBy([]string{"a", "bb", "ccc"}, func(s string) bool {
//	    return len(s) > 5
//	}) // returns false
func ContainsBy[E any](s []E, f func(e E) bool) bool {
	return IndexBy(s, f) >= 0
}

// Clone creates a copy of the slice.
// Returns nil if the input slice is nil (preserves nil vs empty distinction).
// The returned slice has the same elements but different backing array.
//
// Example:
//
//	original := []int{1, 2, 3}
//	copy := slicez.Clone(original)
//	copy[0] = 99
//	// original is still []int{1, 2, 3}
//	// copy is []int{99, 2, 3}
//
//	// Nil preservation
//	var nilSlice []int
//	result := slicez.Clone(nilSlice)  // result is nil, not []int{}
func Clone[E any](s []E) []E {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append([]E{}, s...)
}

// Compare performs lexicographic comparison of two slices.
// Returns -1 if s1 < s2, 0 if s1 == s2, and +1 if s1 > s2.
// Comparison is done element by element using natural ordering.
// If all elements are equal, the shorter slice is considered less.
//
// Example:
//
//	slicez.Compare([]int{1, 2, 3}, []int{1, 2, 4})  // returns -1
//	slicez.Compare([]int{1, 2, 3}, []int{1, 2, 3})  // returns 0
//	slicez.Compare([]int{1, 2, 3}, []int{1, 2})      // returns +1 (longer)
//	slicez.Compare([]string{"a", "b"}, []string{"a", "c"}) // returns -1
func Compare[E compare.Ordered](s1, s2 []E) int {
	return CompareBy(s1, s2, compare.Compare[E])
}

// CompareBy performs lexicographic comparison of two slices using a custom comparison function.
// Returns -1 if s1 < s2, 0 if s1 == s2, and +1 if s1 > s2.
// The comparison function should return negative if a < b, zero if a == b, positive if a > b.
//
// Example:
//
//	// Compare by length
//	slicez.CompareBy([]string{"a", "bb"}, []string{"ccc", "d"}, func(a, b string) int {
//	    return len(a) - len(b)
//	}) // returns -1 (first string shorter overall)
//
//	// Compare different types
//	slicez.CompareBy([]int{1, 2}, []float64{1.0, 2.0, 3.0}, func(i int, f float64) int {
//	    if float64(i) < f { return -1 }
//	    if float64(i) > f { return +1 }
//	    return 0
//	}) // returns -1 (second slice longer)
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

// Concat concatenates multiple slices in order into a single new slice.
// Returns an empty slice if no slices are provided.
// Preserves order: elements from first slice appear first, then second, etc.
//
// Example:
//
//	slicez.Concat([]int{1, 2}, []int{3, 4}, []int{5, 6})
//	// Returns []int{1, 2, 3, 4, 5, 6}
//
//	slicez.Concat([]string{"a"}, []string{"b", "c"})
//	// Returns []string{"a", "b", "c"}
//
//	slicez.Concat[int]() // Returns []int{} (empty)
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

// Reverse returns a new slice with elements in reverse order.
// The first element becomes the last, and the last becomes the first.
// Returns nil for nil input (preserves nil vs empty distinction).
//
// Example:
//
//	slicez.Reverse([]int{1, 2, 3, 4, 5})
//	// Returns []int{5, 4, 3, 2, 1}
//
//	slicez.Reverse([]string{"a", "b", "c"})
//	// Returns []string{"c", "b", "a"}
//
//	slicez.Reverse([]int{}) // Returns []int{} (empty)
func Reverse[A any](slice []A) []A {
	l := len(slice)
	res := make([]A, l)
	for i, val := range slice {
		res[l-i-1] = val
	}
	return res
}

// RepeatBy creates a slice of length n where each element is computed by the provided function.
// The function receives the index (0 to n-1) and returns the value for that position.
// Useful for creating sequences or computed values.
//
// Example:
//
//	// Create squares: [0, 1, 4, 9, 16]
//	slicez.RepeatBy(5, func(i int) int { return i * i })
//
//	// Create pattern: [0, 1, 0, 1, 0]
//	slicez.RepeatBy(5, func(i int) int { return i % 2 })
//
//	// Create initialized structs
//	type Point struct{ X, Y int }
//	slicez.RepeatBy(3, func(i int) Point { return Point{X: i, Y: i * 10} })
//	// Returns [{0 0}, {1 10}, {2 20}]
func RepeatBy[A any](i int, by func(i int) A) []A {
	res := make([]A, 0, i)

	for n := 0; n < i; n++ {
		res = append(res, by(n))
	}
	return res
}

// Head returns the first element of the slice.
// Returns an error if the slice is empty or nil.
//
// Example:
//
//	val, err := slicez.Head([]int{1, 2, 3})
//	// val = 1, err = nil
//
//	val, err := slicez.Head([]string{"a", "b"})
//	// val = "a", err = nil
//
//	val, err := slicez.Head([]int{})
//	// val = 0 (zero value), err = "slice does not have any elements"
func Head[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[0], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

// Tail returns a new slice containing all elements except the first.
// Returns nil if the slice has 0 or 1 elements.
// Useful for recursive processing (head:tail pattern).
//
// Example:
//
//	slicez.Tail([]int{1, 2, 3, 4})
//	// Returns []int{2, 3, 4}
//
//	slicez.Tail([]string{"a", "b"})
//	// Returns []string{"b"}
//
//	slicez.Tail([]int{1}) // Returns nil
func Tail[A any](slice []A) []A {
	return Drop(slice, 1)
}

// Initial returns a new slice containing all elements except the last.
// Returns nil if the slice has 0 or 1 elements.
// Complement of Tail - gets all but the last element.
//
// Example:
//
//	slicez.Initial([]int{1, 2, 3, 4})
//	// Returns []int{1, 2, 3}
//
//	slicez.Initial([]string{"a", "b", "c"})
//	// Returns []string{"a", "b"}
//
//	slicez.Initial([]int{1}) // Returns nil
func Initial[A any](slice []A) []A {
	return DropRight(slice, 1)
}

// Last returns the last element of the slice.
// Returns an error if the slice is empty or nil.
//
// Example:
//
//	val, err := slicez.Last([]int{1, 2, 3})
//	// val = 3, err = nil
//
//	val, err := slicez.Last([]string{"a", "b", "c"})
//	// val = "c", err = nil
//
//	val, err := slicez.Last([]int{})
//	// val = 0 (zero value), err = "slice does not have any elements"
func Last[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[len(slice)-1], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

// Nth returns the element at the given index with modulo/wraparound support.
// Supports negative indices (counting from end: -1 = last, -2 = second to last).
// Supports out-of-bounds positive indices (wraps around using modulo).
// Returns zero value if slice is empty.
//
// Example:
//
//	s := []string{"a", "b", "c", "d"}
//
//	slicez.Nth(s, 0)   // Returns "a" (first element)
//	slicez.Nth(s, 2)   // Returns "c" (third element)
//	slicez.Nth(s, -1)  // Returns "d" (last element)
//	slicez.Nth(s, -2)  // Returns "c" (second to last)
//	slicez.Nth(s, 4)   // Returns "a" (wraps around: 4 % 4 = 0)
//	slicez.Nth(s, 10)  // Returns "b" (wraps around: 10 % 4 = 2)
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

// ForEach applies a function to each element of the slice from left to right.
// The function is called purely for side effects; it returns nothing.
// Similar to a for-range loop but as a higher-order function.
//
// Example:
//
//	// Print each element
//	slicez.ForEach([]int{1, 2, 3}, func(n int) {
//	    fmt.Println(n)
//	})
//	// Output: 1 2 3 (on separate lines)
//
//	// Accumulate sum
//	sum := 0
//	slicez.ForEach([]int{1, 2, 3, 4}, func(n int) {
//	    sum += n
//	})
//	// sum = 10
func ForEach[A any](slice []A, apply func(a A)) {
	for _, a := range slice {
		apply(a)
	}
}

// ForEachRight applies a function to each element from right to left.
// The function is called purely for side effects; it returns nothing.
// Processes elements in reverse order compared to ForEach.
//
// Example:
//
//	// Print in reverse order
//	slicez.ForEachRight([]int{1, 2, 3}, func(n int) {
//	    fmt.Println(n)
//	})
//	// Output: 3 2 1 (on separate lines)
//
//	// Build reversed string
//	var result string
//	slicez.ForEachRight([]string{"a", "b", "c"}, func(s string) {
//	    result += s
//	})
//	// result = "cba"
func ForEachRight[A any](slice []A, apply func(a A)) {
	length := len(slice)
	for i := 0; i < length; i++ {
		apply(slice[length-1-i])
	}
}

// TakeWhile returns elements from the start while the predicate returns true.
// Stops at the first element where the predicate returns false.
// Returns a new slice; the original is not modified.
//
// Example:
//
//	// Take while less than 4
//	slicez.TakeWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
//	    return n < 4
//	})
//	// Returns []int{1, 2, 3}
//
//	// Take while strings start with "a"
//	slicez.TakeWhile([]string{"apple", "avocado", "banana", "apricot"}, func(s string) bool {
//	    return strings.HasPrefix(s, "a")
//	})
//	// Returns []string{"apple", "avocado"}
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

// TakeRightWhile returns elements from the end while the predicate returns true.
// Stops at the first element (from the right) where the predicate returns false.
// Returns a new slice; the original is not modified.
//
// Example:
//
//	// Take from right while greater than 5
//	slicez.TakeRightWhile([]int{1, 2, 6, 7, 8, 3, 4}, func(n int) bool {
//	    return n > 5
//	})
//	// Returns []int{6, 7, 8} (stops at 3)
//
//	// Take strings ending with vowel from right
//	slicez.TakeRightWhile([]string{"cat", "dog", "bee", "flea"}, func(s string) bool {
//	    return strings.ContainsRune("aeiou", rune(s[len(s)-1]))
//	})
//	// Returns []string{"bee", "flea"}
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

// Take returns the first i elements of the slice.
// If i > len(slice), returns a copy of the entire slice.
// If i <= 0, returns an empty slice.
// Returns nil if the slice is nil and i is within bounds.
//
// Example:
//
//	slicez.Take([]int{1, 2, 3, 4, 5}, 3)
//	// Returns []int{1, 2, 3}
//
//	slicez.Take([]string{"a", "b", "c", "d"}, 2)
//	// Returns []string{"a", "b"}
//
//	slicez.Take([]int{1, 2, 3}, 10)
//	// Returns []int{1, 2, 3} (all elements)
//
//	slicez.Take([]int{1, 2, 3}, 0) // Returns []int{}
func Take[A any](slice []A, i int) []A {
	var j int
	return TakeWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}

// TakeRight returns the last i elements of the slice.
// If i > len(slice), returns a copy of the entire slice.
// If i <= 0, returns an empty slice.
//
// Example:
//
//	slicez.TakeRight([]int{1, 2, 3, 4, 5}, 3)
//	// Returns []int{3, 4, 5}
//
//	slicez.TakeRight([]string{"a", "b", "c", "d"}, 2)
//	// Returns []string{"c", "d"}
//
//	slicez.TakeRight([]int{1, 2, 3}, 10)
//	// Returns []int{1, 2, 3} (all elements)
//
//	slicez.TakeRight([]int{1, 2, 3}, 0) // Returns []int{}
func TakeRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return TakeRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

// DropWhile drops elements from the start while the predicate returns true.
// Returns a new slice starting from the first element where the predicate returns false.
// Returns nil if all elements are dropped or if the slice is empty.
//
// Example:
//
//	// Drop while less than 3
//	slicez.DropWhile([]int{1, 2, 3, 4, 5}, func(n int) bool {
//	    return n < 3
//	})
//	// Returns []int{3, 4, 5}
//
//	// Drop strings starting with "a"
//	slicez.DropWhile([]string{"apple", "apricot", "banana", "cherry"}, func(s string) bool {
//	    return strings.HasPrefix(s, "a")
//	})
//	// Returns []string{"banana", "cherry"}
//
//	// When no elements match
//	slicez.DropWhile([]int{1, 2, 3}, func(n int) bool { return n > 10 })
//	// Returns []int{1, 2, 3} (unchanged copy)
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

// DropRightWhile drops elements from the end while the predicate returns true.
// Returns a new slice ending at the last element where the predicate returns false.
// Returns nil if all elements are dropped or if the slice is empty.
//
// Example:
//
//	// Drop from right while greater than 5
//	slicez.DropRightWhile([]int{1, 2, 6, 7, 8}, func(n int) bool {
//	    return n > 5
//	})
//	// Returns []int{1, 2}
//
//	// Drop trailing zeros
//	slicez.DropRightWhile([]int{1, 2, 3, 0, 0, 0}, func(n int) bool {
//	    return n == 0
//	})
//	// Returns []int{1, 2, 3}
//
//	// When no elements match from right
//	slicez.DropRightWhile([]int{1, 2, 3}, func(n int) bool { return n > 10 })
//	// Returns []int{1, 2, 3} (unchanged copy)
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

// Drop returns a slice with the first i elements removed.
// If i <= 0, returns a copy of the entire slice.
// If i >= len(slice), returns nil.
//
// Example:
//
//	slicez.Drop([]int{1, 2, 3, 4, 5}, 3)
//	// Returns []int{4, 5}
//
//	slicez.Drop([]string{"a", "b", "c", "d"}, 2)
//	// Returns []string{"c", "d"}
//
//	slicez.Drop([]int{1, 2, 3}, 10) // Returns nil
//
//	slicez.Drop([]int{1, 2, 3}, 0)  // Returns []int{1, 2, 3} (copy)
func Drop[A any](slice []A, i int) []A {
	var j int
	return DropWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}

// DropRight returns a slice with the last i elements removed.
// If i <= 0, returns a copy of the entire slice.
// If i >= len(slice), returns nil.
//
// Example:
//
//	slicez.DropRight([]int{1, 2, 3, 4, 5}, 3)
//	// Returns []int{1, 2}
//
//	slicez.DropRight([]string{"a", "b", "c", "d"}, 2)
//	// Returns []string{"a", "b"}
//
//	slicez.DropRight([]int{1, 2, 3}, 10) // Returns nil
//
//	slicez.DropRight([]int{1, 2, 3}, 0)  // Returns []int{1, 2, 3} (copy)
func DropRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return DropRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

// Filter returns a new slice containing only elements where the predicate returns true.
// Returns a new slice; the original is not modified.
// Preserves order of elements.
//
// Example:
//
//	// Keep only even numbers
//	slicez.Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
//	    return n%2 == 0
//	})
//	// Returns []int{2, 4, 6}
//
//	// Keep strings longer than 3 chars
//	slicez.Filter([]string{"a", "bb", "ccc", "dddd"}, func(s string) bool {
//	    return len(s) > 3
//	})
//	// Returns []string{"dddd"}
//
//	// Filter with index (using FilterIdx)
//	slicez.FilterIdx([]int{10, 20, 30, 40}, func(i, n int) bool {
//	    return i%2 == 0 && n > 15
//	})
//	// Returns []int{30} (index 2, value > 15)
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

// Reject returns a new slice containing only elements where the predicate returns false.
// Complement of Filter - keeps elements that DON'T match the condition.
// Returns a new slice; the original is not modified.
//
// Example:
//
//	// Remove even numbers (keep odd)
//	slicez.Reject([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
//	    return n%2 == 0
//	})
//	// Returns []int{1, 3, 5}
//
//	// Remove short strings
//	slicez.Reject([]string{"a", "bb", "ccc", "dddd"}, func(s string) bool {
//	    return len(s) <= 2
//	})
//	// Returns []string{"ccc", "dddd"}
func Reject[A any](slice []A, exclude func(a A) bool) []A {
	return Filter(slice, func(a A) bool {
		return !exclude(a)
	})
}

// Without returns a new slice with all specified values removed.
// Excludes all occurrences of each value in exclude.
// Uses equality comparison (==).
// Returns a new slice; the original is not modified.
//
// Example:
//
//	slicez.Without([]int{1, 2, 3, 2, 4, 2, 5}, 2)
//	// Returns []int{1, 3, 4, 5} (all 2s removed)
//
//	slicez.Without([]string{"a", "b", "a", "c", "a"}, "a")
//	// Returns []string{"b", "c"}
//
//	// Remove multiple values
//	slicez.Without([]int{1, 2, 3, 4, 5}, 2, 4)
//	// Returns []int{1, 3, 5}
//
//	slicez.Without([]int{1, 2, 3}) // Returns []int{1, 2, 3} (unchanged copy)
func Without[A comparable](slice []A, exclude ...A) []A {
	set := Set(exclude)
	return Reject(slice, func(a A) bool {
		return set[a]
	})
}

// Every checks if all elements equal the given value.
// Returns true if the slice is empty (vacuous truth).
// Uses equality comparison (==).
//
// Example:
//
//	slicez.Every([]int{5, 5, 5}, 5)   // Returns true
//	slicez.Every([]int{5, 5, 6}, 5)   // Returns false
//	slicez.Every([]string{"a", "a"}, "a") // Returns true
//	slicez.Every([]int{}, 5)         // Returns true (empty slice)
func Every[A comparable](slice []A, needle A) bool {
	return EveryBy(slice, compare.EqualOf[A](needle))

}

// EveryBy checks if the predicate returns true for all elements.
// Returns true if the slice is empty (vacuous truth).
// Useful for validation and type checking.
//
// Example:
//
//	// Check if all numbers are positive
//	slicez.EveryBy([]int{1, 2, 3, 4}, func(n int) bool {
//	    return n > 0
//	}) // Returns true
//
//	// Check if all strings are non-empty
//	slicez.EveryBy([]string{"a", "b", ""}, func(s string) bool {
//	    return len(s) > 0
//	}) // Returns false (empty string at end)
//
//	slicez.EveryBy([]int{}, func(n int) bool { return n > 0 })
//	// Returns true (empty slice)
func EveryBy[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if !predicate(val) {
			return false
		}
	}
	return true
}

// Some checks if any element equals the given value.
// Alias for Contains. Returns false if the slice is empty.
// Uses equality comparison (==).
//
// Example:
//
//	slicez.Some([]int{1, 2, 3}, 2)     // Returns true
//	slicez.Some([]int{1, 2, 3}, 4)     // Returns false
//	slicez.Some([]string{"a", "b"}, "c") // Returns false
//	slicez.Some([]int{}, 1)            // Returns false (empty slice)
func Some[A comparable](slice []A, needle A) bool {
	return SomeBy(slice, compare.EqualOf[A](needle))
}

// SomeBy checks if any element satisfies the predicate.
// Returns false if the slice is empty.
// Alias for ContainsBy.
//
// Example:
//
//	// Check if any number is even
//	slicez.SomeBy([]int{1, 3, 5, 6, 7}, func(n int) bool {
//	    return n%2 == 0
//	}) // Returns true (6 is even)
//
//	// Check if any string is longer than 5 chars
//	slicez.SomeBy([]string{"a", "bb", "ccc"}, func(s string) bool {
//	    return len(s) > 5
//	}) // Returns false
//
//	slicez.SomeBy([]int{}, func(n int) bool { return n > 0 })
//	// Returns false (empty slice)
func SomeBy[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if predicate(val) {
			return true
		}
	}
	return false
}

// None checks if no element equals the given value.
// Negation of Some/Contains. Returns true if the slice is empty.
// Uses equality comparison (==).
//
// Example:
//
//	slicez.None([]int{1, 2, 3}, 4)     // Returns true (4 not in slice)
//	slicez.None([]int{1, 2, 3}, 2)     // Returns false (2 is in slice)
//	slicez.None([]string{"a", "b"}, "c") // Returns true
//	slicez.None([]int{}, 1)             // Returns true (empty slice)
func None[A comparable](slice []A, needle A) bool {
	return !SomeBy(slice, compare.EqualOf[A](needle))
}

// NoneBy checks if no element satisfies the predicate.
// Negation of SomeBy. Returns true if the slice is empty.
// Useful for ensuring no elements match a condition.
//
// Example:
//
//	// Check no numbers are negative
//	slicez.NoneBy([]int{1, 2, 3, 4}, func(n int) bool {
//	    return n < 0
//	}) // Returns true
//
//	// Check no strings are empty
//	slicez.NoneBy([]string{"a", "b", ""}, func(s string) bool {
//	    return len(s) == 0
//	}) // Returns false (empty string exists)
//
//	slicez.NoneBy([]int{}, func(n int) bool { return n < 0 })
//	// Returns true (empty slice)
func NoneBy[A any](slice []A, predicate func(A) bool) bool {
	return !SomeBy(slice, predicate)
}

// Partition splits a slice into two based on a predicate.
// Returns two slices: one with elements where predicate is true, one where false.
// Both slices maintain original relative ordering.
//
// Example:
//
//	// Separate even and odd numbers
//	even, odd := slicez.Partition([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
//	    return n%2 == 0
//	})
//	// even = []int{2, 4, 6}
//	// odd = []int{1, 3, 5}
//
//	// Separate positive and non-positive
//	pos, nonPos := slicez.Partition([]int{-2, -1, 0, 1, 2}, func(n int) bool {
//	    return n > 0
//	})
//	// pos = []int{1, 2}
//	// nonPos = []int{-2, -1, 0}
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

// PartitionBy groups elements into consecutive groups based on a key function.
// Unlike GroupBy which groups all matching elements, PartitionBy only groups
// consecutive elements with the same key. Similar to "chunking" by key.
//
// Example:
//
//	// Group consecutive equal numbers
//	slicez.PartitionBy([]int{1, 1, 2, 2, 2, 3, 3}, func(a int) int { return a })
//	// Returns [][]int{{1, 1}, {2, 2, 2}, {3, 3}}
//
//	// Group by first letter
//	slicez.PartitionBy([]string{"apple", "avocado", "banana", "blueberry", "cherry"},
//	    func(s string) string { return string(s[0]) })
//	// Returns [][]string{{"apple", "avocado"}, {"banana", "blueberry"}, {"cherry"}}
//
//	// Non-consecutive elements with same key get separate groups
//	slicez.PartitionBy([]int{1, 2, 1}, func(a int) int { return a })
//	// Returns [][]int{{1}, {2}, {1}}
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

// Chunk splits a slice into chunks of size n.
// Each chunk (except possibly the last) has exactly n elements.
// The last chunk may have fewer than n elements if the slice length isn't divisible by n.
//
// Example:
//
//	slicez.Chunk([]int{1, 2, 3, 4, 5, 6, 7}, 3)
//	// Returns [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
//
//	slicez.Chunk([]string{"a", "b", "c", "d", "e"}, 2)
//	// Returns [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
//
//	slicez.Chunk([]int{1, 2, 3}, 5)
//	// Returns [][]int{{1, 2, 3}} (single chunk, smaller than n)
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

// Interleave interleaves multiple slices round-robin style.
// Takes first element from each slice, then second element from each, etc.
// If slices have different lengths, shorter slices are skipped once exhausted.
// Useful for zipping/combining multiple data streams.
//
// Example:
//
//	// Round-robin interleave
//	slicez.Interleave([]int{1, 5, 9}, []int{2, 6, 10}, []int{3, 7, 11}, []int{4, 8, 12})
//	// Returns []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//
//	// Different lengths (shorter slices skipped when exhausted)
//	slicez.Interleave([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
//	// Returns []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	// Two slices (like a simple zip)
//	slicez.Interleave([]int{1, 2, 3}, []int{10, 20, 30})
//	// Returns []int{1, 10, 2, 20, 3, 30}
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

// Shuffle returns a new slice with elements in random order.
// Uses Fisher-Yates shuffle algorithm for uniform random distribution.
// Returns a copy; the original slice is not modified.
// The random source is math/rand; consider seeding for non-deterministic results.
//
// Example:
//
//	slicez.Shuffle([]int{1, 2, 3, 4, 5})
//	// Might return []int{3, 1, 5, 2, 4} (random order)
//
//	slicez.Shuffle([]string{"a", "b", "c", "d"})
//	// Returns strings in random order
//
//	// Empty slice
//	slicez.Shuffle([]int{}) // Returns []int{}
func Shuffle[A any](slice []A) []A {
	var ret = append([]A{}, slice...)
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

// Sample returns n random elements from the slice without replacement.
// Uses efficient algorithms based on sample size: Fisher-Yates shuffle for large samples,
// swap-to-end algorithm for medium slices, and set-based selection for very large slices.
// Returns fewer elements if n > len(slice). Returns empty slice if n <= 0.
// The random source is math/rand; consider seeding for non-deterministic results.
//
// Example:
//
//	slicez.Sample([]int{1, 2, 3, 4, 5}, 3)
//	// Might return []int{2, 5, 1} (3 random elements)
//
//	slicez.Sample([]string{"a", "b", "c"}, 2)
//	// Returns 2 random strings from the slice
//
//	slicez.Sample([]int{1, 2, 3}, 5)
//	// Returns []int{1, 2, 3} (all elements, n > len)
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

// Sort returns a new slice with elements sorted in natural ascending order.
// Uses Go's sort.Slice internally. Returns a copy; the original slice is not modified.
// Works with any Ordered type (integers, floats, strings).
//
// Example:
//
//	slicez.Sort([]int{3, 1, 4, 1, 5, 9, 2, 6})
//	// Returns []int{1, 1, 2, 3, 4, 5, 6, 9}
//
//	slicez.Sort([]string{"cherry", "apple", "banana"})
//	// Returns []string{"apple", "banana", "cherry"}
func Sort[A compare.Ordered](slice []A) []A {
	return SortBy(slice, compare.Less[A])
}

// SortBy returns a new slice sorted using a custom comparison function.
// The less function should return true if a should come before b.
// Returns a copy; the original slice is not modified.
//
// Example:
//
//	// Sort by string length
//	slicez.SortBy([]string{"aaa", "bb", "c", "dddd"}, func(a, b string) bool {
//	    return len(a) < len(b)
//	})
//	// Returns []string{"c", "bb", "aaa", "dddd"}
//
//	// Sort structs by age field
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
//	slicez.SortBy(people, func(a, b Person) bool {
//	    return a.Age < b.Age
//	})
//	// Returns [{Bob 25} {Alice 30} {Charlie 35}]
func SortBy[A any](slice []A, less func(a, b A) bool) []A {
	var res = append([]A{}, slice...)
	sort.Slice(res, less)
	return res
}

// Search performs binary search on a sorted slice.
// Returns the smallest index i where f(slice[i]) is true.
// The slice must be sorted in ascending order according to f.
// If no element satisfies f, returns len(slice) and the zero value.
//
// Example:
//
//	// Find first element >= 23
//	idx, val := slicez.Search([]int{10, 20, 30, 40, 50}, func(e int) bool {
//	    return e >= 23
//	})
//	// Returns (2, 30) - index 2 has value 30
//
//	// Find insertion point for 25 (maintaining sorted order)
//	idx, _ := slicez.Search([]int{10, 20, 30, 40}, func(e int) bool {
//	    return e >= 25
//	})
//	// Returns 2 - insert 25 at index 2
func Search[A any](slice []A, f func(e A) bool) (index int, e A) {
	return sort.Search(slice, f)
}

// Compact removes consecutive duplicate elements from a slice.
// Only removes duplicates that are adjacent; non-consecutive duplicates are kept.
// Uses equality comparison (==) for comparable types.
//
// Example:
//
//	slicez.Compact([]int{1, 1, 2, 1, 2, 2, 2})
//	// Returns []int{1, 2, 1, 2}
//
//	slicez.Compact([]string{"a", "a", "b", "b", "a"})
//	// Returns []string{"a", "b", "a"}
//
//	// Removing duplicate spaces
//	slicez.Compact([]rune("hello    world"))
//	// Returns []rune{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
func Compact[A comparable](slice []A) []A {
	return CompactBy(slice, compare.Equal[A])
}

// CompactBy removes consecutive duplicate elements using a custom equality function.
// Only removes duplicates that are adjacent; non-consecutive duplicates are kept.
//
// Example:
//
//	// Remove consecutive equal-length strings
//	slicez.CompactBy([]string{"hi", "go", "no", "up"}, func(a, b string) bool {
//	    return len(a) == len(b)
//	})
//	// Returns []string{"hi", "go", "up"} (2-letter, 2-letter, 2-letter becomes hi, go, up)
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

// Max returns the largest element from the variadic arguments.
// Works with any Ordered type. Returns zero value if no arguments provided.
//
// Example:
//
//	slicez.Max(3, 1, 4, 1, 5)     // Returns 5
//	slicez.Max("cherry", "apple") // Returns "cherry" (lexicographic comparison)
//	slicez.Max([]int{1, 2, 3}...)  // Returns 3 (slice spread)
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

// Min returns the smallest element from the variadic arguments.
// Works with any Ordered type. Returns zero value if no arguments provided.
//
// Example:
//
//	slicez.Min(3, 1, 4, 1, 5)     // Returns 1
//	slicez.Min("cherry", "apple") // Returns "apple" (lexicographic comparison)
//	slicez.Min([]int{3, 2, 1}...)  // Returns 1 (slice spread)
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

// Flatten flattens a 2D slice into a 1D slice.
// Concatenates all sub-slices in order. Pre-allocates capacity for efficiency.
// Returns empty slice for nil or empty input.
//
// Example:
//
//	slicez.Flatten([][]int{{1, 2}, {3, 4}, {5, 6}})
//	// Returns []int{1, 2, 3, 4, 5, 6}
//
//	slicez.Flatten([][]string{{"a", "b"}, {"c"}, {}, {"d", "e"}})
//	// Returns []string{"a", "b", "c", "d", "e"}
//
//	slicez.Flatten([][]int{})
//	// Returns []int{}
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

// Map transforms each element of the slice using the provided function.
// Returns a new slice with the transformed elements. Preserves order.
//
// Example:
//
//	// Convert integers to strings
//	slicez.Map([]int{1, 2, 3}, func(n int) string {
//	    return fmt.Sprintf("num-%d", n)
//	})
//	// Returns []string{"num-1", "num-2", "num-3"}
//
//	// Double each number
//	slicez.Map([]int{1, 2, 3, 4}, func(n int) int { return n * 2 })
//	// Returns []int{2, 4, 6, 8}
func Map[A any, B any](slice []A, f func(a A) B) []B {
	res := make([]B, 0, len(slice))
	for _, a := range slice {
		res = append(res, f(a))
	}
	return res
}

// FlatMap maps each element to a slice and flattens the results.
// Combines Map and Flatten operations. Useful for "expanding" elements.
//
// Example:
//
//	// Expand each number to [n, n*2]
//	slicez.FlatMap([]int{1, 2, 3}, func(n int) []int {
//	    return []int{n, n * 2}
//	})
//	// Returns []int{1, 2, 2, 4, 3, 6}
//
//	// Split strings into words
//	slicez.FlatMap([]string{"hello world", "foo bar"}, func(s string) []string {
//	    return strings.Split(s, " ")
//	})
//	// Returns []string{"hello", "world", "foo", "bar"}
func FlatMap[A any, B any](slice []A, f func(a A) []B) []B {
	return Flatten(Map(slice, f))
}

// Fold reduces a slice from left to right, accumulating a result.
// Applies the combine function to each element with the current accumulator.
// Also known as Reduce or Inject in other languages.
//
// Example:
//
//	// Sum all numbers
//	slicez.Fold([]int{1, 2, 3, 4}, func(acc, val int) int {
//	    return acc + val
//	}, 0)
//	// Returns 10
//
//	// Concatenate strings with separator
//	slicez.Fold([]string{"a", "b", "c"}, func(acc, val string) string {
//	    if acc == "" {
//	        return val
//	    }
//	    return acc + "-" + val
//	}, "")
//	// Returns "a-b-c"
func Fold[I any, A any](slice []I, combined func(accumulator A, val I) A, init A) A {
	for _, val := range slice {
		init = combined(init, val)
	}
	return init
}

// FoldRight reduces a slice from right to left, accumulating a result.
// Like Fold but processes elements from the end to the beginning.
//
// Example:
//
//	// Build string from right (reverses order)
//	slicez.FoldRight([]string{"a", "b", "c"}, func(acc, val string) string {
//	    if acc == "" {
//	        return val
//	    }
//	    return val + "-" + acc
//	}, "")
//	// Returns "a-b-c" (different associativity than Fold)
//
//	// For subtraction: FoldRight([1,2,3], -, 0) = 1-(2-(3-0)) = 2
//	// While Fold([1,2,3], -, 0) = ((0-1)-2)-3 = -6
func FoldRight[I any, A any](slice []I, combined func(accumulator A, val I) A, init A) A {
	l := len(slice)
	for i := range slice {
		i := l - i - 1
		init = combined(init, slice[i])
	}
	return init
}

// SliceToMap converts a slice to a map using a mapper function.
// Each element is transformed into a key-value pair.
// Later elements overwrite earlier ones if keys collide.
//
// Example:
//
//	// Create map from id to name
//	type User struct { ID int; Name string }
//	users := []User{{1, "Alice"}, {2, "Bob"}}
//	slicez.SliceToMap(users, func(u User) (int, string) {
//	    return u.ID, u.Name
//	})
//	// Returns map[int]string{1: "Alice", 2: "Bob"}
//
// Alias for Associate.
func SliceToMap[E any, K comparable, V any](slice []E, mapper func(a E) (key K, value V)) map[K]V {
	return Associate(slice, mapper)
}

// Associate converts a slice to a map using a mapper function.
// Each element is transformed into a key-value pair.
// Later elements overwrite earlier ones if keys collide.
// Pre-allocates map capacity for efficiency.
//
// Example:
//
//	// Index strings by length
//	slicez.Associate([]string{"hi", "hello", "hey"}, func(s string) (int, string) {
//	    return len(s), s
//	})
//	// Returns map[int]string{2: "hi", 3: "hey", 5: "hello"}
func Associate[E any, K comparable, V any](slice []E, mapper func(e E) (key K, value V)) map[K]V {
	acc := make(map[K]V, len(slice))
	for _, e := range slice {
		k, v := mapper(e)
		acc[k] = v
	}
	return acc
}

// Set creates a set (as map[E]bool) from a slice.
// Returns a map where keys are slice elements and values are true.
// Useful for O(1) membership testing.
//
// Example:
//
//	set := slicez.Set([]int{1, 2, 3, 2, 1})
//	// Returns map[int]bool{1: true, 2: true, 3: true}
//
//	// Check membership
//	if set[2] {
//	    fmt.Println("2 is in the set")
//	}
//
// Note: For more set operations, consider using the setz package.
func Set[E comparable](slice []E) map[E]bool {
	return Associate(slice, func(a E) (key E, value bool) {
		return a, true
	})
}

// KeyBy creates a map indexed by a key function.
// Each element is indexed by the result of the key function.
// If multiple elements have the same key, the first one is kept.
//
// Example:
//
//	type User struct { ID int; Name string }
//	users := []User{{1, "Alice"}, {2, "Bob"}, {1, "Charlie"}}
//	slicez.KeyBy(users, func(u User) int { return u.ID })
//	// Returns map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
//	// Note: {1, "Charlie"} is skipped because ID 1 already exists
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

// GroupBy groups slice elements by a key function into a map.
// Returns a map where each key maps to a slice of all elements with that key.
// For ordered iteration, use GroupByOrdered instead.
//
// Example:
//
//	// Group integers by even/odd
//	slicez.GroupBy([]int{1, 2, 3, 4, 5}, func(n int) string {
//	    if n%2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	})
//	// Returns map[string][]int{"odd": {1, 3, 5}, "even": {2, 4}}
//
//	// Group strings by first letter
//	slicez.GroupBy([]string{"apple", "avocado", "banana"}, func(s string) string {
//	    return string(s[0])
//	})
//	// Returns map[string][]string{"a": {"apple", "avocado"}, "b": {"banana"}}
func GroupBy[A any, B comparable](slice []A, key func(a A) B) map[B][]A {
	m := make(map[B][]A)
	for _, v := range slice {
		k := key(v)
		m[k] = append(m[k], v)
	}
	return m
}

// Uniq removes all duplicate elements from a slice, keeping the first occurrence.
// Uses equality comparison (==) for comparable types. Preserves order.
//
// Example:
//
//	slicez.Uniq([]int{1, 2, 2, 3, 1, 4})
//	// Returns []int{1, 2, 3, 4}
//
//	slicez.Uniq([]string{"a", "b", "a", "c", "b"})
//	// Returns []string{"a", "b", "c"}
func Uniq[A comparable](slice []A) []A {
	return UniqBy(slice, compare.Identity[A])
}

// UniqBy removes duplicates using a key function, keeping the first occurrence.
// Elements are considered duplicates if they have the same key. Preserves order.
//
// Example:
//
//	// Uniq by string length
//	slicez.UniqBy([]string{"hi", "go", "no", "hello", "hey"}, func(s string) int {
//	    return len(s)
//	})
//	// Returns []string{"hi", "hello"} (2-letter and 5-letter strings)
//
//	// Uniq by age
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 30}}
//	slicez.UniqBy(people, func(p Person) int { return p.Age })
//	// Returns [{Alice 30} {Bob 25}]
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

// Union returns the set union of multiple slices.
// Combines all unique elements from all slices, preserving order of first appearance.
// Equivalent to Uniq(Concat(slices...)).
//
// Example:
//
//	slicez.Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
//	// Returns []int{1, 2, 3, 4, 5}
//
//	slicez.Union([]string{"a", "b"}, []string{"c"}, []string{"a", "d"})
//	// Returns []string{"a", "b", "c", "d"}
func Union[A comparable](slices ...[]A) []A {
	return UnionBy(compare.Identity[A], slices...)
}

// UnionBy returns the set union using a key function.
// Elements are considered equal if they have the same key.
// Preserves order of first appearance of each unique key.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	a := []Person{{"Alice", 30}, {"Bob", 25}}
//	b := []Person{{"Charlie", 30}, {"Dave", 35}}
//	slicez.UnionBy(func(p Person) int { return p.Age }, a, b)
//	// Returns [{Alice 30} {Bob 25} {Dave 35}]
//	// Charlie is excluded because age 30 already seen
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

// Intersection returns the set intersection of multiple slices.
// Returns elements that appear in ALL slices. Preserves order from first slice.
//
// Example:
//
//	slicez.Intersection([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
//	// Returns []int{3} (only element in all three)
//
//	slicez.Intersection([]string{"a", "b", "c"}, []string{"b", "c", "d"})
//	// Returns []string{"b", "c"}
func Intersection[A comparable](slices ...[]A) []A {
	return IntersectionBy(compare.Identity[A], slices...)
}

// IntersectionBy returns the intersection using a key function.
// Returns elements whose keys appear in ALL slices. Preserves order from first slice.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	a := []Person{{"Alice", 30}, {"Bob", 25}}
//	b := []Person{{"Charlie", 30}, {"Dave", 25}}
//	slicez.IntersectionBy(func(p Person) int { return p.Age }, a, b)
//	// Returns [{Alice 30} {Bob 25}] (ages 30 and 25 in both)
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

// Difference returns elements that are unique across all slices (symmetric difference).
// Returns elements that appear in exactly one of the input slices.
// Removes elements that appear in the intersection of any two or more slices.
//
// Example:
//
//	slicez.Difference([]int{1, 2, 3}, []int{2, 3, 4})
//	// Returns []int{1, 4} (2 and 3 are in both, so removed)
//
//	// Three slices: keep only elements in exactly one slice
//	slicez.Difference([]int{1, 2}, []int{2, 3}, []int{3, 4})
//	// Returns []int{1, 4} (2 is in first two, 3 is in last two)
func Difference[A comparable](slices ...[]A) []A {
	return DifferenceBy(compare.Identity[A], slices...)
}

// DifferenceBy returns elements unique across slices using a key function.
// Returns elements whose keys appear in exactly one slice.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	a := []Person{{"Alice", 30}, {"Bob", 25}}
//	b := []Person{{"Charlie", 30}, {"Dave", 35}}
//	slicez.DifferenceBy(func(p Person) int { return p.Age }, a, b)
//	// Returns [{Bob 25} {Dave 35}] (ages 30 in both, ages 25 and 35 unique)
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

// Complement returns elements in b that are not in a (set difference b - a).
// Returns a copy of b with all elements from a removed.
// Useful for filtering: keep only elements from b not seen in a.
//
// Example:
//
//	allowed := []int{1, 2, 3}
//	all := []int{1, 2, 3, 4, 5, 6}
//	slicez.Complement(allowed, all)
//	// Returns []int{4, 5, 6} (elements in all but not in allowed)
//
//	// Filter out seen items
//	seen := []string{"a", "b"}
//	newItems := []string{"b", "c", "d", "a"}
//	slicez.Complement(seen, newItems)
//	// Returns []string{"c", "d"}
func Complement[A comparable](a, b []A) []A {
	return ComplementBy(compare.Identity[A], a, b)
}

// ComplementBy returns elements in b not in a using a key function.
// Elements are compared by their keys rather than direct equality.
//
// Example:
//
//	type User struct { ID int; Name string }
//	existing := []User{{1, "Alice"}, {2, "Bob"}}
//	newUsers := []User{{2, "Robert"}, {3, "Charlie"}, {4, "Dave"}}
//	slicez.ComplementBy(func(u User) int { return u.ID }, existing, newUsers)
//	// Returns [{3, "Charlie"}, {4, "Dave"}] (ID 2 already exists)
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

// Zip combines two slices element-wise using a zipper function.
// Stops at the length of the shorter slice. Returns empty slice if either input is empty.
//
// Example:
//
//	// Pair elements
//	slicez.Zip([]int{1, 2, 3}, []string{"a", "b", "c"}, func(n int, s string) string {
//	    return fmt.Sprintf("%d:%s", n, s)
//	})
//	// Returns []string{"1:a", "2:b", "3:c"}
//
//	// Add corresponding elements
//	slicez.Zip([]int{1, 2, 3}, []int{10, 20, 30}, func(a, b int) int {
//	    return a + b
//	})
//	// Returns []int{11, 22, 33}
//
//	// Different lengths (truncated)
//	slicez.Zip([]int{1, 2, 3, 4}, []string{"a", "b"}, func(n int, s string) string {
//	    return s + strconv.Itoa(n)
//	})
//	// Returns []string{"a1", "b2"} (extra elements ignored)
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

// Unzip splits a slice of pairs into two separate slices.
// Inverse operation of Zip. Both returned slices have the same length as input.
//
// Example:
//
//	// Split pairs into components
//	pairs := []struct{ X, Y int }{{1, 10}, {2, 20}, {3, 30}}
//	xs, ys := slicez.Unzip(pairs, func(p struct{ X, Y int }) (int, int) {
//	    return p.X, p.Y
//	})
//	// xs = []int{1, 2, 3}, ys = []int{10, 20, 30}
//
//	// Unzip strings to bytes
//	slicez.Unzip([]string{"hi", "go"}, func(s string) (rune, rune) {
//	    return rune(s[0]), rune(s[1])
//	})
//	// Returns ([]rune{'h', 'g'}, []rune{'i', 'o'})
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

// Zip2 combines three slices element-wise using a zipper function.
// Like Zip but for three slices. Stops at length of shortest slice.
//
// Example:
//
//	// Combine three number sequences
//	slicez.Zip2([]int{1, 2}, []int{10, 20}, []int{100, 200}, func(a, b, c int) int {
//	    return a + b + c
//	})
//	// Returns []int{111, 222}
//
//	// Create structured data
//	type Point struct{ X, Y, Z int }
//	slicez.Zip2([]int{1, 2}, []int{3, 4}, []int{5, 6}, func(x, y, z int) Point {
//	    return Point{x, y, z}
//	})
//	// Returns []Point{{1,3,5}, {2,4,6}}
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

// Unzip2 splits a slice of triples into three separate slices.
// Inverse operation of Zip2. All returned slices have the same length.
//
// Example:
//
//	// Split RGB values
//	colors := []struct{ R, G, B uint8 }{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}}
//	rs, gs, bs := slicez.Unzip2(colors, func(c struct{ R, G, B uint8 }) (uint8, uint8, uint8) {
//	    return c.R, c.G, c.B
//	})
//	// rs = []uint8{255, 0, 0}, gs = []uint8{0, 255, 0}, bs = []uint8{0, 0, 255}
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

// Zip3 combines four slices element-wise using a zipper function.
// Like Zip and Zip2 but for four slices. Stops at length of shortest slice.
//
// Example:
//
//	// Combine four number sequences
//	a, b, c, d := []int{1, 2}, []int{10, 20}, []int{100, 200}, []int{1000, 2000}
//	slicez.Zip3(a, b, c, d, func(w, x, y, z int) int {
//	    return w + x + y + z
//	})
//	// Returns []int{1111, 2222}
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

// Unzip3 splits a slice of quadruples into four separate slices.
// Inverse operation of Zip3. All returned slices have the same length.
//
// Example:
//
//	// Split 4D coordinates
//	coords := []struct{ X, Y, Z, W float64 }{{1, 2, 3, 4}, {5, 6, 7, 8}}
//	xs, ys, zs, ws := slicez.Unzip3(coords, func(c struct{ X, Y, Z, W float64 }) (float64, float64, float64, float64) {
//	    return c.X, c.Y, c.Z, c.W
//	})
//	// xs = []float64{1, 5}, ys = []float64{2, 6}, zs = []float64{3, 7}, ws = []float64{4, 8}
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

// XOR returns the symmetric difference of multiple slices.
// Returns elements that appear in exactly one slice (not in multiple slices).
// Also known as symmetric difference. Preserves order of first appearance.
//
// Example:
//
//	slicez.XOR([]int{1, 2, 3}, []int{2, 3, 4})
//	// Returns []int{1, 4} (2 and 3 appear in both, so excluded)
//
//	slicez.XOR([]int{1, 2}, []int{2, 3}, []int{3, 4})
//	// Returns []int{1, 4} (2 in first two, 3 in last two, 1 and 4 in exactly one)
func XOR[A comparable](slices ...[]A) []A {
	return XORBy(compare.Identity[A], slices...)
}

// XORBy returns the symmetric difference using a key function.
// Returns elements whose keys appear in exactly one slice.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	a := []Person{{"Alice", 30}, {"Bob", 25}}
//	b := []Person{{"Charlie", 30}, {"Dave", 35}}
//	slicez.XORBy(func(p Person) int { return p.Age }, a, b)
//	// Returns [{Bob 25} {Dave 35}] (age 30 in both, ages 25 and 35 unique)
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
