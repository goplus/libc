package libc

func Tanl(x float64) float64 {
	return float64(Tan(float64(x)))
}
