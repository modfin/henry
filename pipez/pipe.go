// Package pipez provides a fluent, chainable API for slice operations.
//
// Pipe wraps slices and provides method chaining for functional programming operations.
// This allows for readable, pipeline-style data transformations:
//
//	result := pipez.Of([]int{1, 2, 3, 4, 5}).
//	    Filter(func(n int) bool { return n > 2 }).
//	    Map(func(n int) int { return n * 2 }).
//	    Slice()
//	// result = []int{6, 8, 10}
//
// Most methods return Pipe for continued chaining. Terminal operations like Head(), Last(),
// and Fold() return concrete values and break the chain.
//
// The pipez package is built on top of slicez, providing the same operations with a
// method-chaining interface instead of standalone functions.
package pipez

import (
	"github.com/modfin/henry/slicez"
)

// Of creates a Pipe from a slice, enabling method chaining.
//
// Example:
//
//	pipe := pipez.Of([]int{1, 2, 3, 4, 5})
//	result := pipe.Filter(func(n int) bool { return n%2 == 0 }).Slice()
//	// result = []int{2, 4}
func Of[A any](a []A) Pipe[A] {
	return a
}

// Pipe is a type alias for []A that provides method chaining for slice operations.
// It wraps slicez functions in a fluent API, allowing operations to be chained together.
//
// Pipe is not a new type but a type alias, so there's zero overhead - it's just
// a slice with methods attached. You can convert between []A and Pipe[A] freely.
//
// Example:
//
//	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	result := pipez.Of(data).
//	    Filter(func(n int) bool { return n > 5 }).
//	    Map(func(n int) int { return n * n }).
//	    Take(3).
//	    Slice()
//	// result = []int{36, 49, 64} (6², 7², 8²)
type Pipe[A any] []A

// Slice returns the underlying slice, breaking the chain.
// Use this at the end of a pipeline to get the final []A result.
//
// Example:
//
//	pipe := pipez.Of([]int{1, 2, 3}).Map(func(n int) int { return n * 2 })
//	slice := pipe.Slice()
//	// slice = []int{2, 4, 6}
func (p Pipe[A]) Slice() []A {
	return p
}

// Peek applies a function to each element for side effects, then returns the Pipe for chaining.
// Useful for logging, debugging, or performing actions without transforming data.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3}).
//	    Peek(func(n int) { fmt.Println("Processing:", n) }).
//	    Map(func(n int) int { return n * 2 }).
//	    Slice()
//	// Prints "Processing: 1", "Processing: 2", "Processing: 3" during execution
func (p Pipe[A]) Peek(apply func(a A)) Pipe[A] {
	slicez.ForEach(p, apply)
	return p
}

// Concat appends additional slices to the pipe, returning a new Pipe.
//
// Example:
//
//	pipez.Of([]int{1, 2}).Concat([]int{3, 4}, []int{5, 6}).Slice()
//	// Returns []int{1, 2, 3, 4, 5, 6}
func (p Pipe[A]) Concat(slices ...[]A) Pipe[A] {
	return slicez.Concat(append([][]A{p}, slices...)...)
}

// Tail returns all elements except the first, returning a new Pipe.
// Returns empty pipe if input has 0 or 1 elements.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4}).Tail().Slice()
//	// Returns []int{2, 3, 4}
func (p Pipe[A]) Tail() Pipe[A] {
	return slicez.Tail(p)
}

// Head returns the first element and an error if the pipe is empty.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	first, err := pipez.Of([]int{1, 2, 3}).Head()
//	// first = 1, err = nil
//
//	_, err := pipez.Of([]int{}).Head()
//	// err != nil (empty slice error)
func (p Pipe[A]) Head() (A, error) {
	return slicez.Head(p)
}

// Last returns the last element and an error if the pipe is empty.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	last, err := pipez.Of([]int{1, 2, 3}).Last()
//	// last = 3, err = nil
func (p Pipe[A]) Last() (A, error) {
	return slicez.Last(p)
}

// Initial returns all elements except the last, returning a new Pipe.
// Returns empty pipe if input has 0 or 1 elements.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4}).Initial().Slice()
//	// Returns []int{1, 2, 3}
func (p Pipe[A]) Initial() Pipe[A] {
	return slicez.Initial(p)
}

// Reverse returns a new Pipe with elements in reverse order.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3}).Reverse().Slice()
//	// Returns []int{3, 2, 1}
func (p Pipe[A]) Reverse() Pipe[A] {
	return slicez.Reverse(p)
}

// Nth returns the element at index i (0-based), panicking if out of bounds.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	elem := pipez.Of([]int{10, 20, 30}).Nth(1)
//	// elem = 20
func (p Pipe[A]) Nth(i int) A {
	return slicez.Nth(p, i)
}

