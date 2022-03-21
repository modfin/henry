package chanz

import (
	"github.com/modfin/henry/slicez"
	"sync"
)

// Map will take a chan, in, and executes mapper and put the resulting on to the return chan.
// The return chan has a buffer of 0.
// It will stop once "in" is closed
func Map[A any, B any](in <-chan A, mapper func(a A) B) <-chan B {
	return MapUntil(nil, 0, in, mapper)
}

// Map1 will take a chan, in, and executes mapper and put the resulting on to the return chan.
// The return chan has a buffer of 1.
// It will stop once "in" is closed
func Map1[A any, B any](in <-chan A, mapper func(a A) B) <-chan B {
	return MapUntil(nil, 1, in, mapper)
}

// MapN will take a chan, in, and executes mapper and put the resulting on to the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once "in" is closed
func MapN[A any, B any](buffer int, in <-chan A, mapper func(a A) B) <-chan B {
	return MapUntil(nil, buffer, in, mapper)
}

// MapUntil will take a chan, in, and executes mapper and put the resulting on to the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once "in" or "done" is closed
func MapUntil[A any, B any](done <-chan interface{}, buffer int, in <-chan A, mapper func(a A) B) <-chan B {
	out := make(chan B, buffer)
	go func() {
		defer close(out)
		for e := range in {
			select {
			case <-done:
				return
			case out <- mapper(e):
			}
		}
	}()
	return out
}

// Peek will take a chan, in, and executes apply on every element and then writes the element to the return chan.
// The return chan has a buffer of 0.
// It will stop once "in" is closed
func Peek[A any](in <-chan A, apply func(a A)) <-chan A {
	return MapUntil(nil, 0, in, func(a A) A {
		apply(a)
		return a
	})
}

// Peek1 will take a chan, in, and executes apply on every element and then writes the element to the return chan.
// The return chan has a buffer of 1
// It will stop once "in" is closed
func Peek1[A any](in <-chan A, apply func(a A)) <-chan A {
	return MapUntil(nil, 1, in, func(a A) A {
		apply(a)
		return a
	})
}

// PeekN will take a chan, in, and executes apply on every element and then writes the element to the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once "in" is closed
func PeekN[A any](buffer int, in <-chan A, apply func(a A)) <-chan A {
	return MapUntil(nil, buffer, in, func(a A) A {
		apply(a)
		return a
	})
}

// PeekUntil will take a chan, in, and executes apply on every element and then writes the element to the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once "in" or done is closed
func PeekUntil[A any](done <-chan interface{}, buffer int, in <-chan A, apply func(a A)) <-chan A {
	return MapUntil(done, buffer, in, func(a A) A {
		apply(a)
		return a
	})
}

// Flatten takes a chan of a slice put all items received on the returning chan, one by one
// The return chan has a buffer of 0
// It will stop once the in chan are closed
func Flatten[A any](in <-chan []A) <-chan A {
	return FlattenUntil(nil, 0, in)
}

// Flatten1 takes a chan of a slice put all items received on the returning chan, one by one
// The return chan has a buffer of 1
// It will stop once the in chan are closed
func Flatten1[A any](in <-chan []A) <-chan A {
	return FlattenUntil(nil, 1, in)
}

// FlattenN takes a chan of a slice put all items received on the returning chan, one by one
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed
func FlattenN[A any](buffer int, in <-chan []A) <-chan A {
	return FlattenUntil(nil, buffer, in)
}

