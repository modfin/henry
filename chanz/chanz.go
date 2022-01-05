package chanz

import (
	"sync"
)

func Map0[A any, B any](in <-chan A, mapper func(a A) B) <-chan B {
	return MapDone(nil, 0, in, mapper)
}
func Map1[A any, B any](in <-chan A, mapper func(a A) B) <-chan B {
	return MapDone(nil, 1, in, mapper)
}

func Map[A any, B any](buffer int, in <-chan A, mapper func(a A) B) <-chan B {
	return MapDone(nil, buffer, in, mapper)
}

func MapDone[A any, B any](done <-chan interface{}, buffer int, in <-chan A, mapper func(a A) B) <-chan B {
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

func Generate[A any](buffer int, elements ...A) <-chan A {
	return GenerateDone(nil, buffer, elements...)
}

func GenerateDone[A any](done <-chan interface{}, buffer int, elements ...A) <-chan A {
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

func Merge[A any](buffer int, cs ...<-chan A) <-chan A {
	return MergeDone(nil, buffer, cs...)
}

func MergeDone[A any](done <-chan interface{}, buffer int, cs ...<-chan A) <-chan A {
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

func OrDone(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-OrDone(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
