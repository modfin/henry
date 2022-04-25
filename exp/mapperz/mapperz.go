package mapperz

import (
	"github.com/modfin/henry/compare"
	"github.com/modfin/henry/exp/numberz"
	"github.com/modfin/henry/slicez"
)

type IElm[E any] struct {
	Element E
	Index   int
}

// Indexed wraps each element in a IElm struct and adds the current index in the slice of the element
//  eg.
//  slice := []int{1,2,3,4,5}
//  everyOtherIndexed := slicez.Filter(slices.Indexed(slice), func(e slicez.IElm[int]) bool { return e.Index % 2 == 0})
//  everyOther := slicez.Map(everyOtherIndexed, func(e slicez.IElm[int]) int { return e.Element})
//  // [1 3 5]
func Indexed[A any](slice []A) []IElm[A] {
	var index = -1
	return slicez.Map(slice, func(a A) IElm[A] {
		index++
		return IElm[A]{
			Element: a,
			Index:   index,
		}
	})
}
func Elements[A any](slice []IElm[A]) []A {
	return slicez.Map(slice, func(a IElm[A]) A {
		return a.Element
	})
}

// Negate will return -1*n
func Negate[N compare.Number](slice []N) []N {
	return slicez.Map(slice, numberz.Negate[N])
}

// Abs will return the absolut number of n
func Abs[N compare.Number](slice []N) []N {
	return slicez.Map(slice, numberz.Abs[N])
}

// CastFloat64 will cast n to a float64
func CastFloat64[N compare.Number](slice []N) []float64 {
	return slicez.Map(slice, numberz.CastFloat64[N])
}

// CastFloat32 will cast n to a float32
func CastFloat32[N compare.Number](slice []N) []float32 {
	return slicez.Map(slice, numberz.CastFloat32[N])
}

// CastInt will cast n to a int
func CastInt[N compare.Number](slice []N) []int {
	return slicez.Map(slice, numberz.CastInt[N])
}

// CastInt8 will cast n to a int8
func CastInt8[N compare.Number](slice []N) []int8 {
	return slicez.Map(slice, numberz.CastInt8[N])
}

// CastInt16 will cast n to a int16
func CastInt16[N compare.Number](slice []N) []int16 {
	return slicez.Map(slice, numberz.CastInt16[N])
}

// CastInt32 will cast n to a int32
func CastInt32[N compare.Number](slice []N) []int32 {
	return slicez.Map(slice, numberz.CastInt32[N])
}

// CastInt64 will cast n to a int64
func CastInt64[N compare.Number](slice []N) []int64 {
	return slicez.Map(slice, numberz.CastInt64[N])
}

// CastUInt will cast n to an uint
func CastUInt[N compare.Number](slice []N) []uint {
	return slicez.Map(slice, numberz.CastUInt[N])
}

// CastUInt8 will cast n to an uint8
func CastUInt8[N compare.Number](slice []N) []uint8 {
	return slicez.Map(slice, numberz.CastUInt8[N])
}

// CastUInt16 will cast n to an uint16
func CastUInt16[N compare.Number](slice []N) []uint16 {
	return slicez.Map(slice, numberz.CastUInt16[N])
}

// CastUInt32 will cast n to an uint32
func CastUInt32[N compare.Number](slice []N) []uint32 {
	return slicez.Map(slice, numberz.CastUInt32[N])
}

// CastUInt64 will cast n to an uint64
func CastUInt64[N compare.Number](slice []N) []uint64 {
	return slicez.Map(slice, numberz.CastUInt64[N])
}

// CastByte will cast n to an byte
func CastByte[N compare.Number](slice []N) []byte {
	return slicez.Map(slice, numberz.CastByte[N])
}
