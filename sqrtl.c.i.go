package libc

func Sqrtl(x float64) float64 {
	return float64(Sqrt(float64(x)))
}
