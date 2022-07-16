package libc

func fmodl(x float64, y float64) float64 {
	return float64(fmod(float64(x), float64(y)))
}
