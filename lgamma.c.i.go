package libc

func Lgamma(x float64) float64 {
	return __lgamma_r(x, &__signgam)
}
