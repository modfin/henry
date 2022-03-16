package numberz

type Numbers interface {
	~byte | compare.Integer | compare.Float
}

type SignedNumbers interface {
	compare.Signed | compare.Float
}

func MapNegate[N Numbers](n N) N {
	return -n
}
func MapAbs[N Numbers](n N) N {
	if n < 0 {
		return -n
	}
	return n
}

func MapFloat64[N Numbers](n N) float64 {
	return float64(n)
}
func MapFloat32[N Numbers](n N) float32 {
	return float32(n)
}

func MapInt[N Numbers](n N) int {
	return int(n)
}
func MapInt8[N Numbers](n N) int8 {
	return int8(n)
}
func MapInt16[N Numbers](n N) int16 {
	return int16(n)
}
func MapInt32[N Numbers](n N) int32 {
	return int32(n)
}
func MapInt64[N Numbers](n N) int64 {
	return int64(n)
}

func MapUInt[N Numbers](n N) uint {
	return uint(n)
}
func MapUInt8[N Numbers](n N) uint8 {
	return uint8(n)
}
func MapUInt16[N Numbers](n N) uint16 {
	return uint16(n)
}
func MapUInt32[N Numbers](n N) uint32 {
	return uint32(n)
}
func MapUInt64[N Numbers](n N) uint64 {
	return uint64(n)
}

func MapByte[N Numbers](n N) byte {
	return byte(n)
}
