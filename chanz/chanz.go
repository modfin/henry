// Package chanz provides utility functions for working with Go channels.
//
// The package offers functional-style operations on channels including:
//   - Transformation: Map, Flatten, Zip/Unzip
//   - Filtering: Filter, Compact, Take/Drop variants
//   - Aggregation: FanIn, FanOut, Concat
//   - Generation: Generate, Generator
//   - Utilities: Collect, Partition, Done signal handling
//
// Most functions support functional options for configuration:
//   - OpBuffer(n): Set channel buffer size (default 0)
//   - OpContext(ctx): Stop when context is cancelled
//   - OpDone(ch): Stop when done channel is closed
//
// Functions ending in "With" (e.g., MapWith) return closures that can be reused
// with the same options, useful for pipeline building.
//
// Example pipeline:
//
//	input := chanz.Generate(1, 2, 3, 4, 5)
//	doubled := chanz.Map(input, func(n int) int { return n * 2 })
//	evens := chanz.Filter(doubled, func(n int) bool { return n%2 == 0 })
//	result := chanz.Collect(evens)
//	// result = []int{2, 4, 6, 8, 10}
package chanz

import (
	"context"
	"sync"

	"github.com/modfin/henry/slicez"
)

// settings holds configuration for channel operations.
// Used internally with the Option pattern.
type settings struct {
	done   <-chan struct{} // Signal to stop processing
	buffer int             // Channel buffer size
}

// Option is a functional option for configuring channel operations.
// Use OpBuffer, OpContext, or OpDone to create options.
type Option func(s settings) settings

// OpContext creates an option that stops processing when the context is cancelled.
// Combines with existing done signals using SomeDone.
func OpContext(ctx context.Context) Option {
	return func(s settings) settings {
		s.done = SomeDone(ctx.Done(), s.done)
		return s
	}
}

// OpBuffer creates an option that sets the output channel buffer size.
// Default buffer size is 0 (unbuffered).
func OpBuffer(size int) Option {
	return func(s settings) settings {
		s.buffer = size
		return s
	}
}

// OpDone creates an option that stops processing when the done channel is closed.
// Combines with existing done signals using SomeDone.
func OpDone(done <-chan struct{}) Option {
	return func(s settings) settings {
		s.done = SomeDone(done, s.done)
		return s
	}
}

// Map will take a chan, in, and executes mapper and put the resulting on to the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Map[A any, B any](in <-chan A, mapper func(a A) B, options ...Option) <-chan B {
	var s settings
	for _, o := range options {
		s = o(s)
	}

	out := make(chan B, s.buffer)
	go func() {
		defer close(out)
		for e := range in {
			select {
			case <-s.done:
				return
			case out <- mapper(e):
			}
		}
	}()
	return out
}

// MapWith returns a configured Map function closure.
// Allows creating reusable mappers with preset options.
//
// Example:
//
//	doubler := chanz.MapWith[int, int](OpBuffer(10))
//	input := chanz.Generate(1, 2, 3, 4, 5)
//	result := chanz.Collect(doubler(input, func(n int) int { return n * 2 }))
//	// result = []int{2, 4, 6, 8, 10}
func MapWith[A any, B any](options ...Option) func(in <-chan A, mapper func(a A) B) <-chan B {
	return func(in <-chan A, mapper func(a A) B) <-chan B {
		return Map(in, mapper, options...)
	}
}

// Peek will take a chan, in, and executes apply on every element and then writes the element to the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Peek[A any](in <-chan A, apply func(a A), options ...Option) <-chan A {
	return Map(in, func(a A) A {
		apply(a)
		return a
	}, options...)
}

// PeekWith returns a configured Peek function closure.
// Allows creating reusable peekers with preset options.
//
// Example:
//
//	logger := chanz.PeekWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 2, 3)
//	logged := logger(input, func(n int) { fmt.Println("Processing:", n) })
//	chanz.Collect(logged) // Prints each number as it's processed
func PeekWith[A any](options ...Option) func(in <-chan A, apply func(a A)) <-chan A {
	return func(in <-chan A, apply func(a A)) <-chan A {
		return Peek(in, apply, options...)
	}
}

