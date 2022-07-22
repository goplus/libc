package libc

func Expl(x float64) float64 {
	return float64(Exp(float64(x)))
}
