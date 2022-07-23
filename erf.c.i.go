package libc

import unsafe "unsafe"

var _cgos_erx_erf float64 = 0.84506291151046753
var _cgos_efx8_erf float64 = 1.0270333367641007
var _cgos_pp0_erf float64 = 0.12837916709551256
var _cgos_pp1_erf float64 = -0.3250421072470015
var _cgos_pp2_erf float64 = -0.02848174957559851
var _cgos_pp3_erf float64 = -0.0057702702964894416
var _cgos_pp4_erf float64 = -2.3763016656650163e-5
var _cgos_qq1_erf float64 = 0.39791722395915535
var _cgos_qq2_erf float64 = 0.065022249988767294
var _cgos_qq3_erf float64 = 0.0050813062818757656
var _cgos_qq4_erf float64 = 1.3249473800432164e-4
var _cgos_qq5_erf float64 = -3.9602282787753681e-6
var _cgos_pa0_erf float64 = -0.0023621185607526594
var _cgos_pa1_erf float64 = 0.41485611868374833
var _cgos_pa2_erf float64 = -0.37220787603570132
var _cgos_pa3_erf float64 = 0.31834661990116175
var _cgos_pa4_erf float64 = -0.11089469428239668
var _cgos_pa5_erf float64 = 0.035478304325618236
var _cgos_pa6_erf float64 = -0.0021663755948687908
var _cgos_qa1_erf float64 = 0.10642088040084423
var _cgos_qa2_erf float64 = 0.54039791770217105
var _cgos_qa3_erf float64 = 0.071828654414196266
var _cgos_qa4_erf float64 = 0.12617121980876164
var _cgos_qa5_erf float64 = 0.013637083912029051
var _cgos_qa6_erf float64 = 0.011984499846799107
var _cgos_ra0_erf float64 = -0.0098649440348471482
var _cgos_ra1_erf float64 = -0.69385857270718176
var _cgos_ra2_erf float64 = -10.558626225323291
var _cgos_ra3_erf float64 = -62.375332450326006
var _cgos_ra4_erf float64 = -162.39666946257347
var _cgos_ra5_erf float64 = -184.60509290671104
var _cgos_ra6_erf float64 = -81.287435506306593
var _cgos_ra7_erf float64 = -9.8143293441691454
var _cgos_sa1_erf float64 = 19.651271667439257
var _cgos_sa2_erf float64 = 137.65775414351904
var _cgos_sa3_erf float64 = 434.56587747522923
var _cgos_sa4_erf float64 = 645.38727173326788
var _cgos_sa5_erf float64 = 429.00814002756783
var _cgos_sa6_erf float64 = 108.63500554177944
var _cgos_sa7_erf float64 = 6.5702497703192817
var _cgos_sa8_erf float64 = -0.060424415214858099
var _cgos_rb0_erf float64 = -0.0098649429247000992
var _cgos_rb1_erf float64 = -0.79928323768052301
var _cgos_rb2_erf float64 = -17.757954917754752
var _cgos_rb3_erf float64 = -160.63638485582192
var _cgos_rb4_erf float64 = -637.56644336838963
var _cgos_rb5_erf float64 = -1025.0951316110772
var _cgos_rb6_erf float64 = -483.5191916086514
var _cgos_sb1_erf float64 = 30.338060743482458
var _cgos_sb2_erf float64 = 325.79251299657392
var _cgos_sb3_erf float64 = 1536.729586084437
var _cgos_sb4_erf float64 = 3199.8582195085955
var _cgos_sb5_erf float64 = 2553.0504064331644
var _cgos_sb6_erf float64 = 474.52854120695537
var _cgos_sb7_erf float64 = -22.440952446585818

