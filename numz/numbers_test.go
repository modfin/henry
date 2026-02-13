package numz

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	if Min(1, 2, 3, 4, -2, 3) != -2 {
		t.Log("Expected ", -2)
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	if Max(1, 2, 3, 4, -2, 3) != 4 {
		t.Log("Expected ", -2)
		t.Fail()
	}
}

func TestRange(t *testing.T) {
	r := Range[float64](5.5, 4.9, 4.85, 5.25, 5.05, 6.25, 5)
	if r != 1.4000000000000004 {
		t.Log("Expected 1.4, got", r)
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	if Sum(1, 2, 3, 4, -2, 3) != 11 {
		t.Log("Expected ", 11)
		t.Fail()
	}
}

func TestSum2(t *testing.T) {
	if Sum(1.0, 2.0, 3.1, 4.1, -2.0, 3.0) != 11.2 {
		t.Log("Expected ", 11)
		t.Fail()
	}
}
func TestMean(t *testing.T) {
	if Mean(1, 2, 3, 4, -2, 3) != 11.0/6.0 {
		t.Log("Expected ")
		t.Fail()
	}
}

func TestVDot(t *testing.T) {
	x := []int{4, 6, 8}
	y := []int{3, 2, 5}
	dot := VDot(x, y)
	if dot != 64 {
		t.Log("Expected 64, got", dot)
		t.Fail()
	}
}

func TestMAD(t *testing.T) {
	v := []int{3, 8, 10, 17, 24, 27}
	res := MAD(v...)
	if res != 7.833333333333334 {
		t.Log("Expected ", 1, "got", res)
		t.Fail()
	}
}

func TestMedian(t *testing.T) {
	m := Median(1)
	if m != 1 {
		t.Log("Expected ", 1, "got", m)
		t.Fail()
	}
	m = Median(1, 2)
	if m != 1.5 {
		t.Log("Expected ", 1.5, "got", m)
		t.Fail()
	}
	m = Median(1, 2, 3, 4, 5)
	if m != 3 {
		t.Log("Expected ", 3, "got", m)
		t.Fail()
	}
	m = Median(1, 2, 3, 4)
	if m != 2.5 {
		t.Log("Expected", 2.5, "got", m)
		t.Fail()
	}
}

func TestModes(t *testing.T) {
	m := Modes(1)
	if !reflect.DeepEqual(m, []int{1}) {
		t.Log("Expected [1]", "got", m)
		t.Fail()
	}

	m = Modes(1, 1)
	if !reflect.DeepEqual(m, []int{1}) {
		t.Log("Expected [1]", "got", m)
		t.Fail()
	}

	m = Modes(1, 1, 2)
	if !reflect.DeepEqual(m, []int{1}) {
		t.Log("Expected [1]", "got", m)
		t.Fail()
	}

	m = Modes(1, 1, 2, 3, 3)
	if !reflect.DeepEqual(m, []int{1, 3}) {
		t.Log("Expected [1, 3]", "got", m)
		t.Fail()
	}
	mm := Mode(1, 1, 2, 3, 3)
	if mm != 1 {
		t.Log("Expected 1", "got", mm)
		t.Fail()
	}
	mm = Mode(1, 1, 2, 2, 2, 3, 3, 3)
	if mm != 2 {
		t.Log("Expected 2", "got", mm)
		t.Fail()
	}
	mm = Mode(1, 1, 2, 2, 2, 3, 3, 3, 3)
	if mm != 3 {
		t.Log("Expected 3", "got", mm)
		t.Fail()
	}
}

func TestVariance(t *testing.T) {
	v := Var(1, 2, 3, 4, 5)
	if v != 2.5 {
		t.Log("Expected 2.5", "got", v)
		t.Fail()
	}
}
func TestStdDev(t *testing.T) {
	v := StdDev(1, 2, 3, 4, 5)
	if v != 1.5811388300841898 {
		t.Log("Expected 1.5811388", "got", v)
		t.Fail()
	}
}
func TestStdErr(t *testing.T) {
	v := StdErr(10, 20, 30, 40)
	if v != 6.454972243679028 {
		t.Log("Expected 6.455", "got", v)
		t.Fail()
	}
}

func TestSNR(t *testing.T) {
	v := SNR(3, 8, 10, 17, 24, 27)
	if v != 1.569101279812298 {
		t.Log("Expected 1.5691", "got", v)
		t.Fail()
	}
}

func TestSkew(t *testing.T) {
	v := Skew(3, 8, 10, 17, 24, 27)
	if v != 0.1165799592157562 {
		t.Log("Expected 0.1165799592157562", "got", v)
		t.Fail()
	}
}

func TestCorrelation(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := []int{2, 3, 4, 8, 6, 7, 8}
	corr := Corr(x, y)
	if corr != 0.8854377448471462 {
		t.Log("Expected 0.8854", "got", corr)
		t.Fail()
	}
}

func TestR2(t *testing.T) {
	x := []int{3, 8, 10, 17, 24, 27}
	y := []int{2, 8, 10, 13, 18, 20}
	r2 := R2(x, y)
	if r2 != 0.9728504570950011 {
		t.Log("Expected 0.9728504570950011", "got", r2)
		t.Fail()
	}
}

func TestLinReg(t *testing.T) {
	x := []int{4, 5, 6, 7, 10}
	y := []int{3, 8, 20, 30, 12}
	intercept, slope := LinReg(x, y)
	if intercept != 1.6415094339622642 || slope != 4.09433962264151 {
		t.Log("Expected 1.6415 + 4.0943 x, ", "got", intercept, "+", slope, "x")
		t.Fail()
	}
}

func TestFTest(t *testing.T) {
	x := []int{1, 2, 4, 5, 8}
	y := []int{5, 20, 40, 80, 100}
	test := FTest(x, y)
	if test != 0.004672897196261682 {
		t.Log("Expected 0.0047, ", "got", test)
		t.Fail()
	}
}

func TestZScore(t *testing.T) {
	x := 20
	y := []int{2, 8, 18, 20, 28}
	test := ZScore(x, y)
	if test != 0.4679865455802192 {
		fmt.Println(Mean(y...))
		fmt.Println(StdDev(y...))
		t.Log("Expected 0.46799, ", "got", test)
		t.Fail()
	}
}

func TestCovariance(t *testing.T) {
	x := []int{5, 12, 18, 23, 45}
	y := []int{2, 8, 18, 20, 28}
	cov := Cov(x, y)
	if cov != 116.88 {
		t.Log("Expected 116.88", "got", cov)
		t.Fail()
	}
}

func TestBitOR(t *testing.T) {
	x := []int{1, 2}
	res := BitOR(x)
	if res != 3 {
		t.Log("Expected 3", "got", res)
		t.Fail()
	}
}
func TestBitAND(t *testing.T) {
	x := []int{1, 3}
	res := BitAND(x)
	if res != 1 {
		t.Log("Expected 1", "got", res)
		t.Fail()
	}
}
func TestBitXOR(t *testing.T) {
	x := []int{1, 3}
	res := BitXOR(x)
	if res != 2 {
		t.Log("Expected 2", "got", res)
		t.Fail()
	}
}

func TestPercentile(t *testing.T) {
	x := []int{3, 8, 10, 17, 24, 27, 32, 12}
	res := Percentile(24, x...)
	if res != 0.625 {
		t.Log("Expected 1", "got", res)
		t.Fail()
	}
}

func TestGCD(t *testing.T) {
	res := GCD(8, 40, 100)
	if res != 4 {
		t.Log("Expected 4", "got", res)
		t.Fail()
	}
}
func TestLCM(t *testing.T) {
	res := LCM(10, 15, 20)
	if res != 60 {
		t.Log("Expected 60", "got", res)
		t.Fail()
	}
}

func TestVPow(t *testing.T) {
	vector := []float64{1, 2, 3}
	result := VPow(vector, 2)
	expected := []float64{1, 4, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("VPow() = %v, want %v", result, expected)
	}
}

func TestVMul(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	result := VMul(x, y)
	expected := []int{4, 10, 18}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("VMul() = %v, want %v", result, expected)
	}
	// Test with different lengths - should take min length
	x2 := []int{1, 2, 3, 4}
	y2 := []int{5, 6}
	result2 := VMul(x2, y2)
	if len(result2) != 2 {
		t.Errorf("VMul different lengths: expected length 2, got %d", len(result2))
	}
}

