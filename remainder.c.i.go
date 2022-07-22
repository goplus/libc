package libc

func Remainder(x float64, y float64) float64 {
	var q int32
	return Remquo(x, y, &q)
}
