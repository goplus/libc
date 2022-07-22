package libc

func Fmaxl(x float64, y float64) float64 {
	return float64(Fmax(float64(x), float64(y)))
}
