package libc

func Nextafterl(x float64, y float64) float64 {
	return float64(Nextafter(float64(x), float64(y)))
}
