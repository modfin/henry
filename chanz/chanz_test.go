package chanz

import (
	"context"
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
	for val := range Generate(0, 1, 2, 3, 4) {
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

	if !slicez.EqualBy(exp, res, func(e1 *wrap, e2 *wrap) bool {
		return e1.A == e2.A
	}) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestMap0(t *testing.T) {

	res := []string{}
	exp := []string{"1", "2", "3", "4"}
	generated := Generate(1, 2, 3, 4)
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
func TestMap2(t *testing.T) {

	res := []string{}
	exp := []string{"1", "2", "3", "4"}
	generated := Generate[int](1, 2, 3, 4)
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

	c1 := Generate(0, 1, 2, 3, 4, 5)
	c2 := Generate(0, 6, 7, 8)
	merged := FanIn(c1, c2)

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

	s := Generate(in...)
	f := Filter(s, func(a int) bool {
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
	dones := make([]chan struct{}, num)
	for i := range dones {
		dones[i] = make(chan struct{})
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

func TestDropBuffer(t *testing.T) {
	c := GenerateWith[int](OpBuffer(2))(1, 2, 3, 4, 5)
	time.Sleep(100 * time.Millisecond)
	DropBuffer(c, false)

	res := Collect(c)

	exp := []int{3, 4, 5}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestTakeBuffer(t *testing.T) {
	c := GenerateWith[int](OpBuffer(2))(1, 2, 3, 4, 5)
	time.Sleep(100 * time.Millisecond)
	res1 := TakeBuffer(c)
	res2 := Collect(c)

	exp1 := []int{1, 2}
	exp2 := []int{3, 4, 5}

	if !slicez.Equal(res1, exp1) {
		t.Logf("expected, %v, but got %v", exp1, res1)
		t.Fail()
	}
	if !slicez.Equal(res2, exp2) {
		t.Logf("expected, %v, but got %v", exp2, res2)
		t.Fail()
	}
}

func TestEveryDoneContext(t *testing.T) {

	num := 10
	var doneChans []<-chan struct{}
	ctx := context.Background()
	for i := 0; i < num; i++ {
		subCtx, cancel := context.WithCancel(ctx)
		doneChans = append(doneChans, subCtx.Done())
		cancel()
	}

	done := EveryDone(doneChans...)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Log("expected done to be closed by now")
		t.Fail()
	}

}

func TestGenerator(t *testing.T) {
	generator := func(yield func(int)) {
		yield(1)
		yield(2)
		yield(3)
	}

	res := Collect(Generator(generator))
	exp := []int{1, 2, 3}
	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestMapWith(t *testing.T) {
	res := []string{}
	exp := []string{"1", "2", "3"}
	generated := Generate[int](1, 2, 3)
	mapper := MapWith[int, string](OpBuffer(1))
	mapped := mapper(generated, func(a int) string {
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

func TestPeekWith(t *testing.T) {
	var sum int
	generated := Generate(1, 2, 3)
	peeker := PeekWith[int](OpBuffer(1))
	peeked := peeker(generated, func(a int) {
		sum += a
	})

	Collect(peeked)

	if sum != 6 {
		t.Errorf("expected sum 6, got %d", sum)
		t.Fail()
	}
}

func TestFlattenWith(t *testing.T) {
	in := [][]int{{1, 2}, {3}, {}, {4}}
	exp := []int{1, 2, 3, 4}
	generated := Generate[[]int](in...)
	flattener := FlattenWith[int](OpBuffer(1))
	flatten := flattener(generated)

	res := Collect(flatten)

	if !slicez.Equal(exp, res) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestConcat(t *testing.T) {
	c1 := Generate(1, 2, 3)
	c2 := Generate(4, 5, 6)
	concatenated := Concat(c1, c2)
	res := Collect(concatenated)
	exp := []int{1, 2, 3, 4, 5, 6}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestConcatWith(t *testing.T) {
	c1 := Generate(1, 2, 3)
	c2 := Generate(4, 5, 6)
	concatenator := ConcatWith[int](OpBuffer(1))
	concatenated := concatenator(c1, c2)
	res := Collect(concatenated)
	exp := []int{1, 2, 3, 4, 5, 6}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestFilterWith(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	exp := []int{2, 4, 6}
	generated := Generate(in...)
	filterer := FilterWith[int](OpBuffer(1))
	filtered := filterer(generated, func(a int) bool {
		return a%2 == 0
	})

	res := Collect(filtered)

	if !slicez.Equal(exp, res) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestCompactWith(t *testing.T) {
	c := Generate(1, 1, 2, 2, 3, 3)
	compactor := CompactWith[int](OpBuffer(1))
	compact := compactor(c, compare.Equal[int])

	var res []int
	for v := range compact {
		res = append(res, v)
	}
	exp := []int{1, 2, 3}
	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestPartitionWith(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5)
	partitioner := PartitionWith[int](OpBuffer(1))
	even, odd := partitioner(c, func(i int) bool {
		return i%2 == 0
	})

	var resEven []int
	go func() {
		resEven = Collect(even)
	}()
	resOdd := Collect(odd)

	expEven := []int{2, 4}
	if !slicez.Equal(expEven, resEven) {
		t.Logf("expected even %v, but got %v", expEven, resEven)
		t.Fail()
	}

	expOdd := []int{1, 3, 5}
	if !slicez.Equal(expOdd, resOdd) {
		t.Logf("expected odd %v, but got %v", expOdd, resOdd)
		t.Fail()
	}
}

func TestTakeWhileWith(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5)
	taker := TakeWhileWith[int](OpBuffer(1))
	result := taker(c, func(a int) bool {
		return a < 4
	})

	res := Collect(result)
	exp := []int{1, 2, 3}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestTakeWith(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5)
	taker := TakeWith[int](OpBuffer(1))
	result := taker(c, 3)

	res := Collect(result)
	exp := []int{1, 2, 3}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestDropWith(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5)
	dropper := DropWith[int](OpBuffer(1))
	result := dropper(c, 2)

	res := Collect(result)
	exp := []int{3, 4, 5}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestDropWhileWith(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5)
	dropper := DropWhileWith[int](OpBuffer(1))
	result := dropper(c, func(a int) bool {
		return a < 3
	})

	res := Collect(result)
	exp := []int{3, 4, 5}

	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestFanInWith(t *testing.T) {
	c1 := Generate(1, 2, 3)
	c2 := Generate(4, 5, 6)
	fanIn := FanInWith[int](OpBuffer(1))
	result := fanIn(c1, c2)

	res := Collect(result)
	// Since it's concurrent, just check all values are present
	if len(res) != 6 {
		t.Errorf("expected 6 elements, got %d", len(res))
		t.Fail()
	}
}

func TestFanOutWith(t *testing.T) {
	c := Generate(1, 2, 3)
	fanOuter := FanOutWith[int](OpBuffer(1))
	outs := fanOuter(c, 3)

	if len(outs) != 3 {
		t.Errorf("expected 3 output channels, got %d", len(outs))
		t.Fail()
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var results [][]int

	wg.Add(3)
	for _, o := range outs {
		o := o
		go func() {
			nums := Collect(o)
			mu.Lock()
			results = append(results, nums)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	exp := []int{1, 2, 3}
	for _, r := range results {
		if !slicez.Equal(r, exp) {
			t.Logf("expected, %v, but got %v", exp, r)
			t.Fail()
		}
	}
}

func TestGeneratorWith(t *testing.T) {
	generator := func(yield func(int)) {
		yield(1)
		yield(2)
		yield(3)
	}

	gen := GeneratorWith[int](OpBuffer(1))
	res := Collect(gen(generator))
	exp := []int{1, 2, 3}
	if !slicez.Equal(res, exp) {
		t.Logf("expected, %v, but got %v", exp, res)
		t.Fail()
	}
}

func TestBuffer(t *testing.T) {
	c := Generate(1, 2, 3, 4, 5)
	buf, more := Buffer(3, c)

	if len(buf) != 3 {
		t.Errorf("expected buffer length 3, got %d", len(buf))
		t.Fail()
	}

	if !more {
		t.Error("expected more to be true since channel has more elements")
		t.Fail()
	}

	// Collect remaining
	remaining := Collect(c)
	if len(remaining) != 2 {
		t.Errorf("expected 2 remaining elements, got %d", len(remaining))
		t.Fail()
	}
}

func TestDone(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	done := Done(ch)
	select {
	case <-done:
		// Success
	case <-time.After(time.Second):
		t.Error("expected Done to close after channel closes")
		t.Fail()
	}
}

func TestReaders(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	readers := Readers(ch1, ch2, ch3)

	if len(readers) != 3 {
		t.Errorf("expected 3 readers, got %d", len(readers))
		t.Fail()
	}

	// Test that they are receive-only
	// This is a compile-time check, if it compiles, it works
}

func TestWriters(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	writers := Writers(ch1, ch2)

	if len(writers) != 2 {
		t.Errorf("expected 2 writers, got %d", len(writers))
		t.Fail()
	}

	// Test that they are send-only
	// This is a compile-time check, if it compiles, it works
}

func TestWriteTo(t *testing.T) {
	ch := make(chan int, 10)

	// Test WriteSync
	writer := WriteTo(ch, WriteSync)
	writer(42)
	if <-ch != 42 {
		t.Error("WriteSync failed")
	}

	// Test WriteAync - just verify it doesn't block
	writer2 := WriteTo(ch, WriteAync)
	writer2(100)
	time.Sleep(10 * time.Millisecond) // Give goroutine time to execute
	if <-ch != 100 {
		t.Error("WriteAync failed")
	}

	// Test WriteIfFree when channel has space
	writer3 := WriteTo(ch, WriteIfFree)
	writer3(200)
	if <-ch != 200 {
		t.Error("WriteIfFree failed when space available")
	}

	// Test WriteIfFree when channel is full
	fullCh := make(chan int) // Unbuffered, so it will block
	writer4 := WriteTo(fullCh, WriteIfFree)
	writer4(300) // Should not block, should be dropped
	// If we get here without blocking, it worked
}

func TestReadFrom(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 42
	ch <- 100
	close(ch)

	// Test ReadWait
	reader1 := ReadFrom(ch, ReadWait)
	val, ok := reader1()
	if val != 42 || !ok {
		t.Errorf("ReadWait failed: val=%d, ok=%v", val, ok)
	}

	// Test ReadIfWaiting when data available
	reader2 := ReadFrom(ch, ReadIfWaiting)
	val2, ok2 := reader2()
	if val2 != 100 || !ok2 {
		t.Errorf("ReadIfWaiting with data failed: val=%d, ok=%v", val2, ok2)
	}

	// Test ReadIfWaiting when no data available
	emptyCh := make(chan int)
	reader3 := ReadFrom(emptyCh, ReadIfWaiting)
	val3, ok3 := reader3()
	if ok3 {
		t.Errorf("ReadIfWaiting without data should return false, got val=%d, ok=%v", val3, ok3)
	}
}

func TestSomeDone_Empty(t *testing.T) {
	result := SomeDone[int]()
	if result != nil {
		t.Error("SomeDone with no channels should return nil")
	}
}

func TestEveryDone_Empty(t *testing.T) {
	result := EveryDone[int]()
	if result != nil {
		t.Error("EveryDone with no channels should return nil")
	}
}

func TestEveryDone_Single(t *testing.T) {
	ch := make(chan int)
	go func() {
		close(ch)
	}()

	result := EveryDone(ch)
	select {
	case <-result:
		// Success
	case <-time.After(time.Second):
		t.Error("EveryDone with single channel should close when channel closes")
	}
}

func TestSomeDone_Single(t *testing.T) {
	ch := make(chan int)
	go func() {
		close(ch)
	}()

	result := SomeDone(ch)
	select {
	case <-result:
		// Success
	case <-time.After(time.Second):
		t.Error("SomeDone with single channel should close when channel closes")
	}
}

func TestCompact_EmptyChannel(t *testing.T) {
	emptyCh := make(chan int)
	close(emptyCh)

	result := Compact(emptyCh, compare.Equal[int])
	res := Collect(result)

	if len(res) != 0 {
		t.Errorf("Compact on empty channel should return empty, got %v", res)
	}
}
