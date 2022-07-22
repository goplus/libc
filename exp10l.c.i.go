package libc

func Exp10l(x float64) float64 {
	return float64(Exp10(float64(x)))
}
