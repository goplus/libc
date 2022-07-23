package libc

import unsafe "unsafe"

var _cgos_ln2_hi_log1p float64 = 0.69314718036912382
var _cgos_ln2_lo_log1p float64 = 1.9082149292705877e-10
var _cgos_Lg1_log1p float64 = 0.66666666666667351
var _cgos_Lg2_log1p float64 = 0.39999999999409419
var _cgos_Lg3_log1p float64 = 0.28571428743662391
var _cgos_Lg4_log1p float64 = 0.22222198432149784
var _cgos_Lg5_log1p float64 = 0.1818357216161805
var _cgos_Lg6_log1p float64 = 0.15313837699209373
var _cgos_Lg7_log1p float64 = 0.14798198605116586

func Log1p(x float64) float64 {
	type _cgoa_18_log1p struct {
		f float64
	}
	var u _cgoa_18_log1p
	u.f = x
	var hfsq float64
	var f float64
	var c float64
	var s float64
	var z float64
	var R float64
	var w float64
	var t1 float64
	var t2 float64
	var dk float64
	var hx uint32
	var hu uint32
	var k int32
	hx = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
	k = int32(1)
	if hx < uint32(1071284858) || hx>>int32(31) != 0 {
		if hx >= uint32(3220176896) {
			if x == float64(-1) {
				return x / 0
			}
			return (x - x) / 0
		}
		if hx<<int32(1) < uint32(2034237440) {
			if hx&uint32(2146435072) == uint32(0) {
				for {
					if 4 == 4 {
						fp_force_evalf(float32(x))
					} else if 4 == 8 {
						fp_force_eval(float64(float32(x)))
					} else {
						fp_force_evall(float64(float32(x)))
					}
					if true {
						break
					}
				}
			}
			return x
		}
		if hx <= uint32(3218259652) {
			k = int32(0)
			c = float64(int32(0))
			f = x
		}
	} else if hx >= uint32(2146435072) {
		return x
	}
	if k != 0 {
		u.f = float64(int32(1)) + x
		hu = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
		hu += uint32(614242)
		k = int32(hu>>int32(20)) - int32(1023)
		if k < int32(54) {
			c = func() float64 {
				if k >= int32(2) {
					return float64(int32(1)) - (u.f - x)
				} else {
					return x - (u.f - float64(int32(1)))
				}
			}()
			c /= u.f
		} else {
			c = float64(int32(0))
		}
		hu = hu&uint32(1048575) + uint32(1072079006)
		*(*uint64)(unsafe.Pointer(&u)) = uint64(hu)<<int32(32) | *(*uint64)(unsafe.Pointer(&u))&uint64(4294967295)
		f = u.f - float64(int32(1))
	}
	hfsq = 0.5 * f * f
	s = f / (2 + f)
	z = s * s
	w = z * z
	t1 = w * (_cgos_Lg2_log1p + w*(_cgos_Lg4_log1p+w*_cgos_Lg6_log1p))
	t2 = z * (_cgos_Lg1_log1p + w*(_cgos_Lg3_log1p+w*(_cgos_Lg5_log1p+w*_cgos_Lg7_log1p)))
	R = t2 + t1
	dk = float64(k)
	return s*(hfsq+R) + (dk*_cgos_ln2_lo_log1p + c) - hfsq + f + dk*_cgos_ln2_hi_log1p
}
