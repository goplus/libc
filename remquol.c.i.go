package libc

func Remquol(x float64, y float64, quo *int32) float64 {
	return float64(Remquo(float64(x), float64(y), quo))
}
