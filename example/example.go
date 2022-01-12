package main

import (
	"fmt"
	"github.com/modfin/go18exp/result"
	"github.com/modfin/go18exp/slicez"
	"github.com/modfin/go18exp/slicez/pipe"
	"strconv"
)

func main() {

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}

	squares := pipe.Of(slice).
		Map(func(a int) int { // Squaring numbers
			return a * a
		}).
		Filter(func(a int) bool { // Filtering even
			return a%2 == 0
		}).
		Slice()
	fmt.Println(squares)
	// [4 16 36 64]

	sum := slicez.Fold(squares, func(acc int, val int) int {
		return acc + val
	}, 0)
	// 120

	avg := sum / len(squares)
	// 30

	small, big := slicez.Partition(squares, func(i int) bool {
		return i < avg
	})
	fmt.Println(small, big)
	// [4 16] [36 64]

	zipped := slicez.Zip(small, big, func(s int, b int) string {
		return fmt.Sprintf("(%d, %d)", s, b)
	})
	fmt.Println(zipped)
	// [(4, 36) (16, 64)]

	// This is all well and good, but does not cover error handling.
	// Introducing error handling in map function could give the following api
	//
	// mapped, err := Map(slice, func(i int) (string, error){
	//	  i, err := strconv.ParseInt(str, 10, 64)
	//	  return int(i), err
	//})
	//
	//However, this would make the slicez api harder to work with in the usecases where errors is not needed to be returned.
	//Using something like wrapping, boxing or optional might be a better solution
	//eg.

	strslice := []string{"1", "2", "NaN", "4", "inf"}
	maybeInts := slicez.Map(strslice, func(str string) result.Result[int] {
		i, err := strconv.ParseInt(str, 10, 64)
		return result.From(int(i), err)
	})
	fmt.Println(maybeInts)
	// [{1} {2} {strconv.ParseInt: parsing "NaN": invalid syntax} {4} {strconv.ParseInt: parsing "inf": invalid syntax}]

	fmt.Println(result.SliceOk(maybeInts))
	// false

	fmt.Println(result.ErrorOfSlice(maybeInts))
	// strconv.ParseInt: parsing "NaN": invalid syntax

	fmt.Println(result.ErrorsOfSlice(maybeInts))
	// [strconv.ParseInt: parsing "NaN": invalid syntax strconv.ParseInt: parsing "inf": invalid syntax]

	resultInts := slicez.Filter(maybeInts, result.Ok[int])
	//or
	//resultInts := slicez.Filter(maybeInts, func(a result.Result[int]) bool {
	//	return a.Ok()
	//})
	fmt.Println(resultInts)
	// [{1} {2} {4}]

	ints := result.ValuesOfSlice(resultInts)
	//or
	//int := slicez.Map(resultInts, result.ValueOf[int])
	//or
	//ints := slicez.Map(resultInts, func(a result.Result[int]) int {
	//	return a.ValueOf()
	//})
	fmt.Println(ints)
	// [1 2 4]

}
