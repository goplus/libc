package libc

import unsafe "unsafe"

var _cgos_invsqrtpi__j1 float64 = 0.56418958354775628
var _cgos_tpi__j1 float64 = 0.63661977236758138

func _cgos_common__j1(ix uint32, x float64, y1 int32, sign int32) float64 {
	var z float64
	var s float64
	var c float64
	var ss float64
	var cc float64
	s = Sin(x)
	if y1 != 0 {
		s = -s
	}
	c = Cos(x)
	cc = s - c
	if ix < uint32(2145386496) {
		ss = -s - c
		z = Cos(float64(int32(2)) * x)
		if s*c > float64(int32(0)) {
			cc = z / ss
		} else {
			ss = z / cc
		}
		if ix < uint32(1207959552) {
			if y1 != 0 {
				ss = -ss
			}
			cc = _cgos_pone__j1(x)*cc - _cgos_qone__j1(x)*ss
		}
	}
	if sign != 0 {
		cc = -cc
	}
	return _cgos_invsqrtpi__j1 * cc / Sqrt(x)
}

var _cgos_r00__j1 float64 = -0.0625
var _cgos_r01__j1 float64 = 0.0014070566695518971
var _cgos_r02__j1 float64 = -1.599556310840356e-5
var _cgos_r03__j1 float64 = 4.9672799960958445e-8
var _cgos_s01__j1 float64 = 0.019153759953836346
var _cgos_s02__j1 float64 = 1.8594678558863092e-4
var _cgos_s03__j1 float64 = 1.1771846404262368e-6
var _cgos_s04__j1 float64 = 5.0463625707621704e-9
var _cgos_s05__j1 float64 = 1.2354227442613791e-11

func J1(x float64) float64 {
	var z float64
	var r float64
	var s float64
	var ix uint32
	var sign int32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_j1{x})) >> int32(32))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix >= uint32(2146435072) {
		return float64(int32(1)) / (x * x)
	}
	if ix >= uint32(1073741824) {
		return _cgos_common__j1(ix, Fabs(x), int32(0), sign)
	}
	if ix >= uint32(939524096) {
		z = x * x
		r = z * (_cgos_r00__j1 + z*(_cgos_r01__j1+z*(_cgos_r02__j1+z*_cgos_r03__j1)))
		s = float64(int32(1)) + z*(_cgos_s01__j1+z*(_cgos_s02__j1+z*(_cgos_s03__j1+z*(_cgos_s04__j1+z*_cgos_s05__j1))))
		z = r / s
	} else {
		z = x
	}
	return (0.5 + z) * x
}

type _cgoz_18_j1 struct {
	_f float64
}

var _cgos_U0__j1 [5]float64 = [5]float64{-0.19605709064623894, 0.050443871663981128, -0.0019125689587576355, 2.352526005616105e-5, -9.1909915803987887e-8}
var _cgos_V0__j1 [5]float64 = [5]float64{0.01991673182366499, 2.0255258102513517e-4, 1.3560880109751623e-6, 6.227414523646215e-9, 1.6655924620799208e-11}

func Y1(x float64) float64 {
	var z float64
	var u float64
	var v float64
	var ix uint32
	var lx uint32
	for {
		var __u uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_19_j1{x}))
		ix = uint32(__u >> int32(32))
		lx = uint32(__u)
		if true {
			break
		}
	}
	if ix<<int32(1)|lx == uint32(0) {
		return func() float64 {
			return float64(-1)
		}() / 0
	}
	if ix>>int32(31) != 0 {
		return func() float64 {
			return float64(int32(0))
		}() / 0
	}
	if ix >= uint32(2146435072) {
		return float64(int32(1)) / x
	}
	if ix >= uint32(1073741824) {
		return _cgos_common__j1(ix, x, int32(1), int32(0))
	}
	if ix < uint32(1016070144) {
		return -_cgos_tpi__j1 / x
	}
	z = x * x
	u = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_U0__j1)))) + uintptr(int32(0))*8)) + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_U0__j1)))) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_U0__j1)))) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_U0__j1)))) + uintptr(int32(3))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_U0__j1)))) + uintptr(int32(4))*8)))))
	v = float64(int32(1)) + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_V0__j1)))) + uintptr(int32(0))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_V0__j1)))) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_V0__j1)))) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_V0__j1)))) + uintptr(int32(3))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_V0__j1)))) + uintptr(int32(4))*8))))))
	return x*(u/v) + _cgos_tpi__j1*(J1(x)*Log(x)-float64(int32(1))/x)
}

type _cgoz_19_j1 struct {
	_f float64
}

