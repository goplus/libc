package libc

func Exp2l(x float64) float64 {
	return float64(Exp2(float64(x)))
}
