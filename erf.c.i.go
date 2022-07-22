package libc

import unsafe "unsafe"

var erx_cgo18_erf float64 = 0.84506291151046753
var efx8_cgo19_erf float64 = 1.0270333367641007
var pp0_cgo20_erf float64 = 0.12837916709551256
var pp1_cgo21_erf float64 = -0.3250421072470015
var pp2_cgo22_erf float64 = -0.02848174957559851
var pp3_cgo23_erf float64 = -0.0057702702964894416
var pp4_cgo24_erf float64 = -2.3763016656650163e-5
var qq1_cgo25_erf float64 = 0.39791722395915535
var qq2_cgo26_erf float64 = 0.065022249988767294
var qq3_cgo27_erf float64 = 0.0050813062818757656
var qq4_cgo28_erf float64 = 1.3249473800432164e-4
var qq5_cgo29_erf float64 = -3.9602282787753681e-6
var pa0_cgo30_erf float64 = -0.0023621185607526594
var pa1_cgo31_erf float64 = 0.41485611868374833
var pa2_cgo32_erf float64 = -0.37220787603570132
var pa3_cgo33_erf float64 = 0.31834661990116175
var pa4_cgo34_erf float64 = -0.11089469428239668
var pa5_cgo35_erf float64 = 0.035478304325618236
var pa6_cgo36_erf float64 = -0.0021663755948687908
var qa1_cgo37_erf float64 = 0.10642088040084423
var qa2_cgo38_erf float64 = 0.54039791770217105
var qa3_cgo39_erf float64 = 0.071828654414196266
var qa4_cgo40_erf float64 = 0.12617121980876164
var qa5_cgo41_erf float64 = 0.013637083912029051
var qa6_cgo42_erf float64 = 0.011984499846799107
var ra0_cgo43_erf float64 = -0.0098649440348471482
var ra1_cgo44_erf float64 = -0.69385857270718176
var ra2_cgo45_erf float64 = -10.558626225323291
var ra3_cgo46_erf float64 = -62.375332450326006
var ra4_cgo47_erf float64 = -162.39666946257347
var ra5_cgo48_erf float64 = -184.60509290671104
var ra6_cgo49_erf float64 = -81.287435506306593
var ra7_cgo50_erf float64 = -9.8143293441691454
var sa1_cgo51_erf float64 = 19.651271667439257
var sa2_cgo52_erf float64 = 137.65775414351904
var sa3_cgo53_erf float64 = 434.56587747522923
var sa4_cgo54_erf float64 = 645.38727173326788
var sa5_cgo55_erf float64 = 429.00814002756783
var sa6_cgo56_erf float64 = 108.63500554177944
var sa7_cgo57_erf float64 = 6.5702497703192817
var sa8_cgo58_erf float64 = -0.060424415214858099
var rb0_cgo59_erf float64 = -0.0098649429247000992
var rb1_cgo60_erf float64 = -0.79928323768052301
var rb2_cgo61_erf float64 = -17.757954917754752
var rb3_cgo62_erf float64 = -160.63638485582192
var rb4_cgo63_erf float64 = -637.56644336838963
var rb5_cgo64_erf float64 = -1025.0951316110772
var rb6_cgo65_erf float64 = -483.5191916086514
var sb1_cgo66_erf float64 = 30.338060743482458
var sb2_cgo67_erf float64 = 325.79251299657392
var sb3_cgo68_erf float64 = 1536.729586084437
var sb4_cgo69_erf float64 = 3199.8582195085955
var sb5_cgo70_erf float64 = 2553.0504064331644
var sb6_cgo71_erf float64 = 474.52854120695537
var sb7_cgo72_erf float64 = -22.440952446585818

