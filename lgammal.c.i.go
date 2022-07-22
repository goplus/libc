package libc

func __lgammal_r(x float64, sg *int32) float64 {
	return float64(__lgamma_r(float64(x), sg))
}
func Lgammal(x float64) float64 {
	return __lgammal_r(x, &__signgam)
}