// Flatten takes a chan of a slice put all items received on the returning chan, one by one
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Flatten[A any](in <-chan []A, options ...Option) <-chan A {

	var s settings
	for _, o := range options {
		s = o(s)
	}

	out := make(chan A, s.buffer)
	go func() {
		defer close(out)
		for slice := range in {
			if len(slice) == 0 {
				continue
			}
			select {
			case <-s.done:
				return
			case out <- slice[0]:
				for _, e := range slice[1:] {
					out <- e
				}
			}
		}
	}()
	return out
}

// FlattenWith returns a configured Flatten function closure.
// Allows creating reusable flatteners with preset options.
//
// Example:
//
//	flattener := chanz.FlattenWith[int](OpBuffer(20))
//	input := make(chan []int)
//	go func() { input <- []int{1, 2}; input <- []int{3, 4}; close(input) }()
//	result := chanz.Collect(flattener(input))
//	// result = []int{1, 2, 3, 4}
func FlattenWith[A any](options ...Option) func(in <-chan []A) <-chan A {
	return func(in <-chan []A) <-chan A {
		return Flatten(in, options...)
	}
}

// Generator creates a channel that yields values from a generator function.
// The generator receives a yield function to emit values. Stops when done signal received.
// Useful for creating channels from iterative algorithms.
//
// Example:
//
//	// Generate Fibonacci numbers
//	fib := chanz.Generator(func(yield func(int)) {
//	    a, b := 0, 1
//	    for i := 0; i < 10; i++ {
//	        yield(a)
//	        a, b = b, a+b
//	    }
//	})
//	result := chanz.Collect(fib)
//	// result = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
func Generator[A any](gen func(func(A)), options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}
	out := make(chan A, s.buffer)

	yield := func(a A) {
		select {
		case <-s.done:
			return
		case out <- a:
		}
	}
	go func() {
		defer close(out)
		gen(yield)
	}()
	return out
}

// GeneratorWith returns a configured Generator function closure.
// Allows creating reusable generators with preset options.
//
// Example:
//
//	rangeGen := chanz.GeneratorWith[int](OpBuffer(5))
//	oneto5 := rangeGen(func(yield func(int)) {
//	    for i := 1; i <= 5; i++ {
//	        yield(i)
//	    }
//	})
//	result := chanz.Collect(oneto5)
//	// result = []int{1, 2, 3, 4, 5}
func GeneratorWith[A any](options ...Option) func(gen func(func(A))) <-chan A {
	return func(gen func(func(A))) <-chan A {
		return Generator(gen, options...)
	}
}

// Generate takes a slice of elements, returns a channel and writes the elements to the channel. It closes once all elements in the slice are written
// The return chan has a buffer of 0
func Generate[A any](elements ...A) <-chan A {
	return GenerateWith[A]()(elements...)
}

// GenerateWith returns a configured Generate function closure.
// Allows creating reusable element generators with preset options.
//
// Example:
//
//	bufferedGen := chanz.GenerateWith[int](OpBuffer(10))
//	ch := bufferedGen(1, 2, 3, 4, 5)
//	result := chanz.Collect(ch)
//	// result = []int{1, 2, 3, 4, 5}
func GenerateWith[A any](options ...Option) func(elements ...A) <-chan A {
	return func(elements ...A) <-chan A {
		var s settings
		for _, o := range options {
			s = o(s)
		}
		out := make(chan A, s.buffer)
		go func() {
			defer close(out)
			for _, e := range elements {
				select {
				case <-s.done:
					return
				case out <- e:
				}
			}
		}()
		return out
	}
}

// FanIn merges multiple channels into one output channel.
// Reads from all input channels concurrently, so order of output is non-deterministic.
// Output closes when all input channels are closed.
//
// Example:
//
//	ch1 := chanz.Generate(1, 2, 3)
//	ch2 := chanz.Generate(4, 5, 6)
//	merged := chanz.FanIn(ch1, ch2)
//	result := chanz.Collect(merged)
//	// result contains {1,2,3,4,5,6} in some order
func FanIn[A any](cs ...<-chan A) <-chan A {
	return FanInWith[A]()(cs...)
}