// FlattenUntil takes a chan of a slice put all items received on the returning chan, one by one
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed or done is closed
func FlattenUntil[A any](done <-chan interface{}, buffer int, in <-chan []A) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for slice := range in {
			if len(slice) == 0 {
				continue
			}
			select {
			case <-done:
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

// Generate takes a slice of elements, returns a channel and writes the elements to the channel. It closes once all elements in the slice are written
// The return chan has a buffer of 0
func Generate[A any](elements ...A) <-chan A {
	return GenerateUntil(nil, 0, elements...)
}

// Generate1 takes a slice of elements, returns a channel and writes the elements to the channel. It closes once all elements in the slice are written
// The return chan has a buffer of 1
func Generate1[A any](elements ...A) <-chan A {
	return GenerateUntil(nil, 1, elements...)
}

// GenerateN takes a slice of elements, returns a channel and writes the elements to the channel. It closes once all elements in the slice are written
// The return chan has a buffer of buffer size supplied in input args.
func GenerateN[A any](buffer int, elements ...A) <-chan A {
	return GenerateUntil(nil, buffer, elements...)
}

// GenerateUntil takes a slice of elements, returns a channel and writes the elements to the channel. It closes once all elements in the slice are written or done is closed
// The return chan has a buffer of buffer size supplied in input args.
func GenerateUntil[A any](done <-chan interface{}, buffer int, elements ...A) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for _, e := range elements {
			select {
			case <-done:
				return
			case out <- e:
			}
		}
	}()
	return out
}

// Merge will merge all input from input channels into one output channel. It differs from Flatten in that it reads from all channels concurrently instead of synchronized
// The return chan has a buffer of 0
// The out chan is closed once all input chans are closed
func Merge[A any](cs ...<-chan A) <-chan A {
	return MergeUntil(nil, 0, cs...)
}

// Merge1 will merge all input from input channels into one output channel. It differs from Flatten1 in that it reads from all channels concurrently instead of synchronized
// The return chan has a buffer of 1
// The out chan is closed once all input chans are closed
func Merge1[A any](cs ...<-chan A) <-chan A {
	return MergeUntil(nil, 1, cs...)
}

// MergeN will merge all input from input channels into one output channel. It differs from FlattenN in that it reads from all channels concurrently instead of synchronized
// The return chan has a buffer with the supplied buffer size
// The out chan is closed once all input chans are closed
func MergeN[A any](buffer int, cs ...<-chan A) <-chan A {
	return MergeUntil(nil, buffer, cs...)
}

