package libc

func Logl(x float64) float64 {
	return float64(Log(float64(x)))
}
