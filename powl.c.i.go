package libc

func Powl(x float64, y float64) float64 {
	return float64(Pow(float64(x), float64(y)))
}
