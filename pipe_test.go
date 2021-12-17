package henry

import (
	"reflect"
	"testing"
)

func TestPipe_Concat(t *testing.T) {
	var res = PipeOf([]int{1, 2, 3}).Concat([]int{4, 5, 6}).Slice()
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}
