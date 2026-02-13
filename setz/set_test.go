package setz

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Empty set
	s := New[int]()
	if !s.IsEmpty() {
		t.Error("New() should create empty set")
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}

	// Set with initial elements
	s2 := New(1, 2, 3)
	if s2.Len() != 3 {
		t.Errorf("Len() = %d, want 3", s2.Len())
	}
	if !s2.Contains(1) || !s2.Contains(2) || !s2.Contains(3) {
		t.Error("Set should contain all initial elements")
	}

	// Duplicate elements should be deduplicated
	s3 := New(1, 1, 2, 2, 3)
	if s3.Len() != 3 {
		t.Errorf("Len() with duplicates = %d, want 3", s3.Len())
	}
}

func TestFromSlice(t *testing.T) {
	slice := []int{1, 2, 2, 3, 3, 3}
	s := FromSlice(slice)
	if s.Len() != 3 {
		t.Errorf("Len() = %d, want 3", s.Len())
	}
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Error("Set should contain all unique elements from slice")
	}

	// Empty slice
	empty := FromSlice([]int{})
	if !empty.IsEmpty() {
		t.Error("FromSlice([]) should create empty set")
	}
}

func TestFromMap(t *testing.T) {
	m := map[string]struct{}{"a": {}, "b": {}, "c": {}}
	s := FromMap(m)
	if s.Len() != 3 {
		t.Errorf("Len() = %d, want 3", s.Len())
	}
	if !s.Contains("a") || !s.Contains("b") || !s.Contains("c") {
		t.Error("Set should contain all keys from map")
	}

	// Empty map
	empty := FromMap(map[int]struct{}{})
	if !empty.IsEmpty() {
		t.Error("FromMap({}) should create empty set")
	}
}

func TestLenAndIsEmpty(t *testing.T) {
	s := New[int]()
	if s.Len() != 0 || !s.IsEmpty() {
		t.Error("Empty set should have Len=0 and IsEmpty=true")
	}

	s.Add(1)
	if s.Len() != 1 || s.IsEmpty() {
		t.Error("Set with 1 element should have Len=1 and IsEmpty=false")
	}

	s.Add(2, 3)
	if s.Len() != 3 {
		t.Errorf("Set with 3 elements should have Len=3, got %d", s.Len())
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1).Add(2).Add(3)
	if s.Len() != 3 {
		t.Errorf("After adding 3 elements, Len() = %d", s.Len())
	}

	// Adding duplicate should not change size
	s.Add(1)
	if s.Len() != 3 {
		t.Error("Adding duplicate should not increase size")
	}

	// Add multiple at once
	s2 := New[int]()
	s2.Add(10, 20, 30)
	if s2.Len() != 3 {
		t.Errorf("Add multiple should add all, Len() = %d", s2.Len())
	}
}

func TestRemove(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	s.Remove(3)
	if s.Contains(3) {
		t.Error("After Remove(3), set should not contain 3")
	}
	if s.Len() != 4 {
		t.Errorf("After removing 1 element, Len() = %d, want 4", s.Len())
	}

	// Remove multiple
	s.Remove(1, 2)
	if s.Len() != 2 {
		t.Errorf("After removing 2 more, Len() = %d, want 2", s.Len())
	}

	// Remove non-existent element
	s.Remove(100)
	if s.Len() != 2 {
		t.Error("Removing non-existent should not change size")
	}
}

func TestClear(t *testing.T) {
	s := New(1, 2, 3)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("After Clear(), set should be empty")
	}
	if s.Len() != 0 {
		t.Errorf("After Clear(), Len() = %d", s.Len())
	}
}

func TestContains(t *testing.T) {
	s := New(1, 2, 3)
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}

	// Nil set
	var nilSet *Set[int]
	if nilSet.Contains(1) {
		t.Error("Nil set should not contain anything")
	}
}

func TestContainsAll(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	if !s.ContainsAll(1, 2, 3) {
		t.Error("Set should contain all of 1, 2, 3")
	}
	if s.ContainsAll(1, 2, 10) {
		t.Error("Set should not contain all when 10 is missing")
	}
	if !s.ContainsAll() {
		t.Error("ContainsAll with no args should return true")
	}
}

func TestPop(t *testing.T) {
	s := New(1, 2, 3)
	elem, ok := s.Pop()
	if !ok {
		t.Error("Pop from non-empty set should return ok=true")
	}
	if s.Contains(elem) {
		t.Error("Popped element should be removed from set")
	}
	if s.Len() != 2 {
		t.Errorf("After Pop, Len() = %d, want 2", s.Len())
	}

	// Pop until empty
	for s.Len() > 0 {
		s.Pop()
	}
	_, ok = s.Pop()
	if ok {
		t.Error("Pop from empty set should return ok=false")
	}
}

