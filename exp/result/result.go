package result

import (
	"fmt"
	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/slicez"
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

func (r Result[A]) Unwrap() (A, error) {
	return r.value, r.err
}

func FromValue[A any](a A) Result[A] {
	return Result[A]{
		value: a,
		err:   nil,
	}
}
func FromError[A any](err error) Result[A] {
	var z A
	return Result[A]{
		value: z,
		err:   err,
	}
}

func From[A any](a A, err error) Result[A] {
	return Result[A]{
		value: a,
		err:   err,
	}
}

func Values[A any](results []Result[A]) []A {
	vals := slicez.Filter(results, ok[A])
	return slicez.Map(vals, value[A])
}

func Errors[A any](results []Result[A]) []error {
	errs := slicez.Filter(results, compare.NegateOf(ok[A]))
	return slicez.Map(errs, err[A])
}

func Error[A any](results []Result[A]) error {
	r, found := slicez.Find(results, compare.NegateOf(ok[A]))
	if !found {
		return nil
	}
	return r.err
}

func Unwrap[A any](results []Result[A]) ([]A, error) {
	return Values(results), Error(results)
}

func value[A any](res Result[A]) A {
	return res.Value()
}
func ok[A any](res Result[A]) bool {
	return res.Ok()
}
func err[A any](res Result[A]) error {
	return res.Error()
}
