package slicez

import (
	"fmt"
	"github.com/modfin/go18exp/compare"
	"reflect"
	"strconv"
	"testing"
)

func TestDropLeft0(t *testing.T) {
	var ints []int
	var exp []int
	res := Drop(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}

func TestDropLeftAll(t *testing.T) {
	var ints = []int{1}
	var exp []int
	res := Drop(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}

func TestDropRightAll(t *testing.T) {
	var ints = []int{1}
	var exp []int
	res := DropRight(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}

func TestDropRight0(t *testing.T) {
	var ints []int
	var exp []int
	res := DropRight(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}

func TestDropLeft(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{2, 3}
	res := Drop(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}
func TestDropLeft2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{1, 2, 3}
	res := Drop(ints, 0)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}
func TestDropRight(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{1, 2}
	res := DropRight(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}
func TestDropRight2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{1, 2, 3}
	res := DropRight(ints, 0)
	if !reflect.DeepEqual(res, exp) {
		t.Logf("expected, %v to equal %v\n", res, exp)
		t.Fail()
	}
}

func TestDropRightWhile(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{1}
	res := DropRightWhile(ints, func(a int) bool {
		return a > 1
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestDropLeftWhile(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{3}
	res := DropWhile(ints, func(a int) bool {
		return a < 3
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestTakeRight(t *testing.T) {

	ints := []int{1, 2, 3}
	exp := []int{2, 3}
	res := TakeRight(ints, 2)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestTakeLeft(t *testing.T) {

	ints := []int{1, 2, 3}
	exp := []int{1}
	res := Take(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestFilter(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{1, 3}
	res := Filter(ints, func(a int) bool {
		return a%2 == 1
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestReject(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := []int{1, 3}
	res := Reject(ints, func(a int) bool {
		return a == 2
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestSome(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := true
	res := Some(ints, func(a int) bool {
		return a == 2
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestSome2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := false
	res := Some(ints, func(a int) bool {
		return a == 5
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNone(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := true
	res := None(ints, func(a int) bool {
		return a == 5
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNone2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := false
	res := None(ints, func(a int) bool {
		return a == 2
	})

	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestEvery(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := true
	res := Every(ints, func(a int) bool {
		return a < 5
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestEvery2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := false
	res := Every(ints, func(a int) bool {
		return a < 3
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNth(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := 2
	res, _ := Nth(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNth2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := 3
	res, _ := Nth(ints, -1)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNth3(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := 0
	res, _ := Nth(ints, 10)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestPartition(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	ints := []int{1, 2, 3, 4}
	expEven := []int{2, 4}
	expOdd := []int{1, 3}
	even, odd := Partition(ints, isEven)
	if !reflect.DeepEqual(even, expEven) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", even, expEven)
	}
	if !reflect.DeepEqual(even, expEven) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", odd, expOdd)
	}
}

func TestSort(t *testing.T) {
	ints := []int{3, 2, 1}
	exp := []int{1, 2, 3}
	res := Sort(ints)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}
func TestSortFunc(t *testing.T) {
	ints := []int{3, 2, 1}
	exp := []int{1, 2, 3}
	res := SortFunc(ints, func(a, b int) bool {
		return a < b
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestCompact(t *testing.T) {
	ints := []int{3, 2, 2, 1, 1, 1, 1, 1, 3, 3, 4, 5}
	exp := []int{3, 2, 1, 3, 4, 5}
	res := CompactFunc(ints, func(a, b int) bool {
		return a == b
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestCompact2(t *testing.T) {
	str := []byte("Remove   a  lot    of    white   spaces    !!")
	exp := "Remove a lot of white spaces !!"
	res := CompactFunc(str, func(a, b byte) bool {
		return a == b && a == byte(' ')
	})
	resStr := string(res)
	if resStr != exp {
		t.Fail()
		t.Logf("expected, \"%v\" to equal \"%v\"\n", resStr, exp)
	}
}

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

	res := UniqBy[int, int](compare.EqualBy[int], a)
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

func TestZip(t *testing.T) {
	as := []int{1, 2, 3, 4}
	bs := []string{"a", "b", "c"}
	exp := []string{"1a", "2b", "3c"}
	res := Zip(as, bs, func(a int, b string) string {
		return fmt.Sprintf("%d%s", a, b)
	})
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}

func TestUnzip(t *testing.T) {
	cs := []string{"1a", "2b", "3c"}

	expA := []int{1, 2, 3}
	expB := []string{"a", "b", "c"}

	a, b := Unzip(cs, func(c string) (int, string) {
		a, _ := strconv.Atoi(string(c[0]))
		return a, string(c[1])
	})
	if !reflect.DeepEqual(expA, a) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", expA, a)
	}
	if !reflect.DeepEqual(expB, b) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", expB, b)
	}
}