func TestVAdd(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	result := VAdd(x, y)
	expected := []int{5, 7, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("VAdd() = %v, want %v", result, expected)
	}
}

func TestVSub(t *testing.T) {
	x := []int{10, 20, 30}
	y := []int{1, 2, 3}
	result := VSub(x, y)
	expected := []int{9, 18, 27}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("VSub() = %v, want %v", result, expected)
	}
}

func TestNegate(t *testing.T) {
	if Negate(5) != -5 {
		t.Errorf("Negate(5) = %d, want -5", Negate(5))
	}
	if Negate(-3) != 3 {
		t.Errorf("Negate(-3) = %d, want 3", Negate(-3))
	}
	if Negate(0) != 0 {
		t.Errorf("Negate(0) = %d, want 0", Negate(0))
	}
}

func TestAbs(t *testing.T) {
	if Abs(5) != 5 {
		t.Errorf("Abs(5) = %d, want 5", Abs(5))
	}
	if Abs(-5) != 5 {
		t.Errorf("Abs(-5) = %d, want 5", Abs(-5))
	}
	if Abs(0) != 0 {
		t.Errorf("Abs(0) = %d, want 0", Abs(0))
	}
}

func TestCastFunctions(t *testing.T) {
	n := 42.5
	if CastFloat64(n) != 42.5 {
		t.Errorf("CastFloat64 failed")
	}
	if CastFloat32(n) != 42.5 {
		t.Errorf("CastFloat32 failed")
	}
	if CastInt(n) != 42 {
		t.Errorf("CastInt failed")
	}
	if CastInt8(100) != 100 {
		t.Errorf("CastInt8 failed")
	}
	if CastInt16(1000) != 1000 {
		t.Errorf("CastInt16 failed")
	}
	if CastInt32(100000) != 100000 {
		t.Errorf("CastInt32 failed")
	}
	if CastInt64(10000000000) != 10000000000 {
		t.Errorf("CastInt64 failed")
	}
	if CastUInt(42) != 42 {
		t.Errorf("CastUInt failed")
	}
	if CastUInt8(255) != 255 {
		t.Errorf("CastUInt8 failed")
	}
	if CastUInt16(1000) != 1000 {
		t.Errorf("CastUInt16 failed")
	}
	if CastUInt32(100000) != 100000 {
		t.Errorf("CastUInt32 failed")
	}
	if CastUInt64(10000000000) != 10000000000 {
		t.Errorf("CastUInt64 failed")
	}
	if CastByte(65) != 65 {
		t.Errorf("CastByte failed")
	}
}

