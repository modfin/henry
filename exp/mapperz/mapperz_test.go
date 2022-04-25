package mapperz

import (
	"github.com/modfin/henry/slicez"
	"testing"
)

func TestIndexed(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	indexed := Indexed(slice)
	for i, el := range indexed {
		if el.Element != el.Index && i != el.Index {
			t.Log("expected: ", el.Index, " ", i, " got ", el.Element)
		}
	}

	slice = []int{1, 2, 3, 4, 5}
	everyOtherIndexed := slicez.Filter(Indexed(slice), func(e IElm[int]) bool { return e.Index%2 == 0 })
	everyOther := Elements(everyOtherIndexed)

	if !slicez.Equal(everyOther, []int{1, 3, 5}) {
		t.Log("Expected []int{1,3,5}, got ", everyOther)
	}
}
