package numbers

import (
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
func TestAverage(t *testing.T) {
	if Mean(1, 2, 3, 4, -2, 3) != 11.0/6.0 {
		t.Log("Expected ")
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
	v := Variance(1, 2, 3, 4, 5)
	if v != 2.5 {
		t.Log("Expected 2.5", "got", v)
	}
}
func TestStdDev(t *testing.T) {
	v := StdDev(1, 2, 3, 4, 5)
	if v != 1.5811388300841898 {
		t.Log("Expected 1.5811388", "got", v)
	}
}

func TestCorrelation(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := []int{2, 3, 4, 8, 6, 7, 8}
	corr := Correlation(x, y)
	if corr != 0.8854377448471462 {
		t.Log("Expected 0.8854", "got", corr)
	}
}
