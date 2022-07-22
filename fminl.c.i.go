package libc

func Fminl(x float64, y float64) float64 {
	return float64(Fmin(float64(x), float64(y)))
}
