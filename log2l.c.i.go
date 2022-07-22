package libc

func Log2l(x float64) float64 {
	return float64(Log2(float64(x)))
}
