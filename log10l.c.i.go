package libc

func Log10l(x float64) float64 {
	return float64(Log10(float64(x)))
}
