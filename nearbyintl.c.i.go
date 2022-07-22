package libc

func Nearbyintl(x float64) float64 {
	return float64(Nearbyint(float64(x)))
}
