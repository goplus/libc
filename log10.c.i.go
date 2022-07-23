package libc

import unsafe "unsafe"

var ivln10hi_cgos__log10 float64 = 0.43429448187816888
var ivln10lo_cgos__log10 float64 = 2.5082946711645275e-11
var log10_2hi_cgos__log10 float64 = 0.30102999566361177
var log10_2lo_cgos__log10 float64 = 3.6942390771589308e-13
var Lg1_cgos__log10 float64 = 0.66666666666667351
var Lg2_cgos__log10 float64 = 0.39999999999409419
var Lg3_cgos__log10 float64 = 0.28571428743662391
var Lg4_cgos__log10 float64 = 0.22222198432149784
var Lg5_cgos__log10 float64 = 0.1818357216161805
var Lg6_cgos__log10 float64 = 0.15313837699209373
var Lg7_cgos__log10 float64 = 0.14798198605116586

func Log10(x float64) float64 {
	type _cgoa_18_log10 struct {
		f float64
	}
	var u _cgoa_18_log10
	u.f = x
	var hfsq float64
	var f float64
	var s float64
	var z float64
	var R float64
	var w float64
	var t1 float64
	var t2 float64
	var dk float64
	var y float64
	var hi float64
	var lo float64
	var val_hi float64
	var val_lo float64
	var hx uint32
	var k int32
	hx = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
	k = int32(0)
	if hx < uint32(1048576) || hx>>int32(31) != 0 {
		if *(*uint64)(unsafe.Pointer(&u))<<int32(1) == uint64(0) {
			return float64(-1) / (x * x)
		}
		if hx>>int32(31) != 0 {
			return (x - x) / 0
		}
		k -= int32(54)
		x *= float64(18014398509481984)
		u.f = x
		hx = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
	} else if hx >= uint32(2146435072) {
		return x
	} else if hx == uint32(1072693248) && *(*uint64)(unsafe.Pointer(&u))<<int32(32) == uint64(0) {
		return float64(int32(0))
	}
	hx += uint32(614242)
	k += int32(hx>>int32(20)) - int32(1023)
	hx = hx&uint32(1048575) + uint32(1072079006)
	*(*uint64)(unsafe.Pointer(&u)) = uint64(hx)<<int32(32) | *(*uint64)(unsafe.Pointer(&u))&uint64(4294967295)
	x = u.f
	f = x - 1
	hfsq = 0.5 * f * f
	s = f / (2 + f)
	z = s * s
	w = z * z
	t1 = w * (Lg2_cgos__log10 + w*(Lg4_cgos__log10+w*Lg6_cgos__log10))
	t2 = z * (Lg1_cgos__log10 + w*(Lg3_cgos__log10+w*(Lg5_cgos__log10+w*Lg7_cgos__log10)))
	R = t2 + t1
	hi = f - hfsq
	u.f = hi
	*(*uint64)(unsafe.Pointer(&u)) &= 18446744069414584320
	hi = u.f
	lo = f - hi - hfsq + s*(hfsq+R)
	val_hi = hi * ivln10hi_cgos__log10
	dk = float64(k)
	y = dk * log10_2hi_cgos__log10
	val_lo = dk*log10_2lo_cgos__log10 + (lo+hi)*ivln10lo_cgos__log10 + lo*ivln10hi_cgos__log10
	w = y + val_hi
	val_lo += y - w + val_hi
	val_hi = w
	return val_lo + val_hi
}
