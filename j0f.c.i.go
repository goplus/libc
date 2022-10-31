package libc

import unsafe "unsafe"

var _cgos_invsqrtpi_j0f float32 = float32(0.56418961287000002)
var _cgos_tpi_j0f float32 = float32(0.63661974668999999)

func _cgos_common_j0f(ix uint32, x float32, y0 int32) float32 {
	var z float32
	var s float32
	var c float32
	var ss float32
	var cc float32
	s = Sinf(x)
	c = Cosf(x)
	if y0 != 0 {
		c = -c
	}
	cc = s + c
	if ix < uint32(2130706432) {
		ss = s - c
		z = -Cosf(float32(int32(2)) * x)
		if s*c < float32(int32(0)) {
			cc = z / ss
		} else {
			ss = z / cc
		}
		if ix < uint32(1484783616) {
			if y0 != 0 {
				ss = -ss
			}
			cc = _cgos_pzerof_j0f(x)*cc - _cgos_qzerof_j0f(x)*ss
		}
	}
	return _cgos_invsqrtpi_j0f * cc / Sqrtf(x)
}

var _cgos_R02_j0f float32 = float32(0.015625)
var _cgos_R03_j0f float32 = float32(-1.8997929874e-4)
var _cgos_R04_j0f float32 = float32(1.8295404515999999e-6)
var _cgos_R05_j0f float32 = float32(-4.6183270540999997e-9)
var _cgos_S01_j0f float32 = float32(0.015619102865)
var _cgos_S02_j0f float32 = float32(1.1692678527e-4)
var _cgos_S03_j0f float32 = float32(5.1354652442000005e-7)
var _cgos_S04_j0f float32 = float32(1.1661400733999999e-9)

func J0f(x float32) float32 {
	var z float32
	var r float32
	var s float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_j0f{x}))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(2139095040) {
		return float32(int32(1)) / (x * x)
	}
	x = Fabsf(x)
	if ix >= uint32(1073741824) {
		return _cgos_common_j0f(ix, x, int32(0))
	}
	if ix >= uint32(973078528) {
		z = x * x
		r = z * (_cgos_R02_j0f + z*(_cgos_R03_j0f+z*(_cgos_R04_j0f+z*_cgos_R05_j0f)))
		s = float32(int32(1)) + z*(_cgos_S01_j0f+z*(_cgos_S02_j0f+z*(_cgos_S03_j0f+z*_cgos_S04_j0f)))
		return (float32(int32(1))+x/float32(int32(2)))*(float32(int32(1))-x/float32(int32(2))) + z*(r/s)
	}
	if ix >= uint32(562036736) {
		x = 0.25 * x * x
	}
	return float32(int32(1)) - x
}

type _cgoz_18_j0f struct {
	_f float32
}

var _cgos_u00_j0f float32 = float32(-0.073804296552999998)
var _cgos_u01_j0f float32 = float32(0.17666645348000001)
var _cgos_u02_j0f float32 = float32(-0.013818567618999999)
var _cgos_u03_j0f float32 = float32(3.4745343146000001e-4)
var _cgos_u04_j0f float32 = float32(-3.8140706238000001e-6)
var _cgos_u05_j0f float32 = float32(1.9559013963999999e-8)
var _cgos_u06_j0f float32 = float32(-3.982051841e-11)
var _cgos_v01_j0f float32 = float32(0.012730483896999999)
var _cgos_v02_j0f float32 = float32(7.6006865129000004e-5)
var _cgos_v03_j0f float32 = float32(2.5915085189000002e-7)
var _cgos_v04_j0f float32 = float32(4.4111031493999999e-10)

func Y0f(x float32) float32 {
	var z float32
	var u float32
	var v float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_19_j0f{x}))
		if true {
			break
		}
	}
	if ix&uint32(2147483647) == uint32(0) {
		return func() float32 {
			return float32(-1)
		}() / 0.0
	}
	if ix>>int32(31) != 0 {
		return func() float32 {
			return float32(int32(0))
		}() / 0.0
	}
	if ix >= uint32(2139095040) {
		return float32(int32(1)) / x
	}
	if ix >= uint32(1073741824) {
		return _cgos_common_j0f(ix, x, int32(1))
	}
	if ix >= uint32(956301312) {
		z = x * x
		u = _cgos_u00_j0f + z*(_cgos_u01_j0f+z*(_cgos_u02_j0f+z*(_cgos_u03_j0f+z*(_cgos_u04_j0f+z*(_cgos_u05_j0f+z*_cgos_u06_j0f)))))
		v = float32(int32(1)) + z*(_cgos_v01_j0f+z*(_cgos_v02_j0f+z*(_cgos_v03_j0f+z*_cgos_v04_j0f)))
		return u/v + _cgos_tpi_j0f*(J0f(x)*Logf(x))
	}
	return _cgos_u00_j0f + _cgos_tpi_j0f*Logf(x)
}

type _cgoz_19_j0f struct {
	_f float32
}

