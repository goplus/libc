package libc

func Acoshl(x float64) float64 {
	return float64(Acosh(float64(x)))
}
