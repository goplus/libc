package libc

import unsafe "unsafe"

var invsqrtpi_cgos__j0f float32 = float32(0.56418961287000002)
var tpi_cgos__j0f float32 = float32(0.63661974668999999)

func common_cgos__j0f(ix uint32, x float32, y0 int32) float32 {
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
			cc = pzerof_cgos__j0f(x)*cc - qzerof_cgos__j0f(x)*ss
		}
	}
	return invsqrtpi_cgos__j0f * cc / Sqrtf(x)
}

var R02_cgos__j0f float32 = float32(0.015625)
var R03_cgos__j0f float32 = float32(-1.8997929874e-4)
var R04_cgos__j0f float32 = float32(1.8295404515999999e-6)
var R05_cgos__j0f float32 = float32(-4.6183270540999997e-9)
var S01_cgos__j0f float32 = float32(0.015619102865)
var S02_cgos__j0f float32 = float32(1.1692678527e-4)
var S03_cgos__j0f float32 = float32(5.1354652442000005e-7)
var S04_cgos__j0f float32 = float32(1.1661400733999999e-9)

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
		return common_cgos__j0f(ix, x, int32(0))
	}
	if ix >= uint32(973078528) {
		z = x * x
		r = z * (R02_cgos__j0f + z*(R03_cgos__j0f+z*(R04_cgos__j0f+z*R05_cgos__j0f)))
		s = float32(int32(1)) + z*(S01_cgos__j0f+z*(S02_cgos__j0f+z*(S03_cgos__j0f+z*S04_cgos__j0f)))
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

var u00_cgos__j0f float32 = float32(-0.073804296552999998)
var u01_cgos__j0f float32 = float32(0.17666645348000001)
var u02_cgos__j0f float32 = float32(-0.013818567618999999)
var u03_cgos__j0f float32 = float32(3.4745343146000001e-4)
var u04_cgos__j0f float32 = float32(-3.8140706238000001e-6)
var u05_cgos__j0f float32 = float32(1.9559013963999999e-8)
var u06_cgos__j0f float32 = float32(-3.982051841e-11)
var v01_cgos__j0f float32 = float32(0.012730483896999999)
var v02_cgos__j0f float32 = float32(7.6006865129000004e-5)
var v03_cgos__j0f float32 = float32(2.5915085189000002e-7)
var v04_cgos__j0f float32 = float32(4.4111031493999999e-10)

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
		}() / 0
	}
	if ix>>int32(31) != 0 {
		return func() float32 {
			return float32(int32(0))
		}() / 0
	}
	if ix >= uint32(2139095040) {
		return float32(int32(1)) / x
	}
	if ix >= uint32(1073741824) {
		return common_cgos__j0f(ix, x, int32(1))
	}
	if ix >= uint32(956301312) {
		z = x * x
		u = u00_cgos__j0f + z*(u01_cgos__j0f+z*(u02_cgos__j0f+z*(u03_cgos__j0f+z*(u04_cgos__j0f+z*(u05_cgos__j0f+z*u06_cgos__j0f)))))
		v = float32(int32(1)) + z*(v01_cgos__j0f+z*(v02_cgos__j0f+z*(v03_cgos__j0f+z*v04_cgos__j0f)))
		return u/v + tpi_cgos__j0f*(J0f(x)*Logf(x))
	}
	return u00_cgos__j0f + tpi_cgos__j0f*Logf(x)
}

type _cgoz_19_j0f struct {
	_f float32
}

var pR8_cgos__j0f [6]float32 = [6]float32{float32(0), float32(-0.0703125), float32(-8.0816707610999998), float32(-257.06311034999999), float32(-2485.2163086), float32(-5253.0439452999999)}
var pS8_cgos__j0f [5]float32 = [5]float32{float32(116.53436279), float32(3833.7448730000001), float32(40597.855469000002), float32(116752.96875), float32(47627.726562000003)}
var pR5_cgos__j0f [6]float32 = [6]float32{float32(-1.1412546255e-11), float32(-0.070312492549000002), float32(-4.1596107483000004), float32(-67.674766540999997), float32(-331.23129272), float32(-346.43338012999999)}
var pS5_cgos__j0f [5]float32 = [5]float32{float32(60.753936768000003), float32(1051.2523193), float32(5978.9707030999998), float32(9625.4453125), float32(2406.0581054999998)}
var pR3_cgos__j0f [6]float32 = [6]float32{float32(-2.5470459075e-9), float32(-0.070311963557999999), float32(-2.4090321064000002), float32(-21.965976715), float32(-58.079170226999999), float32(-31.447946548000001)}
var pS3_cgos__j0f [5]float32 = [5]float32{float32(35.856033324999999), float32(361.51397704999999), float32(1193.6077881000001), float32(1127.9968262), float32(173.58093262)}
var pR2_cgos__j0f [6]float32 = [6]float32{float32(-8.8753431271000001e-8), float32(-0.070303097366999995), float32(-1.4507384299999999), float32(-7.6356959343000001), float32(-11.193166733), float32(-3.2336456776000002)}
var pS2_cgos__j0f [5]float32 = [5]float32{float32(22.220300674000001), float32(136.20678710999999), float32(270.47027587999997), float32(153.87539673000001), float32(14.657617568999999)}

