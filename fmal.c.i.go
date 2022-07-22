package libc

func Fmal(x float64, y float64, z float64) float64 {
	return float64(Fma(float64(x), float64(y), float64(z)))
}
