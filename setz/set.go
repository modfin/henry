// Package setz provides generic set data structures and operations.
// It offers both a Set type with methods for fluent API usage and standalone
// functions for interoperability with other packages in the henry library.
package setz

import "fmt"

// Set represents a mathematical set of unique elements.
// It is implemented as a type alias of map[T]struct{} for O(1) operations.
// The Set is mutable - methods like Add, Remove, and Clear modify the set.
// Operations like Union, Intersection return new sets (immutable results).
type Set[T comparable] map[T]struct{}

// New creates a new empty Set with optional initial elements.
//
// Example:
//
//	s := setz.New[int]()           // Empty set
//	s := setz.New(1, 2, 3)         // Set with elements 1, 2, 3
//	s := setz.New("a", "b", "c")     // Set of strings
func New[T comparable](elements ...T) Set[T] {
	s := make(Set[T], len(elements))
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

// FromSlice creates a Set from a slice, removing duplicates.
//
// Example:
//
//	s := setz.FromSlice([]int{1, 2, 2, 3, 3, 3})
//	// s contains {1, 2, 3}
func FromSlice[T comparable](slice []T) Set[T] {
	return New(slice...)
}

// FromMap creates a Set from an existing map[T]struct{}.
// This allows interoperability with slicez.Set() output.
//
// Example:
//
//	m := map[string]struct{}{"a": {}, "b": {}}
//	s := setz.FromMap(m)  // Set with "a", "b"
func FromMap[T comparable](m Set[T]) Set[T] {
	s := make(Set[T], len(m))
	for k := range m {
		s[k] = struct{}{}
	}
	return s
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s)
}