func erfc1_cgo73_erf(x float64) float64 {
	var s float64
	var P float64
	var Q float64
	s = Fabs(x) - float64(int32(1))
	P = pa0_cgo30_erf + s*(pa1_cgo31_erf+s*(pa2_cgo32_erf+s*(pa3_cgo33_erf+s*(pa4_cgo34_erf+s*(pa5_cgo35_erf+s*pa6_cgo36_erf)))))
	Q = float64(int32(1)) + s*(qa1_cgo37_erf+s*(qa2_cgo38_erf+s*(qa3_cgo39_erf+s*(qa4_cgo40_erf+s*(qa5_cgo41_erf+s*qa6_cgo42_erf)))))
	return float64(int32(1)) - erx_cgo18_erf - P/Q
}
func erfc2_cgo74_erf(ix uint32, x float64) float64 {
	var s float64
	var R float64
	var S float64
	var z float64
	if ix < uint32(1072955392) {
		return erfc1_cgo73_erf(x)
	}
	x = Fabs(x)
	s = float64(int32(1)) / (x * x)
	if ix < uint32(1074191213) {
		R = ra0_cgo43_erf + s*(ra1_cgo44_erf+s*(ra2_cgo45_erf+s*(ra3_cgo46_erf+s*(ra4_cgo47_erf+s*(ra5_cgo48_erf+s*(ra6_cgo49_erf+s*ra7_cgo50_erf))))))
		S = 1 + s*(sa1_cgo51_erf+s*(sa2_cgo52_erf+s*(sa3_cgo53_erf+s*(sa4_cgo54_erf+s*(sa5_cgo55_erf+s*(sa6_cgo56_erf+s*(sa7_cgo57_erf+s*sa8_cgo58_erf)))))))
	} else {
		R = rb0_cgo59_erf + s*(rb1_cgo60_erf+s*(rb2_cgo61_erf+s*(rb3_cgo62_erf+s*(rb4_cgo63_erf+s*(rb5_cgo64_erf+s*rb6_cgo65_erf)))))
		S = 1 + s*(sb1_cgo66_erf+s*(sb2_cgo67_erf+s*(sb3_cgo68_erf+s*(sb4_cgo69_erf+s*(sb5_cgo70_erf+s*(sb6_cgo71_erf+s*sb7_cgo72_erf))))))
	}
	z = x
	for {
		z = *(*float64)(unsafe.Pointer(&_cgoz_75_erf{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_76_erf{z}))>>int32(32))<<int32(32) | uint64(0)}))
		if true {
			break
		}
	}
	return Exp(-z*z-0.5625) * Exp((z-x)*(z+x)+R/S) / x
}

type _cgoz_75_erf struct {
	_i uint64
}
type _cgoz_76_erf struct {
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
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_77_erf{x})) >> int32(32))
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
			return 0.125 * (float64(int32(8))*x + efx8_cgo19_erf*x)
		}
		z = x * x
		r = pp0_cgo20_erf + z*(pp1_cgo21_erf+z*(pp2_cgo22_erf+z*(pp3_cgo23_erf+z*pp4_cgo24_erf)))
		s = 1 + z*(qq1_cgo25_erf+z*(qq2_cgo26_erf+z*(qq3_cgo27_erf+z*(qq4_cgo28_erf+z*qq5_cgo29_erf))))
		y = r / s
		return x + x*y
	}
	if ix < uint32(1075314688) {
		y = float64(int32(1)) - erfc2_cgo74_erf(ix, x)
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

type _cgoz_77_erf struct {
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
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_78_erf{x})) >> int32(32))
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
		r = pp0_cgo20_erf + z*(pp1_cgo21_erf+z*(pp2_cgo22_erf+z*(pp3_cgo23_erf+z*pp4_cgo24_erf)))
		s = 1 + z*(qq1_cgo25_erf+z*(qq2_cgo26_erf+z*(qq3_cgo27_erf+z*(qq4_cgo28_erf+z*qq5_cgo29_erf))))
		y = r / s
		if sign != 0 || ix < uint32(1070596096) {
			return 1 - (x + x*y)
		}
		return 0.5 - (x - 0.5 + x*y)
	}
	if ix < uint32(1077673984) {
		return func() float64 {
			if sign != 0 {
				return float64(int32(2)) - erfc2_cgo74_erf(ix, x)
			} else {
				return erfc2_cgo74_erf(ix, x)
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

type _cgoz_78_erf struct {
	_f float64
}
