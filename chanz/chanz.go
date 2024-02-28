package chanz

import (
	"context"
	"github.com/modfin/henry/slicez"
	"sync"
)

type settings struct {
	done   <-chan struct{}
	buffer int
}

type Option func(s settings) settings

func WithContext(ctx context.Context) Option {
	return func(s settings) settings {
		s.done = SomeDone(ctx.Done(), s.done)
		return s
	}
}
func Buffer(size int) Option {
	return func(s settings) settings {
		s.buffer = size
		return s
	}
}
func WithDoneChan(done <-chan struct{}) Option {
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

// MapWith will take a chan, in, and executes mapper and put the resulting on to the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// PeekWith will take a chan, in, and executes apply on every element and then writes the element to the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

func FlattenWith[A any](options ...Option) func(in <-chan []A) <-chan A {
	return func(in <-chan []A) <-chan A {
		return Flatten(in, options...)
	}
}

// Generate takes a slice of elements, returns a channel and writes the elements to the channel. It closes once all elements in the slice are written
// The return chan has a buffer of 0
func Generate[A any](elements ...A) <-chan A {
	return GenerateWith[A]()(elements...)
}

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

// Merge will merge all input from input channels into one output channel.
// It differs from Flattenin that it reads from all channels concurrently instead of synchronized
func Merge[A any](cs ...<-chan A) <-chan A {
	return MergeWith[A]()(cs...)
}

// MergeWith will merge all input from input channels into one output channel. It differs from FlattenUntil in that it reads from all channels concurrently instead of synchronized
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func MergeWith[A any](options ...Option) func(cs ...<-chan A) <-chan A {
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

// FanOut will return a slice of chans on which entries read from the input chan are written to every output chan.
// A new entry won't be read from the input chan until all output chans has consumed entry, if the buffer is full.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// FanOutWith will return a slice of chans on which entries read from the input chan are written to every output chan.
// A new entry won't be read from the input chan until all output chans has consumed entry, if the buffer is full.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
func FanOutWith[A any](options ...Option) func(c <-chan A, size int) []<-chan A {
	return func(c <-chan A, size int) []<-chan A {
		return FanOut(c, size, options...)
	}
}

// Concat takes a slice of chans and put all items received on the returning chan
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

// FilterWith takes a chan and applies the "include" func to every item. If it returns true, the item is out on the output chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// CompactWith takes a chan and applies the "equal" func to every item and its predecessor. If it returns true, the current
// item being the same as the previous item, the current one will not be includes on the output chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// PartitionWith takes a chan and returns two chans. For every item consumed it is passed through the predicate func.
// If it returns true, the item is put on the satisfied chan otherwise it is put on the notSatisfied chan
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// TakeWhileWith takes a chan and returns a chan. It will write all items read from the in chan onto the return chan until the take function returns false, then the out chan will close
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// TakeWith takes a chan and returns a chan. It will write the first "i" items read from the in chan onto the return chan and then close the read chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// DropWith takes a chan and returns a chan. It will drop the first "i" items read from the in chan and write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// DropWhileWith takes a chan and returns a chan. It will drop items until the drop function returns false, and will then write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input Option, default is 0.
// It will stop once "in", "done" channel is closed or the context.Done is closed, which is supplied in Option
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

// DropAll will consume a channel until it closes. If async is false, it will block until the channel is closed and all entries are consumed.
// If async is true it will immediately return and consume all elements in the background
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

// TakeBuffer will take everything in the channels buffer ()
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

// DropBuffer will drop everything in the channels buffer ()
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

// Readers takes a slice of chans and returns the reader version of it
func Readers[A any](chans ...chan A) []<-chan A {
	return slicez.Map(chans, func(a chan A) <-chan A {
		return a
	})
}

// Writers takes a slice of chans and returns the writer version of it
func Writers[A any](chans ...chan A) []chan<- A {
	return slicez.Map(chans, func(a chan A) chan<- A {
		return a
	})
}

const (
	WriteSync = iota
	WriteAync
	WriteIfFree
)

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
	ReadWait      = iota
	ReadIfWaiting = iota
)

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