// MergeUntil will merge all input from input channels into one output channel. It differs from FlattenUntil in that it reads from all channels concurrently instead of synchronized
// The return chan has a buffer with the supplied buffer size
// The out chan is closed once all input chans are closed or done is closed
func MergeUntil[A any](done <-chan interface{}, buffer int, cs ...<-chan A) <-chan A {
	var wg sync.WaitGroup
	out := make(chan A, buffer)
	output := func(c <-chan A) {
		defer wg.Done()
		for e := range c {
			select {
			case <-done:
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

// FanOut will return a slice of chans on which entries read from the input chan are written to every output chan
// A new entry won't be read from the input chan until all output chans has consumed entry.
// The returning chans has a buffer of 0
// The returning chans are closed once the in chan is closed
func FanOut[A any](c <-chan A, size int) []<-chan A {
	return FanOutUntil(nil, 0, c, size)
}

// FanOut1 will return a slice of chans on which entries read from the input chan are written to every output chan
// A new entry won't be read from the input chan until all output chans has consumed entry, if the buffer is full.
// The returning chans has a buffer of 1
// The returning chans are closed once the in chan is closed
func FanOut1[A any](c <-chan A, size int) []<-chan A {
	return FanOutUntil(nil, 1, c, size)
}

// FanOutN will return a slice of chans on which entries read from the input chan are written to every output chan
// A new entry won't be read from the input chan until all output chans has consumed entry, if the buffer is full.
// The returning chans has a buffer with the supplied buffer size
// The returning chans are closed once the in chan is closed
func FanOutN[A any](buffer int, c <-chan A, size int) []<-chan A {
	return FanOutUntil(nil, buffer, c, size)
}

// FanOutUntil will return a slice of chans on which entries read from the input chan are written to every output chan.
// A new entry won't be read from the input chan until all output chans has consumed entry, if the buffer is full.
// The returning chans has a buffer with the supplied buffer size
// The returning chans are closed once the in chan is closed
func FanOutUntil[A any](done <-chan interface{}, buffer int, c <-chan A, size int) []<-chan A {
	outs := make([]chan A, size)
	for i := range outs {
		outs[i] = make(chan A, buffer)
	}

	go func() {
		defer func() {
			for _, o := range outs {
				close(o)
			}
		}()

		for e := range c {
			select {
			case <-done:
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

// Concat takes a slice of chans and put all items received on the returning chan
// The input chans are read until closed in order.
// The return chan has a buffer of 0
// It will stop once all in chan are closed
func Concat[A any](cs ...<-chan A) <-chan A {
	return ConcatUntil(nil, 0, cs...)
}

// Concat1 takes a slice of chans and put all items received on the returning chan
// The input chans are read until closed in order.
// The return chan has a buffer of 1
// It will stop once all in chan are closed
func Concat1[A any](cs ...<-chan A) <-chan A {
	return ConcatUntil(nil, 1, cs...)
}

// ConcatN takes a slice of chans and put all items received on the returning chan
// The input chans are read until closed in order.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once all in chan are closed
func ConcatN[A any](buffer int, cs ...<-chan A) <-chan A {
	return ConcatUntil(nil, buffer, cs...)
}

// ConcatUntil takes a slice of chans and put all items received on the returning chan
// The input chans are read until closed in order.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once all in chan are closed or done is closed
func ConcatUntil[A any](done <-chan interface{}, buffer int, cs ...<-chan A) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for _, c := range cs {
			for e := range c {
				select {
				case <-done:
					return
				case out <- e:
				}
			}
		}
	}()
	return out
}

// Filter takes a chan and applies the "include" func to every item. If it returns true, the item is out on the output chan
// The return chan has a buffer of 0.
// It will stop once the in chan are closed
func Filter[A any](c <-chan A, include func(a A) bool) <-chan A {
	return FilterUntil(nil, 0, c, include)
}

// Filter1 takes a chan and applies the "include" func to every item. If it returns true, the item is out on the output chan
// The return chan has a buffer of 1.
// It will stop once the in chan are closed
func Filter1[A any](c <-chan A, include func(a A) bool) <-chan A {
	return FilterUntil(nil, 1, c, include)
}

// FilterN takes a chan and applies the "include" func to every item. If it returns true, the item is out on the output chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed
func FilterN[A any](buffer int, c <-chan A, include func(a A) bool) <-chan A {
	return FilterUntil(nil, buffer, c, include)
}

// FilterUntil takes a chan and applies the "include" func to every item. If it returns true, the item is out on the output chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed or done is closed
func FilterUntil[A any](done <-chan interface{}, buffer int, c <-chan A, include func(a A) bool) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for e := range c {
			if !include(e) {
				continue
			}
			select {
			case <-done:
				return
			case out <- e:
			}

		}
	}()
	return out
}

// Compact takes a chan and applies the "equal" func to every item and its predecessor. If it returns true, the current
// item being the same as the previous item, the current one will not be includes on the output chan
// The return chan has a buffer 0
// It will stop once the in chan are closed
func Compact[A any](c <-chan A, equal func(a, b A) bool) <-chan A {
	return CompactUntil(nil, 0, c, equal)
}

// Compact1 takes a chan and applies the "equal" func to every item and its predecessor. If it returns true, the current
// item being the same as the previous item, the current one will not be includes on the output chan
// The return chan has a buffer 1
// It will stop once the in chan are closed
func Compact1[A any](c <-chan A, equal func(a, b A) bool) <-chan A {
	return CompactUntil(nil, 1, c, equal)
}

// CompactN takes a chan and applies the "equal" func to every item and its predecessor. If it returns true, the current
// item being the same as the previous item, the current one will not be includes on the output chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed
func CompactN[A any](buffer int, c <-chan A, equal func(a, b A) bool) <-chan A {
	return CompactUntil(nil, buffer, c, equal)
}

// CompactUntil takes a chan and applies the "equal" func to every item and its predecessor. If it returns true, the current
// item being the same as the previous item, the current one will not be includes on the output chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed or done is closed
func CompactUntil[A any](done <-chan interface{}, buffer int, c <-chan A, equal func(a, b A) bool) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)

		var a, ok = <-c
		if !ok {
			return
		}
		select {
		case <-done:
			return
		case out <- a:
		}

		for b := range c {
			if equal(a, b) {
				continue
			}
			a = b
			select {
			case <-done:
				return
			case out <- b:
			}
		}
	}()
	return out
}

// Partition takes a chan and returns two chans. For every item consumed it is passed through the predicate func.
// If it returns true, the item is put on the satisfied chan otherwise it is put on the notSatisfied chan
// The return chan has a buffer of 0
// It will stop once the in chan are closed
func Partition[A any](c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return PartitionUntil(nil, 0, c, predicate)
}

// Partition1 takes a chan and returns two chans. For every item consumed it is passed through the predicate func.
// If it returns true, the item is put on the satisfied chan otherwise it is put on the notSatisfied chan
// The return chan has a buffer of 1
// It will stop once the in chan are closed
func Partition1[A any](c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return PartitionUntil(nil, 1, c, predicate)
}

// PartitionN takes a chan and returns two chans. For every item consumed it is passed through the predicate func.
// If it returns true, the item is put on the satisfied chan otherwise it is put on the notSatisfied chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed
func PartitionN[A any](buffer int, c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return PartitionUntil(nil, buffer, c, predicate)
}

// PartitionUntil takes a chan and returns two chans. For every item consumed it is passed through the predicate func.
// If it returns true, the item is put on the satisfied chan otherwise it is put on the notSatisfied chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed or done is closed
func PartitionUntil[A any](done <-chan interface{}, buffer int, c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	sat := make(chan A, buffer)
	not := make(chan A, buffer)
	go func() {
		defer close(sat)
		defer close(not)

		for e := range c {
			out := sat
			if !predicate(e) {
				out = not
			}
			select {
			case <-done:
				return
			case out <- e:
			}
		}
	}()
	return sat, not
}

// TakeWhile takes a chan and returns a chan. It will write all items read from the in chan onto the return chan until the take function returns false, then the out chan will close
// The return chan has a buffer 0
// It will stop once the in chan are closed or take returns false
func TakeWhile[A any](c <-chan A, take func(a A) bool) <-chan A {
	return TakeWhileUntil(nil, 0, c, take)
}

// TakeWhile1 takes a chan and returns a chan. It will write all items read from the in chan onto the return chan until the take function returns false, then the out chan will close
// The return chan has a buffer 1
// It will stop once the in chan are closed or take returns false
func TakeWhile1[A any](c <-chan A, take func(a A) bool) <-chan A {
	return TakeWhileUntil(nil, 1, c, take)
}

// TakeWhileN takes a chan and returns a chan. It will write all items read from the in chan onto the return chan until the take function returns false, then the out chan will close
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed or take returns false
func TakeWhileN[A any](buffer int, c <-chan A, take func(a A) bool) <-chan A {
	return TakeWhileUntil(nil, buffer, c, take)
}

// TakeWhileUntil takes a chan and returns a chan. It will write all items read from the in chan onto the return chan until the take function returns false, then the out chan will close
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed, done is closed or take returns false
func TakeWhileUntil[A any](done <-chan interface{}, buffer int, c <-chan A, take func(a A) bool) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for e := range c {
			if !take(e) {
				return
			}
			select {
			case <-done:
				return
			case out <- e:
			}

		}
	}()
	return out
}