// IsEmpty returns true if the set has no elements.
func (s Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Add adds elements to the set. Returns the set for chaining.
//
// Example:
//
//	s := setz.New[int]().Add(1, 2, 3).Add(4)
func (s Set[T]) Add(elements ...T) Set[T] {
	if s == nil {
		s = make(Set[T])
	}
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

// Remove removes elements from the set. Returns the set for chaining.
//
// Example:
//
//	s := setz.New(1, 2, 3, 4).Remove(2, 3)  // s contains {1, 4}
func (s Set[T]) Remove(elements ...T) Set[T] {
	if s == nil {
		return s
	}
	for _, e := range elements {
		delete(s, e)
	}
	return s
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	if s == nil {
		return
	}
	for e := range s {
		delete(s, e)
	}
}

// Contains returns true if the element is in the set.
//
// Example:
//
//	s := setz.New(1, 2, 3)
//	s.Contains(2)  // true
//	s.Contains(5)  // false
func (s Set[T]) Contains(element T) bool {
	if s == nil {
		return false
	}
	_, exists := s[element]
	return exists
}

// ContainsAll returns true if all elements are in the set.
func (s Set[T]) ContainsAll(elements ...T) bool {
	for _, e := range elements {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// Pop removes and returns an arbitrary element from the set.
// Returns the zero value and false if the set is empty.
func (s Set[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	for e := range s {
		delete(s, e)
		return e, true
	}
	return zero, false
}

// ToSlice returns all elements as a slice (order not guaranteed).
func (s Set[T]) ToSlice() []T {
	if s.IsEmpty() {
		return []T{}
	}
	result := make([]T, 0, s.Len())
	for e := range s {
		result = append(result, e)
	}
	return result
}

// Copy returns a new set with the same elements.
func (s Set[T]) Copy() Set[T] {
	return FromMap(s)
}

// Union returns a new set with all elements from both sets.
//
// Example:
//
//	s1 := setz.New(1, 2, 3)
//	s2 := setz.New(3, 4, 5)
//	s1.Union(s2)  // {1, 2, 3, 4, 5}
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := s.Copy()
	for e := range other {
		result[e] = struct{}{}
	}
	return result
}

// Intersection returns a new set with elements common to both sets.
//
// Example:
//
//	s1 := setz.New(1, 2, 3, 4)
//	s2 := setz.New(2, 3, 4, 5)
//	s1.Intersection(s2)  // {2, 3, 4}
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	if s.IsEmpty() || other.IsEmpty() {
		return New[T]()
	}

	// Iterate over the smaller set for efficiency
	small, large := s, other
	if s.Len() > other.Len() {
		small, large = other, s
	}

	result := New[T]()
	for e := range small {
		if large.Contains(e) {
			result[e] = struct{}{}
		}
	}
	return result
}

// Difference returns a new set with elements in s but not in other.
//
// Example:
//
//	s1 := setz.New(1, 2, 3, 4)
//	s2 := setz.New(2, 4)
//	s1.Difference(s2)  // {1, 3}
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	if s.IsEmpty() {
		return result
	}

	for e := range s {
		if _, exists := other[e]; !exists {
			result[e] = struct{}{}
		}
	}
	return result
}

// SymmetricDifference returns a new set with elements in either set but not both.
// This is equivalent to (s ∪ other) - (s ∩ other).
//
// Example:
//
//	s1 := setz.New(1, 2, 3)
//	s2 := setz.New(2, 3, 4)
//	s1.SymmetricDifference(s2)  // {1, 4}
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := New[T]()
	for e := range s {
		if _, exists := other[e]; !exists {
			result[e] = struct{}{}
		}
	}
	for e := range other {
		if _, exists := s[e]; !exists {
			result[e] = struct{}{}
		}
	}
	return result
}

// IsSubset returns true if all elements of s are in other.
//
// Example:
//
//	s1 := setz.New(1, 2)
//	s2 := setz.New(1, 2, 3)
//	s1.IsSubset(s2)  // true
//	s2.IsSubset(s1)  // false
func (s Set[T]) IsSubset(other Set[T]) bool {
	if s.IsEmpty() {
		return true
	}
	if other.IsEmpty() {
		return false
	}
	for e := range s {
		if _, exists := other[e]; !exists {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s contains all elements of other.
//
// Example:
//
//	s1 := setz.New(1, 2, 3)
//	s2 := setz.New(1, 2)
//	s1.IsSuperset(s2)  // true
func (s Set[T]) IsSuperset(other Set[T]) bool {
	if other.IsEmpty() {
		return true
	}
	if s.IsEmpty() {
		return false
	}
	return other.IsSubset(s)
}

// IsProperSubset returns true if s is a subset of other and s ≠ other.
func (s Set[T]) IsProperSubset(other Set[T]) bool {
	return s.IsSubset(other) && s.Len() < other.Len()
}

// IsProperSuperset returns true if s is a superset of other and s ≠ other.
func (s Set[T]) IsProperSuperset(other Set[T]) bool {
	return s.IsSuperset(other) && s.Len() > other.Len()
}

// IsDisjoint returns true if s and other have no elements in common.
//
// Example:
//
//	s1 := setz.New(1, 2)
//	s2 := setz.New(3, 4)
//	s1.IsDisjoint(s2)  // true
func (s Set[T]) IsDisjoint(other Set[T]) bool {
	if s.IsEmpty() || other.IsEmpty() {
		return true
	}

	small, large := s, other
	if s.Len() > other.Len() {
		small, large = other, s
	}

	for e := range small {
		if _, exists := large[e]; exists {
			return false
		}
	}
	return true
}

// IsEqual returns true if s and other contain the same elements.
func (s Set[T]) IsEqual(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	return s.IsSubset(other)
}

// String returns a string representation of the set.
// Note: Element order is not guaranteed.
func (s Set[T]) String() string {
	if s.IsEmpty() {
		return "Set{}"
	}
	return fmt.Sprintf("Set%v", s.ToSlice())
}

// Filter returns a new set with elements that satisfy the predicate.
func (s Set[T]) Filter(predicate func(T) bool) Set[T] {
	result := New[T]()
	for e := range s {
		if predicate(e) {
			result[e] = struct{}{}
		}
	}
	return result
}

// Map returns a new set by applying a function to each element.
func Map[T comparable, U comparable](s Set[T], mapper func(T) U) Set[U] {
	result := New[U]()
	if s.IsEmpty() {
		return result
	}
	for e := range s {
		result[mapper(e)] = struct{}{}
	}
	return result
}

// Union returns a new set with all elements from the given sets.
// Standalone function version.
func Union[T comparable](sets ...Set[T]) Set[T] {
	result := New[T]()
	for _, s := range sets {
		for e := range s {
			result[e] = struct{}{}
		}
	}
	return result
}

// Intersection returns a new set with elements common to all sets.
// Standalone function version.
func Intersection[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return New[T]()
	}

	// Find the first non-empty set
	var first Set[T]
	for _, s := range sets {
		if !s.IsEmpty() {
			first = s
			break
		}
	}
	if first.IsEmpty() {
		return New[T]()
	}

	// Start with the first set and intersect with others
	result := first.Copy()
	for _, s := range sets {
		result = result.Intersection(s)
	}
	return result
}

// Difference returns a new set with elements in the first set but not in others.
// Standalone function version.
func Difference[T comparable](first Set[T], others ...Set[T]) Set[T] {
	if first.IsEmpty() {
		return New[T]()
	}

	result := first.Copy()
	for _, s := range others {
		for e := range s {
			delete(result, e)
		}
	}
	return result
}

// Contains returns true if the element is in the set.
// Standalone function version.
func Contains[T comparable](s Set[T], element T) bool {
	_, exists := s[element]
	return exists
}

// ToSlice returns all elements from a set as a slice.
// Standalone function version.
func ToSlice[T comparable](s Set[T]) []T {
	if len(s) == 0 {
		return []T{}
	}
	result := make([]T, 0, len(s))
	for e := range s {
		result = append(result, e)
	}
	return result
}

// IsSubset returns true if all elements of a are in b.
// Standalone function version.
func IsSubset[T comparable](a, b Set[T]) bool {
	if len(a) == 0 {
		return true
	}
	if len(b) == 0 {
		return false
	}
	for e := range a {
		if _, exists := b[e]; !exists {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if a and b have no elements in common.
// Standalone function version.
func IsDisjoint[T comparable](a, b Set[T]) bool {
	if len(a) == 0 || len(b) == 0 {
		return true
	}

	small, large := a, b
	if len(a) > len(b) {
		small, large = b, a
	}

	for e := range small {
		if _, exists := large[e]; exists {
			return false
		}
	}
	return true
}
