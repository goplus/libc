package libc

func __math_uflowf(sign uint32) float32 {
	return __math_xflowf(sign, 2.5243549e-29)
}
