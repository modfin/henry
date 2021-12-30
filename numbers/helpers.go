package numbers

import "constraints"

type Numbers interface {
	~byte | constraints.Integer | constraints.Float
}

func MapFloat64[N Numbers](_ int, n N) float64 {
	return float64(n)
}
func MapFloat32[N Numbers](_ int, n N) float32 {
	return float32(n)
}

func MapInt[N Numbers](_ int, n N) int {
	return int(n)
}
func MapInt8[N Numbers](_ int, n N) int8 {
	return int8(n)
}
func MapInt16[N Numbers](_ int, n N) int16 {
	return int16(n)
}
func MapInt32[N Numbers](_ int, n N) int32 {
	return int32(n)
}
func MapInt64[N Numbers](_ int, n N) int64 {
	return int64(n)
}

func MapUInt[N Numbers](_ int, n N) uint {
	return uint(n)
}
func MapUInt8[N Numbers](_ int, n N) uint8 {
	return uint8(n)
}
func MapUInt16[N Numbers](_ int, n N) uint16 {
	return uint16(n)
}
func MapUInt32[N Numbers](_ int, n N) uint32 {
	return uint32(n)
}
func MapUInt64[N Numbers](_ int, n N) uint64 {
	return uint64(n)
}

func MapByte[N Numbers](_ int, n N) byte {
	return byte(n)
}
