package libc

func Coshl(x float64) float64 {
	return float64(Cosh(float64(x)))
}