var _cgos_pR8_j0f [6]float32 = [6]float32{float32(0.0), float32(-0.0703125), float32(-8.0816707610999998), float32(-257.06311034999999), float32(-2485.2163086), float32(-5253.0439452999999)}
var _cgos_pS8_j0f [5]float32 = [5]float32{float32(116.53436279), float32(3833.7448730000001), float32(40597.855469000002), float32(116752.96875), float32(47627.726562000003)}
var _cgos_pR5_j0f [6]float32 = [6]float32{float32(-1.1412546255e-11), float32(-0.070312492549000002), float32(-4.1596107483000004), float32(-67.674766540999997), float32(-331.23129272), float32(-346.43338012999999)}
var _cgos_pS5_j0f [5]float32 = [5]float32{float32(60.753936768000003), float32(1051.2523193), float32(5978.9707030999998), float32(9625.4453125), float32(2406.0581054999998)}
var _cgos_pR3_j0f [6]float32 = [6]float32{float32(-2.5470459075e-9), float32(-0.070311963557999999), float32(-2.4090321064000002), float32(-21.965976715), float32(-58.079170226999999), float32(-31.447946548000001)}
var _cgos_pS3_j0f [5]float32 = [5]float32{float32(35.856033324999999), float32(361.51397704999999), float32(1193.6077881000001), float32(1127.9968262), float32(173.58093262)}
var _cgos_pR2_j0f [6]float32 = [6]float32{float32(-8.8753431271000001e-8), float32(-0.070303097366999995), float32(-1.4507384299999999), float32(-7.6356959343000001), float32(-11.193166733), float32(-3.2336456776000002)}
var _cgos_pS2_j0f [5]float32 = [5]float32{float32(22.220300674000001), float32(136.20678710999999), float32(270.47027587999997), float32(153.87539673000001), float32(14.657617568999999)}

func _cgos_pzerof_j0f(x float32) float32 {
	var p *float32
	var q *float32
	var z float32
	var r float32
	var s float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_20_j0f{x}))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1090519040) {
		p = (*float32)(unsafe.Pointer(&_cgos_pR8_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_pS8_j0f))
	} else if ix >= uint32(1083274219) {
		p = (*float32)(unsafe.Pointer(&_cgos_pR5_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_pS5_j0f))
	} else if ix >= uint32(1077336343) {
		p = (*float32)(unsafe.Pointer(&_cgos_pR3_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_pS3_j0f))
	} else {
		p = (*float32)(unsafe.Pointer(&_cgos_pR2_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_pS2_j0f))
	}
	z = 1.0 / (x * x)
	r = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*4))))))
	s = 1.0 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*4))))))
	return 1.0 + r/s
}

type _cgoz_20_j0f struct {
	_f float32
}

var _cgos_qR8_j0f [6]float32 = [6]float32{float32(0.0), float32(0.0732421875), float32(11.768206596000001), float32(557.67340088000003), float32(8859.1972655999998), float32(37014.625)}
var _cgos_qS8_j0f [6]float32 = [6]float32{float32(163.77603149000001), float32(8098.3447266000003), float32(142538.29688000001), float32(803309.25), float32(840501.5625), float32(-343899.28125)}
var _cgos_qR5_j0f [6]float32 = [6]float32{float32(1.8408595828e-11), float32(0.073242180049000002), float32(5.8356351852000001), float32(135.11157227000001), float32(1027.2437743999999), float32(1989.9779053)}
var _cgos_qS5_j0f [6]float32 = [6]float32{float32(82.776611328000001), float32(2077.8142090000001), float32(18847.289062), float32(56751.113280999998), float32(35976.753905999998), float32(-5354.3427733999997)}
var _cgos_qR3_j0f [6]float32 = [6]float32{float32(4.3774099900000002e-9), float32(0.073241114615999997), float32(3.3442313670999999), float32(42.621845245000003), float32(170.80809020999999), float32(166.73394775)}
var _cgos_qS3_j0f [6]float32 = [6]float32{float32(48.758872986), float32(709.68920897999999), float32(3704.1481933999999), float32(6460.4252930000002), float32(2516.3337402000002), float32(-149.24745178000001)}
var _cgos_qR2_j0f [6]float32 = [6]float32{float32(1.5044444979e-7), float32(0.073223426938000005), float32(1.9981917143000001), float32(14.495602608), float32(31.666231154999998), float32(16.252708434999999)}
var _cgos_qS2_j0f [6]float32 = [6]float32{float32(30.365585327000002), float32(269.34811401000002), float32(844.78375243999994), float32(882.93585204999999), float32(212.66638184000001), float32(-5.3109550476000003)}

func _cgos_qzerof_j0f(x float32) float32 {
	var p *float32
	var q *float32
	var s float32
	var r float32
	var z float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_21_j0f{x}))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1090519040) {
		p = (*float32)(unsafe.Pointer(&_cgos_qR8_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_qS8_j0f))
	} else if ix >= uint32(1083274219) {
		p = (*float32)(unsafe.Pointer(&_cgos_qR5_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_qS5_j0f))
	} else if ix >= uint32(1077336343) {
		p = (*float32)(unsafe.Pointer(&_cgos_qR3_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_qS3_j0f))
	} else {
		p = (*float32)(unsafe.Pointer(&_cgos_qR2_j0f))
		q = (*float32)(unsafe.Pointer(&_cgos_qS2_j0f))
	}
	z = 1.0 / (x * x)
	r = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*4))))))
	s = 1.0 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(5))*4)))))))
	return (-0.125 + r/s) / x
}

type _cgoz_21_j0f struct {
	_f float32
}
