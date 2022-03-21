package chanz

import (
	"fmt"
	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/slicez"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestGenerated(t *testing.T) {
	var i int
	for val := range GenerateN[int](0, 0, 1, 2, 3, 4) {
		if val != i {
			t.Logf("expected, %v, but got %v", i, val)
			t.Fail()
		}
		i++
	}
}

func TestPeek(t *testing.T) {
	type wrap struct {
		A int
	}
	in := []*wrap{{1}, {2}, {3}}
	exp := []*wrap{{1}, {4}, {9}}
	generated := Generate[*wrap](in...)
	peeked := Peek[*wrap](generated, func(a *wrap) {
		a.A = a.A * a.A
	})

	res := Collect(peeked)

	if !slicez.EqualFunc(exp, res, func(e1 *wrap, e2 *wrap) bool {
		return e1.A == e2.A
	}) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestMap0(t *testing.T) {

	res := []string{}
	exp := []string{"1", "2", "3", "4"}
	generated := GenerateN[int](0, 1, 2, 3, 4)
	mapped := Map[int, string](generated, func(a int) string {
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

func TestFlatten(t *testing.T) {

	in := [][]int{{1, 2}, {3}, {}, {4}}
	exp := []int{1, 2, 3, 4}
	generated := Generate[[]int](in...)
	flatten := Flatten[int](generated)

	res := Collect(flatten)

	if !slicez.Equal(exp, res) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}

}

func TestMerge(t *testing.T) {

	var res []int
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8}

	c1 := GenerateN(0, 1, 2, 3, 4, 5)
	c2 := GenerateN(0, 6, 7, 8)
	merged := MergeN(0, c1, c2)

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

func TestFilter(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	exp := []int{2, 4, 6, 8}

	s := GenerateN(0, in...)
	f := FilterN(0, s, func(a int) bool {
		return a%2 == 0
	})

	res := Collect(f)

	if !slicez.Equal(exp, res) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestFanOut(t *testing.T) {

	c := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	expansion := 11
	outs := FanOut(c, expansion)

	var wg sync.WaitGroup

	var mu sync.Mutex

	var res [][]int

	wg.Add(expansion)
	for _, o := range outs {
		o := o
		go func() {
			nums := Collect(o)
			mu.Lock()
			res = append(res, nums)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if len(res) != expansion {
		t.Logf("expected, %v, but got %v", expansion, len(res))
		t.Fail()
	}

	for i, r := range res {
		if !slicez.Equal(r, exp) {
			t.Logf("expected, %v, but got %v, for i %d", exp, r, i)
			t.Fail()
		}
	}

}

func TestCompact(t *testing.T) {

	c := Compact(Generate(1, 2, 3, 4, 5, 5, 5, 6, 7, 7, 7, 8, 8, 8, 9), compare.Equal[int])

	var res []int
	for v := range c {
		res = append(res, v)
	}
	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}

}

func TestPartition(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9)

	even, odd := Partition(c, func(i int) bool {
		return i%2 == 0
	})

	var resEven []int
	go func() {
		resEven = Collect(even)
	}()
	resOdd := Collect(odd)

	expEven := []int{2, 4, 6, 8}
	if !slicez.Equal(expEven, resEven) {
		t.Logf("expected, %v, but got %v", expEven, resEven)
		t.Fail()
	}

	expOdd := []int{1, 3, 5, 7, 9}
	if !slicez.Equal(expOdd, resOdd) {
		t.Logf("expected, %v, but got %v", expOdd, resOdd)
		t.Fail()
	}
}

func TestTakeWhile(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	taker := TakeWhile(c, func(a int) bool {
		return a < 5
	})

	res := Collect(taker)

	exp := []int{1, 2, 3, 4}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestDropWhile(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	dropper := DropWhile(c, func(a int) bool {
		return a < 5
	})

	res := Collect(dropper)

	exp := []int{5, 6, 7, 8, 9}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestTake(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	taker := Take(c, 3)

	res := Collect(taker)
	exp := []int{1, 2, 3}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}

	exp = []int{4, 5, 6, 7, 8, 9}

	rest := Collect(c)
	if !slicez.Equal(rest, exp) {
		t.Logf("expected, %v, but got %v", exp, rest)
		t.Fail()
	}
}

func TestDrop(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	dropper := Drop(c, 3)

	res := Collect(dropper)

	exp := []int{4, 5, 6, 7, 8, 9}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestDropAll(t *testing.T) {
	c := Generate(1, 2, 3, 4)
	DropAll(c, false)
	_, ok := <-c
	if ok {
		t.Log("expected channel to be closed, but was open")
		t.Fail()
	}
}

func TestZip1(t *testing.T) {
	ac := Generate(1, 2, 3, 4)
	bc := Generate("a", "b", "c")
	z := Zip(ac, bc, func(a int, b string) string {
		return fmt.Sprintf("%d%s", a, b)
	})
	res := Collect(z)
	exp := []string{"1a", "2b", "3c"}
	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}
func TestZip2(t *testing.T) {
	ac := Generate(1, 2, 3)
	bc := Generate("a", "b", "c", "something")
	z := Zip(ac, bc, func(a int, b string) string {
		return fmt.Sprintf("%d%s", a, b)
	})
	res := Collect(z)
	exp := []string{"1a", "2b", "3c"}
	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestUnzip(t *testing.T) {
	z := Generate("a1", "b2", "c3")
	sc, ic := Unzip(z, func(c string) (string, int) {
		i, _ := strconv.ParseInt(string(c[1]), 10, 64)
		return string(c[0]), int(i)
	})

	var ints []int
	var strings []string

	go func() {
		strings = Collect(sc)
	}()
	ints = Collect(ic)

	iexp := []int{1, 2, 3}

	if !slicez.Equal(ints, iexp) {
		t.Logf("expected, %v, but got %v", iexp, ints)
		t.Fail()
	}

	sexp := []string{"a", "b", "c"}
	if !slicez.Equal(strings, sexp) {
		t.Logf("expected, %v, but got %v", sexp, strings)
		t.Fail()
	}

}

func TestSomeDone(t *testing.T) {

	num := 10
	dones := make([]chan interface{}, num)
	for i := range dones {
		dones[i] = make(chan interface{})
	}

	done := SomeDone(Readers(dones...)...)

	close(dones[rand.Intn(num)])

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Log("expected done to be closed by now")
		t.Fail()
	}

}

func TestEveryDone(t *testing.T) {

	num := 10
	dones := make([]chan interface{}, num)
	for i := range dones {
		dones[i] = make(chan interface{})
	}

	done := EveryDone(Readers(dones...)...)

	for _, d := range slicez.Shuffle(dones) {
		close(d)
		select {
		case <-done:
			t.Log("did not expect early close")
			t.Fail()
		default:
		}
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Log("expected done to be closed by now")
		t.Fail()
	}

}
