package libc

func X__fpclassifyl(x float64) int32 {
	return X__fpclassify(float64(x))
}