func TestToSlice(t *testing.T) {
	s := New(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("ToSlice() length = %d, want 3", len(slice))
	}

	// Check all elements present (order not guaranteed)
	elements := make(map[int]bool)
	for _, e := range slice {
		elements[e] = true
	}
	if !elements[1] || !elements[2] || !elements[3] {
		t.Error("ToSlice should contain all set elements")
	}

	// Empty set
	empty := New[int]()
	if len(empty.ToSlice()) != 0 {
		t.Error("Empty set ToSlice should be empty")
	}
}

func TestToMap(t *testing.T) {
	s := New("a", "b", "c")
	m := s.ToMap()
	if len(m) != 3 {
		t.Errorf("ToMap() length = %d, want 3", len(m))
	}
	if _, ok := m["a"]; !ok {
		t.Error("ToMap should contain 'a'")
	}

	// Modifying returned map should not affect set
	delete(m, "a")
	if !s.Contains("a") {
		t.Error("Modifying ToMap result should not affect original set")
	}
}

func TestCopy(t *testing.T) {
	s := New(1, 2, 3)
	sCopy := s.Copy()

	// Should have same elements
	if !sCopy.IsEqual(s) {
		t.Error("Copy should have same elements as original")
	}

	// Modifying copy should not affect original
	sCopy.Add(4)
	if s.Contains(4) {
		t.Error("Modifying copy should not affect original")
	}
}

func TestUnion(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(3, 4, 5)
	result := s1.Union(s2)

	if result.Len() != 5 {
		t.Errorf("Union of {1,2,3} and {3,4,5} should have 5 elements, got %d", result.Len())
	}
	if !result.ContainsAll(1, 2, 3, 4, 5) {
		t.Error("Union should contain all elements from both sets")
	}

	// Union with nil
	result2 := s1.Union(nil)
	if !result2.IsEqual(s1) {
		t.Error("Union with nil should return copy of first set")
	}

	// Original sets should not be modified
	if s1.Len() != 3 || s2.Len() != 3 {
		t.Error("Union should not modify original sets")
	}
}

func TestIntersection(t *testing.T) {
	s1 := New(1, 2, 3, 4)
	s2 := New(3, 4, 5, 6)
	result := s1.Intersection(s2)

	if result.Len() != 2 {
		t.Errorf("Intersection of {1,2,3,4} and {3,4,5,6} should have 2 elements, got %d", result.Len())
	}
	if !result.ContainsAll(3, 4) {
		t.Error("Intersection should contain common elements 3 and 4")
	}
	if result.Contains(1) || result.Contains(5) {
		t.Error("Intersection should not contain non-common elements")
	}

	// Intersection with nil
	result2 := s1.Intersection(nil)
	if !result2.IsEmpty() {
		t.Error("Intersection with nil should return empty set")
	}

	// Empty intersection
	s3 := New(1, 2)
	s4 := New(3, 4)
	result3 := s3.Intersection(s4)
	if !result3.IsEmpty() {
		t.Error("Intersection of disjoint sets should be empty")
	}
}

