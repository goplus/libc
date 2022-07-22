package libc

func Asinhl(x float64) float64 {
	return float64(Asinh(float64(x)))
}
