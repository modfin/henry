package set

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	if s.Exists(2) {
		t.Log("expected not containing 2")
		t.Fail()
	}
	s.Add(2)
	if !s.Exists(2) {
		t.Log("expected to containing 2")
		t.Fail()
	}
	s.Delete(2)
	if s.Exists(2) {
		t.Log("expected not containing 2")
		t.Fail()
	}

	func() {
		defer func() {
			r := recover()
			if r == nil {
				t.Log("expected a panic")
				t.Fail()
			}
		}()
		s.Immutable()
		s.Add(10)
	}()

}
