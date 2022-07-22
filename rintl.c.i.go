package libc

func Rintl(x float64) float64 {
	return float64(Rint(float64(x)))
}
