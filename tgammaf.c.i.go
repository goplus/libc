package libc

func Tgammaf(x float32) float32 {
	return float32(Tgamma(float64(x)))
}
