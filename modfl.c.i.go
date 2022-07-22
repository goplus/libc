package libc

func Modfl(x float64, iptr *float64) float64 {
	var d float64
	var r float64
	r = float64(Modf(float64(x), &d))
	*iptr = float64(d)
	return r
}
