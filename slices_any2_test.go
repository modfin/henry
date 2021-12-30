package henry

import (
	"fmt"
	"github.com/crholm/henry/numbers"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []string{"1", "2", "3"}

	res := Map(ints, func(_ int, i int) string {
		return fmt.Sprintf("%d", i)
	})
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v", res, exp)
		t.Fail()
	}

}

func TestMap2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []float64{1.0, 2.0, 3.0}

	res := Map(ints, numbers.MapFloat64[int])
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v", res, exp)
		t.Fail()
	}

}

func TestFoldLeft(t *testing.T) {

	ints := []int{1, 2, 3}
	exp := "123"
	res := FoldLeft(ints, func(_ int, acc string, i int) string {
		return fmt.Sprintf("%s%d", acc, i)
	}, "")
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v", res, exp)
		t.Fail()
	}
}

func TestFoldRight(t *testing.T) {

	ints := []int{1, 2, 3}
	exp := "321"
	res := FoldRight(ints, func(_ int, acc string, i int) string {
		return fmt.Sprintf("%s%d", acc, i)
	}, "")
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v", res, exp)
		t.Fail()
	}
}

func TestFlatMap(t *testing.T) {
	ints := []int{1, 2}
	exp := []string{"1", "1", "2", "2"}

	res := FlatMap(ints, func(_ int, i int) []string {
		str := fmt.Sprintf("%d", i)
		return []string{str, str}
	})
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v", res, exp)
		t.Fail()
	}
}

func TestKeyBy(t *testing.T) {
	a := []int{1, 2, 3}

	exp := map[string]int{
		"[1]": 1,
		"[2]": 2,
		"[3]": 3,
	}

	m := KeyBy(a, func(_ int, a int) string {
		return fmt.Sprintf("[%d]", a)
	})
	if !reflect.DeepEqual(exp, m) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, m)
	}
}

func TestGroupBy(t *testing.T) {
	a := []int{1, 2, 3}

	exp := map[string][]int{
		"0": {2},
		"1": {1, 3},
	}

	m := GroupBy(a, func(_ int, a int) string {
		return fmt.Sprintf("%d", a%2)
	})
	if !reflect.DeepEqual(exp, m) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, m)
	}
}
