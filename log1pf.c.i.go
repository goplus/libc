package libc

import unsafe "unsafe"

var _cgos_ln2_hi_log1pf float32 = float32(0.69313812255999996)
var _cgos_ln2_lo_log1pf float32 = float32(9.0580006144999996e-6)
var _cgos_Lg1_log1pf float32 = float32(0.66666662693023682)
var _cgos_Lg2_log1pf float32 = float32(0.40000972151756287)
var _cgos_Lg3_log1pf float32 = float32(0.28498786687850952)
var _cgos_Lg4_log1pf float32 = float32(0.24279078841209412)

func Log1pf(x float32) float32 {
	type _cgoa_18_log1pf struct {
		f float32
	}
	var u _cgoa_18_log1pf
	u.f = x
	var hfsq float32
	var f float32
	var c float32
	var s float32
	var z float32
	var R float32
	var w float32
	var t1 float32
	var t2 float32
	var dk float32
	var ix uint32
	var iu uint32
	var k int32
	ix = *(*uint32)(unsafe.Pointer(&u))
	k = int32(1)
	if ix < uint32(1054086096) || ix>>int32(31) != 0 {
		if ix >= uint32(3212836864) {
			if x == float32(-1) {
				return x / 0.0
			}
			return (x - x) / 0.0
		}
		if ix<<int32(1) < uint32(1728053248) {
			if ix&uint32(2139095040) == uint32(0) {
				for {
					if 4 == 4 {
						fp_force_evalf(x * x)
					} else if 4 == 8 {
						fp_force_eval(float64(x * x))
					} else {
						fp_force_evall(float64(x * x))
					}
					if true {
						break
					}
				}
			}
			return x
		}
		if ix <= uint32(3197498905) {
			k = int32(0)
			c = float32(int32(0))
			f = x
		}
	} else if ix >= uint32(2139095040) {
		return x
	}
	if k != 0 {
		u.f = float32(int32(1)) + x
		iu = *(*uint32)(unsafe.Pointer(&u))
		iu += uint32(4913933)
		k = int32(iu>>int32(23)) - int32(127)
		if k < int32(25) {
			c = func() float32 {
				if k >= int32(2) {
					return float32(int32(1)) - (u.f - x)
				} else {
					return x - (u.f - float32(int32(1)))
				}
			}()
			c /= u.f
		} else {
			c = float32(int32(0))
		}
		iu = iu&uint32(8388607) + uint32(1060439283)
		*(*uint32)(unsafe.Pointer(&u)) = iu
		f = u.f - float32(int32(1))
	}
	s = f / (2.0 + f)
	z = s * s
	w = z * z
	t1 = w * (_cgos_Lg2_log1pf + w*_cgos_Lg4_log1pf)
	t2 = z * (_cgos_Lg1_log1pf + w*_cgos_Lg3_log1pf)
	R = t2 + t1
	hfsq = 0.5 * f * f
	dk = float32(k)
	return s*(hfsq+R) + (dk*_cgos_ln2_lo_log1pf + c) - hfsq + f + dk*_cgos_ln2_hi_log1pf
}
