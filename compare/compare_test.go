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
	// With integers
	result := Coalesce(0, 0, 42, 100)
	if result != 42 {
		t.Errorf("Expected first non-zero value 42, got %v", result)
	}

	// All zeros
	result = Coalesce(0, 0, 0)
	if result != 0 {
		t.Errorf("Expected zero value 0, got %v", result)
	}

	// With strings
	resultStr := Coalesce("", "", "hello", "world")
	if resultStr != "hello" {
		t.Errorf("Expected first non-empty string 'hello', got %v", resultStr)
	}

	// All empty strings
	resultStr = Coalesce("", "", "")
	if resultStr != "" {
		t.Errorf("Expected empty string, got %v", resultStr)
	}

	// First is non-zero
	result = Coalesce(1, 2, 3)
	if result != 1 {
		t.Errorf("Expected first value 1, got %v", result)
	}
}
