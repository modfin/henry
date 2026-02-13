# compare

> Comparison utilities and constraints for Go

The `compare` package provides comparison functions, constraints, and predicates for generic programming. It defines the `Ordered` constraint and offers utilities for comparing, sorting, and filtering.

## Quick Reference

**By Category:**
- [Core Comparisons](#core-comparisons) - Compare, Less, Greater, Equal
- [Range Operations](#range-operations) - Between, Clamp
- [Predicate Constructors](#predicate-constructors) - EqualOf, IsZero, NegateOf
- [Function Utilities](#function-utilities) - Negate, Identity, Ternary, Coalesce

## Installation

```bash
go get github.com/modfin/henry/compare
```

## Usage

```go
import "github.com/modfin/henry/compare"

// Compare values
result := compare.Compare(5, 10)  // -1 (5 < 10)
result = compare.Compare(10, 10)    // 0 (equal)
result = compare.Compare(10, 5)      // 1 (10 > 5)

// Check range
inRange := compare.Between(5, 1, 10)  // true

// Create predicates
isEven := func(n int) bool { return n%2 == 0 }
isOdd := compare.NegateOf(isEven)
```

## Core Comparisons

### Compare
Three-way comparison with NaN handling.

```go
// Standard comparison
compare.Compare(5, 10)   // -1
compare.Compare(10, 5)   // 1
compare.Compare(5, 5)    // 0

// With NaN (floats)
compare.Compare(math.NaN(), 5.0)      // -1 (NaN < everything)
compare.Compare(5.0, math.NaN())      // 1
twoNaNs := compare.Compare(math.NaN(), math.NaN()) // 0 (NaN == NaN for comparison)
```

### Less / LessOrEqual

```go
compare.Less(5, 10)          // true
compare.LessOrEqual(5, 5)    // true

// Use with sort
sorted := slicez.SortBy(nums, compare.Less[int])
```

### Greater / GreaterOrEqual

```go
compare.Greater(10, 5)          // true
compare.GreaterOrEqual(5, 5)    // true
```

### Equal

```go
compare.Equal(5, 5)    // true
compare.Equal(5, 3)    // false
```

## Range Operations

### Between
Check if value is in range.

```go
// Inclusive (default)
compare.Between(5, 1, 10)                    // true (1 <= 5 <= 10)
compare.Between(1, 1, 10)                    // true
compare.Between(10, 1, 10)                   // true

// Exclusive
compare.Between(5, 1, 10, compare.BetweenExclusive)    // true (1 < 5 < 10)
compare.Between(1, 1, 10, compare.BetweenExclusive)    // false

// Left-inclusive
compare.Between(1, 1, 10, compare.BetweenLeftInclusive)   // true (1 <= 1 < 10)
compare.Between(10, 1, 10, compare.BetweenLeftInclusive) // false

// Right-inclusive
compare.Between(10, 1, 10, compare.BetweenRightInclusive) // true (1 < 10 <= 10)
```

### Clamp
Constrain value to range.

```go
compare.Clamp(50, 0, 100)    // 50 (within range)
compare.Clamp(-10, 0, 100)   // 0 (clamped to min)
compare.Clamp(150, 0, 100)   // 100 (clamped to max)

// Practical use: constrain percentage
pct := compare.Clamp(userInput, 0, 100)
```

## Predicate Constructors

### EqualOf
Create equality predicate.

```go
isTwo := compare.EqualOf(2)
isTwo(2)    // true
isTwo(3)    // false

// Use with filter
twos := slicez.Filter(nums, compare.EqualOf(2))
```

### IsZero
Check for zero value.

```go
isZeroInt := compare.IsZero[int]()
isZeroInt(0)     // true
isZeroInt(5)     // false

emptyStrings := slicez.Filter(strings, compare.IsZero[string]())
```

### IsNotZero
Check for non-zero value.

```go
nonEmpty := slicez.Filter(strings, compare.IsNotZero[string]())
// Keeps only non-empty strings
```

### NegateOf
Negate a predicate.

```go
isEven := func(n int) bool { return n%2 == 0 }
isOdd := compare.NegateOf(isEven)
isOdd(3)     // true
isOdd(2)     // false
```

## Function Utilities

### Negate
Negate a comparison function.

```go
// Create descending sort
nums := []int{1, 2, 3, 4, 5}
descending := slicez.SortBy(nums, compare.Negate(compare.Less[int]))
// descending = []int{5, 4, 3, 2, 1}
```

### Identity
Identity function.

```go
// Use when no transformation needed
result := slicez.Map(nums, compare.Identity[int])
```

### Ternary
Conditional expression.

```go
// If-else in one line
sign := compare.Ternary(x >= 0, "positive", "negative")
max := compare.Ternary(a > b, a, b)
```

### Coalesce
First non-zero value.

```go
// Provide fallback values
name := compare.Coalesce(userName, nickname, "Anonymous")
// Returns first non-empty string

port := compare.Coalesce(envPort, configPort, 8080)
// Returns first non-zero port
```

## The Ordered Constraint

The `Ordered` constraint is defined in `ordered.go`:

```go
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}
```

Use it in your generic functions:

```go
func Max[T compare.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

## See Also

- [slicez](../slicez/) - Use comparison functions with SortBy, Filter, etc.
- [setz](../setz/) - Set operations using comparisons
