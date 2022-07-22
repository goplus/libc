package libc

func Acosl(x float64) float64 {
	return float64(Acos(float64(x)))
}
