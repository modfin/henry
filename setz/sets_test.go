package setz

import (
	"testing"

	"github.com/modfin/henry/slicez"
)

func TestSet_Of(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 5}
	set := Of(a)
	if len(set) != 5 {
		t.Fatalf("Set should have length 5, got %d", len(set))
	}
	for _, v := range slicez.Uniq(a) {
		if !Has(set, v) {
			t.Fatalf("Set should contain %d", v)
		}
	}
}

func TestSet_Has(t *testing.T) {
	set := Of([]int{1, 2, 3, 4, 5})
	if !Has(set, 3) {
		t.Fatalf("Set should contain 3")
	}
	if Has(set, 6) {
		t.Fatalf("Set should not contain 6")
	}
}

func TestSet_Add(t *testing.T) {
	set := Of([]int{1, 2, 3, 4, 5})
	set = Add(set, 6)
	if !Has(set, 6) {
		t.Fatalf("Set should contain 6")
	}
}

func TestSet_Remove(t *testing.T) {
	set := Of([]int{1, 2, 3, 4, 5})
	set = Remove(set, 3)
	if Has(set, 3) {
		t.Fatalf("Set should not contain 3")
	}
}

func TestSet_Union(t *testing.T) {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{4, 5, 6, 7, 8}
	set1 := Of(a1)
	set2 := Of(a2)
	set := Union(set1, set2)
	if len(set) != 8 {
		t.Fatalf("Set should have length 8, got %d", len(set))
	}
	for _, v := range slicez.Uniq(append(slicez.Uniq(a1), slicez.Uniq(a2)...)) {
		if !Has(set, v) {
			t.Fatalf("Set should contain %d", v)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	set1 := Of([]int{1, 2, 3, 4, 5})
	set2 := Of([]int{4, 5, 6, 7, 8})
	set := Intersection(set1, set2)
	if len(set) != 2 {
		t.Fatalf("Set should have length 2, got %d", len(set))
	}
	for _, v := range []int{4, 5} {
		if !Has(set, v) {
			t.Fatalf("Set should contain %d", v)
		}
	}
}

func TestSet_Difference(t *testing.T) {
	set1 := Of([]int{1, 2, 3, 4, 5})
	set2 := Of([]int{4, 5, 6, 7, 8})
	diff12 := Difference(set1, set2)
	if len(diff12) != 3 {
		t.Fatalf("Set should have length 3, got %d", len(diff12))
	}
	for _, v := range []int{1, 2, 3} {
		if !Has(diff12, v) {
			t.Fatalf("Set should contain %d", v)
		}
	}

	diff21 := Difference(set2, set1)
	if len(diff21) != 3 {
		t.Fatalf("Set should have length 3, got %d", len(diff21))
	}
	for _, v := range []int{6, 7, 8} {
		if !Has(diff21, v) {
			t.Fatalf("Set should contain %d", v)
		}
	}
}
