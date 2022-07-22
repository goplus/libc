package libc

func Hypotl(x float64, y float64) float64 {
	return float64(Hypot(float64(x), float64(y)))
}
