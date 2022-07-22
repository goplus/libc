package libc

var S1_cgo18___sindf float64 = -0.16666666641626524
var S2_cgo19___sindf float64 = 0.0083333293858894632
var S3_cgo20___sindf float64 = -1.9839334836096632e-4
var S4_cgo21___sindf float64 = 2.7183114939898219e-6

func __sindf(x float64) float32 {
	var r float64
	var s float64
	var w float64
	var z float64
	z = x * x
	w = z * z
	r = S3_cgo20___sindf + z*S4_cgo21___sindf
	s = z * x
	return float32(x + s*(S1_cgo18___sindf+z*S2_cgo19___sindf) + s*w*r)
}