// Take takes a chan and returns a chan. It will write the first "i" items read from the in chan onto the return chan and then close the read chan.
// The return chan has a buffer of 0
// It will stop once the in chan are closed or "i" items are read
func Take[A any](c <-chan A, i int) <-chan A {
	return TakeUntil(nil, 0, c, i)
}

// Take1 takes a chan and returns a chan. It will write the first "i" items read from the in chan onto the return chan and then close the read chan.
// The return chan has a buffer of 1
// It will stop once the in chan are closed or "i" items are read
func Take1[A any](c <-chan A, i int) <-chan A {
	return TakeUntil(nil, 1, c, i)
}

// TakeN takes a chan and returns a chan. It will write the first "i" items read from the in chan onto the return chan and then close the read chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed or "i" items are read
func TakeN[A any](buffer int, c <-chan A, i int) <-chan A {
	return TakeUntil(nil, buffer, c, i)
}

// TakeUntil takes a chan and returns a chan. It will write the first "i" items read from the in chan onto the return chan and then close the read chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed, done is closed or "i" items are read
func TakeUntil[A any](done <-chan interface{}, buffer int, c <-chan A, i int) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		if i < 1 {
			return
		}

		for e := range c {
			select {
			case <-done:
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

// Drop takes a chan and returns a chan. It will drop the first "i" items read from the in chan and write the remaining items onto the return chan.
// The return chan has a buffer of 0
// It will stop once the in chan are closed
func Drop[A any](c <-chan A, i int) <-chan A {
	return DropUntil(nil, 0, c, i)
}

// Drop1 takes a chan and returns a chan. It will drop the first "i" items read from the in chan and write the remaining items onto the return chan.
// The return chan has a buffer of 1
// It will stop once the in chan are closed
func Drop1[A any](c <-chan A, i int) <-chan A {
	return DropUntil(nil, 1, c, i)
}

// DropN takes a chan and returns a chan. It will drop the first "i" items read from the in chan and write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed
func DropN[A any](buffer int, c <-chan A, i int) <-chan A {
	return DropUntil(nil, buffer, c, i)
}

// DropUntil takes a chan and returns a chan. It will drop the first "i" items read from the in chan and write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed, done is closed
func DropUntil[A any](done <-chan interface{}, buffer int, c <-chan A, i int) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for e := range c {
			if i > 0 {
				i -= 1
				continue
			}

			select {
			case <-done:
				return
			case out <- e:
			}

		}
	}()
	return out
}

