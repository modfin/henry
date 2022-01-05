package slicez

import (
	"fmt"
	"github.com/crholm/go18exp/compare"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []string{"1", "2", "3"}

	res := Map(ints, func(i int) string {
		return fmt.Sprintf("%d", i)
	})
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v", res, exp)
		t.Fail()
	}

}

//func TestMap2(t *testing.T) {
//	ints := []int{1, 2, 3}
//	exp := []float64{1.0, 2.0, 3.0}
//
//	res := Map(ints, numbers.MapFloat64[int])
//	if !reflect.DeepEqual(res, exp) {
//		t.Logf("expected, %v to equal %v", res, exp)
//		t.Fail()
//	}
//
//}

func TestFoldLeft(t *testing.T) {

	ints := []int{1, 2, 3}
	exp := "123"
	res := Fold(ints, func(acc string, i int) string {
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
	res := FoldRight(ints, func(acc string, i int) string {
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

	res := FlatMap(ints, func(i int) []string {
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

	m := KeyBy(a, func(a int) string {
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

	m := GroupBy(a, func(a int) string {
		return fmt.Sprintf("%d", a%2)
	})
	if !reflect.DeepEqual(exp, m) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, m)
	}
}

func TestUniq(t *testing.T) {
	a := []int{1, 2, 3, 3, 3, 4, 5, 6, 6, 6, 6}
	exp := []int{1, 2, 3, 4, 5, 6}

	res := Uniq[int, int](compare.EqualBy[int], a)
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}
func TestUnion(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3}
	c := []int{3, 4, 5}
	d := []int{4, 5, -1}
	exp := []int{1, 2, 3, 4, 5, -1}

	res := Union[int, int](compare.EqualBy[int], a, b, c, d)
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2}
	c := []int{3, 4, 5, 2}
	exp := []int{2, 3}

	res := Intersection[int, int](compare.EqualBy[int], a, b, c)
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}

func TestDifference(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2}
	c := []int{3, 4, 5, 2}
	exp := []int{1, 4, 5}

	res := Difference[int, int](compare.EqualBy[int], a, b, c)
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}
func TestComplement(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2, 5, 5, 6, 1}
	exp := []int{5, 6}

	res := Complement[int, int](compare.EqualBy[int], a, b)
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}
