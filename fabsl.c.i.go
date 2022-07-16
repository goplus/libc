package libc

func Fabsl(x float64) float64 {
	return float64(Fabs(float64(x)))
}