func TestEmptyInput(t *testing.T) {
	// Test various functions with empty input
	emptySlice := []int{}

	// GCD with empty slice
	if GCD(emptySlice...) != 0 {
		t.Errorf("GCD(empty) = %d, want 0", GCD(emptySlice...))
	}

	// Mean with empty slice
	if Mean(emptySlice...) != 0 {
		t.Errorf("Mean(empty) = %v, want 0", Mean(emptySlice...))
	}

	// Min with empty slice
	if Min(emptySlice...) != 0 {
		t.Errorf("Min(empty) = %d, want 0", Min(emptySlice...))
	}

	// Max with empty slice
	if Max(emptySlice...) != 0 {
		t.Errorf("Max(empty) = %d, want 0", Max(emptySlice...))
	}
}

func TestGCD_Empty(t *testing.T) {
	// Test empty input
	var empty []int
	if GCD(empty...) != 0 {
		t.Errorf("GCD with empty input should return 0")
	}
	// Test single element
	if GCD(42) != 42 {
		t.Errorf("GCD(42) = %d, want 42", GCD(42))
	}
}

func TestLCM_Single(t *testing.T) {
	// Test with single element
	if LCM(10, 5) != 10 {
		t.Errorf("LCM(10, 5) = %d, want 10", LCM(10, 5))
	}
}

func TestRange_Empty(t *testing.T) {
	// Range with no args should return 0
	if Range[int]() != 0 {
		t.Errorf("Range() with no args should return 0")
	}
}

func TestSNR_EdgeCases(t *testing.T) {
	// SNR when all values are the same (std dev = 0) will be +Inf
	result := SNR(5, 5, 5)
	if !math.IsInf(result, 1) {
		t.Errorf("SNR with identical values should be +Inf, got %v", result)
	}
}

func TestSkew_Symmetric(t *testing.T) {
	// Symmetric distribution should have skew close to 0
	data := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	skew := Skew(data...)
	// This is approximately symmetric, skew should be close to 0
	if skew > 0.5 || skew < -0.5 {
		t.Errorf("Skew of symmetric distribution = %v, should be close to 0", skew)
	}
}

func TestPercentile_EdgeCases(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	// 1 is the minimum, should be at 0th percentile
	p := Percentile(1, data...)
	if p != 0 {
		t.Errorf("Percentile(1, min) = %v, want 0", p)
	}
	// 5 is greater than everything except itself
	p2 := Percentile(5, data...)
	if p2 != 0.8 {
		t.Errorf("Percentile(5) = %v, want 0.8", p2)
	}
}
