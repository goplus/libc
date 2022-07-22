package libc

func Sinhl(x float64) float64 {
	return float64(Sinh(float64(x)))
}
