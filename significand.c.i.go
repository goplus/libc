package libc

func Significand(x float64) float64 {
	return Scalbn(x, -Ilogb(x))
}
