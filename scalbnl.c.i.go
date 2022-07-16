package libc

func scalbnl(x float64, n int32) float64 {
	return float64(scalbn(float64(x), n))
}
