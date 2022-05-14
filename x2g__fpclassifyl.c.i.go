package libc

func __fpclassifyl(x float64) int32 {
	return __fpclassify(float64(x))
}
