package libc

func Tgammal(x float64) float64 {
	return float64(Tgamma(float64(x)))
}