func _cgos_erfc1_erf(x float64) float64 {
	var s float64
	var P float64
	var Q float64
	s = Fabs(x) - float64(int32(1))
	P = _cgos_pa0_erf + s*(_cgos_pa1_erf+s*(_cgos_pa2_erf+s*(_cgos_pa3_erf+s*(_cgos_pa4_erf+s*(_cgos_pa5_erf+s*_cgos_pa6_erf)))))
	Q = float64(int32(1)) + s*(_cgos_qa1_erf+s*(_cgos_qa2_erf+s*(_cgos_qa3_erf+s*(_cgos_qa4_erf+s*(_cgos_qa5_erf+s*_cgos_qa6_erf)))))
	return float64(int32(1)) - _cgos_erx_erf - P/Q
}
func _cgos_erfc2_erf(ix uint32, x float64) float64 {
	var s float64
	var R float64
	var S float64
	var z float64
	if ix < uint32(1072955392) {
		return _cgos_erfc1_erf(x)
	}
	x = Fabs(x)
	s = float64(int32(1)) / (x * x)
	if ix < uint32(1074191213) {
		R = _cgos_ra0_erf + s*(_cgos_ra1_erf+s*(_cgos_ra2_erf+s*(_cgos_ra3_erf+s*(_cgos_ra4_erf+s*(_cgos_ra5_erf+s*(_cgos_ra6_erf+s*_cgos_ra7_erf))))))
		S = 1 + s*(_cgos_sa1_erf+s*(_cgos_sa2_erf+s*(_cgos_sa3_erf+s*(_cgos_sa4_erf+s*(_cgos_sa5_erf+s*(_cgos_sa6_erf+s*(_cgos_sa7_erf+s*_cgos_sa8_erf)))))))
	} else {
		R = _cgos_rb0_erf + s*(_cgos_rb1_erf+s*(_cgos_rb2_erf+s*(_cgos_rb3_erf+s*(_cgos_rb4_erf+s*(_cgos_rb5_erf+s*_cgos_rb6_erf)))))
		S = 1 + s*(_cgos_sb1_erf+s*(_cgos_sb2_erf+s*(_cgos_sb3_erf+s*(_cgos_sb4_erf+s*(_cgos_sb5_erf+s*(_cgos_sb6_erf+s*_cgos_sb7_erf))))))
	}
	z = x
	for {
		z = *(*float64)(unsafe.Pointer(&_cgoz_18_erf{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_19_erf{z}))>>int32(32))<<int32(32) | uint64(0)}))
		if true {
			break
		}
	}
	return Exp(-z*z-0.5625) * Exp((z-x)*(z+x)+R/S) / x
}

type _cgoz_18_erf struct {
	_i uint64
}
type _cgoz_19_erf struct {
	_f float64
}

func Erf(x float64) float64 {
	var r float64
	var s float64
	var z float64
	var y float64
	var ix uint32
	var sign int32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_erf{x})) >> int32(32))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix >= uint32(2146435072) {
		return float64(int32(1)-int32(2)*sign) + float64(int32(1))/x
	}
	if ix < uint32(1072365568) {
		if ix < uint32(1043333120) {
			return 0.125 * (float64(int32(8))*x + _cgos_efx8_erf*x)
		}
		z = x * x
		r = _cgos_pp0_erf + z*(_cgos_pp1_erf+z*(_cgos_pp2_erf+z*(_cgos_pp3_erf+z*_cgos_pp4_erf)))
		s = 1 + z*(_cgos_qq1_erf+z*(_cgos_qq2_erf+z*(_cgos_qq3_erf+z*(_cgos_qq4_erf+z*_cgos_qq5_erf))))
		y = r / s
		return x + x*y
	}
	if ix < uint32(1075314688) {
		y = float64(int32(1)) - _cgos_erfc2_erf(ix, x)
	} else {
		y = float64(int32(1)) - 2.2250738585072014e-308
	}
	return func() float64 {
		if sign != 0 {
			return -y
		} else {
			return y
		}
	}()
}

type _cgoz_20_erf struct {
	_f float64
}

func Erfc(x float64) float64 {
	var r float64
	var s float64
	var z float64
	var y float64
	var ix uint32
	var sign int32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_21_erf{x})) >> int32(32))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix >= uint32(2146435072) {
		return float64(int32(2)*sign) + float64(int32(1))/x
	}
	if ix < uint32(1072365568) {
		if ix < uint32(1013972992) {
			return 1 - x
		}
		z = x * x
		r = _cgos_pp0_erf + z*(_cgos_pp1_erf+z*(_cgos_pp2_erf+z*(_cgos_pp3_erf+z*_cgos_pp4_erf)))
		s = 1 + z*(_cgos_qq1_erf+z*(_cgos_qq2_erf+z*(_cgos_qq3_erf+z*(_cgos_qq4_erf+z*_cgos_qq5_erf))))
		y = r / s
		if sign != 0 || ix < uint32(1070596096) {
			return 1 - (x + x*y)
		}
		return 0.5 - (x - 0.5 + x*y)
	}
	if ix < uint32(1077673984) {
		return func() float64 {
			if sign != 0 {
				return float64(int32(2)) - _cgos_erfc2_erf(ix, x)
			} else {
				return _cgos_erfc2_erf(ix, x)
			}
		}()
	}
	return func() float64 {
		if sign != 0 {
			return float64(int32(2)) - 2.2250738585072014e-308
		} else {
			return 2.2250738585072014e-308 * 2.2250738585072014e-308
		}
	}()
}

type _cgoz_21_erf struct {
	_f float64
}
