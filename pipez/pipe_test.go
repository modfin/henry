package pipez

import (
	"reflect"
	"testing"
)

func TestPipe_Concat(t *testing.T) {
	var res = Of([]int{1, 2, 3}).Concat([]int{4, 5, 6}).Slice()
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}

func TestPipe_Concat_withoutWrap(t *testing.T) {
	var firstSlice Pipe[int] = []int{1, 2, 3}
	var secondSlice Pipe[int] = []int{4, 5, 6}
	var res []int = firstSlice.Concat(secondSlice)
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}

func TestOf(t *testing.T) {
	input := []int{1, 2, 3}
	p := Of(input)
	if !reflect.DeepEqual(p.Slice(), input) {
		t.Errorf("Of() = %v, want %v", p.Slice(), input)
	}
}

func TestPipe_Slice(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	if !reflect.DeepEqual(p.Slice(), []int{1, 2, 3}) {
		t.Errorf("Slice() = %v, want %v", p.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_Peek(t *testing.T) {
	var sum int
	p := Pipe[int]{1, 2, 3}
	result := p.Peek(func(a int) {
		sum += a
	})
	if sum != 6 {
		t.Errorf("Peek sum = %d, want 6", sum)
	}
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("Peek() returned %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_Tail(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4}
	result := p.Tail()
	if !reflect.DeepEqual(result.Slice(), []int{2, 3, 4}) {
		t.Errorf("Tail() = %v, want %v", result.Slice(), []int{2, 3, 4})
	}
	// Empty pipe
	empty := Pipe[int]{}
	if empty.Tail() != nil {
		t.Error("Tail() on empty pipe should return nil")
	}
}

func TestPipe_Head(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	head, err := p.Head()
	if err != nil {
		t.Errorf("Head() error = %v, want nil", err)
	}
	if head != 1 {
		t.Errorf("Head() = %d, want 1", head)
	}
	// Empty pipe
	empty := Pipe[int]{}
	_, err = empty.Head()
	if err == nil {
		t.Error("Head() on empty pipe should return error")
	}
}

func TestPipe_Last(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	last, err := p.Last()
	if err != nil {
		t.Errorf("Last() error = %v, want nil", err)
	}
	if last != 3 {
		t.Errorf("Last() = %d, want 3", last)
	}
	// Empty pipe
	empty := Pipe[int]{}
	_, err = empty.Last()
	if err == nil {
		t.Error("Last() on empty pipe should return error")
	}
}

func TestPipe_Initial(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4}
	result := p.Initial()
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("Initial() = %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_Reverse(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	result := p.Reverse()
	if !reflect.DeepEqual(result.Slice(), []int{3, 2, 1}) {
		t.Errorf("Reverse() = %v, want %v", result.Slice(), []int{3, 2, 1})
	}
}

func TestPipe_Nth(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	if p.Nth(0) != 1 {
		t.Errorf("Nth(0) = %d, want 1", p.Nth(0))
	}
	if p.Nth(1) != 2 {
		t.Errorf("Nth(1) = %d, want 2", p.Nth(1))
	}
	if p.Nth(-1) != 3 {
		t.Errorf("Nth(-1) = %d, want 3", p.Nth(-1))
	}
}

func TestPipe_Take(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.Take(3)
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("Take(3) = %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_TakeRight(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.TakeRight(3)
	if !reflect.DeepEqual(result.Slice(), []int{3, 4, 5}) {
		t.Errorf("TakeRight(3) = %v, want %v", result.Slice(), []int{3, 4, 5})
	}
}

func TestPipe_TakeWhile(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.TakeWhile(func(a int) bool {
		return a < 4
	})
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("TakeWhile(<4) = %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_TakeRightWhile(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.TakeRightWhile(func(a int) bool {
		return a > 2
	})
	if !reflect.DeepEqual(result.Slice(), []int{3, 4, 5}) {
		t.Errorf("TakeRightWhile(>2) = %v, want %v", result.Slice(), []int{3, 4, 5})
	}
}

func TestPipe_Drop(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.Drop(2)
	if !reflect.DeepEqual(result.Slice(), []int{3, 4, 5}) {
		t.Errorf("Drop(2) = %v, want %v", result.Slice(), []int{3, 4, 5})
	}
}

func TestPipe_DropRight(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.DropRight(2)
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("DropRight(2) = %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_DropWhile(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.DropWhile(func(a int) bool {
		return a < 3
	})
	if !reflect.DeepEqual(result.Slice(), []int{3, 4, 5}) {
		t.Errorf("DropWhile(<3) = %v, want %v", result.Slice(), []int{3, 4, 5})
	}
}

func TestPipe_DropRightWhile(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.DropRightWhile(func(a int) bool {
		return a > 3
	})
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("DropRightWhile(>3) = %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_Filter(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.Filter(func(a int) bool {
		return a%2 == 1
	})
	if !reflect.DeepEqual(result.Slice(), []int{1, 3, 5}) {
		t.Errorf("Filter(odd) = %v, want %v", result.Slice(), []int{1, 3, 5})
	}
}

func TestPipe_Reject(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.Reject(func(a int) bool {
		return a%2 == 0
	})
	if !reflect.DeepEqual(result.Slice(), []int{1, 3, 5}) {
		t.Errorf("Reject(even) = %v, want %v", result.Slice(), []int{1, 3, 5})
	}
}

func TestPipe_Map(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	result := p.Map(func(a int) int {
		return a * 2
	})
	if !reflect.DeepEqual(result.Slice(), []int{2, 4, 6}) {
		t.Errorf("Map(*2) = %v, want %v", result.Slice(), []int{2, 4, 6})
	}
}

func TestPipe_Fold(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4}
	result := p.Fold(func(acc, val int) int {
		return acc + val
	}, 0)
	if result != 10 {
		t.Errorf("Fold(sum) = %d, want 10", result)
	}
}

func TestPipe_FoldRight(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	result := p.FoldRight(func(acc, val int) int {
		return acc*10 + val
	}, 0)
	// 0*10+3=3, 3*10+2=32, 32*10+1=321
	if result != 321 {
		t.Errorf("FoldRight = %d, want 321", result)
	}
}

func TestPipe_Every(t *testing.T) {
	p := Pipe[int]{2, 4, 6}
	if !p.Every(func(a int) bool { return a%2 == 0 }) {
		t.Error("Every(even) should return true for [2,4,6]")
	}
	p2 := Pipe[int]{1, 2, 3}
	if p2.Every(func(a int) bool { return a%2 == 0 }) {
		t.Error("Every(even) should return false for [1,2,3]")
	}
}

func TestPipe_Some(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	if !p.Some(func(a int) bool { return a%2 == 0 }) {
		t.Error("Some(even) should return true for [1,2,3]")
	}
	p2 := Pipe[int]{1, 3, 5}
	if p2.Some(func(a int) bool { return a%2 == 0 }) {
		t.Error("Some(even) should return false for [1,3,5]")
	}
}

func TestPipe_None(t *testing.T) {
	p := Pipe[int]{1, 3, 5}
	if !p.None(func(a int) bool { return a%2 == 0 }) {
		t.Error("None(even) should return true for [1,3,5]")
	}
	p2 := Pipe[int]{1, 2, 3}
	if p2.None(func(a int) bool { return a%2 == 0 }) {
		t.Error("None(even) should return false for [1,2,3]")
	}
}

func TestPipe_Partition(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	satisfied, notSatisfied := p.Partition(func(a int) bool {
		return a%2 == 0
	})
	if !reflect.DeepEqual(satisfied, []int{2, 4}) {
		t.Errorf("Partition satisfied = %v, want %v", satisfied, []int{2, 4})
	}
	if !reflect.DeepEqual(notSatisfied, []int{1, 3, 5}) {
		t.Errorf("Partition notSatisfied = %v, want %v", notSatisfied, []int{1, 3, 5})
	}
}

func TestPipe_Sample(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.Sample(3)
	if len(result) != 3 {
		t.Errorf("Sample(3) length = %d, want 3", len(result))
	}
	// Check all elements are from original slice
	for _, v := range result {
		if v < 1 || v > 5 {
			t.Errorf("Sample returned invalid value: %d", v)
		}
	}
}

func TestPipe_Shuffle(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	result := p.Shuffle()
	// Just verify length and all elements present
	if len(result) != 5 {
		t.Errorf("Shuffle() length = %d, want 5", len(result))
	}
	// Check all elements from original are present
	original := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	for _, v := range result {
		if !original[v] {
			t.Errorf("Shuffle returned invalid value: %d", v)
		}
	}
}

func TestPipe_SortFunc(t *testing.T) {
	p := Pipe[int]{3, 1, 4, 1, 5}
	result := p.SortFunc(func(a, b int) bool { return a < b })
	if !reflect.DeepEqual(result.Slice(), []int{1, 1, 3, 4, 5}) {
		t.Errorf("SortFunc() = %v, want %v", result.Slice(), []int{1, 1, 3, 4, 5})
	}
}

func TestPipe_Compact(t *testing.T) {
	p := Pipe[int]{1, 1, 2, 2, 3, 3, 3}
	result := p.Compact(func(a, b int) bool { return a == b })
	if !reflect.DeepEqual(result.Slice(), []int{1, 2, 3}) {
		t.Errorf("Compact() = %v, want %v", result.Slice(), []int{1, 2, 3})
	}
}

func TestPipe_Count(t *testing.T) {
	p := Pipe[int]{1, 2, 3, 4, 5}
	if p.Count() != 5 {
		t.Errorf("Count() = %d, want 5", p.Count())
	}
}

func TestPipe_Zip(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	result := p.Zip([]int{10, 20, 30}, func(a, b int) int { return a + b })
	if !reflect.DeepEqual(result.Slice(), []int{11, 22, 33}) {
		t.Errorf("Zip() = %v, want %v", result.Slice(), []int{11, 22, 33})
	}
}

func TestPipe_Unzip(t *testing.T) {
	p := Pipe[string]{"a1", "b2", "c3"}
	a, b := p.Unzip(func(s string) (string, string) {
		return string(s[0]), string(s[1])
	})
	if !reflect.DeepEqual(a, []string{"a", "b", "c"}) {
		t.Errorf("Unzip first = %v, want %v", a, []string{"a", "b", "c"})
	}
	if !reflect.DeepEqual(b, []string{"1", "2", "3"}) {
		t.Errorf("Unzip second = %v, want %v", b, []string{"1", "2", "3"})
	}
}

func TestPipe_Interleave(t *testing.T) {
	p := Pipe[int]{1, 2, 3}
	result := p.Interleave([]int{4, 5}, []int{6, 7, 8})
	if !reflect.DeepEqual(result.Slice(), []int{1, 4, 6, 2, 5, 7, 3, 8}) {
		t.Errorf("Interleave() = %v, want %v", result.Slice(), []int{1, 4, 6, 2, 5, 7, 3, 8})
	}
}
