package libc

func Lgammaf(x float32) float32 {
	return __lgammaf_r(x, &__signgam)
}
