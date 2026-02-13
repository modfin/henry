package slicez

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/modfin/henry/compare"
)

func TestCut(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expLeft := []int{1, 2}
	expRight := []int{4, 5}

	left, right, _ := Cut(a, 3)
	if !Equal(left, expLeft) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", expLeft, left)
	}
	if !Equal(right, expRight) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", expRight, right)
	}
}

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

func TestLastIndex(t *testing.T) {
	ints := []int{1, 4, 3, 4, 5}
	i := LastIndex(ints, 4)

	if i != 3 {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", i, 3)
	}

	i = LastIndex(ints, 10)

	if i != -1 {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", i, -1)
	}
}

func TestFind(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	res, found := Find(ints, func(ii int) bool {
		return ii == 3
	})
	exp := 3
	if res != exp {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
	if !found {
		t.Fail()
		t.Logf("expected to be found")
	}

	res, found = Find(ints, func(ii int) bool {
		return ii == 10
	})
	exp = 0
	if res != exp {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
	if found {
		t.Fail()
		t.Logf("expected not to be found")
	}
}

func TestJoin(t *testing.T) {
	ints := [][]int{{1, 2, 3}, {4, 5}, {6}}
	exp := []int{1, 2, 3, 0, 0, 4, 5, 0, 0, 6}
	res := Join(ints, []int{0, 0})
	if !Equal(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestTakeRight(t *testing.T) {
	ints := []int{1, 2, 3}
	check := func(res []int, exp []int) {
		if !reflect.DeepEqual(res, exp) {
			t.Fail()
			t.Logf("expected, %v to equal %v\n", res, exp)
		}
	}
	t.Run("take zero", func(t *testing.T) {
		res := TakeRight(ints, 0)
		exp := []int{}
		check(res, exp)
	})
	t.Run("take last", func(t *testing.T) {
		res := TakeRight(ints, 1)
		exp := []int{3}
		check(res, exp)
	})
	t.Run("take two last", func(t *testing.T) {
		res := TakeRight(ints, 2)
		exp := []int{2, 3}
		check(res, exp)
	})
	t.Run("take all", func(t *testing.T) {
		res := TakeRight(ints, len(ints))
		check(res, ints)
	})

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
	res := SomeBy(ints, func(a int) bool {
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
	res := SomeBy(ints, func(a int) bool {
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
	res := None(ints, 5)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNone2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := false
	res := None(ints, 2)

	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestEvery(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := true
	res := EveryBy(ints, func(a int) bool {
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
	res := EveryBy(ints, func(a int) bool {
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
	res := Nth(ints, 1)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNth2(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := 3
	res := Nth(ints, -1)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
	res = Nth(ints, -4)
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestNth3(t *testing.T) {
	ints := []int{1, 2, 3}
	exp := 2
	res := Nth(ints, 4)
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
func TestSortBy(t *testing.T) {
	ints := []int{3, 2, 1}
	exp := []int{1, 2, 3}
	res := SortBy(ints, func(a, b int) bool {
		return a < b
	})
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}
}

func TestOrderBy(t *testing.T) {
	// Test sorting by extracted key (ascending by default)
	words := []string{"banana", "pie", "apple", "kiwi"}
	exp := []string{"pie", "kiwi", "apple", "banana"}
	res := OrderBy(words, func(s string) int { return len(s) })
	if !reflect.DeepEqual(res, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", res, exp)
	}

	// Test sorting with compare.Asc explicitly
	resAsc := OrderBy(words, func(s string) int { return len(s) }, compare.Asc[int])
	if !reflect.DeepEqual(resAsc, exp) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", resAsc, exp)
	}

	// Test sorting with compare.Desc
	expDesc := []string{"banana", "apple", "kiwi", "pie"}
	resDesc := OrderBy(words, func(s string) int { return len(s) }, compare.Desc[int])
	if !reflect.DeepEqual(resDesc, expDesc) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", resDesc, expDesc)
	}

	// Test sorting structs
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
	expPeople := []Person{{"Bob", 25}, {"Alice", 30}, {"Charlie", 35}}
	resPeople := OrderBy(people, func(p Person) int { return p.Age })
	if !reflect.DeepEqual(resPeople, expPeople) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", resPeople, expPeople)
	}

	// Test empty slice
	empty := []int{}
	resEmpty := OrderBy(empty, func(i int) int { return i })
	if len(resEmpty) != 0 {
		t.Fail()
		t.Logf("expected empty slice, got %v\n", resEmpty)
	}

	// Test nil slice (becomes empty slice, consistent with SortBy behavior)
	var nilSlice []int
	resNil := OrderBy(nilSlice, func(i int) int { return i })
	if resNil == nil || len(resNil) != 0 {
		t.Fail()
		t.Logf("expected empty slice, got %v\n", resNil)
	}

	// Test with floats
	floats := []float64{3.5, 1.2, 2.8, 0.5}
	expFloats := []float64{0.5, 1.2, 2.8, 3.5}
	resFloats := OrderBy(floats, func(f float64) float64 { return f })
	if !reflect.DeepEqual(resFloats, expFloats) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", resFloats, expFloats)
	}
}

func TestCompact(t *testing.T) {
	ints := []int{3, 2, 2, 1, 1, 1, 1, 1, 3, 3, 4, 5}
	exp := []int{3, 2, 1, 3, 4, 5}
	res := CompactBy(ints, func(a, b int) bool {
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
	res := CompactBy(str, func(a, b byte) bool {
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

func TestAssociate(t *testing.T) {
	type foo struct {
		baz string
		bar int
	}
	in := []foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}
	m := Associate(in, func(a foo) (key string, value int) {
		return a.baz, a.bar
	})

	exp := map[string]int{
		"apple":  1,
		"banana": 2,
	}

	if !reflect.DeepEqual(exp, m) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, m)
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

	res := UniqBy[int, int](a, compare.Identity[int])
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}
func TestUniq2(t *testing.T) {
	a := []int{1, 2, 3, 3, 3, 4, 5, 6, 6, 6, 6}
	exp := []int{1, 2, 3, 4, 5, 6}

	res := Uniq[int](a)
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

	res := UnionBy[int, int](compare.Identity[int], a, b, c, d)
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

	res := IntersectionBy[int, int](compare.Identity[int], a, b, c)
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

	res := DifferenceBy[int, int](compare.Identity[int], a, b, c)
	if !reflect.DeepEqual(exp, res) {
		t.Fail()
		t.Logf("expected, %v to equal %v\n", exp, res)
	}
}
func TestComplement(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2, 5, 5, 6, 1}
	exp := []int{5, 6}

	res := ComplementBy[int, int](compare.Identity[int], a, b)
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

func TestFlatten(t *testing.T) {
	s := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7},
		{8, 9, 10},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ret := Flatten(s)
	if !reflect.DeepEqual(ret, expected) {
		t.Logf("expected %v, but got %v", expected, ret)
		t.Fail()
	}

}

func TestInterleave(t *testing.T) {
	type args[A any] struct {
		slices [][]A
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want []A
	}
	tests := []testCase[int]{
		{name: "basic",
			args: struct{ slices [][]int }{slices: [][]int{[]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10}}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interleave(tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	type args[E comparable] struct {
		haystack    []E
		needle      E
		replacement E
		n           int
	}
	type testCase[E comparable] struct {
		name string
		args args[E]
		want []E
	}
	tests := []testCase[int]{
		{
			name: "none",
			args: args[int]{[]int{1, 2, 1, 2, 1}, 3, 1, 2},
			want: []int{1, 2, 1, 2, 1},
		},
		{
			name: "first",
			args: args[int]{[]int{1, 2, 2, 2, 1}, 2, 1, 1},
			want: []int{1, 1, 2, 2, 1},
		},
		{
			name: "two",
			args: args[int]{[]int{1, 2, 1, 2, 1}, 2, 1, 2},
			want: []int{1, 1, 1, 1, 1},
		},
		{
			name: "all",
			args: args[int]{[]int{1, 2, 2, 2, 1}, 2, 1, -1},
			want: []int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.args.haystack, tt.args.needle, tt.args.replacement, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args[A any] struct {
		slice []A
		n     int
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want [][]A
	}
	tests := []testCase[int]{
		{
			name: "simple1",
			args: args[int]{[]int{0, 1, 2, 3, 4, 5, 6}, -1},
			want: [][]int{{0, 1, 2, 3, 4, 5, 6}},
		},
		{
			name: "simple1",
			args: args[int]{[]int{0, 1, 2, 3, 4, 5, 6}, 2},
			want: [][]int{{0, 1}, {2, 3}, {4, 5}, {6}},
		},
		{
			name: "simple2",
			args: args[int]{[]int{0, 1, 2, 3, 4, 5, 6}, 3},
			want: [][]int{{0, 1, 2}, {3, 4, 5}, {6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithout(t *testing.T) {
	type args[A comparable] struct {
		slice   []A
		exclude []A
	}
	type testCase[A comparable] struct {
		name string
		args args[A]
		want []A
	}
	tests := []testCase[int]{
		{
			name: "simple",
			args: args[int]{
				slice:   []int{1, 2, 3, 4, 5, 3, 3, 2, 1},
				exclude: []int{2, 3},
			},
			want: []int{1, 4, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Without(tt.args.slice, tt.args.exclude...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Without() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitial(t *testing.T) {
	type args[A any] struct {
		slice []A
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want []A
	}
	tests := []testCase[int]{
		{
			name: "",
			args: args[int]{
				slice: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Initial(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXOR(t *testing.T) {
	type args[A comparable] struct {
		slices [][]A
	}
	type testCase[A comparable] struct {
		name string
		args args[A]
		want []A
	}
	tests := []testCase[int]{
		{
			name: "simple",
			args: args[int]{
				slices: [][]int{{2, 1}, {2, 3}},
			},
			want: []int{1, 3},
		},
		{
			name: "simple2",
			args: args[int]{
				slices: [][]int{{2, 1}, {2, 3}, {4, 5}, {5, 6}},
			},
			want: []int{1, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XOR(tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXORBy(t *testing.T) {
	type args[A any, B comparable] struct {
		by     func(A) B
		slices [][]A
	}
	type testCase[A any, B comparable] struct {
		name string
		args args[A, B]
		want []A
	}
	tests := []testCase[float64, float64]{
		{
			name: "floor",
			args: args[float64, float64]{
				by:     math.Floor,
				slices: [][]float64{{2.1, 1.2}, {2.3, 3.4}},
			},
			want: []float64{1.2, 3.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XORBy(tt.args.by, tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XORBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	ints := []int{1, 2, 3, 2, 1}
	if Index(ints, 3) != 2 {
		t.Errorf("Index(ints, 3) = %d, want 2", Index(ints, 3))
	}
	if Index(ints, 5) != -1 {
		t.Errorf("Index(ints, 5) = %d, want -1", Index(ints, 5))
	}
}

func TestCutBy_NotFound(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	left, right, found := CutBy(ints, func(e int) bool { return e == 10 })
	if found {
		t.Error("Expected found = false for non-existent element")
	}
	if !reflect.DeepEqual(left, ints) {
		t.Errorf("Expected left = original slice when not found")
	}
	if right != nil {
		t.Error("Expected right = nil when not found")
	}
}

func TestReplaceFirstAndAll(t *testing.T) {
	ints := []int{1, 2, 3, 2, 1}

	// ReplaceFirst
	result := ReplaceFirst(ints, 2, 9)
	expected := []int{1, 9, 3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReplaceFirst() = %v, want %v", result, expected)
	}

	// ReplaceAll
	result2 := ReplaceAll(ints, 2, 9)
	expected2 := []int{1, 9, 3, 9, 1}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("ReplaceAll() = %v, want %v", result2, expected2)
	}
}

func TestFindLast(t *testing.T) {
	ints := []int{1, 2, 3, 2, 1}
	val, found := FindLast(ints, func(e int) bool { return e == 2 })
	if !found {
		t.Error("Expected found = true")
	}
	if val != 2 {
		t.Errorf("FindLast() = %d, want 2", val)
	}

	// Not found
	_, notFound := FindLast(ints, func(e int) bool { return e == 10 })
	if notFound {
		t.Error("Expected not found for non-existent element")
	}
}

func TestContains(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	if !Contains(ints, 3) {
		t.Error("Contains(ints, 3) should return true")
	}
	if Contains(ints, 10) {
		t.Error("Contains(ints, 10) should return false")
	}
}

func TestContainsBy(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	if !ContainsBy(ints, func(e int) bool { return e > 3 }) {
		t.Error("ContainsBy(ints, >3) should return true")
	}
	if ContainsBy(ints, func(e int) bool { return e > 10 }) {
		t.Error("ContainsBy(ints, >10) should return false")
	}
}

func TestForEach(t *testing.T) {
	ints := []int{1, 2, 3}
	var sum int
	ForEach(ints, func(a int) {
		sum += a
	})
	if sum != 6 {
		t.Errorf("ForEach sum = %d, want 6", sum)
	}
}

func TestForEachRight(t *testing.T) {
	ints := []int{1, 2, 3}
	var result []int
	ForEachRight(ints, func(a int) {
		result = append(result, a)
	})
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ForEachRight() = %v, want %v", result, expected)
	}
}

func TestClone_Nil(t *testing.T) {
	var nilSlice []int
	cloned := Clone(nilSlice)
	if cloned != nil {
		t.Error("Clone(nil) should return nil")
	}

	// Non-nil slice
	nonNil := []int{1, 2, 3}
	cloned2 := Clone(nonNil)
	if !reflect.DeepEqual(cloned2, nonNil) {
		t.Errorf("Clone() = %v, want %v", cloned2, nonNil)
	}
}

func TestCompare(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	c := []int{1, 2}
	d := []int{1, 2, 3}

	if Compare(a, b) != -1 {
		t.Error("Compare([1,2,3], [1,2,4]) should return -1")
	}
	if Compare(b, a) != 1 {
		t.Error("Compare([1,2,4], [1,2,3]) should return 1")
	}
	if Compare(a, d) != 0 {
		t.Error("Compare([1,2,3], [1,2,3]) should return 0")
	}
	if Compare(a, c) != 1 {
		t.Error("Compare([1,2,3], [1,2]) should return 1")
	}
}

func TestCompareBy(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	c := []int{1, 2}

	// Simple comparison function
	cmp := func(x, y int) int {
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return 0
	}

	if CompareBy(a, b, cmp) != -1 {
		t.Error("CompareBy([1,2,3], [1,2,4]) should return -1")
	}
	if CompareBy(b, a, cmp) != 1 {
		t.Error("CompareBy([1,2,4], [1,2,3]) should return 1")
	}
	if CompareBy(a, a, cmp) != 0 {
		t.Error("CompareBy([1,2,3], [1,2,3]) should return 0")
	}
	if CompareBy(a, c, cmp) != 1 {
		t.Error("CompareBy([1,2,3], [1,2]) should return 1")
	}
}

func TestRepeatBy(t *testing.T) {
	result := RepeatBy(5, func(i int) int { return i * i })
	expected := []int{0, 1, 4, 9, 16}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RepeatBy() = %v, want %v", result, expected)
	}
}

func TestZip2(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{10, 20, 30}
	c := []int{100, 200, 300}

	result := Zip2(a, b, c, func(x, y, z int) int { return x + y + z })
	expected := []int{111, 222, 333}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zip2() = %v, want %v", result, expected)
	}

	// Test with different lengths
	a2 := []int{1, 2}
	b2 := []int{10, 20, 30}
	c2 := []int{100}
	result2 := Zip2(a2, b2, c2, func(x, y, z int) int { return x + y + z })
	if len(result2) != 1 {
		t.Errorf("Zip2 with different lengths: expected length 1, got %d", len(result2))
	}
}

func TestUnzip2(t *testing.T) {
	input := []string{"111", "222", "333"}
	a, b, c := Unzip2(input, func(s string) (int, int, int) {
		return int(s[0] - '0'), int(s[1] - '0'), int(s[2] - '0')
	})

	if !reflect.DeepEqual(a, []int{1, 2, 3}) {
		t.Errorf("Unzip2 first = %v, want %v", a, []int{1, 2, 3})
	}
	if !reflect.DeepEqual(b, []int{1, 2, 3}) {
		t.Errorf("Unzip2 second = %v, want %v", b, []int{1, 2, 3})
	}
	if !reflect.DeepEqual(c, []int{1, 2, 3}) {
		t.Errorf("Unzip2 third = %v, want %v", c, []int{1, 2, 3})
	}
}

func TestZip3(t *testing.T) {
	a := []int{1, 2}
	b := []int{10, 20}
	c := []int{100, 200}
	d := []int{1000, 2000}

	result := Zip3(a, b, c, d, func(w, x, y, z int) int { return w + x + y + z })
	expected := []int{1111, 2222}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zip3() = %v, want %v", result, expected)
	}
}

func TestUnzip3(t *testing.T) {
	input := []string{"1234", "5678"}
	a, b, c, d := Unzip3(input, func(s string) (int, int, int, int) {
		return int(s[0] - '0'), int(s[1] - '0'), int(s[2] - '0'), int(s[3] - '0')
	})

	if !reflect.DeepEqual(a, []int{1, 5}) {
		t.Errorf("Unzip3 first = %v, want %v", a, []int{1, 5})
	}
	if !reflect.DeepEqual(b, []int{2, 6}) {
		t.Errorf("Unzip3 second = %v, want %v", b, []int{2, 6})
	}
	if !reflect.DeepEqual(c, []int{3, 7}) {
		t.Errorf("Unzip3 third = %v, want %v", c, []int{3, 7})
	}
	if !reflect.DeepEqual(d, []int{4, 8}) {
		t.Errorf("Unzip3 fourth = %v, want %v", d, []int{4, 8})
	}
}

func TestSliceToMap(t *testing.T) {
	slice := []string{"a:1", "b:2", "c:3"}
	result := SliceToMap(slice, func(s string) (string, int) {
		parts := strings.Split(s, ":")
		val, _ := strconv.Atoi(parts[1])
		return parts[0], val
	})
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SliceToMap() = %v, want %v", result, expected)
	}
}

func TestScanLeft(t *testing.T) {
	// Running sums starting from 0
	result := ScanLeft([]int{1, 2, 3}, func(acc, val int) int { return acc + val }, 0)
	expected := []int{0, 1, 3, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ScanLeft() = %v, want %v", result, expected)
	}

	// String concatenation
	strings := ScanLeft([]string{"a", "b", "c"}, func(acc, val string) string { return acc + val }, "")
	expectedStr := []string{"", "a", "ab", "abc"}
	if !reflect.DeepEqual(strings, expectedStr) {
		t.Errorf("ScanLeft() strings = %v, want %v", strings, expectedStr)
	}
}

func TestScanRight(t *testing.T) {
	// Running sums from right
	result := ScanRight([]int{1, 2, 3}, func(acc, val int) int { return acc + val }, 0)
	expected := []int{6, 5, 3, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ScanRight() = %v, want %v", result, expected)
	}
}

func TestScan(t *testing.T) {
	// Scan is alias for ScanLeft
	result := Scan([]int{1, 2, 3}, func(acc, val int) int { return acc + val }, 0)
	expected := []int{0, 1, 3, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Scan() = %v, want %v", result, expected)
	}
}

func TestSlidingWindow(t *testing.T) {
	result := SlidingWindow([]int{1, 2, 3, 4, 5}, 3)
	expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SlidingWindow() = %v, want %v", result, expected)
	}

	// Edge cases
	if SlidingWindow([]int{1, 2}, 3) != nil {
		t.Error("SlidingWindow with n > len should return nil")
	}
	if SlidingWindow([]int{1, 2, 3}, 0) != nil {
		t.Error("SlidingWindow with n <= 0 should return nil")
	}
}

func TestSlidingWindow2(t *testing.T) {
	result := SlidingWindow([]int{1, 2, 3, 4, 5}, 4)
	expected := [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SlidingWindow() = %v, want %v", result, expected)
	}
}

func TestTranspose(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	result := Transpose(matrix)
	expected := [][]int{{1, 4}, {2, 5}, {3, 6}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Transpose() = %v, want %v", result, expected)
	}

	// Edge cases
	if Transpose([][]int{}) != nil {
		t.Error("Transpose of empty matrix should return nil")
	}
	if Transpose([][]int{{}}) != nil {
		t.Error("Transpose of matrix with empty rows should return nil")
	}
}

func TestIntersperse(t *testing.T) {
	result := Intersperse([]int{1, 2, 3}, 0)
	expected := []int{1, 0, 2, 0, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Intersperse() = %v, want %v", result, expected)
	}

	// Edge cases
	if !reflect.DeepEqual(Intersperse([]int{1}, 0), []int{1}) {
		t.Error("Intersperse with single element should return clone")
	}
	if !reflect.DeepEqual(Intersperse([]int{}, 0), []int{}) {
		t.Error("Intersperse with empty slice should return empty slice")
	}
}

func TestSplitAt(t *testing.T) {
	before, after := SplitAt([]int{1, 2, 3, 4, 5}, 2)
	if !reflect.DeepEqual(before, []int{1, 2}) {
		t.Errorf("SplitAt before = %v, want %v", before, []int{1, 2})
	}
	if !reflect.DeepEqual(after, []int{3, 4, 5}) {
		t.Errorf("SplitAt after = %v, want %v", after, []int{3, 4, 5})
	}

	// Edge cases
	b1, a1 := SplitAt([]int{1, 2, 3}, 0)
	if len(b1) != 0 || !reflect.DeepEqual(a1, []int{1, 2, 3}) {
		t.Error("SplitAt at index 0 should return empty before and full after")
	}

	b2, a2 := SplitAt([]int{1, 2, 3}, 5)
	if !reflect.DeepEqual(b2, []int{1, 2, 3}) || len(a2) != 0 {
		t.Error("SplitAt at index >= len should return full before and empty after")
	}
}

func TestSpan(t *testing.T) {
	init, rest := Span([]int{1, 2, 3, 4, 5}, func(n int) bool { return n < 4 })
	if !reflect.DeepEqual(init, []int{1, 2, 3}) {
		t.Errorf("Span init = %v, want %v", init, []int{1, 2, 3})
	}
	if !reflect.DeepEqual(rest, []int{4, 5}) {
		t.Errorf("Span rest = %v, want %v", rest, []int{4, 5})
	}

	// All satisfy predicate
	i1, r1 := Span([]int{1, 2, 3}, func(n int) bool { return n > 0 })
	if !reflect.DeepEqual(i1, []int{1, 2, 3}) || len(r1) != 0 {
		t.Error("Span when all satisfy should return full init and empty rest")
	}

	// None satisfy predicate
	i2, r2 := Span([]int{4, 5, 6}, func(n int) bool { return n < 4 })
	if len(i2) != 0 || !reflect.DeepEqual(r2, []int{4, 5, 6}) {
		t.Error("Span when none satisfy should return empty init and full rest")
	}
}

func TestMapIdx(t *testing.T) {
	result := MapIdx([]string{"a", "b", "c"}, func(i int, s string) string {
		return fmt.Sprintf("%d:%s", i, s)
	})
	expected := []string{"0:a", "1:b", "2:c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MapIdx() = %v, want %v", result, expected)
	}
}

func TestFilterIdx(t *testing.T) {
	// Keep elements at even indices with value > 10
	result := FilterIdx([]int{5, 20, 15, 25, 10}, func(i int, n int) bool {
		return i%2 == 0 && n > 10
	})
	expected := []int{15} // index 2, value 15
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterIdx() = %v, want %v", result, expected)
	}
}

func TestRejectIdx(t *testing.T) {
	// Reject elements at even indices
	result := RejectIdx([]int{1, 2, 3, 4, 5}, func(i int, n int) bool {
		return i%2 == 0
	})
	expected := []int{2, 4} // elements at odd indices
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RejectIdx() = %v, want %v", result, expected)
	}
}

func TestIsAllUnique(t *testing.T) {
	if !IsAllUnique([]int{1, 2, 3, 4, 5}) {
		t.Error("IsAllUnique should return true for unique elements")
	}
	if IsAllUnique([]int{1, 2, 2, 3}) {
		t.Error("IsAllUnique should return false for duplicates")
	}
	if !IsAllUnique([]string{"a", "b", "c"}) {
		t.Error("IsAllUnique should work with strings")
	}
	if !IsAllUnique([]int{}) {
		t.Error("IsAllUnique should return true for empty slice")
	}
}

func TestIsSorted(t *testing.T) {
	if !IsSorted([]int{1, 2, 3, 4, 5}) {
		t.Error("IsSorted should return true for sorted slice")
	}
	if !IsSorted([]int{1, 2, 2, 3}) {
		t.Error("IsSorted should return true for sorted slice with duplicates")
	}
	if IsSorted([]int{3, 2, 1}) {
		t.Error("IsSorted should return false for reverse sorted slice")
	}
	if !IsSorted([]int{}) {
		t.Error("IsSorted should return true for empty slice")
	}
	if !IsSorted([]int{42}) {
		t.Error("IsSorted should return true for single element")
	}
}

func TestIsSortedBy(t *testing.T) {
	// Sort by length
	strings := []string{"a", "bb", "ccc"}
	if !IsSortedBy(strings, func(a, b string) bool { return len(a) < len(b) }) {
		t.Error("IsSortedBy should return true when sorted by custom criteria")
	}

	// Not sorted by first character
	if IsSortedBy([]string{"b", "a", "c"}, func(a, b string) bool { return a < b }) {
		t.Error("IsSortedBy should return false when not sorted")
	}
}

func TestGroupByOrdered(t *testing.T) {
	words := []string{"apple", "banana", "avocado", "blueberry", "cherry"}
	result := GroupByOrdered(words, func(s string) string { return string(s[0]) })

	if len(result) != 3 {
		t.Errorf("Expected 3 groups, got %d", len(result))
	}

	// Check order is preserved (a, b, c)
	if result[0].Key != "a" || len(result[0].Values) != 2 {
		t.Errorf("First group should be 'a' with 2 values, got %s with %d", result[0].Key, len(result[0].Values))
	}
	if result[1].Key != "b" || len(result[1].Values) != 2 {
		t.Errorf("Second group should be 'b' with 2 values, got %s with %d", result[1].Key, len(result[1].Values))
	}
	if result[2].Key != "c" || len(result[2].Values) != 1 {
		t.Errorf("Third group should be 'c' with 1 value, got %s with %d", result[2].Key, len(result[2].Values))
	}

	// Empty slice
	if GroupByOrdered([]string{}, func(s string) string { return s }) != nil {
		t.Error("GroupByOrdered on empty slice should return nil")
	}
}

func TestChunkBy(t *testing.T) {
	// Group by equality
	result := ChunkBy([]int{1, 1, 1, 2, 2, 3, 3, 3}, func(a, b int) bool { return a == b })
	expected := [][]int{{1, 1, 1}, {2, 2}, {3, 3, 3}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkBy() = %v, want %v", result, expected)
	}

	// Group by increasing sequence
	result2 := ChunkBy([]int{1, 2, 3, 2, 2, 1}, func(a, b int) bool { return a <= b })
	expected2 := [][]int{{1, 2, 3}, {2, 2}, {1}}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("ChunkBy() = %v, want %v", result2, expected2)
	}

	// Empty slice
	if ChunkBy([]int{}, func(a, b int) bool { return a == b }) != nil {
		t.Error("ChunkBy on empty slice should return nil")
	}

	// Single element
	single := ChunkBy([]int{42}, func(a, b int) bool { return a == b })
	if len(single) != 1 || !reflect.DeepEqual(single[0], []int{42}) {
		t.Error("ChunkBy with single element should return one chunk")
	}
}

func TestDeduplicate(t *testing.T) {
	// Remove consecutive duplicates
	result := Deduplicate([]int{1, 1, 2, 2, 2, 3, 3})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Deduplicate() = %v, want %v", result, expected)
	}

	// Non-consecutive duplicates are preserved
	result2 := Deduplicate([]int{1, 2, 1, 2, 1})
	expected2 := []int{1, 2, 1, 2, 1}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Deduplicate() = %v, want %v", result2, expected2)
	}

	// Empty slice
	if !reflect.DeepEqual(Deduplicate([]int{}), []int{}) {
		t.Error("Deduplicate on empty slice should return empty slice")
	}

	// Single element
	if !reflect.DeepEqual(Deduplicate([]int{42}), []int{42}) {
		t.Error("Deduplicate with single element should return same element")
	}

	// All same
	result3 := Deduplicate([]int{5, 5, 5, 5})
	if !reflect.DeepEqual(result3, []int{5}) {
		t.Errorf("Deduplicate all same = %v, want [5]", result3)
	}
}

func TestFill(t *testing.T) {
	result := Fill(5, 42)
	expected := []int{42, 42, 42, 42, 42}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Fill() = %v, want %v", result, expected)
	}

	// Strings
	strResult := Fill(3, "x")
	strExpected := []string{"x", "x", "x"}
	if !reflect.DeepEqual(strResult, strExpected) {
		t.Errorf("Fill() strings = %v, want %v", strResult, strExpected)
	}

	// Zero or negative
	if !reflect.DeepEqual(Fill(0, 42), []int{}) {
		t.Error("Fill with n=0 should return empty slice")
	}
	if !reflect.DeepEqual(Fill(-5, 42), []int{}) {
		t.Error("Fill with negative n should return empty slice")
	}
}

func TestRange(t *testing.T) {
	result := Range(1, 5)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Range(1, 5) = %v, want %v", result, expected)
	}

	// Same start and end
	result2 := Range(3, 3)
	if !reflect.DeepEqual(result2, []int{3}) {
		t.Errorf("Range(3, 3) = %v, want [3]", result2)
	}

	// Start > end
	result3 := Range(5, 1)
	if !reflect.DeepEqual(result3, []int{}) {
		t.Errorf("Range(5, 1) = %v, want empty", result3)
	}

	// Negative numbers
	result4 := Range(-3, 2)
	expected4 := []int{-3, -2, -1, 0, 1, 2}
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Range(-3, 2) = %v, want %v", result4, expected4)
	}
}

func TestRangeFrom(t *testing.T) {
	result := RangeFrom(10, 5)
	expected := []int{10, 11, 12, 13, 14}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RangeFrom(10, 5) = %v, want %v", result, expected)
	}

	// Zero count
	if !reflect.DeepEqual(RangeFrom(5, 0), []int{}) {
		t.Error("RangeFrom with n=0 should return empty slice")
	}

	// Negative count
	if !reflect.DeepEqual(RangeFrom(5, -3), []int{}) {
		t.Error("RangeFrom with negative n should return empty slice")
	}
}

func TestRangeStep(t *testing.T) {
	// Positive step
	result := RangeStep(0, 10, 2)
	expected := []int{0, 2, 4, 6, 8, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RangeStep(0, 10, 2) = %v, want %v", result, expected)
	}

	// Negative step
	result2 := RangeStep(10, 0, -2)
	expected2 := []int{10, 8, 6, 4, 2, 0}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("RangeStep(10, 0, -2) = %v, want %v", result2, expected2)
	}

	// Step that doesn't land exactly on end
	result3 := RangeStep(0, 10, 3)
	expected3 := []int{0, 3, 6, 9}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("RangeStep(0, 10, 3) = %v, want %v", result3, expected3)
	}

	// Zero step
	if !reflect.DeepEqual(RangeStep(0, 10, 0), []int{}) {
		t.Error("RangeStep with step=0 should return empty slice")
	}

	// Incompatible direction
	if !reflect.DeepEqual(RangeStep(0, 10, -1), []int{}) {
		t.Error("RangeStep with wrong direction should return empty slice")
	}
}

func TestRepeat(t *testing.T) {
	result := Repeat([]int{1, 2}, 3)
	expected := []int{1, 2, 1, 2, 1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Repeat() = %v, want %v", result, expected)
	}

	// Zero or negative repeats
	if !reflect.DeepEqual(Repeat([]int{1, 2}, 0), []int{}) {
		t.Error("Repeat with n=0 should return empty slice")
	}
	if !reflect.DeepEqual(Repeat([]int{1, 2}, -1), []int{}) {
		t.Error("Repeat with negative n should return empty slice")
	}

	// Empty slice
	if !reflect.DeepEqual(Repeat([]int{}, 5), []int{}) {
		t.Error("Repeat with empty slice should return empty slice")
	}

	// Single element
	single := Repeat([]int{42}, 4)
	if !reflect.DeepEqual(single, []int{42, 42, 42, 42}) {
		t.Errorf("Repeat single element = %v, want [42, 42, 42, 42]", single)
	}
}

// ============================================================================
// MISSING FUNCTION TESTS
// ============================================================================

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "take while less than 4",
			input:     []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool { return n < 4 },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "empty slice",
			input:     []int{},
			predicate: func(n int) bool { return n < 4 },
			expected:  nil,
		},
		{
			name:      "nil slice",
			input:     nil,
			predicate: func(n int) bool { return n < 4 },
			expected:  nil,
		},
		{
			name:      "all satisfy predicate",
			input:     []int{1, 2, 3},
			predicate: func(n int) bool { return n > 0 },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "none satisfy predicate",
			input:     []int{5, 6, 7},
			predicate: func(n int) bool { return n < 4 },
			expected:  nil,
		},
		{
			name:      "first element fails",
			input:     []int{5, 1, 2, 3},
			predicate: func(n int) bool { return n < 4 },
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TakeWhile(tt.input, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeWhile() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNoneBy(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "no elements satisfy",
			input:     []int{1, 2, 3},
			predicate: func(n int) bool { return n < 0 },
			expected:  true,
		},
		{
			name:      "some elements satisfy",
			input:     []int{1, 2, -3},
			predicate: func(n int) bool { return n < 0 },
			expected:  false,
		},
		{
			name:      "empty slice",
			input:     []int{},
			predicate: func(n int) bool { return n < 0 },
			expected:  true,
		},
		{
			name:      "nil slice",
			input:     nil,
			predicate: func(n int) bool { return n < 0 },
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NoneBy(tt.input, tt.predicate)
			if result != tt.expected {
				t.Errorf("NoneBy() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		isZero   bool
	}{
		{
			name:     "basic max",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: 9,
			isZero:   false,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: 42,
			isZero:   false,
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: 2,
			isZero:   false,
		},
		{
			name:     "all same",
			input:    []int{5, 5, 5},
			expected: 5,
			isZero:   false,
		},
		{
			name:     "negative numbers",
			input:    []int{-10, -5, -20},
			expected: -5,
			isZero:   false,
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-5, 0, 5},
			expected: 5,
			isZero:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.input...)
			if result != tt.expected {
				t.Errorf("Max() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test empty slice (should return zero value)
	t.Run("empty slice", func(t *testing.T) {
		result := Max[int]()
		if result != 0 {
			t.Errorf("Max() with empty slice = %v, want 0", result)
		}
	})

	// Test with strings
	t.Run("strings", func(t *testing.T) {
		result := Max("apple", "banana", "cherry")
		if result != "cherry" {
			t.Errorf("Max(strings) = %v, want cherry", result)
		}
	})
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "basic min",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: 1,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "two elements",
			input:    []int{2, 1},
			expected: 1,
		},
		{
			name:     "all same",
			input:    []int{5, 5, 5},
			expected: 5,
		},
		{
			name:     "negative numbers",
			input:    []int{-10, -5, -20},
			expected: -20,
		},
		{
			name:     "mixed positive and negative",
			input:    []int{-5, 0, 5},
			expected: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.input...)
			if result != tt.expected {
				t.Errorf("Min() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test empty slice (should return zero value)
	t.Run("empty slice", func(t *testing.T) {
		result := Min[int]()
		if result != 0 {
			t.Errorf("Min() with empty slice = %v, want 0", result)
		}
	})

	// Test with strings
	t.Run("strings", func(t *testing.T) {
		result := Min("cherry", "apple", "banana")
		if result != "apple" {
			t.Errorf("Min(strings) = %v, want apple", result)
		}
	})
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		predicate     func(int) bool
		expectedIdx   int
		expectedVal   int
		expectedFound bool
	}{
		{
			name:          "find first >= 23",
			input:         []int{10, 20, 30, 40, 50},
			predicate:     func(e int) bool { return e >= 23 },
			expectedIdx:   2,
			expectedVal:   30,
			expectedFound: true,
		},
		{
			name:          "find first >= 0",
			input:         []int{10, 20, 30},
			predicate:     func(e int) bool { return e >= 0 },
			expectedIdx:   0,
			expectedVal:   10,
			expectedFound: true,
		},
		{
			name:          "find insertion point for large value",
			input:         []int{10, 20, 30},
			predicate:     func(e int) bool { return e >= 100 },
			expectedIdx:   3,
			expectedVal:   0,
			expectedFound: false,
		},
		{
			name:          "empty slice",
			input:         []int{},
			predicate:     func(e int) bool { return e >= 10 },
			expectedIdx:   0,
			expectedVal:   0,
			expectedFound: false,
		},
		{
			name:          "single element - found",
			input:         []int{42},
			predicate:     func(e int) bool { return e >= 10 },
			expectedIdx:   0,
			expectedVal:   42,
			expectedFound: true,
		},
		{
			name:          "single element - not found",
			input:         []int{5},
			predicate:     func(e int) bool { return e >= 10 },
			expectedIdx:   1,
			expectedVal:   0,
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, val := Search(tt.input, tt.predicate)
			if idx != tt.expectedIdx {
				t.Errorf("Search() index = %v, want %v", idx, tt.expectedIdx)
			}
			found := idx < len(tt.input)
			if found != tt.expectedFound {
				t.Errorf("Search() found = %v, want %v", found, tt.expectedFound)
			}
			if found && val != tt.expectedVal {
				t.Errorf("Search() value = %v, want %v", val, tt.expectedVal)
			}
		})
	}
}

// ============================================================================
// EDGE CASE TESTS
// ============================================================================

func TestNilVsEmptySlice(t *testing.T) {
	// Test that nil and empty slices are handled correctly

	// Equal - nil and empty both have length 0, so they're equal in terms of elements
	t.Run("Equal nil vs empty", func(t *testing.T) {
		var nilSlice []int
		emptySlice := []int{}

		// Both nil and empty have len=0, so they compare as equal
		if !Equal(nilSlice, emptySlice) {
			t.Error("Equal(nil, empty) should return true (both have len=0)")
		}
		if !Equal(nilSlice, nilSlice) {
			t.Error("Equal(nil, nil) should return true")
		}
		if !Equal(emptySlice, emptySlice) {
			t.Error("Equal(empty, empty) should return true")
		}
	})

	// Clone - should preserve nil
	t.Run("Clone preserves nil", func(t *testing.T) {
		var nilSlice []int
		cloned := Clone(nilSlice)
		if cloned != nil {
			t.Error("Clone(nil) should return nil")
		}

		emptySlice := []int{}
		clonedEmpty := Clone(emptySlice)
		if clonedEmpty == nil {
			t.Error("Clone(empty) should return empty slice, not nil")
		}
		if len(clonedEmpty) != 0 {
			t.Error("Clone(empty) should return slice with length 0")
		}
	})

	// Drop - nil returns nil
	t.Run("Drop nil", func(t *testing.T) {
		var nilSlice []int
		result := Drop(nilSlice, 1)
		if result != nil {
			t.Errorf("Drop(nil, 1) should return nil, got %v", result)
		}
	})

	// Take - nil returns nil
	t.Run("Take nil", func(t *testing.T) {
		var nilSlice []int
		result := Take(nilSlice, 1)
		if result != nil {
			t.Errorf("Take(nil, 1) should return nil, got %v", result)
		}
	})

	// Filter - nil returns empty slice (not nil, since Filter creates a new slice)
	t.Run("Filter nil", func(t *testing.T) {
		var nilSlice []int
		result := Filter(nilSlice, func(n int) bool { return true })
		// Filter creates a new slice with make([]A, 0, len(slice)), so it returns [] not nil
		if result == nil {
			t.Error("Filter(nil) should return empty slice [], not nil")
		}
		if len(result) != 0 {
			t.Errorf("Filter(nil) should return slice with length 0, got %d", len(result))
		}
	})

	// Map - nil returns empty
	t.Run("Map nil", func(t *testing.T) {
		var nilSlice []int
		result := Map(nilSlice, func(n int) string { return "" })
		if result == nil {
			t.Error("Map(nil) should return empty slice, not nil")
		}
		if len(result) != 0 {
			t.Errorf("Map(nil) should return slice with length 0, got %d", len(result))
		}
	})

	// Contains - nil returns false
	t.Run("Contains nil", func(t *testing.T) {
		var nilSlice []int
		if Contains(nilSlice, 1) {
			t.Error("Contains(nil, 1) should return false")
		}
	})

	// Head - nil returns error
	t.Run("Head nil", func(t *testing.T) {
		var nilSlice []int
		_, err := Head(nilSlice)
		if err == nil {
			t.Error("Head(nil) should return error")
		}
	})

	// Last - nil returns error
	t.Run("Last nil", func(t *testing.T) {
		var nilSlice []int
		_, err := Last(nilSlice)
		if err == nil {
			t.Error("Last(nil) should return error")
		}
	})
}

func TestNegativeIndices(t *testing.T) {
	// Test negative index handling
	slice := []int{1, 2, 3, 4, 5}

	// Nth with negative indices
	t.Run("Nth negative", func(t *testing.T) {
		if Nth(slice, -1) != 5 {
			t.Errorf("Nth(-1) should return last element")
		}
		if Nth(slice, -2) != 4 {
			t.Errorf("Nth(-2) should return second to last")
		}
		if Nth(slice, -5) != 1 {
			t.Errorf("Nth(-5) should return first element")
		}
	})
}

func TestBoundaryConditions(t *testing.T) {
	// Test boundary conditions for various functions

	slice := []int{1, 2, 3, 4, 5}

	// Drop - boundary conditions
	t.Run("Drop boundary", func(t *testing.T) {
		// Drop 0 should return copy of slice
		result := Drop(slice, 0)
		if !Equal(result, slice) {
			t.Error("Drop(slice, 0) should return copy of slice")
		}

		// Drop all elements
		result = Drop(slice, 5)
		if result != nil {
			t.Error("Drop(slice, len) should return nil")
		}

		// Drop more than length
		result = Drop(slice, 10)
		if result != nil {
			t.Error("Drop(slice, >len) should return nil")
		}

		// Drop negative
		result = Drop(slice, -1)
		if !Equal(result, slice) {
			t.Error("Drop(slice, -1) should return copy of slice")
		}
	})

	// Take - boundary conditions
	t.Run("Take boundary", func(t *testing.T) {
		// Take 0
		result := Take(slice, 0)
		if len(result) != 0 {
			t.Error("Take(slice, 0) should return empty slice")
		}

		// Take all elements
		result = Take(slice, 5)
		if !Equal(result, slice) {
			t.Error("Take(slice, len) should return copy of slice")
		}

		// Take more than length
		result = Take(slice, 10)
		if !Equal(result, slice) {
			t.Error("Take(slice, >len) should return copy of slice")
		}

		// Take negative
		result = Take(slice, -1)
		if len(result) != 0 {
			t.Error("Take(slice, -1) should return empty slice")
		}
	})

	// DropRight - boundary conditions
	t.Run("DropRight boundary", func(t *testing.T) {
		// DropRight 0 should return copy of slice
		result := DropRight(slice, 0)
		if !Equal(result, slice) {
			t.Error("DropRight(slice, 0) should return copy of slice")
		}

		// DropRight all elements
		result = DropRight(slice, 5)
		if result != nil {
			t.Error("DropRight(slice, len) should return nil")
		}
	})

	// TakeRight - boundary conditions
	t.Run("TakeRight boundary", func(t *testing.T) {
		// TakeRight 0
		result := TakeRight(slice, 0)
		if len(result) != 0 {
			t.Error("TakeRight(slice, 0) should return empty slice")
		}

		// TakeRight all elements
		result = TakeRight(slice, 5)
		if !Equal(result, slice) {
			t.Error("TakeRight(slice, len) should return copy of slice")
		}

		// TakeRight more than length
		result = TakeRight(slice, 10)
		if !Equal(result, slice) {
			t.Error("TakeRight(slice, >len) should return copy of slice")
		}
	})
}

// ============================================================================
// IMMUTABILITY VERIFICATION TESTS
// ============================================================================

func TestImmutability(t *testing.T) {
	// Test that functions don't mutate the original slice

	original := []int{1, 2, 3, 4, 5}
	originalCopy := []int{1, 2, 3, 4, 5}

	// Reverse
	t.Run("Reverse immutability", func(t *testing.T) {
		result := Reverse(original)
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Reverse mutated original slice")
		}
	})

	// Sort
	t.Run("Sort immutability", func(t *testing.T) {
		unsorted := []int{3, 1, 4, 1, 5}
		unsortedCopy := []int{3, 1, 4, 1, 5}
		result := Sort(unsorted)
		_ = result
		if !Equal(unsorted, unsortedCopy) {
			t.Error("Sort mutated original slice")
		}
	})

	// Filter
	t.Run("Filter immutability", func(t *testing.T) {
		result := Filter(original, func(n int) bool { return n%2 == 0 })
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Filter mutated original slice")
		}
	})

	// Map
	t.Run("Map immutability", func(t *testing.T) {
		result := Map(original, func(n int) int { return n * 2 })
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Map mutated original slice")
		}
	})

	// Take
	t.Run("Take immutability", func(t *testing.T) {
		result := Take(original, 3)
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Take mutated original slice")
		}
	})

	// Drop
	t.Run("Drop immutability", func(t *testing.T) {
		result := Drop(original, 2)
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Drop mutated original slice")
		}
	})

	// Compact
	t.Run("Compact immutability", func(t *testing.T) {
		dups := []int{1, 1, 2, 2, 3}
		dupsCopy := []int{1, 1, 2, 2, 3}
		result := Compact(dups)
		_ = result
		if !Equal(dups, dupsCopy) {
			t.Error("Compact mutated original slice")
		}
	})

	// Replace
	t.Run("Replace immutability", func(t *testing.T) {
		result := Replace(original, 1, 99, -1)
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Replace mutated original slice")
		}
	})

	// Without
	t.Run("Without immutability", func(t *testing.T) {
		result := Without(original, 1, 2)
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Without mutated original slice")
		}
	})

	// Uniq
	t.Run("Uniq immutability", func(t *testing.T) {
		dups := []int{1, 2, 2, 3, 3, 3}
		dupsCopy := []int{1, 2, 2, 3, 3, 3}
		result := Uniq(dups)
		_ = result
		if !Equal(dups, dupsCopy) {
			t.Error("Uniq mutated original slice")
		}
	})

	// Shuffle - test that original is not modified
	t.Run("Shuffle immutability", func(t *testing.T) {
		result := Shuffle(original)
		_ = result
		if !Equal(original, originalCopy) {
			t.Error("Shuffle mutated original slice")
		}
	})
}

// ============================================================================
// PROPERTY-BASED TESTS (INVARIANTS)
// ============================================================================

func TestInvariants(t *testing.T) {
	// Test mathematical invariants that should always hold

	// Filter then Filter with same predicate should equal single Filter
	t.Run("Filter idempotent", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		isEven := func(n int) bool { return n%2 == 0 }

		once := Filter(data, isEven)
		twice := Filter(once, isEven)

		if !Equal(once, twice) {
			t.Error("Filter is not idempotent")
		}
	})

	// Map of identity should return same elements
	t.Run("Map identity", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		identity := func(n int) int { return n }

		result := Map(data, identity)
		if !Equal(result, data) {
			t.Error("Map with identity function changed elements")
		}
	})

	// Reverse of reverse is identity
	t.Run("Reverse involution", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		once := Reverse(data)
		twice := Reverse(once)

		if !Equal(twice, data) {
			t.Error("Reverse(Reverse(x)) != x")
		}
	})

	// Take(n) + Drop(n) = original (for n <= len)
	t.Run("Take/Drop complement", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		n := 3

		taken := Take(data, n)
		dropped := Drop(data, n)
		combined := Concat(taken, dropped)

		if !Equal(combined, data) {
			t.Error("Take(n) + Drop(n) != original")
		}
	})

	// Filter + Reject = original
	t.Run("Filter/Reject complement", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6}
		isEven := func(n int) bool { return n%2 == 0 }

		filtered := Filter(data, isEven)
		rejected := Reject(data, isEven)
		combined := Concat(filtered, rejected)
		sortedCombined := SortBy(combined, func(a, b int) bool { return a < b })

		if !Equal(sortedCombined, data) {
			t.Error("Filter + Reject elements don't match original")
		}
	})

	// Uniq preserves all unique elements
	t.Run("Uniq preserves elements", func(t *testing.T) {
		data := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
		result := Uniq(data)

		for _, v := range []int{1, 2, 3, 4} {
			if !Contains(result, v) {
				t.Errorf("Uniq lost element %d", v)
			}
		}
	})

	// Sort is idempotent
	t.Run("Sort idempotent", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6}
		once := Sort(data)
		twice := Sort(once)

		if !Equal(once, twice) {
			t.Error("Sort is not idempotent")
		}
	})

	// Contains(x, y) == Some(x, y)
	t.Run("Contains equals Some", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}

		for i := 1; i <= 5; i++ {
			if Contains(data, i) != Some(data, i) {
				t.Errorf("Contains and Some differ for element %d", i)
			}
		}
		if Contains(data, 10) != Some(data, 10) {
			t.Error("Contains and Some differ for non-existent element")
		}
	})

	// Max >= all elements
	t.Run("Max bounds", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6}
		m := Max(data...)

		for _, v := range data {
			if v > m {
				t.Errorf("Max %d is less than element %d", m, v)
			}
		}
	})

	// Min <= all elements
	t.Run("Min bounds", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6}
		m := Min(data...)

		for _, v := range data {
			if v < m {
				t.Errorf("Min %d is greater than element %d", m, v)
			}
		}
	})
}

// ============================================================================
// COMPREHENSIVE EDGE CASE TESTS FOR EXISTING FUNCTIONS
// ============================================================================

func TestCutEdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		needle        int
		expectedLeft  []int
		expectedRight []int
		expectedFound bool
	}{
		{
			name:          "empty slice",
			input:         []int{},
			needle:        1,
			expectedLeft:  []int{},
			expectedRight: nil,
			expectedFound: false,
		},
		{
			name:          "nil slice",
			input:         nil,
			needle:        1,
			expectedLeft:  nil,
			expectedRight: nil,
			expectedFound: false,
		},
		{
			name:          "single element - found",
			input:         []int{1},
			needle:        1,
			expectedLeft:  []int{},
			expectedRight: []int{},
			expectedFound: true,
		},
		{
			name:          "single element - not found",
			input:         []int{1},
			needle:        2,
			expectedLeft:  []int{1},
			expectedRight: nil,
			expectedFound: false,
		},
		{
			name:          "needle at start",
			input:         []int{1, 2, 3},
			needle:        1,
			expectedLeft:  []int{},
			expectedRight: []int{2, 3},
			expectedFound: true,
		},
		{
			name:          "needle at end",
			input:         []int{1, 2, 3},
			needle:        3,
			expectedLeft:  []int{1, 2},
			expectedRight: []int{},
			expectedFound: true,
		},
		{
			name:          "multiple occurrences - cuts at first",
			input:         []int{1, 2, 1, 3},
			needle:        1,
			expectedLeft:  []int{},
			expectedRight: []int{2, 1, 3},
			expectedFound: true,
		},
		{
			name:          "all same elements",
			input:         []int{5, 5, 5},
			needle:        5,
			expectedLeft:  []int{},
			expectedRight: []int{5, 5},
			expectedFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, right, found := Cut(tt.input, tt.needle)
			if !reflect.DeepEqual(left, tt.expectedLeft) {
				t.Errorf("Cut() left = %v, want %v", left, tt.expectedLeft)
			}
			if !reflect.DeepEqual(right, tt.expectedRight) {
				t.Errorf("Cut() right = %v, want %v", right, tt.expectedRight)
			}
			if found != tt.expectedFound {
				t.Errorf("Cut() found = %v, want %v", found, tt.expectedFound)
			}
		})
	}
}

func TestCutByEdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		predicate     func(int) bool
		expectedLeft  []int
		expectedRight []int
		expectedFound bool
	}{
		{
			name:          "empty slice",
			input:         []int{},
			predicate:     func(n int) bool { return n > 2 },
			expectedLeft:  []int{},
			expectedRight: nil,
			expectedFound: false,
		},
		{
			name:          "nil slice",
			input:         nil,
			predicate:     func(n int) bool { return n > 2 },
			expectedLeft:  nil,
			expectedRight: nil,
			expectedFound: false,
		},
		{
			name:          "first element matches",
			input:         []int{5, 1, 2},
			predicate:     func(n int) bool { return n > 4 },
			expectedLeft:  []int{},
			expectedRight: []int{1, 2},
			expectedFound: true,
		},
		{
			name:          "no element matches",
			input:         []int{1, 2, 3},
			predicate:     func(n int) bool { return n > 10 },
			expectedLeft:  []int{1, 2, 3},
			expectedRight: nil,
			expectedFound: false,
		},
		{
			name:          "last element matches",
			input:         []int{1, 2, 5},
			predicate:     func(n int) bool { return n > 4 },
			expectedLeft:  []int{1, 2},
			expectedRight: []int{},
			expectedFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, right, found := CutBy(tt.input, tt.predicate)
			if !reflect.DeepEqual(left, tt.expectedLeft) {
				t.Errorf("CutBy() left = %v, want %v", left, tt.expectedLeft)
			}
			if !reflect.DeepEqual(right, tt.expectedRight) {
				t.Errorf("CutBy() right = %v, want %v", right, tt.expectedRight)
			}
			if found != tt.expectedFound {
				t.Errorf("CutBy() found = %v, want %v", found, tt.expectedFound)
			}
		})
	}
}

func TestIndexEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		expected int
	}{
		{"empty slice", []int{}, 1, -1},
		{"nil slice", nil, 1, -1},
		{"single element - found", []int{1}, 1, 0},
		{"single element - not found", []int{1}, 2, -1},
		{"first element", []int{1, 2, 3}, 1, 0},
		{"last element", []int{1, 2, 3}, 3, 2},
		{"middle element", []int{1, 2, 3}, 2, 1},
		{"multiple occurrences - first", []int{1, 2, 1, 3}, 1, 0},
		{"not found", []int{1, 2, 3}, 4, -1},
		{"all same - found", []int{5, 5, 5}, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Index(tt.input, tt.needle)
			if result != tt.expected {
				t.Errorf("Index() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLastIndexEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		expected int
	}{
		{"empty slice", []int{}, 1, -1},
		{"nil slice", nil, 1, -1},
		{"single element - found", []int{1}, 1, 0},
		{"single element - not found", []int{1}, 2, -1},
		{"first element", []int{1, 2, 3}, 1, 0},
		{"last element", []int{1, 2, 3}, 3, 2},
		{"multiple occurrences - last", []int{1, 2, 1, 3}, 1, 2},
		{"not found", []int{1, 2, 3}, 4, -1},
		{"all same", []int{5, 5, 5}, 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastIndex(tt.input, tt.needle)
			if result != tt.expected {
				t.Errorf("LastIndex() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReplaceEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		replace  int
		n        int
		expected []int
	}{
		{"empty slice", []int{}, 1, 9, -1, []int{}},
		{"nil slice", nil, 1, 9, -1, []int{}},
		{"single element - replace", []int{1}, 1, 9, -1, []int{9}},
		{"single element - no match", []int{1}, 2, 9, -1, []int{1}},
		{"replace first 0", []int{1, 1, 1}, 1, 9, 0, []int{1, 1, 1}},
		{"replace first 1", []int{1, 1, 1}, 1, 9, 1, []int{9, 1, 1}},
		{"replace first 2", []int{1, 1, 1}, 1, 9, 2, []int{9, 9, 1}},
		{"replace all", []int{1, 1, 1}, 1, 9, -1, []int{9, 9, 9}},
		{"replace more than exist", []int{1, 2}, 1, 9, 5, []int{9, 2}},
		{"no matches", []int{1, 2, 3}, 4, 9, -1, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Replace(tt.input, tt.needle, tt.replace, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Replace() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindEdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		predicate     func(int) bool
		expectedVal   int
		expectedFound bool
	}{
		{"empty slice", []int{}, func(n int) bool { return n > 0 }, 0, false},
		{"nil slice", nil, func(n int) bool { return n > 0 }, 0, false},
		{"single - found", []int{5}, func(n int) bool { return n > 0 }, 5, true},
		{"single - not found", []int{5}, func(n int) bool { return n > 10 }, 0, false},
		{"first element", []int{5, 10, 15}, func(n int) bool { return n > 0 }, 5, true},
		{"last element", []int{5, 10, 15}, func(n int) bool { return n > 12 }, 15, true},
		{"no match", []int{1, 2, 3}, func(n int) bool { return n > 10 }, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, found := Find(tt.input, tt.predicate)
			if val != tt.expectedVal {
				t.Errorf("Find() val = %v, want %v", val, tt.expectedVal)
			}
			if found != tt.expectedFound {
				t.Errorf("Find() found = %v, want %v", found, tt.expectedFound)
			}
		})
	}
}

func TestJoinEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		glue     []int
		expected []int
	}{
		{"empty slices", [][]int{}, []int{0}, []int{}},
		{"nil slices", nil, []int{0}, []int{}},
		{"single empty slice", [][]int{{}}, []int{0}, nil},
		{"single slice", [][]int{{1, 2, 3}}, []int{0}, []int{1, 2, 3}},
		{"two slices", [][]int{{1, 2}, {3, 4}}, []int{0}, []int{1, 2, 0, 3, 4}},
		{"empty glue", [][]int{{1, 2}, {3, 4}}, []int{}, []int{1, 2, 3, 4}},
		{"empty middle slice", [][]int{{1, 2}, {}, {3, 4}}, []int{0}, []int{1, 2, 0, 0, 3, 4}},
		{"all empty slices", [][]int{{}, {}}, []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Join(tt.input, tt.glue)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Join() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConcatEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{"no slices", [][]int{}, []int{}},
		{"nil slices", nil, []int{}},
		{"single empty slice", [][]int{{}}, []int{}},
		{"single slice", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"two slices", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{"empty middle slice", [][]int{{1, 2}, {}, {3, 4}}, []int{1, 2, 3, 4}},
		{"all empty", [][]int{{}, {}}, []int{}},
		{"nil in middle", [][]int{{1, 2}, nil, {3, 4}}, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Concat(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Concat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReverseEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"nil", nil, []int{}},
		{"single", []int{1}, []int{1}},
		{"two", []int{1, 2}, []int{2, 1}},
		{"three", []int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Reverse() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestShuffleDeterministic(t *testing.T) {
	// Shuffle with same seed should produce consistent results
	// Note: We can only test that shuffle preserves length and elements

	tests := []struct {
		name  string
		input []int
	}{
		{"empty", []int{}},
		{"nil", nil},
		{"single", []int{1}},
		{"two", []int{1, 2}},
		{"many", []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Shuffle(tt.input)

			// Length preserved
			if len(result) != len(tt.input) {
				t.Errorf("Shuffle() len = %v, want %v", len(result), len(tt.input))
			}

			// All elements preserved (skip for nil/empty)
			if len(tt.input) > 0 {
				sortedResult := Sort(result)
				sortedInput := Sort(tt.input)
				if !reflect.DeepEqual(sortedResult, sortedInput) {
					t.Errorf("Shuffle() elements changed")
				}
			}

			// Nil input returns empty slice
			if tt.input == nil && len(result) != 0 {
				t.Errorf("Shuffle(nil) should return empty slice, got %v", result)
			}
		})
	}
}

func TestSampleEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected int // expected length
	}{
		{"empty slice", []int{}, 3, 0},
		{"nil slice", nil, 3, 0},
		{"sample 0", []int{1, 2, 3}, 0, 0},
		{"sample negative", []int{1, 2, 3}, -1, 0},
		{"sample 1", []int{1, 2, 3}, 1, 1},
		{"sample all", []int{1, 2, 3}, 3, 3},
		{"sample more than len", []int{1, 2}, 5, 2},
		{"single element", []int{42}, 1, 1},
		{"sample more from single", []int{42}, 5, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sample(tt.input, tt.n)
			if len(result) != tt.expected {
				t.Errorf("Sample() len = %v, want %v", len(result), tt.expected)
			}

			// Verify all sampled elements are from input
			for _, v := range result {
				if !Contains(tt.input, v) {
					t.Errorf("Sample() returned element %d not in input", v)
				}
			}
		})
	}
}

func TestFlattenEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{"empty", [][]int{}, []int{}},
		{"nil", nil, []int{}},
		{"single empty", [][]int{{}}, []int{}},
		{"single", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"multiple", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{"empty middle", [][]int{{1, 2}, {}, {3, 4}}, []int{1, 2, 3, 4}},
		{"all empty", [][]int{{}, {}}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Flatten(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Flatten() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestZipEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []string
		expected []string
	}{
		{"empty a", []int{}, []string{"a", "b"}, []string{}},
		{"empty b", []int{1, 2}, []string{}, []string{}},
		{"both empty", []int{}, []string{}, []string{}},
		{"a shorter", []int{1}, []string{"a", "b"}, []string{"1a"}},
		{"b shorter", []int{1, 2}, []string{"a"}, []string{"1a"}},
		{"same length", []int{1, 2}, []string{"a", "b"}, []string{"1a", "2b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Zip(tt.a, tt.b, func(x int, y string) string {
				return fmt.Sprintf("%d%s", x, y)
			})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Zip() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUnzipEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		expectedA []int
		expectedB []string
	}{
		{"empty", []string{}, []int{}, []string{}},
		{"nil", nil, []int{}, []string{}},
		{"single", []string{"1a"}, []int{1}, []string{"a"}},
		{"multiple", []string{"1a", "2b"}, []int{1, 2}, []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := Unzip(tt.input, func(s string) (int, string) {
				x, _ := strconv.Atoi(string(s[0]))
				return x, string(s[1])
			})
			if !reflect.DeepEqual(a, tt.expectedA) {
				t.Errorf("Unzip() a = %v, want %v", a, tt.expectedA)
			}
			if !reflect.DeepEqual(b, tt.expectedB) {
				t.Errorf("Unzip() b = %v, want %v", b, tt.expectedB)
			}
		})
	}
}

func TestUnionEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [][]int
		expected []int
	}{
		{"no slices", [][]int{}, []int{}},
		{"single empty", [][]int{{}}, []int{}},
		{"single slice", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"two empty", [][]int{{}, {}}, []int{}},
		{"with duplicates", [][]int{{1, 2}, {2, 3}}, []int{1, 2, 3}},
		{"all same", [][]int{{1, 1}, {1, 1}}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Union(tt.inputs...)
			// Sort both for comparison since Union doesn't guarantee order
			sortedResult := Sort(result)
			sortedExpected := Sort(tt.expected)
			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("Union() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIntersectionEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [][]int
		expected []int
	}{
		{"no slices", [][]int{}, []int{}},
		{"single empty", [][]int{{}}, []int{}},
		{"no intersection", [][]int{{1, 2}, {3, 4}}, []int{}},
		{"single element", [][]int{{1}, {1}}, []int{1}},
		{"multiple elements", [][]int{{1, 2, 3}, {2, 3, 4}}, []int{2, 3}},
		{"three slices", [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, []int{3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.inputs...)
			sortedResult := Sort(result)
			sortedExpected := Sort(tt.expected)
			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("Intersection() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDifferenceEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{"single slice - all in intersection", [][]int{{1, 2, 3}}, []int{}},
		{"two identical - all in intersection", [][]int{{1, 2}, {1, 2}}, []int{}},
		{"two - unique elements", [][]int{{1, 2, 3}, {2, 4, 5}}, []int{1, 3, 4, 5}},
		{"three slices - nothing common", [][]int{{1, 2, 3, 4}, {2, 5}, {4, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"three slices with common", [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, []int{1, 2, 4, 5}},
		{"with empty", [][]int{{1, 2}, {}}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slices...)
			sortedResult := Sort(result)
			sortedExpected := Sort(tt.expected)
			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("Difference() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestComplementEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{"empty a", []int{}, []int{1, 2}, []int{1, 2}},
		{"empty b", []int{1, 2}, []int{}, []int{}},
		{"both empty", []int{}, []int{}, []int{}},
		{"no complement", []int{1, 2}, []int{1, 2}, []int{}},
		{"partial", []int{1, 2}, []int{1}, []int{}},
		{"full", []int{1, 2}, []int{3, 4}, []int{3, 4}},
		{"b larger", []int{1}, []int{1, 2, 3}, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Complement(tt.a, tt.b)
			sortedResult := Sort(result)
			sortedExpected := Sort(tt.expected)
			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("Complement() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGroupByEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		key      func(int) string
		expected map[string][]int
	}{
		{"empty", []int{}, func(n int) string { return "key" }, map[string][]int{}},
		{"nil", nil, func(n int) string { return "key" }, map[string][]int{}},
		{"single", []int{1}, func(n int) string { return "odd" }, map[string][]int{"odd": {1}}},
		{"multiple groups", []int{1, 2, 3, 4}, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		}, map[string][]int{"odd": {1, 3}, "even": {2, 4}}},
		{"all same group", []int{1, 3, 5}, func(n int) string { return "odd" }, map[string][]int{"odd": {1, 3, 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupBy(tt.input, tt.key)
			if len(result) != len(tt.expected) {
				t.Errorf("GroupBy() len = %v, want %v", len(result), len(tt.expected))
			}
			for k, v := range tt.expected {
				if !reflect.DeepEqual(result[k], v) {
					t.Errorf("GroupBy()[%s] = %v, want %v", k, result[k], v)
				}
			}
		})
	}
}

func TestPartitionEdgeCases(t *testing.T) {
	tests := []struct {
		name              string
		input             []int
		predicate         func(int) bool
		expectedSatisfied []int
		expectedNot       []int
	}{
		{"empty", []int{}, func(n int) bool { return n > 0 }, []int{}, []int{}},
		{"nil", nil, func(n int) bool { return n > 0 }, []int{}, []int{}},
		{"all true", []int{1, 2, 3}, func(n int) bool { return n > 0 }, []int{1, 2, 3}, []int{}},
		{"all false", []int{-1, -2, -3}, func(n int) bool { return n > 0 }, []int{}, []int{-1, -2, -3}},
		{"mixed", []int{1, -2, 3, -4}, func(n int) bool { return n > 0 }, []int{1, 3}, []int{-2, -4}},
		{"single true", []int{5}, func(n int) bool { return n > 0 }, []int{5}, []int{}},
		{"single false", []int{-5}, func(n int) bool { return n > 0 }, []int{}, []int{-5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			satisfied, notSatisfied := Partition(tt.input, tt.predicate)
			if !reflect.DeepEqual(satisfied, tt.expectedSatisfied) {
				t.Errorf("Partition() satisfied = %v, want %v", satisfied, tt.expectedSatisfied)
			}
			if !reflect.DeepEqual(notSatisfied, tt.expectedNot) {
				t.Errorf("Partition() notSatisfied = %v, want %v", notSatisfied, tt.expectedNot)
			}
		})
	}
}

func TestChunkEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected [][]int
	}{
		{"empty", []int{}, 2, nil},
		{"nil", nil, 2, nil},
		{"size 0", []int{1, 2, 3}, 0, [][]int{{1, 2, 3}}},
		{"negative size", []int{1, 2, 3}, -1, [][]int{{1, 2, 3}}},
		{"size 1", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"size 2", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"uneven", []int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}},
		{"larger than slice", []int{1, 2}, 5, [][]int{{1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Chunk(tt.input, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Chunk() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestScanEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		init     int
		expected []int
	}{
		{"empty", []int{}, 0, []int{0}},
		{"nil", nil, 0, []int{0}},
		{"single", []int{5}, 0, []int{0, 5}},
		{"two elements", []int{1, 2}, 0, []int{0, 1, 3}},
		{"with init", []int{1, 2, 3}, 10, []int{10, 11, 13, 16}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Scan(tt.input, func(acc, val int) int { return acc + val }, tt.init)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Scan() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSlidingWindowEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected [][]int
	}{
		{"empty", []int{}, 2, nil},
		{"nil", nil, 2, nil},
		{"n=0", []int{1, 2, 3}, 0, nil},
		{"n negative", []int{1, 2, 3}, -1, nil},
		{"n=1", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"n=2", []int{1, 2, 3}, 2, [][]int{{1, 2}, {2, 3}}},
		{"n=len", []int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
		{"n>len", []int{1, 2}, 3, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SlidingWindow(tt.input, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SlidingWindow() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTransposeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{"empty matrix", [][]int{}, nil},
		{"nil matrix", nil, nil},
		{"single row", [][]int{{1, 2, 3}}, [][]int{{1}, {2}, {3}}},
		{"single column", [][]int{{1}, {2}, {3}}, [][]int{{1, 2, 3}}},
		{"2x2", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}}},
		{"3x2", [][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
		{"empty rows", [][]int{{}, {}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Transpose(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Transpose() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRangeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{"start=end", 5, 5, []int{5}},
		{"start>end", 5, 1, []int{}},
		{"adjacent", 1, 2, []int{1, 2}},
		{"negative range", -3, 2, []int{-3, -2, -1, 0, 1, 2}},
		{"both negative", -5, -3, []int{-5, -4, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Range(tt.start, tt.end)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Range(%d, %d) = %v, want %v", tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestRangeStepEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		step     int
		expected []int
	}{
		{"step=0", 1, 5, 0, []int{}},
		{"wrong direction", 1, 5, -1, []int{}},
		{"wrong direction negative", 5, 1, 1, []int{}},
		{"step doesn't divide evenly", 1, 10, 3, []int{1, 4, 7, 10}},
		{"single step", 1, 5, 5, []int{1}},
		{"large step", 1, 5, 10, []int{1}},
		{"negative step", 10, 0, -2, []int{10, 8, 6, 4, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RangeStep(tt.start, tt.end, tt.step)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RangeStep(%d, %d, %d) = %v, want %v", tt.start, tt.end, tt.step, result, tt.expected)
			}
		})
	}
}

func TestFillEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		value    int
		expected []int
	}{
		{"n=0", 0, 42, []int{}},
		{"n=1", 1, 42, []int{42}},
		{"n=5", 5, 42, []int{42, 42, 42, 42, 42}},
		{"negative n", -1, 42, []int{}},
		{"large n", 1000, 42, nil}, // Don't check exact, just that it works
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fill(tt.n, tt.value)
			if tt.expected != nil && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Fill(%d, %d) = %v, want %v", tt.n, tt.value, result, tt.expected)
			}
			// Just check length for large n
			if tt.expected == nil && len(result) != 1000 {
				t.Errorf("Fill(%d, %d) len = %v, want %v", tt.n, tt.value, len(result), 1000)
			}
		})
	}
}

func TestRepeatEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{"empty slice", []int{}, 3, []int{}},
		{"nil slice", nil, 3, []int{}},
		{"n=0", []int{1, 2}, 0, []int{}},
		{"n=1", []int{1, 2}, 1, []int{1, 2}},
		{"n=3", []int{1, 2}, 3, []int{1, 2, 1, 2, 1, 2}},
		{"negative n", []int{1, 2}, -1, []int{}},
		{"single element", []int{42}, 4, []int{42, 42, 42, 42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Repeat(tt.input, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Repeat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEveryEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		expected bool
	}{
		{"empty", []int{}, 1, true},
		{"nil", nil, 1, true},
		{"all match", []int{5, 5, 5}, 5, true},
		{"none match", []int{1, 2, 3}, 5, false},
		{"some match", []int{5, 5, 6}, 5, false},
		{"single match", []int{5}, 5, true},
		{"single no match", []int{5}, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Every(tt.input, tt.needle)
			if result != tt.expected {
				t.Errorf("Every() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSomeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		expected bool
	}{
		{"empty", []int{}, 1, false},
		{"nil", nil, 1, false},
		{"found", []int{1, 2, 3}, 2, true},
		{"not found", []int{1, 2, 3}, 5, false},
		{"single found", []int{5}, 5, true},
		{"single not found", []int{5}, 1, false},
		{"multiple occurrences", []int{2, 2, 2}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Some(tt.input, tt.needle)
			if result != tt.expected {
				t.Errorf("Some() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNoneEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		needle   int
		expected bool
	}{
		{"empty", []int{}, 1, true},
		{"nil", nil, 1, true},
		{"none found", []int{1, 2, 3}, 5, true},
		{"found", []int{1, 2, 3}, 2, false},
		{"single not found", []int{5}, 1, true},
		{"single found", []int{5}, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := None(tt.input, tt.needle)
			if result != tt.expected {
				t.Errorf("None() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHeadTailLastInitialEdgeCases(t *testing.T) {
	t.Run("Head", func(t *testing.T) {
		// Empty
		_, err := Head([]int{})
		if err == nil {
			t.Error("Head(empty) should return error")
		}

		// Single element
		val, err := Head([]int{42})
		if err != nil {
			t.Error("Head(single) should not return error")
		}
		if val != 42 {
			t.Errorf("Head() = %v, want 42", val)
		}

		// Multiple elements
		val, err = Head([]int{1, 2, 3})
		if val != 1 {
			t.Errorf("Head() = %v, want 1", val)
		}
	})

	t.Run("Tail", func(t *testing.T) {
		// Empty
		if Tail([]int{}) != nil {
			t.Error("Tail(empty) should return nil")
		}

		// Single element
		if Tail([]int{42}) != nil {
			t.Error("Tail(single) should return nil")
		}

		// Multiple elements
		result := Tail([]int{1, 2, 3})
		if !reflect.DeepEqual(result, []int{2, 3}) {
			t.Errorf("Tail() = %v, want [2, 3]", result)
		}
	})

	t.Run("Last", func(t *testing.T) {
		// Empty
		_, err := Last([]int{})
		if err == nil {
			t.Error("Last(empty) should return error")
		}

		// Single element
		val, err := Last([]int{42})
		if err != nil {
			t.Error("Last(single) should not return error")
		}
		if val != 42 {
			t.Errorf("Last() = %v, want 42", val)
		}

		// Multiple elements
		val, err = Last([]int{1, 2, 3})
		if val != 3 {
			t.Errorf("Last() = %v, want 3", val)
		}
	})

	t.Run("Initial", func(t *testing.T) {
		// Empty
		if Initial([]int{}) != nil {
			t.Error("Initial(empty) should return nil")
		}

		// Single element
		if Initial([]int{42}) != nil {
			t.Error("Initial(single) should return nil")
		}

		// Multiple elements
		result := Initial([]int{1, 2, 3})
		if !reflect.DeepEqual(result, []int{1, 2}) {
			t.Errorf("Initial() = %v, want [1, 2]", result)
		}
	})
}
