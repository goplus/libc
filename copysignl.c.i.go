package libc

func Copysignl(x float64, y float64) float64 {
	return float64(Copysign(float64(x), float64(y)))
}
