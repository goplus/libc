package libc

func Fdiml(x float64, y float64) float64 {
	return float64(Fdim(float64(x), float64(y)))
}
