package libc

func Fmodl(x float64, y float64) float64 {
	return float64(Fmod(float64(x), float64(y)))
}
