package libc

func Roundl(x float64) float64 {
	return float64(Round(float64(x)))
}
