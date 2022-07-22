package libc

func Atan2l(y float64, x float64) float64 {
	return float64(Atan2(float64(y), float64(x)))
}
