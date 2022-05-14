package libc

func __signbitl(x float64) int32 {
	return __signbit(float64(x))
}
