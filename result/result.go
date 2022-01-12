package result

import (
	"fmt"
	"github.com/modfin/go18exp/compare"
	"github.com/modfin/go18exp/slicez"
)

type Result[A any] struct {
	value A
	err   error
}

func (r Result[A]) Value() A {
	return r.value
}

func (r Result[A]) ValueOr(or A) A {
	if r.err == nil {
		return r.value
	}
	return or
}

func (r Result[A]) String() string {
	if r.err == nil {
		return fmt.Sprintf("{%v}", r.value)
	}
	return fmt.Sprintf("{%v}", r.err)
}

func (r Result[A]) Error() error {
	return r.err
}

func (r Result[A]) Ok() bool {
	return r.err == nil
}

func SliceOf[A any](slice []A) []Result[A] {
	return slicez.Map(slice, Of[A])
}

func Of[A any](a A) Result[A] {
	return Result[A]{
		value: a,
		err:   nil,
	}
}

func From[A any](a A, err error) Result[A] {
	return Result[A]{
		value: a,
		err:   err,
	}
}

func ValuesOfSlice[A any](results []Result[A]) []A {
	vals := slicez.Filter(results, Ok[A])
	return slicez.Map(vals, ValueOf[A])
}

func ErrorsOfSlice[A any](results []Result[A]) []error {
	errs := slicez.Filter(results, compare.NegateOf(Ok[A]))
	return slicez.Map(errs, ErrorOf[A])
}

func ErrorOfSlice[A any](results []Result[A]) error {
	r, found := slicez.Find(results, compare.NegateOf(Ok[A]))
	if !found {
		return nil
	}
	return r.err
}
func SliceOk[A any](results []Result[A]) bool {
	return slicez.EveryFunc(results, Ok[A])
}

func ValueOf[A any](res Result[A]) A {
	return res.Value()
}
func Ok[A any](res Result[A]) bool {
	return res.Ok()
}

func ErrorOf[A any](res Result[A]) error {
	return res.Error()
}
