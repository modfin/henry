package sort

import (
	"github.com/crholm/go18exp/compare"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	in := []int{2, 3, 5, 1, 12, 3, 6, 7, 34, 123, 65, 4631, 1, 1323}
	exp := []int{1, 1, 2, 3, 3, 5, 6, 7, 12, 34, 65, 123, 1323, 4631}
	Slice(in, compare.Less[int])
	if !reflect.DeepEqual(exp, in) {
		t.Log("Expected", exp)
		t.Log("     got", in)
		t.Fail()
	}
}

func TestStableSlice(t *testing.T) {
	in := []int{2, 3, 5, 1, 12, 3, 6, 7, 34, 123, 65, 4631, 1, 1323}
	exp := []int{1, 1, 2, 3, 3, 5, 6, 7, 12, 34, 65, 123, 1323, 4631}
	StableSlice(in, compare.Less[int])
	if !reflect.DeepEqual(exp, in) {
		t.Log("Expected", exp)
		t.Log("     got", in)
		t.Fail()
	}
}

func TestSliceReverse(t *testing.T) {
	in := []int{2, 3, 5, 1, 12, 3, 6, 7, 34, 123, 65, 4631, 1, 1323}
	exp := []int{4631, 1323, 123, 65, 34, 12, 7, 6, 5, 3, 3, 2, 1, 1}
	Slice(in, compare.Reverse(compare.Less[int]))
	if !reflect.DeepEqual(exp, in) {
		t.Log("Expected", exp)
		t.Log("     got", in)
		t.Fail()
	}
}

func TestIsSorted(t *testing.T) {
	in := []int{1, 1, 2, 3, 3, 5, 6, 7, 12, 34, 65, 123, 1323, 4631}
	if !IsSorted(in, compare.Less[int]) {
		t.Log("Expected to be sorted")
		t.Fail()
	}
	if IsSorted(in, compare.Reverse(compare.Less[int])) {
		t.Log("Expected to be sorted")
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	in := []int{1, 1, 2, 3, 4, 5, 6, 7, 12, 34, 65, 123, 1323, 4631}
	i, e := Search(in, func(e int) bool {
		return 5 <= e
	})
	if i != 5 {
		t.Log("Expected index 5, got", i)
		t.Fail()
	}
	if e != 5 {
		t.Log("Expected value 5, got", i)
		t.Fail()
	}
}
