package libc

func Atanl(x float64) float64 {
	return float64(Atan(float64(x)))
}
