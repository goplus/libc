package libc

func Ilogbl(x float64) int32 {
	return Ilogb(float64(x))
}
