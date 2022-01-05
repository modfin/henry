package numberz

import (
	"fmt"
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
