// Package compare provides comparison utilities and constraints for generic types.
//
// The package defines the Ordered constraint for types that support ordering operations
// (integers, floats, strings) and provides comparison functions that work with any Ordered type.
//
// Core comparison functions:
//   - Compare: Three-way comparison returning -1, 0, or +1
//   - Less/LessOrEqual/Greater/GreaterOrEqual: Standard comparisons
//   - Equal: Equality comparison for comparable types
//
// Predicate constructors (useful for filtering/matching):
//   - EqualOf: Creates a predicate checking equality to a specific value
//   - IsZero/IsNotZero: Check for zero values
//   - NegateOf: Negates a predicate function
//
// Utility functions:
//   - Between: Check if a value is within a range (with inclusive/exclusive modes)
//   - Clamp: Constrain a value to a range
//   - Ternary: Conditional expression equivalent
//   - Coalesce: Return first non-zero value from variadic arguments
//   - Negate: Negate a comparison function
//   - Identity: Identity function for comparable types
//
// The Ordered constraint is defined in ordered.go and includes all ordered types
// (integers, unsigned integers, floats, and strings).
package compare

// Compare performs a three-way comparison between two Ordered values.
// Returns -1 if a < b, 0 if a == b, and +1 if a > b.
//
// Handles NaN values correctly for floats: NaN is considered less than any
// non-NaN value, and two NaN values are considered equal to each other.
//
// Example:
//
//	compare.Compare(5, 10)   // Returns -1
//	compare.Compare(10, 10)  // Returns 0
//	compare.Compare(10, 5)   // Returns +1
//
//	// With floats including NaN
//	compare.Compare(math.NaN(), 5.0)  // Returns -1 (NaN < everything)
//	compare.Compare(math.NaN(), math.NaN()) // Returns 0 (NaN == NaN for comparison)
func Compare[N Ordered](a N, b N) int {
	// Handle NaN cases - NaN is not comparable and should not equal anything
	if isNaN(a) && isNaN(b) {
		return 0 // NaN == NaN for comparison purposes
	}
	if isNaN(a) {
		return -1 // NaN is less than everything (except NaN)
	}
	if isNaN(b) {
		return +1 // everything is greater than NaN
	}
	switch {
	case a < b:
		return -1
	case a > b:
		return +1
	default:
		return 0
	}
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T Ordered](x T) bool {
	return x != x
}

// Identity returns its argument unchanged.
// Useful as a default mapper function or when a function type is required
// but no transformation is needed.
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	evens := slicez.Filter(nums, func(n int) bool { return n%2 == 0 })
//	identity := slicez.Map(evens, compare.Identity[int])
//	// identity = []int{2, 4}
func Identity[N comparable](n N) N {
	return n
}

// Equal returns true if a == b.
// Works with any comparable type. Often used as a predicate function.
//
// Example:
//
//	slicez.Filter([]int{1, 2, 3, 2, 4}, func(n int) bool {
//	    return compare.Equal(n, 2)
//	})
//	// Returns []int{2, 2}
func Equal[N comparable](a, b N) bool {
	return a == b
}

// Less returns true if a < b.
// Works with any Ordered type (integers, floats, strings).
//
// Example:
//
//	sorted := slicez.SortBy([]int{3, 1, 4, 1, 5}, compare.Less[int])
//	// sorted = []int{1, 1, 3, 4, 5}
func Less[E Ordered](a, b E) bool {
	return a < b
}

// LessOrEqual returns true if a <= b.
// Works with any Ordered type (integers, floats, strings).
//
// Example:
//
//	compare.LessOrEqual(5, 10)  // true
//	compare.LessOrEqual(5, 5)   // true
//	compare.LessOrEqual(5, 3)   // false
func LessOrEqual[E Ordered](a, b E) bool {
	return a <= b
}

// Negate returns a function that negates the result of the given comparison function.
// Useful for reversing sort orders or inverting predicates.
//
// Example:
//
//	// Create a "greater than" comparison from "less than"
//	greater := compare.Negate(compare.Less[int])
//	greater(10, 5) // true (equivalent to 10 > 5)
//
//	// Reverse sort order
//	nums := []int{1, 2, 3, 4, 5}
//	descending := slicez.SortBy(nums, compare.Negate(compare.Less[int]))
//	// descending = []int{5, 4, 3, 2, 1}
func Negate[A any](f func(a, b A) bool) func(A, A) bool {
	return func(a, b A) bool {
		return !f(a, b)
	}
}

// EqualOf returns a predicate function that checks if its argument equals the needle value.
// Useful for creating equality predicates for filtering or searching.
//
// Example:
//
//	isTwo := compare.EqualOf(2)
//	isTwo(2) // true
//	isTwo(3) // false
//
//	// Use with slicez functions
//	nums := []int{1, 2, 3, 2, 4, 2}
//	filtered := slicez.Filter(nums, compare.EqualOf(2))
//	// filtered = []int{2, 2, 2}
func EqualOf[N comparable](needle N) func(b N) bool {
	return func(b N) bool {
		return needle == b
	}
}