// FanInWith will merge all input from input channels into one output channel. It differs from FlattenUntil in that it reads from all channels concurrently instead of synchronized
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func FanInWith[A any](options ...Option) func(cs ...<-chan A) <-chan A {
	return func(cs ...<-chan A) <-chan A {

		var s settings
		for _, o := range options {
			s = o(s)
		}

		var wg sync.WaitGroup
		out := make(chan A, s.buffer)
		output := func(c <-chan A) {
			defer wg.Done()
			for e := range c {
				select {
				case <-s.done:
					return
				case out <- e:
				}
			}
		}
		wg.Add(len(cs))
		for _, c := range cs {
			go output(c)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}
}

// FanOut splits one input channel into multiple output channels.
// Each value from the input is sent to all output channels (broadcast pattern).
// A value won't be read from input until all outputs have consumed the previous value
// (if buffers are full). Output channels are closed when input closes.
//
// Example:
//
//	input := chanz.Generate(1, 2, 3)
//	outputs := chanz.FanOut(input, 2, OpBuffer(1))
//	ch1, ch2 := outputs[0], outputs[1]
//	// Both ch1 and ch2 receive: 1, 2, 3
func FanOut[A any](c <-chan A, size int, options ...Option) []<-chan A {

	var s settings
	for _, o := range options {
		s = o(s)
	}

	outs := make([]chan A, size)
	for i := range outs {
		outs[i] = make(chan A, s.buffer)
	}

	go func() {
		defer func() {
			for _, o := range outs {
				close(o)
			}
		}()

		for e := range c {
			select {
			case <-s.done:
				return
			case outs[0] <- e: // Might want to do this concurrently somehow?
				for _, o := range outs[1:] {
					o <- e
				}
			}
		}
	}()
	return Readers(outs...)
}

// FanOutWith returns a configured FanOut function closure.
// Allows creating reusable broadcasters with preset options.
//
// Example:
//
//	broadcaster := chanz.FanOutWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 2, 3)
//	outputs := broadcaster(input, 3) // Split into 3 channels
func FanOutWith[A any](options ...Option) func(c <-chan A, size int) []<-chan A {
	return func(c <-chan A, size int) []<-chan A {
		return FanOut(c, size, options...)
	}
}

// Concat concatenates multiple channels sequentially.
// Reads from channels in order: waits for first channel to close,
// then starts reading from next. Unlike FanIn, this preserves order.
// Output closes when all input channels are closed.
//
// Example:
//
//	ch1 := chanz.Generate(1, 2, 3)
//	ch2 := chanz.Generate(4, 5, 6)
//	combined := chanz.Concat(ch1, ch2)
//	result := chanz.Collect(combined)
//	// result = []int{1, 2, 3, 4, 5, 6} (order preserved)
func Concat[A any](cs ...<-chan A) <-chan A {
	return ConcatWith[A]()(cs...)
}

// ConcatWith takes a slice of chans and put all items received on the returning chan
// The input chans are read until closed in order.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func ConcatWith[A any](options ...Option) func(cs ...<-chan A) <-chan A {
	return func(cs ...<-chan A) <-chan A {
		var s settings
		for _, o := range options {
			s = o(s)
		}

		out := make(chan A, s.buffer)
		go func() {
			defer close(out)
			for _, c := range cs {
				for e := range c {
					select {
					case <-s.done:
						return
					case out <- e:
					}
				}
			}
		}()
		return out
	}
}

// Filter takes a chan and applies the "include" func to every item. If it returns true, the item is out on the output chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Filter[A any](c <-chan A, include func(a A) bool, options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}

	out := make(chan A, s.buffer)
	go func() {
		defer close(out)
		for e := range c {
			if !include(e) {
				continue
			}
			select {
			case <-s.done:
				return
			case out <- e:
			}
		}
	}()
	return out
}

