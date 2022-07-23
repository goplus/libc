package libc

var _cgos_C1____cos float64 = 0.041666666666666602
var _cgos_C2____cos float64 = -0.001388888888887411
var _cgos_C3____cos float64 = 2.4801587289476729e-5
var _cgos_C4____cos float64 = -2.7557314351390663e-7
var _cgos_C5____cos float64 = 2.0875723212981748e-9
var _cgos_C6____cos float64 = -1.1359647557788195e-11

func __cos(x float64, y float64) float64 {
	var hz float64
	var z float64
	var r float64
	var w float64
	z = x * x
	w = z * z
	r = z*(_cgos_C1____cos+z*(_cgos_C2____cos+z*_cgos_C3____cos)) + w*w*(_cgos_C4____cos+z*(_cgos_C5____cos+z*_cgos_C6____cos))
	hz = 0.5 * z
	w = 1 - hz
	return w + (1 - w - hz + (z*r - x*y))
}
