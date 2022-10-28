package libc

import unsafe "unsafe"

var _cgos_ivln10hi_log10f float32 = float32(0.43432617188)
var _cgos_ivln10lo_log10f float32 = float32(-3.1689971364999999e-5)
var _cgos_log10_2hi_log10f float32 = float32(0.30102920531999999)
var _cgos_log10_2lo_log10f float32 = float32(7.9034151667999997e-7)
var _cgos_Lg1_log10f float32 = float32(0.66666662693023682)
var _cgos_Lg2_log10f float32 = float32(0.40000972151756287)
var _cgos_Lg3_log10f float32 = float32(0.28498786687850952)
var _cgos_Lg4_log10f float32 = float32(0.24279078841209412)

func Log10f(x float32) float32 {
	type _cgoa_18_log10f struct {
		f float32
	}
	var u _cgoa_18_log10f
	u.f = x
	var hfsq float32
	var f float32
	var s float32
	var z float32
	var R float32
	var w float32
	var t1 float32
	var t2 float32
	var dk float32
	var hi float32
	var lo float32
	var ix uint32
	var k int32
	ix = *(*uint32)(unsafe.Pointer(&u))
	k = int32(0)
	if ix < uint32(8388608) || ix>>int32(31) != 0 {
		if ix<<int32(1) == uint32(0) {
			return float32(-1) / (x * x)
		}
		if ix>>int32(31) != 0 {
			return (x - x) / 0.0
		}
		k -= int32(25)
		x *= float32(33554432.0)
		u.f = x
		ix = *(*uint32)(unsafe.Pointer(&u))
	} else if ix >= uint32(2139095040) {
		return x
	} else if ix == uint32(1065353216) {
		return float32(int32(0))
	}
	ix += uint32(4913933)
	k += int32(ix>>int32(23)) - int32(127)
	ix = ix&uint32(8388607) + uint32(1060439283)
	*(*uint32)(unsafe.Pointer(&u)) = ix
	x = u.f
	f = x - 1.0
	s = f / (2.0 + f)
	z = s * s
	w = z * z
	t1 = w * (_cgos_Lg2_log10f + w*_cgos_Lg4_log10f)
	t2 = z * (_cgos_Lg1_log10f + w*_cgos_Lg3_log10f)
	R = t2 + t1
	hfsq = 0.5 * f * f
	hi = f - hfsq
	u.f = hi
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(4294963200)
	hi = u.f
	lo = f - hi - hfsq + s*(hfsq+R)
	dk = float32(k)
	return dk*_cgos_log10_2lo_log10f + (lo+hi)*_cgos_ivln10lo_log10f + lo*_cgos_ivln10hi_log10f + hi*_cgos_ivln10hi_log10f + dk*_cgos_log10_2hi_log10f
}