// FilterWith returns a configured Filter function closure.
// Allows creating reusable filters with preset options.
//
// Example:
//
//	evens := chanz.FilterWith[int](OpBuffer(10))
//	input := chanz.Generate(1, 2, 3, 4, 5)
//	result := chanz.Collect(evens(input, func(n int) bool { return n%2 == 0 }))
//	// result = []int{2, 4}
func FilterWith[A any](options ...Option) func(c <-chan A, include func(a A) bool) <-chan A {
	return func(c <-chan A, include func(a A) bool) <-chan A {
		return Filter(c, include, options...)
	}
}

// Compact takes a chan and applies the "equal" func to every item and its predecessor. If it returns true, the current
// item being the same as the previous item, the current one will not be includes on the output chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Compact[A any](c <-chan A, equal func(a, b A) bool, options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}
	out := make(chan A, s.buffer)
	go func() {
		defer close(out)

		var a, ok = <-c
		if !ok {
			return
		}
		select {
		case <-s.done:
			return
		case out <- a:
		}

		for b := range c {
			if equal(a, b) {
				continue
			}
			a = b
			select {
			case <-s.done:
				return
			case out <- b:
			}
		}
	}()
	return out
}

// CompactWith returns a configured Compact function closure.
// Allows creating reusable compacters with preset options.
//
// Example:
//
//	dedup := chanz.CompactWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 1, 2, 2, 2, 3, 3)
//	result := chanz.Collect(dedup(input, func(a, b int) bool { return a == b }))
//	// result = []int{1, 2, 3}
func CompactWith[A any](options ...Option) func(c <-chan A, equal func(a, b A) bool) <-chan A {
	return func(c <-chan A, equal func(a A, b A) bool) <-chan A {
		return Compact(c, equal, options...)
	}
}

// Partition takes a chan and returns two chans. For every item consumed it is passed through the predicate func.
// If it returns true, the item is put on the satisfied chan otherwise it is put on the notSatisfied chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Partition[A any](c <-chan A, predicate func(a A) bool, options ...Option) (satisfied, notSatisfied <-chan A) {

	var s settings
	for _, o := range options {
		s = o(s)
	}

	sat := make(chan A, s.buffer)
	not := make(chan A, s.buffer)
	go func() {
		defer close(sat)
		defer close(not)

		for e := range c {
			out := sat
			if !predicate(e) {
				out = not
			}
			select {
			case <-s.done:
				return
			case out <- e:
			}
		}
	}()
	return sat, not
}

// PartitionWith returns a configured Partition function closure.
// Allows creating reusable partitioners with preset options.
//
// Example:
//
//	splitter := chanz.PartitionWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 2, 3, 4, 5)
//	evens, odds := splitter(input, func(n int) bool { return n%2 == 0 })
//	// evens receives 2, 4; odds receives 1, 3, 5
func PartitionWith[A any](options ...Option) func(c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return func(c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
		return Partition(c, predicate, options...)
	}
}

// TakeWhile takes a chan and returns a chan. It will write all items read from the in chan onto the return chan until the take function returns false, then the out chan will close
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func TakeWhile[A any](c <-chan A, take func(a A) bool, options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}
	out := make(chan A, s.buffer)
	go func() {
		defer close(out)
		for e := range c {
			if !take(e) {
				return
			}
			select {
			case <-s.done:
				return
			case out <- e:
			}

		}
	}()
	return out
}

// TakeWhileWith returns a configured TakeWhile function closure.
// Allows creating reusable takers with preset options.
//
// Example:
//
//	takeUnder10 := chanz.TakeWhileWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 5, 8, 12, 3, 4)
//	result := chanz.Collect(takeUnder10(input, func(n int) bool { return n < 10 }))
//	// result = []int{1, 5, 8} (stops at 12)
func TakeWhileWith[A any](options ...Option) func(c <-chan A, take func(a A) bool) <-chan A {
	return func(c <-chan A, take func(a A) bool) <-chan A {
		return TakeWhile(c, take, options...)
	}
}

