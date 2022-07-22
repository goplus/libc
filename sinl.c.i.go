package libc

func Sinl(x float64) float64 {
	return float64(Sin(float64(x)))
}