// Take returns the first i elements, returning a new Pipe.
// Returns all elements if i > length.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).Take(3).Slice()
//	// Returns []int{1, 2, 3}
func (p Pipe[A]) Take(i int) Pipe[A] {
	return slicez.Take(p, i)
}

// TakeRight returns the last i elements, returning a new Pipe.
// Returns all elements if i > length.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).TakeRight(3).Slice()
//	// Returns []int{3, 4, 5}
func (p Pipe[A]) TakeRight(i int) Pipe[A] {
	return slicez.TakeRight(p, i)
}

// TakeWhile returns elements from the start while the predicate is true.
// Stops at the first element where predicate returns false.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).TakeWhile(func(n int) bool { return n < 4 }).Slice()
//	// Returns []int{1, 2, 3}
func (p Pipe[A]) TakeWhile(take func(a A) bool) Pipe[A] {
	return slicez.TakeWhile(p, take)
}

// TakeRightWhile returns elements from the end while the predicate is true.
// Stops at the first element from the end where predicate returns false.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).TakeRightWhile(func(n int) bool { return n > 3 }).Slice()
//	// Returns []int{4, 5}
func (p Pipe[A]) TakeRightWhile(take func(a A) bool) Pipe[A] {
	return slicez.TakeRightWhile(p, take)
}

// Drop removes the first i elements, returning a new Pipe with the remainder.
// Returns empty pipe if i >= length.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).Drop(2).Slice()
//	// Returns []int{3, 4, 5}
func (p Pipe[A]) Drop(i int) Pipe[A] {
	return slicez.Drop(p, i)
}

// DropRight removes the last i elements, returning a new Pipe with the remainder.
// Returns empty pipe if i >= length.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).DropRight(2).Slice()
//	// Returns []int{1, 2, 3}
func (p Pipe[A]) DropRight(i int) Pipe[A] {
	return slicez.DropRight(p, i)
}

// DropWhile removes elements from the start while the predicate is true.
// Returns elements starting from the first one where predicate returns false.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).DropWhile(func(n int) bool { return n < 4 }).Slice()
//	// Returns []int{4, 5}
func (p Pipe[A]) DropWhile(drop func(a A) bool) Pipe[A] {
	return slicez.DropWhile(p, drop)
}

// DropRightWhile removes elements from the end while the predicate is true.
// Returns elements up to the last one where predicate returns false.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).DropRightWhile(func(n int) bool { return n > 3 }).Slice()
//	// Returns []int{1, 2, 3}
func (p Pipe[A]) DropRightWhile(drop func(a A) bool) Pipe[A] {
	return slicez.DropRightWhile(p, drop)
}

// Filter returns elements that satisfy the predicate, removing others.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).Filter(func(n int) bool { return n%2 == 0 }).Slice()
//	// Returns []int{2, 4} (only even numbers)
func (p Pipe[A]) Filter(include func(a A) bool) Pipe[A] {
	return slicez.Filter(p, include)
}

// Reject returns elements that do NOT satisfy the predicate (inverse of Filter).
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).Reject(func(n int) bool { return n%2 == 0 }).Slice()
//	// Returns []int{1, 3, 5} (only odd numbers)
func (p Pipe[A]) Reject(exclude func(a A) bool) Pipe[A] {
	return slicez.Reject(p, exclude)
}

// Map transforms each element using the provided function.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3}).Map(func(n int) int { return n * n }).Slice()
//	// Returns []int{1, 4, 9}
func (p Pipe[A]) Map(f func(a A) A) Pipe[A] {
	return slicez.Map(p, f)
}

// Fold reduces the pipe from left to right, accumulating a result.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	sum := pipez.Of([]int{1, 2, 3, 4}).Fold(func(acc, val int) int { return acc + val }, 0)
//	// sum = 10
func (p Pipe[A]) Fold(combined func(accumulator A, val A) A, accumulator A) A {
	return slicez.Fold(p, combined, accumulator)
}

// FoldRight reduces the pipe from right to left, accumulating a result.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	result := pipez.Of([]string{"a", "b", "c"}).FoldRight(func(acc, val string) string {
//	    if acc == "" { return val }
//	    return val + "-" + acc
//	}, "")
//	// result = "a-b-c"
func (p Pipe[A]) FoldRight(combined func(accumulator A, val A) A, accumulator A) A {
	return slicez.FoldRight(p, combined, accumulator)
}

// Every returns true if all elements satisfy the predicate.
// Returns true for empty pipes (vacuous truth).
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	allPositive := pipez.Of([]int{1, 2, 3}).Every(func(n int) bool { return n > 0 })
//	// allPositive = true
//
//	allEven := pipez.Of([]int{1, 2, 3}).Every(func(n int) bool { return n%2 == 0 })
//	// allEven = false
func (p Pipe[A]) Every(predicate func(a A) bool) bool {
	return slicez.EveryBy(p, predicate)
}

