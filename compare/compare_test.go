package compare

import (
	"math"
	"reflect"
	"testing"
)

func TestTernary(t *testing.T) {
	type args struct {
		boolean bool
		ifTrue  string
		ifFalse string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "true",
		args: args{
			true,
			"true",
			"false",
		},
		want: "true",
	},
		{
			name: "false",
			args: args{
				false,
				"true",
				"false",
			},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.boolean, tt.args.ifTrue, tt.args.ifFalse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	// Integers
	if Compare(5, 3) != 1 {
		t.Error("Expected 5 > 3")
	}
	if Compare(3, 5) != -1 {
		t.Error("Expected 3 < 5")
	}
	if Compare(5, 5) != 0 {
		t.Error("Expected 5 == 5")
	}

	// Floats
	if Compare(5.5, 3.3) != 1 {
		t.Error("Expected 5.5 > 3.3")
	}
	if Compare(3.3, 5.5) != -1 {
		t.Error("Expected 3.3 < 5.5")
	}
	if Compare(5.5, 5.5) != 0 {
		t.Error("Expected 5.5 == 5.5")
	}

	// Strings
	if Compare("b", "a") != 1 {
		t.Error("Expected 'b' > 'a'")
	}
	if Compare("a", "b") != -1 {
		t.Error("Expected 'a' < 'b'")
	}
	if Compare("a", "a") != 0 {
		t.Error("Expected 'a' == 'a'")
	}

	// NaN handling
	nan := math.NaN()
	// NaN compared to regular number should return -1 (NaN is less)
	if Compare(nan, 1.0) != -1 {
		t.Errorf("Expected NaN < 1.0, got %d", Compare(nan, 1.0))
	}
	// Regular number compared to NaN should return +1 (number is greater)
	if Compare(1.0, nan) != +1 {
		t.Errorf("Expected 1.0 > NaN, got %d", Compare(1.0, nan))
	}
	// NaN compared to NaN should return 0
	if Compare(nan, nan) != 0 {
		t.Errorf("Expected NaN == NaN for comparison, got %d", Compare(nan, nan))
	}
}

func TestIsNaN(t *testing.T) {
	// Test with NaN
	if !isNaN(math.NaN()) {
		t.Error("Expected isNaN(NaN) to return true")
	}
	// Test with regular float
	if isNaN(1.0) {
		t.Error("Expected isNaN(1.0) to return false")
	}
	// Test with integer (should always return false)
	if isNaN(1) {
		t.Error("Expected isNaN(1) to return false")
	}
	// Test with string (should always return false)
	if isNaN("hello") {
		t.Error("Expected isNaN(\"hello\") to return false")
	}
}

func TestIdentity(t *testing.T) {
	if Identity(42) != 42 {
		t.Error("Expected Identity to return same value")
	}
	if Identity("hello") != "hello" {
		t.Error("Expected Identity to return same string")
	}
}

func TestEqual(t *testing.T) {
	if !Equal(42, 42) {
		t.Error("Expected 42 to equal 42")
	}
	if Equal(42, 43) {
		t.Error("Expected 42 not to equal 43")
	}
	if !Equal("hello", "hello") {
		t.Error("Expected 'hello' to equal 'hello'")
	}
	if Equal("hello", "world") {
		t.Error("Expected 'hello' not to equal 'world'")
	}
}

func TestLess(t *testing.T) {
	if !Less(3, 5) {
		t.Error("Expected 3 < 5")
	}
	if Less(5, 3) {
		t.Error("Expected not 5 < 3")
	}
	if Less(5, 5) {
		t.Error("Expected not 5 < 5")
	}

	if !Less("a", "b") {
		t.Error("Expected 'a' < 'b'")
	}
	if Less("b", "a") {
		t.Error("Expected not 'b' < 'a'")
	}
}

func TestLessOrEqual(t *testing.T) {
	if !LessOrEqual(3, 5) {
		t.Error("Expected 3 <= 5")
	}
	if !LessOrEqual(5, 5) {
		t.Error("Expected 5 <= 5")
	}
	if LessOrEqual(5, 3) {
		t.Error("Expected not 5 <= 3")
	}

	if !LessOrEqual("a", "b") {
		t.Error("Expected 'a' <= 'b'")
	}
	if !LessOrEqual("b", "b") {
		t.Error("Expected 'b' <= 'b'")
	}
}

func TestNegate(t *testing.T) {
	lessThan := func(a, b int) bool { return a < b }
	greaterOrEqual := Negate(lessThan)

	if !greaterOrEqual(5, 3) {
		t.Error("Expected 5 >= 3")
	}
	if !greaterOrEqual(5, 5) {
		t.Error("Expected 5 >= 5")
	}
	if greaterOrEqual(3, 5) {
		t.Error("Expected not 3 >= 5")
	}
}

func TestEqualOf(t *testing.T) {
	is42 := EqualOf(42)
	if !is42(42) {
		t.Error("Expected is42(42) to be true")
	}
	if is42(43) {
		t.Error("Expected is42(43) to be false")
	}

	isHello := EqualOf("hello")
	if !isHello("hello") {
		t.Error("Expected isHello('hello') to be true")
	}
	if isHello("world") {
		t.Error("Expected isHello('world') to be false")
	}
}

func TestIsZero(t *testing.T) {
	isZeroInt := IsZero[int]()
	if !isZeroInt(0) {
		t.Error("Expected isZeroInt(0) to be true")
	}
	if isZeroInt(42) {
		t.Error("Expected isZeroInt(42) to be false")
	}

	isZeroString := IsZero[string]()
	if !isZeroString("") {
		t.Error("Expected isZeroString('') to be true")
	}
	if isZeroString("hello") {
		t.Error("Expected isZeroString('hello') to be false")
	}
}

func TestIsNotZero(t *testing.T) {
	isNotZeroInt := IsNotZero[int]()
	if !isNotZeroInt(42) {
		t.Error("Expected isNotZeroInt(42) to be true")
	}
	if isNotZeroInt(0) {
		t.Error("Expected isNotZeroInt(0) to be false")
	}

	isNotZeroString := IsNotZero[string]()
	if !isNotZeroString("hello") {
		t.Error("Expected isNotZeroString('hello') to be true")
	}
	if isNotZeroString("") {
		t.Error("Expected isNotZeroString('') to be false")
	}
}

func TestNegateOf(t *testing.T) {
	isPositive := func(n int) bool { return n > 0 }
	isNotPositive := NegateOf(isPositive)

	if !isNotPositive(0) {
		t.Error("Expected isNotPositive(0) to be true")
	}
	if !isNotPositive(-5) {
		t.Error("Expected isNotPositive(-5) to be true")
	}
	if isNotPositive(5) {
		t.Error("Expected isNotPositive(5) to be false")
	}
}

func TestCoalesce(t *testing.T) {
	// Test with integers
	result := Coalesce(0, 0, 3, 4)
	if result != 3 {
		t.Errorf("Coalesce(0, 0, 3, 4) = %v, want 3", result)
	}

	// Test with first non-zero
	result = Coalesce(1, 2, 3)
	if result != 1 {
		t.Errorf("Coalesce(1, 2, 3) = %v, want 1", result)
	}

	// Test with strings
	strResult := Coalesce("", "", "third")
	if strResult != "third" {
		t.Errorf("Coalesce(\"\", \"\", \"third\") = %v, want \"third\"", strResult)
	}

	// Test with all zero values
	zeroResult := Coalesce(0, 0, 0)
	if zeroResult != 0 {
		t.Errorf("Coalesce(0, 0, 0) = %v, want 0", zeroResult)
	}
}

func TestGreater(t *testing.T) {
	if !Greater(5, 3) {
		t.Error("Expected Greater(5, 3) to be true")
	}
	if Greater(3, 5) {
		t.Error("Expected Greater(3, 5) to be false")
	}
	if Greater(5, 5) {
		t.Error("Expected Greater(5, 5) to be false")
	}
}

func TestGreaterOrEqual(t *testing.T) {
	if !GreaterOrEqual(5, 3) {
		t.Error("Expected GreaterOrEqual(5, 3) to be true")
	}
	if !GreaterOrEqual(5, 5) {
		t.Error("Expected GreaterOrEqual(5, 5) to be true")
	}
	if GreaterOrEqual(3, 5) {
		t.Error("Expected GreaterOrEqual(3, 5) to be false")
	}
}

func TestBetween(t *testing.T) {
	// Inclusive (default)
	if !Between(5, 1, 10) {
		t.Error("Expected Between(5, 1, 10) to be true")
	}
	if !Between(1, 1, 10) {
		t.Error("Expected Between(1, 1, 10) to be true (inclusive lower)")
	}
	if !Between(10, 1, 10) {
		t.Error("Expected Between(10, 1, 10) to be true (inclusive upper)")
	}
	if Between(11, 1, 10) {
		t.Error("Expected Between(11, 1, 10) to be false")
	}

	// Exclusive
	if Between(1, 1, 10, BetweenExclusive) {
		t.Error("Expected Between(1, 1, 10, BetweenExclusive) to be false")
	}
	if Between(10, 1, 10, BetweenExclusive) {
		t.Error("Expected Between(10, 1, 10, BetweenExclusive) to be false")
	}
	if !Between(5, 1, 10, BetweenExclusive) {
		t.Error("Expected Between(5, 1, 10, BetweenExclusive) to be true")
	}

	// Left inclusive only
	if !Between(1, 1, 10, BetweenLeftInclusive) {
		t.Error("Expected Between(1, 1, 10, BetweenLeftInclusive) to be true")
	}
	if Between(10, 1, 10, BetweenLeftInclusive) {
		t.Error("Expected Between(10, 1, 10, BetweenLeftInclusive) to be false")
	}

	// Right inclusive only
	if Between(1, 1, 10, BetweenRightInclusive) {
		t.Error("Expected Between(1, 1, 10, BetweenRightInclusive) to be false")
	}
	if !Between(10, 1, 10, BetweenRightInclusive) {
		t.Error("Expected Between(10, 1, 10, BetweenRightInclusive) to be true")
	}

	// Strings
	if !Between("b", "a", "c") {
		t.Error("Expected Between(\"b\", \"a\", \"c\") to be true")
	}
}

func TestClamp(t *testing.T) {
	// Within range
	if Clamp(50, 0, 100) != 50 {
		t.Errorf("Clamp(50, 0, 100) = %d, want 50", Clamp(50, 0, 100))
	}

	// Below min
	if Clamp(-10, 0, 100) != 0 {
		t.Errorf("Clamp(-10, 0, 100) = %d, want 0", Clamp(-10, 0, 100))
	}

	// Above max
	if Clamp(150, 0, 100) != 100 {
		t.Errorf("Clamp(150, 0, 100) = %d, want 100", Clamp(150, 0, 100))
	}

	// At boundaries
	if Clamp(0, 0, 100) != 0 {
		t.Errorf("Clamp(0, 0, 100) = %d, want 0", Clamp(0, 0, 100))
	}
	if Clamp(100, 0, 100) != 100 {
		t.Errorf("Clamp(100, 0, 100) = %d, want 100", Clamp(100, 0, 100))
	}

	// With negative range
	if Clamp(-5, -10, 10) != -5 {
		t.Errorf("Clamp(-5, -10, 10) = %d, want -5", Clamp(-5, -10, 10))
	}

	// Floats
	if Clamp(1.5, 0.0, 1.0) != 1.0 {
		t.Errorf("Clamp(1.5, 0.0, 1.0) = %f, want 1.0", Clamp(1.5, 0.0, 1.0))
	}

	// Strings
	if Clamp("m", "a", "z") != "m" {
		t.Errorf("Clamp(\"m\", \"a\", \"z\") = %s, want \"m\"", Clamp("m", "a", "z"))
	}
	if Clamp("A", "a", "z") != "a" {
		t.Errorf("Clamp(\"A\", \"a\", \"z\") = %s, want \"a\"", Clamp("A", "a", "z"))
	}
}