func TestDifference(t *testing.T) {
	s1 := New(1, 2, 3, 4)
	s2 := New(3, 4, 5)
	result := s1.Difference(s2)

	if result.Len() != 2 {
		t.Errorf("Difference should have 2 elements, got %d", result.Len())
	}
	if !result.ContainsAll(1, 2) {
		t.Error("Difference should contain elements only in s1")
	}
	if result.Contains(3) || result.Contains(4) {
		t.Error("Difference should not contain elements in s2")
	}

	// Difference with nil
	result2 := s1.Difference(nil)
	if !result2.IsEqual(s1) {
		t.Error("Difference with nil should return copy of first set")
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(2, 3, 4)
	result := s1.SymmetricDifference(s2)

	if result.Len() != 2 {
		t.Errorf("SymmetricDifference should have 2 elements, got %d", result.Len())
	}
	if !result.ContainsAll(1, 4) {
		t.Error("SymmetricDifference should contain elements in exactly one set")
	}
	if result.Contains(2) || result.Contains(3) {
		t.Error("SymmetricDifference should not contain common elements")
	}
}

func TestIsSubset(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(1, 2, 3, 4)

	if !s1.IsSubset(s2) {
		t.Error("{1,2} should be subset of {1,2,3,4}")
	}
	if s2.IsSubset(s1) {
		t.Error("{1,2,3,4} should not be subset of {1,2}")
	}

	// Empty set is subset of everything
	empty := New[int]()
	if !empty.IsSubset(s2) {
		t.Error("Empty set should be subset of any set")
	}

	// Set is subset of itself
	if !s1.IsSubset(s1) {
		t.Error("Set should be subset of itself")
	}
}

func TestIsSuperset(t *testing.T) {
	s1 := New(1, 2, 3, 4)
	s2 := New(1, 2)

	if !s1.IsSuperset(s2) {
		t.Error("{1,2,3,4} should be superset of {1,2}")
	}
	if s2.IsSuperset(s1) {
		t.Error("{1,2} should not be superset of {1,2,3,4}")
	}

	// Every set is superset of empty set
	empty := New[int]()
	if !s1.IsSuperset(empty) {
		t.Error("Any set should be superset of empty set")
	}

	// Set is superset of itself
	if !s1.IsSuperset(s1) {
		t.Error("Set should be superset of itself")
	}
}

func TestIsProperSubset(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(1, 2, 3)

	if !s1.IsProperSubset(s2) {
		t.Error("{1,2} should be proper subset of {1,2,3}")
	}
	if s1.IsProperSubset(s1) {
		t.Error("Set should not be proper subset of itself")
	}
}

func TestIsProperSuperset(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(1, 2)

	if !s1.IsProperSuperset(s2) {
		t.Error("{1,2,3} should be proper superset of {1,2}")
	}
	if s1.IsProperSuperset(s1) {
		t.Error("Set should not be proper superset of itself")
	}
}

func TestIsDisjoint(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(4, 5, 6)
	s3 := New(3, 4, 5)

	if !s1.IsDisjoint(s2) {
		t.Error("{1,2,3} and {4,5,6} should be disjoint")
	}
	if s1.IsDisjoint(s3) {
		t.Error("{1,2,3} and {3,4,5} should not be disjoint (share 3)")
	}

	// Empty set is disjoint with everything
	empty := New[int]()
	if !empty.IsDisjoint(s1) {
		t.Error("Empty set should be disjoint with any set")
	}
	if !s1.IsDisjoint(empty) {
		t.Error("Any set should be disjoint with empty set")
	}
}

func TestIsEqual(t *testing.T) {
	s1 := New(1, 2, 3)
	s2 := New(3, 2, 1) // Same elements, different order
	s3 := New(1, 2, 3, 4)

	if !s1.IsEqual(s2) {
		t.Error("Sets with same elements should be equal")
	}
	if s1.IsEqual(s3) {
		t.Error("Sets with different elements should not be equal")
	}

	// Empty sets are equal
	empty1 := New[int]()
	empty2 := New[int]()
	if !empty1.IsEqual(empty2) {
		t.Error("Two empty sets should be equal")
	}

	// Empty vs non-empty
	if s1.IsEqual(empty1) {
		t.Error("Empty and non-empty should not be equal")
	}

	// Nil comparison
	if !empty1.IsEqual(nil) {
		t.Error("Empty set should equal nil set")
	}
	if s1.IsEqual(nil) {
		t.Error("Non-empty set should not equal nil")
	}
}

func TestString(t *testing.T) {
	s := New(1, 2, 3)
	str := s.String()
	if str == "" {
		t.Error("String() should not return empty")
	}
	if len(str) < 3 {
		t.Error("String representation should be meaningful")
	}

	empty := New[int]()
	if empty.String() != "Set{}" {
		t.Errorf("Empty set String() = %s, want Set{}", empty.String())
	}
}

func TestFilter(t *testing.T) {
	s := New(1, 2, 3, 4, 5, 6)
	evens := s.Filter(func(n int) bool { return n%2 == 0 })

	if evens.Len() != 3 {
		t.Errorf("Filter should return 3 even numbers, got %d", evens.Len())
	}
	if !evens.ContainsAll(2, 4, 6) {
		t.Error("Filtered set should contain 2, 4, 6")
	}
}

func TestMap(t *testing.T) {
	s := New(1, 2, 3)
	doubled := Map(s, func(n int) int { return n * 2 })

	if doubled.Len() != 3 {
		t.Errorf("Map should return set with 3 elements, got %d", doubled.Len())
	}
	if !doubled.ContainsAll(2, 4, 6) {
		t.Error("Mapped set should contain doubled values")
	}
}

// Standalone function tests

func TestUnionFunc(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	s3 := New(3, 4)

	result := Union(s1, s2, s3)
	if result.Len() != 4 {
		t.Errorf("Union of 3 sets should have 4 elements, got %d", result.Len())
	}
	if !result.ContainsAll(1, 2, 3, 4) {
		t.Error("Union should contain all elements from all sets")
	}

	// Union with no sets
	empty := Union[int]()
	if !empty.IsEmpty() {
		t.Error("Union with no args should return empty set")
	}
}

func TestIntersectionFunc(t *testing.T) {
	s1 := New(1, 2, 3, 4)
	s2 := New(2, 3, 4, 5)
	s3 := New(3, 4, 5, 6)

	result := Intersection(s1, s2, s3)
	if result.Len() != 2 {
		t.Errorf("Intersection of 3 sets should have 2 elements, got %d", result.Len())
	}
	if !result.ContainsAll(3, 4) {
		t.Error("Intersection should contain common elements from all sets")
	}

	// Intersection with no sets
	empty := Intersection[int]()
	if !empty.IsEmpty() {
		t.Error("Intersection with no args should return empty set")
	}
}

func TestDifferenceFunc(t *testing.T) {
	s1 := New(1, 2, 3, 4, 5)
	s2 := New(2, 4)
	s3 := New(5)

	result := Difference(s1, s2, s3)
	if result.Len() != 2 {
		t.Errorf("Difference should have 2 elements, got %d", result.Len())
	}
	if !result.ContainsAll(1, 3) {
		t.Error("Difference should contain elements only in first set")
	}
}

func TestContainsFunc(t *testing.T) {
	m := map[int]struct{}{1: {}, 2: {}, 3: {}}
	if !Contains(m, 1) {
		t.Error("Contains should return true for existing element")
	}
	if Contains(m, 4) {
		t.Error("Contains should return false for non-existing element")
	}
}

func TestToSliceFunc(t *testing.T) {
	m := map[string]struct{}{"a": {}, "b": {}, "c": {}}
	slice := ToSlice(m)
	if len(slice) != 3 {
		t.Errorf("ToSlice should return 3 elements, got %d", len(slice))
	}

	// Check all elements present
	elements := make(map[string]bool)
	for _, e := range slice {
		elements[e] = true
	}
	if !elements["a"] || !elements["b"] || !elements["c"] {
		t.Error("ToSlice should contain all keys")
	}

	// Empty map
	empty := ToSlice(map[int]struct{}{})
	if len(empty) != 0 {
		t.Error("ToSlice of empty map should return empty slice")
	}
}

func TestIsSubsetFunc(t *testing.T) {
	a := map[int]struct{}{1: {}, 2: {}}
	b := map[int]struct{}{1: {}, 2: {}, 3: {}}
	c := map[int]struct{}{4: {}, 5: {}}

	if !IsSubset(a, b) {
		t.Error("{1,2} should be subset of {1,2,3}")
	}
	if IsSubset(b, a) {
		t.Error("{1,2,3} should not be subset of {1,2}")
	}
	if IsSubset(a, c) {
		t.Error("{1,2} should not be subset of {4,5}")
	}
}

func TestIsDisjointFunc(t *testing.T) {
	a := map[int]struct{}{1: {}, 2: {}}
	b := map[int]struct{}{3: {}, 4: {}}
	c := map[int]struct{}{2: {}, 3: {}}

	if !IsDisjoint(a, b) {
		t.Error("{1,2} and {3,4} should be disjoint")
	}
	if IsDisjoint(a, c) {
		t.Error("{1,2} and {2,3} should not be disjoint")
	}

	// Empty maps are disjoint
	empty := map[int]struct{}{}
	if !IsDisjoint(a, empty) {
		t.Error("Any set should be disjoint with empty set")
	}
}

// Integration test with slicez
func TestIntegrationWithSlicez(t *testing.T) {
	// Create set from slice
	slice := []int{1, 2, 2, 3, 3, 3}
	s := FromSlice(slice)

	// Convert back to slice
	result := s.ToSlice()

	// Verify unique elements
	if len(result) != 3 {
		t.Errorf("Integration: should have 3 unique elements, got %d", len(result))
	}

	// Test with slicez Set
	// (Assuming slicez.Set exists and creates map[T]struct{})
	m := map[int]struct{}{10: {}, 20: {}, 30: {}}
	s2 := FromMap(m)
	if s2.Len() != 3 {
		t.Error("FromMap should create correct set from map")
	}
}
