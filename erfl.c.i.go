package libc

func Erfl(x float64) float64 {
	return float64(Erf(float64(x)))
}
func Erfcl(x float64) float64 {
	return float64(Erfc(float64(x)))
}
