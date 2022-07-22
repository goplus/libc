package libc

func Nearbyint(x float64) float64 {
	x = Rint(x)
	return x
}
