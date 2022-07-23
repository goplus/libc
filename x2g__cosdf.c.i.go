package libc

var C0_cgos____cosdf float64 = -0.499999997251031
var C1_cgos____cosdf float64 = 0.041666623323739063
var C2_cgos____cosdf float64 = -0.0013886763774609929
var C3_cgos____cosdf float64 = 2.4390448796277409e-5

func __cosdf(x float64) float32 {
	var r float64
	var w float64
	var z float64
	z = x * x
	w = z * z
	r = C2_cgos____cosdf + z*C3_cgos____cosdf
	return float32(1 + z*C0_cgos____cosdf + w*C1_cgos____cosdf + w*z*r)
}
