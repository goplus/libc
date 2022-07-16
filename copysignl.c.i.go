package libc

func copysignl(x float64, y float64) float64 {
	return float64(copysign(float64(x), float64(y)))
}
