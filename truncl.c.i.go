package libc

func Truncl(x float64) float64 {
	return float64(Trunc(float64(x)))
}
