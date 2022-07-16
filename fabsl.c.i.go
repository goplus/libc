package libc

func fabsl(x float64) float64 {
	return float64(fabs(float64(x)))
}
