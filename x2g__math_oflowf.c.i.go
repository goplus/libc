package libc

func __math_oflowf(sign uint32) float32 {
	return __math_xflowf(sign, 1.58456325e+29)
}
