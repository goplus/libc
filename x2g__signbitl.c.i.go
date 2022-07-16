package libc

func X__signbitl(x float64) int32 {
	return X__signbit(float64(x))
}
