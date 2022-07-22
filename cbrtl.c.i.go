package libc

func Cbrtl(x float64) float64 {
	return float64(Cbrt(float64(x)))
}
