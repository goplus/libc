package libc

func Log1pl(x float64) float64 {
	return float64(Log1p(float64(x)))
}
