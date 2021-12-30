package pipe

import (
	"github.com/crholm/henry"
	"github.com/crholm/henry/maybe"
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

func (p Pipe[A]) Each(apply func(i int, a A)) Pipe[A] {
	henry.Each(p.in, apply)
	return p
}

func (p Pipe[A]) Concat(slices ...[]A) Pipe[A] {
	return Of(henry.Concat(append([][]A{p.in}, slices...)...))
}

func (p Pipe[A]) Tail() Pipe[A] {
	return Of(henry.Tail(p.in))
}
func (p Pipe[A]) Head() maybe.Value[A] {
	return henry.Head(p.in)
}
func (p Pipe[A]) Last() maybe.Value[A] {
	return henry.Last(p.in)
}

func (p Pipe[A]) Reverse() Pipe[A] {
	return Of(henry.Reverse(p.in))
}

func (p Pipe[A]) Nth(i int) maybe.Value[A] {
	return henry.Nth(p.in, i)
}

func (p Pipe[A]) TakeLeft(i int) Pipe[A] {
	return Of(henry.TakeLeft(p.in, i))
}
func (p Pipe[A]) TakeRight(i int) Pipe[A] {
	return Of(henry.TakeRight(p.in, i))
}

func (p Pipe[A]) TakeLeftWhile(take func(i int, a A) bool) Pipe[A] {
	return Of(henry.TakeLeftWhile(p.in, take))
}
func (p Pipe[A]) TakeRightWhile(take func(i int, a A) bool) Pipe[A] {
	return Of(henry.TakeRightWhile(p.in, take))
}

func (p Pipe[A]) DropLeft(i int) Pipe[A] {
	return Of(henry.DropLeft(p.in, i))
}

func (p Pipe[A]) DropRight(i int) Pipe[A] {
	return Of(henry.DropRight(p.in, i))
}

func (p Pipe[A]) DropLeftWhile(drop func(i int, a A) bool) Pipe[A] {
	return Of(henry.DropLeftWhile(p.in, drop))
}

func (p Pipe[A]) DropRightWhile(drop func(i int, a A) bool) Pipe[A] {
	return Of(henry.DropRightWhile(p.in, drop))
}

func (p Pipe[A]) Filter(include func(i int, a A) bool) Pipe[A] {
	return Of(henry.Filter(p.in, include))
}

func (p Pipe[A]) Reject(exclude func(i int, a A) bool) Pipe[A] {
	return Of(henry.Reject(p.in, exclude))
}

func (p Pipe[A]) Map(f func(i int, a A) A) Pipe[A] {
	return Of(henry.Map(p.in, f))
}

func (p Pipe[A]) FoldLeft(combined func(i int, accumulator A, val A) A, accumulator A) A {
	return henry.FoldLeft(p.in, combined, accumulator)
}

func (p Pipe[A]) FoldRight(combined func(i int, accumulator A, val A) A, accumulator A) A {
	return henry.FoldLeft(p.in, combined, accumulator)
}

func (p Pipe[A]) Every(predicate func(a A) bool) bool {
	return henry.Every(p.in, predicate)
}
func (p Pipe[A]) Some(predicate func(a A) bool) bool {
	return henry.Some(p.in, predicate)
}
func (p Pipe[A]) None(predicate func(a A) bool) bool {
	return henry.None(p.in, predicate)
}
func (p Pipe[A]) Has(needle A, predicate func(a A, b A) bool) bool {
	return henry.Has(p.in, needle, predicate)
}

func (p Pipe[A]) Partition(predicate func(i int, a A) bool) (satisfied, notSatisfied []A) {
	return henry.Partition(p.in, predicate)
}

func (p Pipe[A]) Sample(n int) Pipe[A] {
	return Of(henry.Sample(p.in, n))
}
func (p Pipe[A]) Shuffle() Pipe[A] {
	return Of(henry.Shuffle(p.in))
}
func (p Pipe[A]) Sort(less func(a, b A) bool) Pipe[A] {
	return Of(henry.Sort(p.in, less))
}

func (p Pipe[A]) Compact(equal func(a, b A) bool) Pipe[A] {
	return Of(henry.Compact(p.in, equal))
}

func (p Pipe[A]) Count() int {
	return len(p.in)
}
