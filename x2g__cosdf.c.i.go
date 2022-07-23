package libc

var _cgos_C0____cosdf float64 = -0.499999997251031
var _cgos_C1____cosdf float64 = 0.041666623323739063
var _cgos_C2____cosdf float64 = -0.0013886763774609929
var _cgos_C3____cosdf float64 = 2.4390448796277409e-5

func __cosdf(x float64) float32 {
	var r float64
	var w float64
	var z float64
	z = x * x
	w = z * z
	r = _cgos_C2____cosdf + z*_cgos_C3____cosdf
	return float32(1 + z*_cgos_C0____cosdf + w*_cgos_C1____cosdf + w*z*r)
}