// Take takes a chan and returns a chan. It will write the first "i" items read from the in chan onto the return chan and then close the read chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Take[A any](c <-chan A, i int, options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}
	out := make(chan A, s.buffer)
	go func() {
		defer close(out)
		if i < 1 {
			return
		}

		for e := range c {
			select {
			case <-s.done:
				return
			case out <- e:
			}
			i -= 1
			if i == 0 {
				return
			}
		}
	}()
	return out
}

// TakeWith returns a configured Take function closure.
// Allows creating reusable takers with preset options.
//
// Example:
//
//	take5 := chanz.TakeWith[int](OpBuffer(3))
//	input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
//	result := chanz.Collect(take5(input, 5))
//	// result = []int{1, 2, 3, 4, 5}
func TakeWith[A any](option ...Option) func(c <-chan A, i int) <-chan A {
	return func(c <-chan A, i int) <-chan A {
		return Take(c, i, option...)
	}
}

// Drop takes a chan and returns a chan. It will drop the first "i" items read from the in chan and write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Drop[A any](c <-chan A, i int, options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}

	out := make(chan A, s.buffer)
	go func() {
		defer close(out)
		for e := range c {
			if i > 0 {
				i -= 1
				continue
			}

			select {
			case <-s.done:
				return
			case out <- e:
			}

		}
	}()
	return out
}

// DropWith returns a configured Drop function closure.
// Allows creating reusable droppers with preset options.
//
// Example:
//
//	drop3 := chanz.DropWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
//	result := chanz.Collect(drop3(input, 3))
//	// result = []int{4, 5, 6, 7, 8, 9, 10}
func DropWith[A any](options ...Option) func(c <-chan A, i int) <-chan A {
	return func(c <-chan A, i int) <-chan A {
		return Drop(c, i, options...)
	}
}

// DropWhile takes a chan and returns a chan. It will drop items until the drop function returns false, and will then write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func DropWhile[A any](c <-chan A, drop func(a A) bool, options ...Option) <-chan A {
	var s settings
	for _, o := range options {
		s = o(s)
	}

	out := make(chan A, s.buffer)
	go func() {
		defer close(out)
		var dropping = true
		for e := range c {
			if dropping && drop(e) {
				continue
			}
			dropping = false
			select {
			case <-s.done:
				return
			case out <- e:
			}
		}
	}()
	return out
}

// DropWhileWith returns a configured DropWhile function closure.
// Allows creating reusable droppers with preset options.
//
// Example:
//
//	dropSmall := chanz.DropWhileWith[int](OpBuffer(5))
//	input := chanz.Generate(1, 2, 3, 5, 8, 4, 3, 2, 1)
//	result := chanz.Collect(dropSmall(input, func(n int) bool { return n < 5 }))
//	// result = []int{5, 8, 4, 3, 2, 1} (drops while n < 5)
func DropWhileWith[A any](options ...Option) func(c <-chan A, drop func(a A) bool) <-chan A {
	return func(c <-chan A, drop func(a A) bool) <-chan A {
		return DropWhile(c, drop, options...)
	}
}

// Zip takes two chans and returns a chan. it will read a A item and a B item. Apply the zipper to these and output the result on the returning chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Zip[A any, B any, C any](ac <-chan A, bc <-chan B, zipper func(a A, b B) C, options ...Option) <-chan C {
	var s settings
	for _, o := range options {
		s = o(s)
	}

	out := make(chan C, s.buffer)
	go func() {
		defer close(out)
		for a := range ac {
			b, ok := <-bc
			if !ok {
				return
			}
			select {
			case <-s.done:
				return
			case out <- zipper(a, b):
			}
		}
	}()
	return out
}

// Unzip takes one chan and returns a chan. It will read a C item from the input chan, apply the unzipper to the two resulting items on the output chans
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the any in chan are closed, done is closed
func Unzip[A any, B any, C any](zipped <-chan C, unzipper func(c C) (A, B), options ...Option) (<-chan A, <-chan B) {
	var s settings
	for _, o := range options {
		s = o(s)
	}

	ac := make(chan A, s.buffer)
	bc := make(chan B, s.buffer)
	go func() {
		defer close(ac)
		defer close(bc)
		for c := range zipped {
			a, b := unzipper(c)
			select {
			case <-s.done:
				return
			case ac <- a:
				bc <- b
			}
		}
	}()
	return ac, bc
}

