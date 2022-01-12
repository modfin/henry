package pipe

import (
	"github.com/modfin/go18exp/slicez"
)

func Of[A any](a []A) Pipe[A] {
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

func (p Pipe[A]) Peek(apply func(a A)) Pipe[A] {
	slicez.Each(p.in, apply)
	return p
}

func (p Pipe[A]) Concat(slices ...[]A) Pipe[A] {
	return Of(slicez.Concat(append([][]A{p.in}, slices...)...))
}

func (p Pipe[A]) Tail() Pipe[A] {
	return Of(slicez.Tail(p.in))
}
func (p Pipe[A]) Head() (A, error) {
	return slicez.Head(p.in)
}
func (p Pipe[A]) Last() (A, error) {
	return slicez.Last(p.in)
}

func (p Pipe[A]) Reverse() Pipe[A] {
	return Of(slicez.Reverse(p.in))
}

func (p Pipe[A]) Nth(i int) A {
	return slicez.Nth(p.in, i)
}

func (p Pipe[A]) Take(i int) Pipe[A] {
	return Of(slicez.Take(p.in, i))
}
func (p Pipe[A]) TakeRight(i int) Pipe[A] {
	return Of(slicez.TakeRight(p.in, i))
}

func (p Pipe[A]) TakeWhile(take func(a A) bool) Pipe[A] {
	return Of(slicez.TakeWhile(p.in, take))
}
func (p Pipe[A]) TakeRightWhile(take func(a A) bool) Pipe[A] {
	return Of(slicez.TakeRightWhile(p.in, take))
}

func (p Pipe[A]) Drop(i int) Pipe[A] {
	return Of(slicez.Drop(p.in, i))
}

func (p Pipe[A]) DropRight(i int) Pipe[A] {
	return Of(slicez.DropRight(p.in, i))
}

func (p Pipe[A]) DropWhile(drop func(a A) bool) Pipe[A] {
	return Of(slicez.DropWhile(p.in, drop))
}

func (p Pipe[A]) DropRightWhile(drop func(a A) bool) Pipe[A] {
	return Of(slicez.DropRightWhile(p.in, drop))
}

func (p Pipe[A]) Filter(include func(a A) bool) Pipe[A] {
	return Of(slicez.Filter(p.in, include))
}

func (p Pipe[A]) Reject(exclude func(a A) bool) Pipe[A] {
	return Of(slicez.Reject(p.in, exclude))
}

func (p Pipe[A]) Map(f func(a A) A) Pipe[A] {
	return Of(slicez.Map(p.in, f))
}

func (p Pipe[A]) Fold(combined func(accumulator A, val A) A, accumulator A) A {
	return slicez.Fold(p.in, combined, accumulator)
}

func (p Pipe[A]) FoldRight(combined func(accumulator A, val A) A, accumulator A) A {
	return slicez.Fold(p.in, combined, accumulator)
}

func (p Pipe[A]) Every(predicate func(a A) bool) bool {
	return slicez.EveryFunc(p.in, predicate)
}
func (p Pipe[A]) Some(predicate func(a A) bool) bool {
	return slicez.SomeFunc(p.in, predicate)
}
func (p Pipe[A]) None(predicate func(a A) bool) bool {
	return slicez.NoneFunc(p.in, predicate)
}

func (p Pipe[A]) Partition(predicate func(a A) bool) (satisfied, notSatisfied []A) {
	return slicez.Partition(p.in, predicate)
}

func (p Pipe[A]) Sample(n int) Pipe[A] {
	return Of(slicez.Sample(p.in, n))
}
func (p Pipe[A]) Shuffle() Pipe[A] {
	return Of(slicez.Shuffle(p.in))
}
func (p Pipe[A]) SortFunc(less func(a, b A) bool) Pipe[A] {
	return Of(slicez.SortFunc(p.in, less))
}

func (p Pipe[A]) Compact(equal func(a, b A) bool) Pipe[A] {
	return Of(slicez.CompactFunc(p.in, equal))
}

func (p Pipe[A]) Count() int {
	return len(p.in)
}
