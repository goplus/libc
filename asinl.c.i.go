package libc

func Asinl(x float64) float64 {
	return float64(Asin(float64(x)))
}