// Collect will collect all enteries in a channel into a slice and return it. It stops and returns when done or c is closed
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func Collect[A any](c <-chan A, options ...Option) []A {
	var s settings
	for _, o := range options {
		s = o(s)
	}
	var out []A
	for val := range c {
		out = append(out, val)
		select {
		case <-s.done:
			return out
		default:
		}
	}
	return out
}

// DropAll consumes and discards all values from a channel until it closes.
// If async is false, blocks until the channel is closed.
// If async is true, returns immediately and consumes in background.
//
// Example:
//
//	ch := make(chan int)
//	go func() { ch <- 1; ch <- 2; close(ch) }()
//	chanz.DropAll(ch, false) // Blocks until ch is closed
func DropAll[A any](c <-chan A, async bool) {
	dropper := func() {
		for range c {
		}
	}
	if async {
		go dropper()
		return
	}
	dropper()
}

// TakeBuffer reads all values currently in a buffered channel's buffer.
// Non-blocking - only takes what's already buffered, doesn't wait for more.
// Returns values already received in the buffer.
//
// Example:
//
//	ch := make(chan int, 10)
//	ch <- 1; ch <- 2; ch <- 3
//	buff := chanz.TakeBuffer(ch)
//	// buff = []int{1, 2, 3}, ch is now empty
func TakeBuffer[A any](c <-chan A) []A {
	var buff []A
	taker := func() {
		l := len(c)
		for i := 0; i < l; i++ {
			select {
			case val, open := <-c:
				if !open {
					return
				}
				buff = append(buff, val)
			default:
				return
			}
		}
	}
	taker()
	return buff
}

// DropBuffer discards all values currently in a buffered channel's buffer.
// Non-blocking - only drops what's already buffered, doesn't wait for more.
// If async is true, runs in background.
//
// Example:
//
//	ch := make(chan int, 10)
//	ch <- 1; ch <- 2; ch <- 3
//	chanz.DropBuffer(ch, false) // ch is now empty
func DropBuffer[A any](c <-chan A, async bool) {
	dropper := func() {
		l := len(c)
		for i := 0; i < l; i++ {
			select {
			case _, open := <-c:
				if !open {
					return
				}
			default:
				return
			}
		}
	}
	if async {
		go dropper()
		return
	}
	dropper()
}

// Buffer collects up to size elements from a channel into a slice.
// Returns the slice and a boolean indicating if buffer was filled (true) or channel closed (false).
// Stops early if done signal received.
//
// Example:
//
//	input := chanz.Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
//	batch, filled := chanz.Buffer(3, input)
//	// batch = []int{1, 2, 3}, filled = true
//	batch, filled = chanz.Buffer(100, input)
//	// batch = []int{4, 5, 6, 7, 8, 9, 10}, filled = false (channel closed)
func Buffer[A any](size int, in <-chan A, options ...Option) ([]A, bool) {
	var s settings
	for _, o := range options {
		s = o(s)
	}
	var out = make([]A, 0, size)

	for {

		val, more := <-in
		out = append(out, val)

		if !more {
			return out, false
		}
		if len(out) == size {
			return out, true
		}
		select {
		case <-s.done:
			return out, true
		default:
		}
	}

}

// Done takes a channel, c, that is ment to indicate that something is done and returns a chan struct{} that closes once c does
// It is ment to convert a channel of any type to a channel that aligns with context.Context.Done()
// if data is passed on c, Done will drain it
func Done[T any](c chan T) <-chan struct{} {
	ret := make(chan struct{})
	go func() {
		for range c {
		}
		close(ret)
	}()
	return ret
}

// EveryDone returns a channel that closes when all channels from the input arguments are closed
func EveryDone[T any](done ...<-chan T) <-chan T {

	switch len(done) {
	case 0:
		return nil
	case 1:
		return done[0]
	}

	allDone := make(chan T)
	go func() {
		defer close(allDone)
		for _, d := range done {
			<-d
		}
	}()

	return allDone
}

