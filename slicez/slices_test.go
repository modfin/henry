package slicez

import (
	"fmt"
	"github.com/modfin/henry/compare"
	"math"
	"reflect"
	"strconv"
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
