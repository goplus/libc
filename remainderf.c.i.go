package libc

func Remainderf(x float32, y float32) float32 {
	var q int32
	return Remquof(x, y, &q)
}
