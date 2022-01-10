package result

import (
	"fmt"
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

func Error[A any](err error) Result[A] {
	var zero A
	return Result[A]{
		value: zero,
		err:   err,
	}
}

func SliceValues[A any](results []Result[A]) []A {
	values := slicez.Filter(results, ValueFilter[A])
	return slicez.Map(values, ValueMapper[A])
}

func SliceErrors[A any](results []Result[A]) []error {
	errs := slicez.Filter(results, ErrorFilter[A])
	return slicez.Map(errs, ErrorMapper[A])
}

func SliceError[A any](results []Result[A]) error {
	err, _ := slicez.Head(SliceErrors(results))
	return err
}
func SliceOk[A any](results []Result[A]) bool {
	return slicez.None(results, func(a Result[A]) bool {
		return !a.Ok()
	})
}

func ValueMapper[A any](res Result[A]) A {
	return res.Value()
}
func ValueFilter[A any](res Result[A]) bool {
	return res.Ok()
}

func ErrorMapper[A any](res Result[A]) error {
	return res.Error()
}
func ErrorFilter[A any](res Result[A]) bool {
	return !res.Ok()
}
