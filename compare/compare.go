package compare

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

// Identity is the identity function for something that is comparable
func Identity[N comparable](n N) N {
	return n
}

// Equal is the equality function for something that is comparable
func Equal[N comparable](a, b N) bool {
	return a == b
}

// Less returns true if a < b
func Less[E Ordered](a, b E) bool {
	return a < b
}

// LessOrEqual returns true if a <= b
func LessOrEqual[E Ordered](a, b E) bool {
	return a <= b
}

// Negate will return a function negating the result from other function. The use case for this is reversing a comparison
func Negate[A any](f func(a, b A) bool) func(A, A) bool {
	return func(a, b A) bool {
		return !f(a, b)
	}
}

// EqualOf returns a function that compares the input to the original input
func EqualOf[N comparable](needle N) func(b N) bool {
	return func(b N) bool {
		return needle == b
	}
}

// IsZero returns function that looks at the input that returns true if passed in argument is zero
func IsZero[N comparable]() func(b N) bool {
	var n N
	return EqualOf(n)
}

// IsNotZero returns function that looks at the input that returns true if passed in argument is *Not* zero
func IsNotZero[N comparable]() func(b N) bool {
	return NegateOf(IsZero[N]())
}

// NegateOf returns a function the negates the result of the original function passed in
func NegateOf[A any](f func(A) bool) func(A) bool {
	return func(a A) bool {
		return !f(a)
	}
}

// Ternary is equivalent to "expression ? a : b" ternary notation and returns ifTrue if true and ifFalse if false
func Ternary[A any](boolean bool, ifTrue A, ifFalse A) A {
	if boolean {
		return ifTrue
	}
	return ifFalse
}

// Coalesce returns the first of its arguments that is not equal to the zero value.
// If no argument is non-zero, it returns the zero value.
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