// SomeDone returns a channel that closes as soon as any channels from the input arguments are closed
func SomeDone[T any](done ...<-chan T) <-chan T {
	switch len(done) {
	case 0:
		return nil
	case 1:
		return done[0]
	}

	someDone := make(chan T)
	go func() {
		defer close(someDone)
		switch len(done) {
		case 2:
			select {
			case <-done[0]:
			case <-done[1]:
			}
		default:
			select {
			case <-done[0]:
			case <-done[1]:
			case <-done[2]:
			case <-SomeDone(append(done[3:], someDone)...):
			}
		}
	}()
	return someDone
}

// Readers converts a slice of bidirectional channels to read-only channels.
// Useful for type safety when passing channels to consumers.
//
// Example:
//
//	chans := []chan int{make(chan int), make(chan int)}
//	readers := chanz.Readers(chans...)
//	// readers is []<-chan int, can only receive
func Readers[A any](chans ...chan A) []<-chan A {
	return slicez.Map(chans, func(a chan A) <-chan A {
		return a
	})
}

// Writers converts a slice of bidirectional channels to write-only channels.
// Useful for type safety when passing channels to producers.
//
// Example:
//
//	chans := []chan int{make(chan int), make(chan int)}
//	writers := chanz.Writers(chans...)
//	// writers is []chan<- int, can only send
func Writers[A any](chans ...chan A) []chan<- A {
	return slicez.Map(chans, func(a chan A) chan<- A {
		return a
	})
}

const (
	// WriteSync blocks until the write completes (standard channel send).
	WriteSync = iota
	// WriteAync performs the write in a goroutine (non-blocking).
	WriteAync
	// WriteIfFree writes only if the channel buffer has space (non-blocking).
	WriteIfFree
)

// WriteTo returns a function that writes to a channel with specified mode.
// Modes: WriteSync (block), WriteAync (goroutine), WriteIfFree (non-blocking).
//
// Example:
//
//	ch := make(chan int, 1)
//	writeSync := chanz.WriteTo[int](ch, chanz.WriteSync)
//	writeSync(42) // Blocks until written
//
//	writeAsync := chanz.WriteTo[int](ch, chanz.WriteAync)
//	writeAsync(42) // Returns immediately, writes in background
//
//	writeIfFree := chanz.WriteTo[int](ch, chanz.WriteIfFree)
//	writeIfFree(42) // Only writes if buffer has space
func WriteTo[A any](c chan<- A, mode int) func(m A) {
	return func(m A) {
		switch mode {
		case WriteSync:
			c <- m
		case WriteAync:
			go WriteTo(c, WriteSync)(m)
		case WriteIfFree:
			select {
			case c <- m:
			default:
			}
		}
	}
}

const (
	// ReadWait blocks until a value is received (standard channel receive).
	ReadWait = iota
	// ReadIfWaiting receives only if a value is immediately available (non-blocking).
	ReadIfWaiting
)

// ReadFrom returns a function that reads from a channel with specified mode.
// Modes: ReadWait (block), ReadIfWaiting (non-blocking).
// Returns value and ok (false if channel closed or no value available in non-blocking mode).
//
// Example:
//
//	ch := make(chan int, 1)
//	ch <- 42
//
//	readWait := chanz.ReadFrom[int](ch, chanz.ReadWait)
//	val, ok := readWait() // Blocks, returns (42, true)
//
//	readIfWaiting := chanz.ReadFrom[int](ch, chanz.ReadIfWaiting)
//	val, ok := readIfWaiting() // Returns immediately, (0, false) if empty
func ReadFrom[A any](c chan A, mode int) func() (m A, ok bool) {
	return func() (A, bool) {
		switch mode {
		case ReadWait:
			return <-c, true
		case ReadIfWaiting:
			select {
			case m := <-c:
				return m, true
			default:
			}
		default:
			panic("Read mode does not exist")
		}
		var m A
		return m, false

	}
}
