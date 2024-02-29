package mapz

import (
	"fmt"
	"github.com/modfin/henry/slicez"
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3}
	m2 := map[int]int{1: 1, 3: 3}
	if !Equal(m1, m2) {
		t.Log("Expected maps to be equal")
		t.Fail()
	}

	m2[4] = 4
	if Equal(m1, m2) {
		t.Log("Expected maps not to be equal")
		t.Fail()
	}

	m1[4] = 3
	if Equal(m1, m2) {
		t.Log("Expected maps not to be equal")
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3}
	m2 := map[int]int{}
	Clear(m1)
	if !Equal(m1, m2) {
		t.Log("Expected maps be equal")
		t.Fail()
	}
}

func TestClone(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3}
	m2 := Clone(m1)
	if &m1 == &m2 {
		t.Log("should not be the same")
		t.Fail()
	}
	if !Equal(m1, m2) {
		t.Log("Expected maps be equal")
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	src := map[int]int{1: 1, 3: 3}
	dst := map[int]int{3: 33, 4: 4}
	exp := map[int]int{1: 1, 3: 3, 4: 4}
	Copy(dst, src)
	if !Equal(exp, dst) {
		t.Log("Expected maps be equal")
		t.Fail()
	}
}

func TestValues(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	v := Values(m)

	if len(v) != 2 {
		t.Log("Expected size of 2")
		t.Fail()
	}

	if !slicez.Contains(v, 33) {
		t.Log("Expected value 11")
		t.Fail()
	}
	if !slicez.Contains(v, 11) {
		t.Log("Expected value 31")
		t.Fail()
	}
}

func TestMerge(t *testing.T) {
	m1 := map[int]int{1: 11, 3: 33}
	m2 := map[int]int{2: 22, 4: 44}
	exp := map[int]int{1: 11, 3: 33, 2: 22, 4: 44}
	m := Merge(m1, m2)

	if !Equal(m, exp) {
		t.Log("Expected equality")
		t.Fail()
	}

	m1[4] = 444
	m = Merge(m1, m2)
	if !Equal(m, exp) {
		t.Log("Expected equality")
		t.Fail()
	}

	exp[4] = 444
	m = Merge(m2, m1)
	if !Equal(m, exp) {
		t.Log("Expected equality")
		t.Fail()
	}

}

func TestRemap(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	exp := map[string]int{"1": 110, "3": 330}

	res := Remap(m, func(k int, v int) (string, int) {
		return fmt.Sprint(k), v * 10
	})

	if !Equal(res, exp) {
		t.Log("Expected equality")
		t.Fail()
	}
}

func TestKeys(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	v := Keys(m)

	if len(v) != 2 {
		t.Log("Expected size of 2")
		t.Fail()
	}

	if !slicez.Contains(v, 1) {
		t.Log("Expected value 1")
		t.Fail()
	}
	if !slicez.Contains(v, 3) {
		t.Log("Expected value 3")
		t.Fail()
	}
}

func TestDeleteValue(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	DeleteValues(m, 33)

	if len(m) != 1 {
		t.Log("Expected size of 1")
		t.Fail()
	}

	if m[1] != 11 {
		t.Log("Expected to contain key 1")
		t.Fail()
	}
}

func TestDeleteFunc(t *testing.T) {
	m := map[int]int{1: 11, 3: 33}
	Delete(m, func(k int, v int) bool {
		return k == 3
	})

	if len(m) != 1 {
		t.Log("Expected size of 1")
		t.Fail()
	}

	if m[1] != 11 {
		t.Log("Expected to contain key 1")
		t.Fail()
	}
}

func TestPickByKeys(t *testing.T) {
	m := FilterByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	exp := map[string]int{"foo": 1, "baz": 3}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
func TestPickByValues(t *testing.T) {
	m := FilterByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{23, 2})
	exp := map[string]int{"bar": 2}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
func TestOmitByKeys(t *testing.T) {
	m := RejectByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	exp := map[string]int{"bar": 2}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
func TestOmitByValues(t *testing.T) {
	m := RejectByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{23, 2})
	exp := map[string]int{"foo": 1, "baz": 3}
	if !reflect.DeepEqual(m, exp) {
		t.Log("Expected", exp, "got", m)
		t.Fail()
	}
}