func pzerof_cgos__j0f(x float32) float32 {
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
		p = (*float32)(unsafe.Pointer(&pR8_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&pS8_cgos__j0f))
	} else if ix >= uint32(1083274219) {
		p = (*float32)(unsafe.Pointer(&pR5_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&pS5_cgos__j0f))
	} else if ix >= uint32(1077336343) {
		p = (*float32)(unsafe.Pointer(&pR3_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&pS3_cgos__j0f))
	} else {
		p = (*float32)(unsafe.Pointer(&pR2_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&pS2_cgos__j0f))
	}
	z = 1 / (x * x)
	r = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*4))))))
	s = 1 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*4))))))
	return 1 + r/s
}

type _cgoz_20_j0f struct {
	_f float32
}

var qR8_cgos__j0f [6]float32 = [6]float32{float32(0), float32(0.0732421875), float32(11.768206596000001), float32(557.67340088000003), float32(8859.1972655999998), float32(37014.625)}
var qS8_cgos__j0f [6]float32 = [6]float32{float32(163.77603149000001), float32(8098.3447266000003), float32(142538.29688000001), float32(803309.25), float32(840501.5625), float32(-343899.28125)}
var qR5_cgos__j0f [6]float32 = [6]float32{float32(1.8408595828e-11), float32(0.073242180049000002), float32(5.8356351852000001), float32(135.11157227000001), float32(1027.2437743999999), float32(1989.9779053)}
var qS5_cgos__j0f [6]float32 = [6]float32{float32(82.776611328000001), float32(2077.8142090000001), float32(18847.289062), float32(56751.113280999998), float32(35976.753905999998), float32(-5354.3427733999997)}
var qR3_cgos__j0f [6]float32 = [6]float32{float32(4.3774099900000002e-9), float32(0.073241114615999997), float32(3.3442313670999999), float32(42.621845245000003), float32(170.80809020999999), float32(166.73394775)}
var qS3_cgos__j0f [6]float32 = [6]float32{float32(48.758872986), float32(709.68920897999999), float32(3704.1481933999999), float32(6460.4252930000002), float32(2516.3337402000002), float32(-149.24745178000001)}
var qR2_cgos__j0f [6]float32 = [6]float32{float32(1.5044444979e-7), float32(0.073223426938000005), float32(1.9981917143000001), float32(14.495602608), float32(31.666231154999998), float32(16.252708434999999)}
var qS2_cgos__j0f [6]float32 = [6]float32{float32(30.365585327000002), float32(269.34811401000002), float32(844.78375243999994), float32(882.93585204999999), float32(212.66638184000001), float32(-5.3109550476000003)}

func qzerof_cgos__j0f(x float32) float32 {
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
		p = (*float32)(unsafe.Pointer(&qR8_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&qS8_cgos__j0f))
	} else if ix >= uint32(1083274219) {
		p = (*float32)(unsafe.Pointer(&qR5_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&qS5_cgos__j0f))
	} else if ix >= uint32(1077336343) {
		p = (*float32)(unsafe.Pointer(&qR3_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&qS3_cgos__j0f))
	} else {
		p = (*float32)(unsafe.Pointer(&qR2_cgos__j0f))
		q = (*float32)(unsafe.Pointer(&qS2_cgos__j0f))
	}
	z = 1 / (x * x)
	r = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*4))))))
	s = 1 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(5))*4)))))))
	return (-0.125 + r/s) / x
}

type _cgoz_21_j0f struct {
	_f float32
}
