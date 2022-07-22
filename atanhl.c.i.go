package libc

func Atanhl(x float64) float64 {
	return float64(Atanh(float64(x)))
}
