package chanz

import (
	"sync"
)

func Peek[A any](in <-chan A, apply func(a A)) <-chan A {
	return PeekUntil(nil, 0, in, apply)
}
func Peek1[A any](in <-chan A, apply func(a A)) <-chan A {
	return PeekUntil(nil, 1, in, apply)
}

func PeekN[A any](buffer int, in <-chan A, apply func(a A)) <-chan A {
	return PeekUntil(nil, buffer, in, apply)
}

func PeekUntil[A any](done <-chan interface{}, buffer int, in <-chan A, apply func(a A)) <-chan A {
	out := make(chan A, buffer)
	go func() {
		defer close(out)
		for e := range in {
			apply(e)
			select {
			case <-done:
				return
			case out <- e:
			}
		}
	}()
	return out
}

func Map[A any, B any](in <-chan A, mapper func(a A) B) <-chan B {
	return MapUntil(nil, 0, in, mapper)
}
func Map1[A any, B any](in <-chan A, mapper func(a A) B) <-chan B {
	return MapUntil(nil, 1, in, mapper)
}

func MapN[A any, B any](buffer int, in <-chan A, mapper func(a A) B) <-chan B {
	return MapUntil(nil, buffer, in, mapper)
}

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

func Flatten[A any](in <-chan []A) <-chan A {
	return FlattenUntil(nil, 0, in)
}

func Flatten1[A any](in <-chan []A) <-chan A {
	return FlattenUntil(nil, 1, in)
}

func FlattenN[A any](buffer int, in <-chan []A) <-chan A {
	return FlattenUntil(nil, buffer, in)
}

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

func Generate[A any](elements ...A) <-chan A {
	return GenerateUntil(nil, 0, elements...)
}

func Generate1[A any](elements ...A) <-chan A {
	return GenerateUntil(nil, 1, elements...)
}

func GenerateN[A any](buffer int, elements ...A) <-chan A {
	return GenerateUntil(nil, buffer, elements...)
}

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

// Fan-In?
func Merge[A any](cs ...<-chan A) <-chan A {
	return MergeUntil(nil, 0, cs...)
}

func Merge1[A any](cs ...<-chan A) <-chan A {
	return MergeUntil(nil, 1, cs...)
}

func MergeN[A any](buffer int, cs ...<-chan A) <-chan A {
	return MergeUntil(nil, buffer, cs...)
}

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

func FanOut[A any](c <-chan A, size int) []<-chan A {
	return FanOutUntil(nil, 0, c, size)
}

func FanOut1[A any](c <-chan A, size int) []<-chan A {
	return FanOutUntil(nil, 1, c, size)
}

func FanOutN[A any](buffer int, c <-chan A, size int) []<-chan A {
	return FanOutUntil(nil, buffer, c, size)
}
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

func Concat[A any](cs ...<-chan A) <-chan A {
	return ConcatUntil(nil, 0, cs...)
}
func Concat1[A any](cs ...<-chan A) <-chan A {
	return ConcatUntil(nil, 1, cs...)
}
func ConcatN[A any](buffer int, cs ...<-chan A) <-chan A {
	return ConcatUntil(nil, buffer, cs...)
}

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

func Filter[A any](c <-chan A, include func(a A) bool) <-chan A {
	return FilterUntil(nil, 0, c, include)
}

func Filter1[A any](c <-chan A, include func(a A) bool) <-chan A {
	return FilterUntil(nil, 1, c, include)
}

func FilterN[A any](buffer int, c <-chan A, include func(a A) bool) <-chan A {
	return FilterUntil(nil, buffer, c, include)
}
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

func Compact[A any](c <-chan A, equal func(a, b A) bool) <-chan A {
	return CompactUntil(nil, 0, c, equal)
}

func Compact1[A any](c <-chan A, equal func(a, b A) bool) <-chan A {
	return CompactUntil(nil, 1, c, equal)
}

func CompactN[A any](buffer int, c <-chan A, equal func(a, b A) bool) <-chan A {
	return CompactUntil(nil, buffer, c, equal)
}
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

func Partition[A any](c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return PartitionUntil(nil, 0, c, predicate)
}

func Partition1[A any](c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return PartitionUntil(nil, 1, c, predicate)
}

func PartitionN[A any](buffer int, c <-chan A, predicate func(a A) bool) (satisfied, notSatisfied <-chan A) {
	return PartitionUntil(nil, buffer, c, predicate)
}

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

func TakeWhile[A any](c <-chan A, take func(a A) bool) <-chan A {
	return TakeWhileUntil(nil, 0, c, take)
}

func TakeWhile1[A any](c <-chan A, take func(a A) bool) <-chan A {
	return TakeWhileUntil(nil, 1, c, take)
}

func TakeWhileN[A any](buffer int, c <-chan A, take func(a A) bool) <-chan A {
	return TakeWhileUntil(nil, buffer, c, take)
}

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

func Take[A any](c <-chan A, i int) <-chan A {
	return TakeUntil(nil, 0, c, i)
}

func Take1[A any](c <-chan A, i int) <-chan A {
	return TakeUntil(nil, 1, c, i)
}

func TakeN[A any](buffer int, c <-chan A, i int) <-chan A {
	return TakeUntil(nil, buffer, c, i)
}

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

func Drop[A any](c <-chan A, i int) <-chan A {
	return DropUntil(nil, 0, c, i)
}
func Drop1[A any](c <-chan A, i int) <-chan A {
	return DropUntil(nil, 1, c, i)
}

func DropN[A any](buffer int, c <-chan A, i int) <-chan A {
	return DropUntil(nil, buffer, c, i)
}

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

func DropWhile[A any](c <-chan A, drop func(a A) bool) <-chan A {
	return DropWhileUntil(nil, 0, c, drop)
}
func DropWhile1[A any](c <-chan A, drop func(a A) bool) <-chan A {
	return DropWhileUntil(nil, 1, c, drop)
}

func DropWhileN[A any](buffer int, c <-chan A, drop func(a A) bool) <-chan A {
	return DropWhileUntil(nil, buffer, c, drop)

}

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

func Zip[A any, B any, C any](ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	return ZipUntil(nil, 0, ac, bc, zipper)
}
func Zip1[A any, B any, C any](ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	return ZipUntil(nil, 1, ac, bc, zipper)
}
func ZipN[A any, B any, C any](buffer int, ac <-chan A, bc <-chan B, zipper func(a A, b B) C) <-chan C {
	return ZipUntil(nil, buffer, ac, bc, zipper)
}

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

func Unzip[A any, B any, C any](zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	return UnzipUntil(nil, 0, zipped, unzipper)
}

func Unzip1[A any, B any, C any](zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	return UnzipUntil(nil, 1, zipped, unzipper)
}

func UnzipN[A any, B any, C any](buffer int, zipped <-chan C, unzipper func(c C) (A, B)) (<-chan A, <-chan B) {
	return UnzipUntil(nil, buffer, zipped, unzipper)
}
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

func Collect[A any](c <-chan A) []A {
	return CollectUntil(nil, c)
}

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

func Readers[A any](chans ...chan A) []<-chan A {
	readers := make([]<-chan A, len(chans))
	for i := range chans {
		readers[i] = chans[i]
	}
	return readers
}

func Writers[A any](chans ...chan A) []chan<- A {
	readers := make([]chan<- A, len(chans))
	for i := range chans {
		readers[i] = chans[i]
	}
	return readers
}
