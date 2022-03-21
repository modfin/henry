package numberz

import "github.com/modfin/henry/compare"

// Negate will return -1*n
func Negate[N compare.Number](n N) N {
	return -n
}

// Abs will return the absolut number of n
func Abs[N compare.Number](n N) N {
	if n < 0 {
		return -n
	}
	return n
}

// CastFloat64 will cast n to a float64
func CastFloat64[N compare.Number](n N) float64 {
	return float64(n)
}

// CastFloat32 will cast n to a float32
func CastFloat32[N compare.Number](n N) float32 {
	return float32(n)
}

// CastInt will cast n to a int
func CastInt[N compare.Number](n N) int {
	return int(n)
}

// CastInt8 will cast n to a int8
func CastInt8[N compare.Number](n N) int8 {
	return int8(n)
}

// CastInt16 will cast n to a int16
func CastInt16[N compare.Number](n N) int16 {
	return int16(n)
}

// CastInt32 will cast n to a int32
func CastInt32[N compare.Number](n N) int32 {
	return int32(n)
}

// CastInt64 will cast n to a int64
func CastInt64[N compare.Number](n N) int64 {
	return int64(n)
}

// CastUInt will cast n to an uint
func CastUInt[N compare.Number](n N) uint {
	return uint(n)
}

// CastUInt8 will cast n to an uint8
func CastUInt8[N compare.Number](n N) uint8 {
	return uint8(n)
}

// CastUInt16 will cast n to an uint16
func CastUInt16[N compare.Number](n N) uint16 {
	return uint16(n)
}

// CastUInt32 will cast n to an uint32
func CastUInt32[N compare.Number](n N) uint32 {
	return uint32(n)
}

// CastUInt64 will cast n to an uint64
func CastUInt64[N compare.Number](n N) uint64 {
	return uint64(n)
}

// CastByte will cast n to an byte
func CastByte[N compare.Number](n N) byte {
	return byte(n)
}
