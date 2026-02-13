package slicez

import (
	"fmt"
	"github.com/modfin/henry/compare"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
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
