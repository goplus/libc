package libc

func Expm1l(x float64) float64 {
	return float64(Expm1(float64(x)))
}
