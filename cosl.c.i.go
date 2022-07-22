package libc

func Cosl(x float64) float64 {
	return float64(Cos(float64(x)))
}
