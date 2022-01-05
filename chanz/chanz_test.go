package chanz

import (
	"fmt"
	"github.com/modfin/go18exp/slicez"
	"testing"
)

func TestGenerated(t *testing.T) {
	var i int
	for val := range Generate[int](0, 0, 1, 2, 3, 4) {
		if val != i {
			t.Logf("expected, %v, but got %v", i, val)
			t.Fail()
		}
		i++
	}
}

func TestMap0(t *testing.T) {

	res := []string{}
	exp := []string{"1", "2", "3", "4"}
	generated := Generate[int](0, 1, 2, 3, 4)
	mapped := Map0[int, string](generated, func(a int) string {
		return fmt.Sprintf("%d", a)
	})

	for s := range mapped {
		res = append(res, s)
	}

	if !slicez.Equal(exp, res) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}

}

func TestMerge(t *testing.T) {

	var res []int
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8}

	c1 := Generate(0, 1, 2, 3, 4, 5)
	c2 := Generate(0, 6, 7, 8)
	merged := Merge(0, c1, c2)

	for v := range merged {
		res = append(res, v)
	}

	for _, v := range exp {
		if !slicez.Contains(res, v) {
			t.Logf("expected result, %v, to contain %v", exp, v)
			t.Fail()
		}
	}

}
