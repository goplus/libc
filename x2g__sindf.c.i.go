package libc

var _cgos_S1____sindf float64 = -0.16666666641626524
var _cgos_S2____sindf float64 = 0.0083333293858894632
var _cgos_S3____sindf float64 = -1.9839334836096632e-4
var _cgos_S4____sindf float64 = 2.7183114939898219e-6

func __sindf(x float64) float32 {
	var r float64
	var s float64
	var w float64
	var z float64
	z = x * x
	w = z * z
	r = _cgos_S3____sindf + z*_cgos_S4____sindf
	s = z * x
	return float32(x + s*(_cgos_S1____sindf+z*_cgos_S2____sindf) + s*w*r)
}