var _cgos_pr8__j1 [6]float64 = [6]float64{0, 0.11718749999998865, 13.239480659307358, 412.05185430737856, 3874.7453891396053, 7914.4795403189173}
var _cgos_ps8__j1 [5]float64 = [5]float64{114.20737037567841, 3650.9308342085346, 36956.206026903346, 97602.79359349508, 30804.272062788881}
var _cgos_pr5__j1 [6]float64 = [6]float64{1.3199051955624352e-11, 0.1171874931906141, 6.8027512786843287, 108.30818299018911, 517.63613953319975, 528.71520136333754}
var _cgos_ps5__j1 [5]float64 = [5]float64{59.280598722113133, 991.40141873361438, 5353.2669529148798, 7844.6903174955123, 1504.0468881036106}
var _cgos_pr3__j1 [6]float64 = [6]float64{3.0250391613737362e-9, 0.11718686556725359, 3.9329775003331564, 35.119403559163693, 91.055011075078127, 48.559068519736492}
var _cgos_ps3__j1 [5]float64 = [5]float64{34.791309500125152, 336.76245874782575, 1046.8713997577513, 890.81134639825643, 103.78793243963928}
var _cgos_pr2__j1 [6]float64 = [6]float64{1.0771083010687374e-7, 0.11717621946268335, 2.3685149666760879, 12.242610914826123, 17.693971127168773, 5.073523125888185}
var _cgos_ps2__j1 [5]float64 = [5]float64{21.436485936382141, 125.29022716840275, 232.27646905716281, 117.6793732871471, 8.3646389337161828}

func _cgos_pone__j1(x float64) float64 {
	var p *float64
	var q *float64
	var z float64
	var r float64
	var s float64
	var ix uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_j1{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1075838976) {
		p = (*float64)(unsafe.Pointer(&_cgos_pr8__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_ps8__j1))
	} else if ix >= uint32(1074933387) {
		p = (*float64)(unsafe.Pointer(&_cgos_pr5__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_ps5__j1))
	} else if ix >= uint32(1074191213) {
		p = (*float64)(unsafe.Pointer(&_cgos_pr3__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_ps3__j1))
	} else {
		p = (*float64)(unsafe.Pointer(&_cgos_pr2__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_ps2__j1))
	}
	z = 1 / (x * x)
	r = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*8))))))
	s = 1 + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*8))))))
	return 1 + r/s
}

type _cgoz_20_j1 struct {
	_f float64
}

var _cgos_qr8__j1 [6]float64 = [6]float64{0, -0.10253906249999271, -16.271753454458999, -759.60172251395011, -11849.806670242959, -48438.512428575035}
var _cgos_qs8__j1 [6]float64 = [6]float64{161.39536970072291, 7825.3859992334847, 133875.33628724958, 719657.72368324094, 666601.23261777638, -294490.26430383464}
var _cgos_qr5__j1 [6]float64 = [6]float64{-2.089799311417641e-11, -0.10253905024137543, -8.0564482812393603, -183.66960747488838, -1373.1937606550816, -2612.4444045321566}
var _cgos_qs5__j1 [6]float64 = [6]float64{81.276550138433578, 1991.7987346048596, 17468.485192490891, 49851.427091035228, 27948.075163891812, -4719.1835479512847}
var _cgos_qr3__j1 [6]float64 = [6]float64{-5.0783122646176656e-9, -0.10253782982083709, -4.610115811394734, -57.847221656278364, -228.2445407376317, -219.21012847890933}
var _cgos_qs3__j1 [6]float64 = [6]float64{47.665155032372951, 673.86511267669971, 3380.1528667952634, 5547.7290972072278, 1903.119193388108, -135.20119144430734}
var _cgos_qr2__j1 [6]float64 = [6]float64{-1.7838172751095887e-7, -0.10251704260798555, -2.7522056827818746, -19.663616264370372, -42.325313337283049, -21.371921170370406}
var _cgos_qs2__j1 [6]float64 = [6]float64{29.533362906052385, 252.98154998219053, 757.50283486864544, 739.39320532046725, 155.94900333666612, -4.9594989882262821}

func _cgos_qone__j1(x float64) float64 {
	var p *float64
	var q *float64
	var s float64
	var r float64
	var z float64
	var ix uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_21_j1{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1075838976) {
		p = (*float64)(unsafe.Pointer(&_cgos_qr8__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_qs8__j1))
	} else if ix >= uint32(1074933387) {
		p = (*float64)(unsafe.Pointer(&_cgos_qr5__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_qs5__j1))
	} else if ix >= uint32(1074191213) {
		p = (*float64)(unsafe.Pointer(&_cgos_qr3__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_qs3__j1))
	} else {
		p = (*float64)(unsafe.Pointer(&_cgos_qr2__j1))
		q = (*float64)(unsafe.Pointer(&_cgos_qs2__j1))
	}
	z = 1 / (x * x)
	r = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*8))))))
	s = 1 + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(5))*8)))))))
	return (0.375 + r/s) / x
}

type _cgoz_21_j1 struct {
	_f float64
}
