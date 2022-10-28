package libc

var _cgos_C0___cosdf float64 = -0.499999997251031
var _cgos_C1___cosdf float64 = 0.041666623323739063
var _cgos_C2___cosdf float64 = -0.0013886763774609929
var _cgos_C3___cosdf float64 = 2.4390448796277409e-5

func __cosdf(x float64) float32 {
	var r float64
	var w float64
	var z float64
	z = x * x
	w = z * z
	r = _cgos_C2___cosdf + z*_cgos_C3___cosdf
	return float32(1.0 + z*_cgos_C0___cosdf + w*_cgos_C1___cosdf + w*z*r)
}