// IsZero returns a predicate function that checks if its argument equals the zero value.
// The zero value is type-dependent (0 for numbers, "" for strings, nil for pointers, etc.).
//
// Example:
//
//	filterZeros := compare.IsZero[int]()
//	filterZeros(0) // true
//	filterZeros(5) // false
//
//	strings := []string{"hello", "", "world", ""}
//	empty := slicez.Filter(strings, compare.IsZero[string]())
//	// empty = []string{"", ""}
func IsZero[N comparable]() func(b N) bool {
	var n N
	return EqualOf(n)
}

// IsNotZero returns a predicate function that checks if its argument does NOT equal the zero value.
// Complement of IsZero. Useful for filtering out empty/null values.
//
// Example:
//
//	filterNonZero := compare.IsNotZero[int]()
//	filterNonZero(0) // false
//	filterNonZero(5) // true
//
//	strings := []string{"hello", "", "world", ""}
//	nonEmpty := slicez.Filter(strings, compare.IsNotZero[string]())
//	// nonEmpty = []string{"hello", "world"}
func IsNotZero[N comparable]() func(b N) bool {
	return NegateOf(IsZero[N]())
}

// NegateOf returns a function that negates the result of the given predicate function.
// Complement of Negate but works with single-argument predicates instead of comparisons.
//
// Example:
//
//	isEven := func(n int) bool { return n%2 == 0 }
//	isOdd := compare.NegateOf(isEven)
//	isOdd(3) // true
//	isOdd(2) // false
func NegateOf[A any](f func(A) bool) func(A) bool {
	return func(a A) bool {
		return !f(a)
	}
}

// Ternary returns ifTrue if boolean is true, otherwise returns ifFalse.
// Equivalent to the ternary operator (cond ? a : b) found in many languages.
// All arguments are evaluated (not short-circuited).
//
// Example:
//
//	sign := compare.Ternary(x >= 0, "positive", "negative")
//	// Returns "positive" if x >= 0, otherwise "negative"
//
//	max := compare.Ternary(a > b, a, b)
//	// Returns the larger of a and b
func Ternary[A any](boolean bool, ifTrue A, ifFalse A) A {
	if boolean {
		return ifTrue
	}
	return ifFalse
}

// Coalesce returns the first non-zero value from its arguments.
// Returns the zero value if all arguments are zero.
// Useful for providing fallback/default values.
//
// Example:
//
//	// Use first non-empty string
//	name := compare.Coalesce(userName, nickName, "Anonymous")
//	// Returns userName if set, otherwise nickName, otherwise "Anonymous"
//
//	// Use first non-zero number
//	port := compare.Coalesce(configPort, envPort, 8080)
//	// Returns first non-zero port number, or 8080 as default
func Coalesce[T comparable](vals ...T) (t T) {
	zero := IsZero[T]()
	for _, val := range vals {
		if !zero(val) {
			return val
		}
	}
	return t
}

// Greater returns true if a > b
func Greater[E Ordered](a, b E) bool {
	return a > b
}

// GreaterOrEqual returns true if a >= b
func GreaterOrEqual[E Ordered](a, b E) bool {
	return a >= b
}

// BetweenMode specifies whether bounds are inclusive or exclusive
// for the Between function.
type BetweenMode int

const (
	// BetweenInclusive includes both start and end bounds (start <= x <= end)
	BetweenInclusive BetweenMode = iota
	// BetweenExclusive excludes both start and end bounds (start < x < end)
	BetweenExclusive
	// BetweenLeftInclusive includes only start bound (start <= x < end)
	BetweenLeftInclusive
	// BetweenRightInclusive includes only end bound (start < x <= end)
	BetweenRightInclusive
)

// Between returns true if value is between start and end.
// By default, bounds are inclusive (start <= value <= end).
// Use BetweenMode options to control inclusivity.
//
// Examples:
//
//	Between(5, 1, 10)                    // true (1 <= 5 <= 10)
//	Between(1, 1, 10)                    // true (bounds inclusive)
//	Between(1, 1, 10, BetweenExclusive)  // false (1 is not > 1)
//	Between(10, 1, 10, BetweenRightInclusive) // true (10 <= 10)
func Between[E Ordered](value, start, end E, mode ...BetweenMode) bool {
	m := BetweenInclusive
	if len(mode) > 0 {
		m = mode[0]
	}

	switch m {
	case BetweenInclusive:
		return value >= start && value <= end
	case BetweenExclusive:
		return value > start && value < end
	case BetweenLeftInclusive:
		return value >= start && value < end
	case BetweenRightInclusive:
		return value > start && value <= end
	default:
		return value >= start && value <= end
	}
}

// Clamp constrains a value to be within the range [min, max].
// If value is less than min, returns min.
// If value is greater than max, returns max.
// Otherwise returns the value unchanged.
//
// Example:
//
//	Clamp(50, 0, 100)   // Returns 50 (within range)
//	Clamp(-10, 0, 100) // Returns 0 (clamped to min)
//	Clamp(150, 0, 100) // Returns 100 (clamped to max)
//	Clamp(5, 10, 20)   // Returns 10 (clamped to min)
func Clamp[E Ordered](value, min, max E) E {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
