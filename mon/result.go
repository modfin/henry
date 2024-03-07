package mon

// https://github.com/samber/mo

import "fmt"

// Ok builds a Result when value is valid.
func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		isErr: false,
	}
}

// Err builds a Result when value is invalid.
func Err[T any](err error) Result[T] {
	return Result[T]{
		err:   err,
		isErr: true,
	}
}

// Errf builds a Result when value is invalid.
// Errf formats according to a format specifier and returns the error as a value that satisfies Result[T].
func Errf[T any](format string, a ...any) Result[T] {
	return Err[T](fmt.Errorf(format, a...))
}

// TupleToResult convert a pair of T and error into a Result.
// Play: https://go.dev/play/p/KWjfqQDHQwa
func TupleToResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// Try returns either a Ok or Err object.
func Try[T any](f func() (T, error)) Result[T] {
	return TupleToResult(f())
}

// Result represents a result of an action having one
// of the following output: success or failure.
// An instance of Result is an instance of either Ok or Err.
type Result[T any] struct {
	isErr bool
	value T
	err   error
}

// Ok returns true when value is valid.
func (r Result[T]) Ok() bool {
	return !r.isErr
}

// Error returns error when value is invalid or nil.
func (r Result[T]) Error() error {
	return r.err
}

// Get returns value and error.
func (r Result[T]) Get() (T, error) {
	if r.isErr {
		return empty[T](), r.err
	}

	return r.value, nil
}

// MustGet returns value when Result is valid or panics.
func (r Result[T]) MustGet() T {
	if r.isErr {
		panic(r.err)
	}

	return r.value
}

// OrElse returns value when Result is valid or default value.
func (r Result[T]) OrElse(fallback T) T {
	if r.isErr {
		return fallback
	}

	return r.value
}

// OrEmpty returns value when Result is valid or empty value.
func (r Result[T]) OrEmpty() T {
	return r.value
}

// ForEach executes the given side-effecting function if Result is valid.
func (r Result[T]) ForEach(mapper func(value T)) {
	if !r.isErr {
		mapper(r.value)
	}
}

// Match executes the first function if Result is valid and second function if invalid.
// It returns a new Result.
func (r Result[T]) Match(onSuccess func(value T) (T, error), onError func(err error) (T, error)) Result[T] {
	if r.isErr {
		return TupleToResult(onError(r.err))
	}
	return TupleToResult(onSuccess(r.value))
}

// Map executes the mapper function if Result is valid. It returns a new Result.
func (r Result[T]) Map(mapper func(value T) (T, error)) Result[T] {
	if !r.isErr {
		return TupleToResult(mapper(r.value))
	}

	return Err[T](r.err)
}

// MapErr executes the mapper function if Result is invalid. It returns a new Result.
func (r Result[T]) MapErr(mapper func(error) (T, error)) Result[T] {
	if r.isErr {
		return TupleToResult(mapper(r.err))
	}

	return Ok(r.value)
}

// FlatMap executes the mapper function if Result is valid. It returns a new Result.
func (r Result[T]) FlatMap(mapper func(value T) Result[T]) Result[T] {
	if !r.isErr {
		return mapper(r.value)
	}

	return Err[T](r.err)
}