// DropWhile takes a chan and returns a chan. It will drop items until the drop function returns false, and will then write the remaining items onto the return chan.
// The return chan has a buffer of 0
// It will stop once the in chan are closed
func DropWhile[A any](c <-chan A, drop func(a A) bool) <-chan A {
	return DropWhileUntil(nil, 0, c, drop)
}

// DropWhile1 takes a chan and returns a chan. It will drop items until the drop function returns false, and will then write the remaining items onto the return chan.
// The return chan has a buffer of 1
// It will stop once the in chan are closed
func DropWhile1[A any](c <-chan A, drop func(a A) bool) <-chan A {
	return DropWhileUntil(nil, 1, c, drop)
}

// DropWhileN takes a chan and returns a chan. It will drop items until the drop function returns false, and will then write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed
func DropWhileN[A any](buffer int, c <-chan A, drop func(a A) bool) <-chan A {
	return DropWhileUntil(nil, buffer, c, drop)

}

// DropWhileUntil takes a chan and returns a chan. It will drop items until the drop function returns false, and will then write the remaining items onto the return chan.
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the in chan are closed, done is closed
func DropWhileUntil[A any](done <-chan interface{}, buffer int, c <-chan A, drop func(a A) bool) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		var dropping = true
		for e := range c {
			if dropping && drop(e) {
				continue
			}
			dropping = false
			select {
			case <-done:
				return
			case out <- e:
			}
		}
	}()
	return out
}

// Zip takes two chans and returns a chan. it will read a A item and a B item. Apply the zipper to these and output the result on the returning chan
// The return chan has a buffer of 0.
// It will stop once the any in chan are closed
func Zip[A any, B any, C any](ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	return ZipUntil(nil, 0, ac, bc, zipper)
}

