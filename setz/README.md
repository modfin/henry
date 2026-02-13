# setz

> Generic set data structure for Go

The `setz` package provides a `Set[T]` type with O(1) operations for set membership, union, intersection, and more. Implemented as a type alias of `map[T]struct{}` for zero overhead.

## Quick Reference

**By Category:**
- [Construction](#construction) - New, FromSlice, FromMap
- [Basic Operations](#basic-operations) - Add, Remove, Contains, Pop, Clear
- [Set Operations](#set-operations) - Union, Intersection, Difference
- [Predicates](#predicates) - IsSubset, IsSuperset, IsDisjoint, IsEqual
- [Conversion](#conversion) - ToSlice, Copy, Filter, Map

## Installation

```bash
go get github.com/modfin/henry/setz
```

## Usage

```go
import "github.com/modfin/henry/setz"

// Create sets
s1 := setz.New(1, 2, 3, 4, 5)
s2 := setz.New(4, 5, 6, 7, 8)

// Set operations
union := s1.Union(s2)           // {1, 2, 3, 4, 5, 6, 7, 8}
intersection := s1.Intersection(s2)  // {4, 5}
difference := s1.Difference(s2)    // {1, 2, 3}

// Check membership
if s1.Contains(3) {
    fmt.Println("3 is in the set")
}
```

## Construction

### New
Create a set from elements.

```go
s := setz.New(1, 2, 3, 4, 5)
// s = Set{1, 2, 3, 4, 5}
```

### FromSlice
Create from a slice.

```go
nums := []int{1, 2, 2, 3, 3, 3}
s := setz.FromSlice(nums)
// s = Set{1, 2, 3} (duplicates removed)
```

### FromMap
Create from an existing map.

```go
m := map[string]struct{}{"a": {}, "b": {}}
s := setz.FromMap(m)
// s = Set{"a", "b"}
```

## Basic Operations

### Add
Add elements (mutates).

```go
s := setz.New(1, 2)
s.Add(3, 4, 5)
// s = Set{1, 2, 3, 4, 5}
```

### Remove
Remove elements (mutates).

```go
s := setz.New(1, 2, 3, 4, 5)
s.Remove(2, 4)
// s = Set{1, 3, 5}
```

### Contains
Check membership.

```go
s := setz.New(1, 2, 3)
s.Contains(2)    // true
s.Contains(5)    // false
```

### ContainsAll
Check if all elements are present.

```go
s := setz.New(1, 2, 3, 4, 5)
s.ContainsAll(2, 4)     // true
s.ContainsAll(2, 6)     // false (6 not in set)
```

### Pop
Remove and return arbitrary element.

```go
s := setz.New(1, 2, 3)
elem, ok := s.Pop()
// elem = 1 (or 2 or 3), ok = true
// s now has 2 elements

// Empty set
empty := setz.New[int]()
elem, ok := empty.Pop()
// elem = 0 (zero value), ok = false
```

### Clear
Remove all elements (mutates).

```go
s := setz.New(1, 2, 3)
s.Clear()
// s = Set{}
```

### Len
Get number of elements.

```go
s := setz.New(1, 2, 3)
n := s.Len()   // n = 3
```

### IsEmpty
Check if empty.

```go
s := setz.New[int]()
s.IsEmpty()   // true
```

## Set Operations

### Union
Combine two sets (new set).

```go
s1 := setz.New(1, 2, 3)
s2 := setz.New(2, 3, 4)
union := s1.Union(s2)
// union = Set{1, 2, 3, 4}
```

### Intersection
Common elements (new set).

```go
s1 := setz.New(1, 2, 3, 4)
s2 := setz.New(3, 4, 5, 6)
common := s1.Intersection(s2)
// common = Set{3, 4}
```

### Difference
Elements in first but not second (new set).

```go
s1 := setz.New(1, 2, 3, 4)
s2 := setz.New(3, 4, 5, 6)
diff := s1.Difference(s2)
// diff = Set{1, 2}
```

### SymmetricDifference
Elements in exactly one set (XOR).

```go
s1 := setz.New(1, 2, 3)
s2 := setz.New(2, 3, 4)
xor := s1.SymmetricDifference(s2)
// xor = Set{1, 4}
```

## Predicates

### IsSubset
Check if all elements are in another set.

```go
small := setz.New(1, 2)
large := setz.New(1, 2, 3, 4)
small.IsSubset(large)    // true
large.IsSubset(small)    // false
```

### IsSuperset
Check if contains all elements of another set.

```go
large := setz.New(1, 2, 3, 4)
small := setz.New(1, 2)
large.IsSuperset(small)  // true
```

### IsProperSubset
Subset but not equal.

```go
a := setz.New(1, 2)
b := setz.New(1, 2, 3)
a.IsProperSubset(b)   // true
b.IsProperSubset(a)   // false
```

### IsProperSuperset
Superset but not equal.

```go
a := setz.New(1, 2, 3)
b := setz.New(1, 2)
a.IsProperSuperset(b)  // true
```

### IsDisjoint
No common elements.

```go
s1 := setz.New(1, 2, 3)
s2 := setz.New(4, 5, 6)
s1.IsDisjoint(s2)   // true

s3 := setz.New(3, 4, 5)
s1.IsDisjoint(s3)   // false (3 in common)
```

### IsEqual
Same elements.

```go
s1 := setz.New(1, 2, 3)
s2 := setz.New(3, 2, 1)  // Order doesn't matter
s1.IsEqual(s2)   // true
```

## Conversion

### ToSlice
Convert to slice (order not guaranteed).

```go
s := setz.New(1, 2, 3)
slice := s.ToSlice()
// slice = []int{1, 2, 3} (or any order)
```

### Copy
Create a copy.

```go
s1 := setz.New(1, 2, 3)
s2 := s1.Copy()
// s2 is independent copy
```

### Filter
Filter elements (new set).

```go
s := setz.New(1, 2, 3, 4, 5, 6)
evens := s.Filter(func(n int) bool {
    return n%2 == 0
})
// evens = Set{2, 4, 6}
```

### Map
Transform elements (new set).

```go
s := setz.New(1, 2, 3)
doubled := setz.Map(s, func(n int) int {
    return n * 2
})
// doubled = Set{2, 4, 6}
```

## Standalone Functions

### Union (variadic)
Union of multiple sets.

```go
s1 := setz.New(1, 2)
s2 := setz.New(2, 3)
s3 := setz.New(3, 4)
combined := setz.Union(s1, s2, s3)
// combined = Set{1, 2, 3, 4}
```

### Intersection (variadic)
Intersection of multiple sets.

```go
s1 := setz.New(1, 2, 3, 4)
s2 := setz.New(2, 3, 4, 5)
s3 := setz.New(3, 4, 5, 6)
common := setz.Intersection(s1, s2, s3)
// common = Set{3, 4}
```

### Difference (variadic)
Difference with multiple sets.

```go
s1 := setz.New(1, 2, 3, 4, 5)
s2 := setz.New(2, 3)
s3 := setz.New(4)
result := setz.Difference(s1, s2, s3)
// result = Set{1, 5}
```

### Contains (standalone)
Check membership (functional style).

```go
s := setz.New(1, 2, 3)
setz.Contains(s, 2)   // true
```

### ToSlice (standalone)
Convert to slice.

```go
s := setz.New(1, 2, 3)
slice := setz.ToSlice(s)
```

### IsSubset (standalone)
Check subset relation.

```go
small := setz.New(1, 2)
large := setz.New(1, 2, 3)
setz.IsSubset(small, large)   // true
```

### IsDisjoint (standalone)
Check disjointness.

```go
s1 := setz.New(1, 2)
s2 := setz.New(3, 4)
setz.IsDisjoint(s1, s2)   // true
```

## String Representation

```go
s := setz.New(1, 2, 3)
fmt.Println(s.String())
// Output: Set[1 2 3] (or any order)

empty := setz.New[int]()
fmt.Println(empty.String())
// Output: Set{}
```

## Performance

- **O(1)** operations: Add, Remove, Contains, Len
- **O(n)** operations: Union, Intersection, ToSlice
- **Zero overhead**: `Set[T]` is just `map[T]struct{}`
- **Memory efficient**: Uses empty struct (0 bytes) for values

## Comparison with slicez

**Use setz when:**
- You need O(1) membership testing
- Doing set operations (union, intersection)
- Deduplication is the primary concern

**Use slicez set functions when:**
- Order matters
- You need indexed access
- Memory is constrained (sets use 2x memory)

```go
// setz version - O(1) lookup
s := setz.New(1, 2, 3, 4, 5)
exists := s.Contains(3)

// slicez version - O(n) lookup
nums := []int{1, 2, 3, 4, 5}
exists := slicez.Contains(nums, 3)
```

## See Also

- [slicez](../slicez/) - Set operations on slices (Uniq, Union, Intersection, etc.)
- [mapz](../mapz/) - Underlying map operations
