package libc

import unsafe "unsafe"

var _cgos_pi__tgamma float64 = 3.1415926535897931

func _cgos_sinpi__tgamma(x float64) float64 {
	var n int32
	x = x * 0.5
	x = float64(int32(2)) * (x - Floor(x))
	n = int32(float64(int32(4)) * x)
	n = (n + int32(1)) / int32(2)
	x -= float64(n) * 0.5
	x *= _cgos_pi__tgamma
	switch n {
	default:
		fallthrough
	case int32(0):
		return __sin(x, float64(int32(0)), int32(0))
	case int32(1):
		return __cos(x, float64(int32(0)))
	case int32(2):
		return __sin(-x, float64(int32(0)), int32(0))
	case int32(3):
		return -__cos(x, float64(int32(0)))
	}
	return 0
}

var _cgos_gmhalf__tgamma float64 = 5.5246800407767296
var _cgos_Snum__tgamma [13]float64 = [13]float64{23531376880.410759, 42919803642.649101, 35711959237.355667, 17921034426.037209, 6039542586.3520279, 1439720407.3117216, 248874557.86205417, 31426415.585400194, 2876370.6289353725, 186056.26539522348, 8071.6720023658163, 210.82427775157936, 2.5066282746310002}
var _cgos_Sden__tgamma [13]float64 = [13]float64{float64(int32(0)), float64(int32(39916800)), float64(int32(120543840)), float64(int32(150917976)), float64(int32(105258076)), float64(int32(45995730)), float64(int32(13339535)), float64(int32(2637558)), float64(int32(357423)), float64(int32(32670)), float64(int32(1925)), float64(int32(66)), float64(int32(1))}
var _cgos_fact__tgamma [23]float64 = [23]float64{float64(int32(1)), float64(int32(1)), float64(int32(2)), float64(int32(6)), float64(int32(24)), float64(int32(120)), float64(int32(720)), 5040, 40320, 362880, 3628800, 39916800, 479001600, 6227020800, 87178291200, 1307674368000, 20922789888000, 355687428096000, 6402373705728000, 1.21645100408832e+17, 2.43290200817664e+18, 5.109094217170944e+19, 1.1240007277776077e+21}

func _cgos_S__tgamma(x float64) float64 {
	var num float64 = float64(int32(0))
	var den float64 = float64(int32(0))
	var i int32
	if x < float64(int32(8)) {
		for i = int32(12); i >= int32(0); i-- {
			num = num*x + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_Snum__tgamma)))) + uintptr(i)*8))
			den = den*x + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_Sden__tgamma)))) + uintptr(i)*8))
		}
	} else {
		for i = int32(0); i <= int32(12); i++ {
			num = num/x + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_Snum__tgamma)))) + uintptr(i)*8))
			den = den/x + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_Sden__tgamma)))) + uintptr(i)*8))
		}
	}
	return num / den
}
func Tgamma(x float64) float64 {
	type _cgoa_18_tgamma struct {
		f float64
	}
	var u _cgoa_18_tgamma
	u.f = x
	var absx float64
	var y float64
	var dy float64
	var z float64
	var r float64
	var ix uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32) & uint64(2147483647))
	var sign int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	if ix >= uint32(2146435072) {
		return x + float64(X__builtin_inff())
	}
	if ix < uint32(1016070144) {
		return float64(int32(1)) / x
	}
	if x == Floor(x) {
		if sign != 0 {
			return func() float64 {
				return float64(int32(0))
			}() / 0
		}
		if x <= float64(184/8) {
			return *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_fact__tgamma)))) + uintptr(int32(x)-int32(1))*8))
		}
	}
	if ix >= uint32(1080492032) {
		if sign != 0 {
			for {
				if 4 == 4 {
					fp_force_evalf(float32(1.1754943508222875e-38 / x))
				} else if 4 == 8 {
					fp_force_eval(float64(float32(1.1754943508222875e-38 / x)))
				} else {
					fp_force_evall(float64(float32(1.1754943508222875e-38 / x)))
				}
				if true {
					break
				}
			}
			if Floor(x)*0.5 == Floor(x*0.5) {
				return float64(int32(0))
			}
			return float64(-0)
		}
		x *= float64(8.9884656743115795e+307)
		return x
	}
	absx = func() float64 {
		if sign != 0 {
			return -x
		} else {
			return x
		}
	}()
	y = absx + _cgos_gmhalf__tgamma
	if absx > _cgos_gmhalf__tgamma {
		dy = y - absx
		dy -= _cgos_gmhalf__tgamma
	} else {
		dy = y - _cgos_gmhalf__tgamma
		dy -= absx
	}
	z = absx - 0.5
	r = _cgos_S__tgamma(absx) * Exp(-y)
	if x < float64(int32(0)) {
		r = -_cgos_pi__tgamma / (_cgos_sinpi__tgamma(absx) * absx * r)
		dy = -dy
		z = -z
	}
	r += dy * (_cgos_gmhalf__tgamma + 0.5) * r / y
	z = Pow(y, 0.5*z)
	y = r * z * z
	return y
}
