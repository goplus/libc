package libc

func Remainderl(x float64, y float64) float64 {
	return float64(Remainder(float64(x), float64(y)))
}
