package henry

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

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