// Some returns true if at least one element satisfies the predicate.
// Returns false for empty pipes.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	hasEven := pipez.Of([]int{1, 2, 3}).Some(func(n int) bool { return n%2 == 0 })
//	// hasEven = true
func (p Pipe[A]) Some(predicate func(a A) bool) bool {
	return slicez.SomeBy(p, predicate)
}

// None returns true if no elements satisfy the predicate.
// Returns true for empty pipes.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	noNegatives := pipez.Of([]int{1, 2, 3}).None(func(n int) bool { return n < 0 })
//	// noNegatives = true
func (p Pipe[A]) None(predicate func(a A) bool) bool {
	return slicez.NoneBy(p, predicate)
}

// Partition splits the pipe into two slices based on a predicate.
// First slice contains elements where predicate is true, second where false.
// This is a terminal operation that returns two slices (not a Pipe).
//
// Example:
//
//	even, odd := pipez.Of([]int{1, 2, 3, 4, 5}).Partition(func(n int) bool { return n%2 == 0 })
//	// even = []int{2, 4}, odd = []int{1, 3, 5}
func (p Pipe[A]) Partition(predicate func(a A) bool) (satisfied, notSatisfied []A) {
	return slicez.Partition(p, predicate)
}

// Sample returns n random elements from the pipe.
// Uses efficient algorithms (Fisher-Yates, swap-to-end) for uniform sampling.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).Sample(3).Slice()
//	// Might return []int{2, 5, 1} (3 random elements)
func (p Pipe[A]) Sample(n int) Pipe[A] {
	return slicez.Sample(p, n)
}

// Shuffle returns a new Pipe with elements in random order.
// Uses Fisher-Yates shuffle for uniform distribution.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3, 4, 5}).Shuffle().Slice()
//	// Might return []int{3, 1, 5, 2, 4}
func (p Pipe[A]) Shuffle() Pipe[A] {
	return slicez.Shuffle(p)
}

// SortFunc sorts the pipe using a custom comparison function.
// The less function should return true if a should come before b.
//
// Example:
//
//	pipez.Of([]int{3, 1, 4, 1, 5}).SortFunc(func(a, b int) bool { return a < b }).Slice()
//	// Returns []int{1, 1, 3, 4, 5}
func (p Pipe[A]) SortFunc(less func(a, b A) bool) Pipe[A] {
	return slicez.SortBy(p, less)
}

// Compact removes consecutive duplicate elements using the equality function.
// Only removes duplicates that are adjacent; non-consecutive duplicates are kept.
//
// Example:
//
//	pipez.Of([]int{1, 1, 2, 2, 2, 3, 3}).Compact(func(a, b int) bool { return a == b }).Slice()
//	// Returns []int{1, 2, 3}
func (p Pipe[A]) Compact(equal func(a, b A) bool) Pipe[A] {
	return slicez.CompactBy(p, equal)
}

// Count returns the number of elements in the pipe.
// This is a terminal operation that breaks the chain.
//
// Example:
//
//	count := pipez.Of([]int{1, 2, 3, 4, 5}).Count()
//	// count = 5
func (p Pipe[A]) Count() int {
	return len(p)
}

// Zip combines this pipe with another slice element-wise using a zipper function.
// Stops at the length of the shorter collection.
//
// Example:
//
//	pipez.Of([]int{1, 2, 3}).Zip([]int{10, 20, 30}, func(a, b int) int { return a + b }).Slice()
//	// Returns []int{11, 22, 33}
func (p Pipe[A]) Zip(b []A, zipper func(a, b A) A) Pipe[A] {
	return slicez.Zip(p, b, zipper)
}

// Unzip splits this pipe of pairs into two separate slices.
// This is a terminal operation that returns two slices (not a Pipe).
//
// Example:
//
//	pairs := []struct{ X, Y int }{{1, 10}, {2, 20}, {3, 30}}
//	xs, ys := pipez.Of(pairs).Unzip(func(p struct{ X, Y int }) (int, int) { return p.X, p.Y })
//	// xs = []int{1, 2, 3}, ys = []int{10, 20, 30}
func (p Pipe[A]) Unzip(unzipper func(a A) (A, A)) ([]A, []A) {
	return slicez.Unzip(p, unzipper)
}

// Interleave interleaves this pipe with additional slices round-robin style.
// Takes first element from each, then second from each, etc.
//
// Example:
//
//	pipez.Of([]int{1, 5, 9}).Interleave([]int{2, 6, 10}, []int{3, 7, 11}).Slice()
//	// Returns []int{1, 2, 3, 5, 6, 7, 9, 10, 11}
func (p Pipe[A]) Interleave(a ...[]A) Pipe[A] {
	return slicez.Interleave[A](append([][]A{p}, a...)...)
}