// Zip1 takes two chans and returns a chan. it will read a A item and a B item. Apply the zipper to these and output the result on the returning chan
// The return chan has a buffer of 1.
// It will stop once the any in chan are closed
func Zip1[A any, B any, C any](ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	return ZipUntil(nil, 1, ac, bc, zipper)
}

// ZipN takes two chans and returns a chan. it will read a A item and a B item. Apply the zipper to these and output the result on the returning chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the any in chan are closed
func ZipN[A any, B any, C any](buffer int, ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	return ZipUntil(nil, buffer, ac, bc, zipper)
}

// ZipUntil takes two chans and returns a chan. it will read a A item and a B item. Apply the zipper to these and output the result on the returning chan
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the any in chan are closed, done is closed
func ZipUntil[A any, B any, C any](done <-chan interface{}, buffer int, ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	out := make(chan C, buffer)
	go func() {
		defer close(out)
		for a := range ac {
			b, ok := <-bc
			if !ok {
				return
			}
			select {
			case <-done:
				return
			case out <- zipper(a, b):
			}
		}
	}()
	return out
}

// Unzip takes one chan and returns a chan. It will read a C item from the input chan, apply the unzipper to the two resulting items on the output chans
// The return chan has a buffer of 0
// It will stop once the any in chan are closed
func Unzip[A any, B any, C any](zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	return UnzipUntil(nil, 0, zipped, unzipper)
}

// Unzip1 takes one chan and returns a chan. It will read a C item from the input chan, apply the unzipper to the two resulting items on the output chans
// The return chan has a buffer of 1
// It will stop once the any in chan are closed
func Unzip1[A any, B any, C any](zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	return UnzipUntil(nil, 1, zipped, unzipper)
}

// UnzipN takes one chan and returns a chan. It will read a C item from the input chan, apply the unzipper to the two resulting items on the output chans
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the any in chan are closed
func UnzipN[A any, B any, C any](buffer int, zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	return UnzipUntil(nil, buffer, zipped, unzipper)
}

// UnzipUntil takes one chan and returns a chan. It will read a C item from the input chan, apply the unzipper to the two resulting items on the output chans
// The return chan has a buffer of buffer size supplied in input args.
// It will stop once the any in chan are closed, done is closed
func UnzipUntil[A any, B any, C any](done <-chan interface{}, buffer int, zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	ac := make(chan A, buffer)
	bc := make(chan B, buffer)
	go func() {
		defer close(ac)
		defer close(bc)
		for c := range zipped {
			a, b := unzipper(c)
			select {
			case <-done:
				return
			case ac <- a:
				bc <- b
			}
		}
	}()
	return ac, bc
}

// Collect will collect all entries in a channel into a slice and return it. It stops and returns c is closed
func Collect[A any](c <-chan A) []A {
	return CollectUntil(nil, c)
}

// CollectUntil will collect all enteries in a channel into a slice and return it. It stops and returns when done or c is closed
func CollectUntil[A any](done <-chan interface{}, c <-chan A) []A {
	var out []A
	for val := range c {
		out = append(out, val)
		select {
		case <-done:
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

// EveryDone returns a channel that closes when all channels from the input arguments are closed
func EveryDone(done ...<-chan interface{}) <-chan interface{} {

	switch len(done) {
	case 0:
		return nil
	case 1:
		return done[0]
	}

	allDone := make(chan interface{})
	go func() {
		defer close(allDone)
		for _, d := range done {
			<-d
		}
	}()

	return allDone
}

// SomeDone returns a channel that closes as soon as any channels from the input arguments are closed
func SomeDone(done ...<-chan interface{}) <-chan interface{} {
	switch len(done) {
	case 0:
		return nil
	case 1:
		return done[0]
	}

	someDone := make(chan interface{})
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
