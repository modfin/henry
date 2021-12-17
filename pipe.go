package henry

import "github.com/crholm/henry/maybe"

func PipeOf[A any](a []A) Pipe[A] {
	return Pipe[A]{
		in: a,
	}
}

type Pipe[A any] struct {
	in []A
}

func (p Pipe[A]) Slice() []A {
	return p.in
}

func (p Pipe[A]) Each(apply func(i int, a A)) Pipe[A] {
	Each(p.in, apply)
	return p
}

func (p Pipe[A]) Concat(slices ...[]A) Pipe[A] {
	return PipeOf(Concat(append([][]A{p.in}, slices...)...))
}

func (p Pipe[A]) Tail() Pipe[A] {
	return PipeOf(Tail(p.in))
}
func (p Pipe[A]) Head() maybe.Value[A] {
	return Head(p.in)
}
func (p Pipe[A]) Last() maybe.Value[A] {
	return Last(p.in)
}

func (p Pipe[A]) Reverse() Pipe[A] {
	return PipeOf(Reverse(p.in))
}

func (p Pipe[A]) Nth(i int) maybe.Value[A] {
	return Nth(p.in, i)
}

func (p Pipe[A]) TakeLeft(i int) Pipe[A] {
	return PipeOf(TakeLeft(p.in, i))
}
func (p Pipe[A]) TakeRight(i int) Pipe[A] {
	return PipeOf(TakeRight(p.in, i))
}

func (p Pipe[A]) TakeLeftWhile(take func(i int, a A) bool) Pipe[A] {
	return PipeOf(TakeLeftWhile(p.in, take))
}
func (p Pipe[A]) TakeRightWhile(take func(i int, a A) bool) Pipe[A] {
	return PipeOf(TakeRightWhile(p.in, take))
}

func (p Pipe[A]) DropLeft(i int) Pipe[A] {
	return PipeOf(DropLeft(p.in, i))
}

func (p Pipe[A]) DropRight(i int) Pipe[A] {
	return PipeOf(DropRight(p.in, i))
}

func (p Pipe[A]) DropLeftWhile(drop func(i int, a A) bool) Pipe[A] {
	return PipeOf(DropLeftWhile(p.in, drop))
}

func (p Pipe[A]) DropRightWhile(drop func(i int, a A) bool) Pipe[A] {
	return PipeOf(DropRightWhile(p.in, drop))
}

func (p Pipe[A]) Filter(include func(i int, a A) bool) Pipe[A] {
	return PipeOf(Filter(p.in, include))
}

func (p Pipe[A]) Reject(exclude func(i int, a A) bool) Pipe[A] {
	return PipeOf(Reject(p.in, exclude))
}

func (p Pipe[A]) Every(predicate func(a A) bool) bool {
	return Every(p.in, predicate)
}
func (p Pipe[A]) Some(predicate func(a A) bool) bool {
	return Some(p.in, predicate)
}
func (p Pipe[A]) None(predicate func(a A) bool) bool {
	return None(p.in, predicate)
}
func (p Pipe[A]) Has(needle A, predicate func(a A, b A) bool) bool {
	return Has(p.in, needle, predicate)
}
