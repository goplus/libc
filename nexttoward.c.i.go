package libc

func Nexttoward(x float64, y float64) float64 {
	return Nextafter(x, float64(y))
}
