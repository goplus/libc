package libc

import unsafe "unsafe"

var _cgos_toint___rem_pio2 float64 = 1.5 / 2.2204460492503131e-16
var _cgos_pio4___rem_pio2 float64 = 0.78539816339744828
var _cgos_invpio2___rem_pio2 float64 = 0.63661977236758138
var _cgos_pio2_1___rem_pio2 float64 = 1.5707963267341256
var _cgos_pio2_1t___rem_pio2 float64 = 6.0771005065061922e-11
var _cgos_pio2_2___rem_pio2 float64 = 6.077100506303966e-11
var _cgos_pio2_2t___rem_pio2 float64 = 2.0222662487959506e-21
var _cgos_pio2_3___rem_pio2 float64 = 2.0222662487111665e-21
var _cgos_pio2_3t___rem_pio2 float64 = 8.4784276603688995e-32

func __rem_pio2(x float64, y *float64) int32 {

	type _cgoa_18___rem_pio2 struct {
		f float64
	}
	var u _cgoa_18___rem_pio2
	u.f = x
	var z float64
	var w float64
	var t float64
	var r float64
	var fn float64
	var tx [3]float64
	var ty [2]float64
	var ix uint32
	var sign int32
	var n int32
	var ex int32
	var ey int32
	var i int32
	sign = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	ix = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32) & uint64(2147483647))
	if !(ix <= uint32(1074752122)) {
		goto _cgol_1
	}
	if !(ix&uint32(1048575) == uint32(598523)) {
		goto _cgol_2
	}
	goto medium
_cgol_2:
	if ix <= uint32(1073928572) {
		if !(sign != 0) {
			z = x - _cgos_pio2_1___rem_pio2
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z - _cgos_pio2_1t___rem_pio2
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) - _cgos_pio2_1t___rem_pio2
			return int32(1)
		} else {
			z = x + _cgos_pio2_1___rem_pio2
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z + _cgos_pio2_1t___rem_pio2
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) + _cgos_pio2_1t___rem_pio2
			return -1
		}
	} else if !(sign != 0) {
		z = x - float64(int32(2))*_cgos_pio2_1___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z - float64(int32(2))*_cgos_pio2_1t___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) - float64(int32(2))*_cgos_pio2_1t___rem_pio2
		return int32(2)
	} else {
		z = x + float64(int32(2))*_cgos_pio2_1___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z + float64(int32(2))*_cgos_pio2_1t___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) + float64(int32(2))*_cgos_pio2_1t___rem_pio2
		return -2
	}
_cgol_1:
	if !(ix <= uint32(1075594811)) {
		goto _cgol_3
	}
	if !(ix <= uint32(1075183036)) {
		goto _cgol_5
	}
	if !(ix == uint32(1074977148)) {
		goto _cgol_6
	}
	goto medium
_cgol_6:
	if !(sign != 0) {
		z = x - float64(int32(3))*_cgos_pio2_1___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z - float64(int32(3))*_cgos_pio2_1t___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) - float64(int32(3))*_cgos_pio2_1t___rem_pio2
		return int32(3)
	} else {
		z = x + float64(int32(3))*_cgos_pio2_1___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z + float64(int32(3))*_cgos_pio2_1t___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) + float64(int32(3))*_cgos_pio2_1t___rem_pio2
		return -3
	}
	goto _cgol_4
_cgol_5:
	if !(ix == uint32(1075388923)) {
		goto _cgol_7
	}
	goto medium
_cgol_7:
	if !(sign != 0) {
		z = x - float64(int32(4))*_cgos_pio2_1___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z - float64(int32(4))*_cgos_pio2_1t___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) - float64(int32(4))*_cgos_pio2_1t___rem_pio2
		return int32(4)
	} else {
		z = x + float64(int32(4))*_cgos_pio2_1___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = z + float64(int32(4))*_cgos_pio2_1t___rem_pio2
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) + float64(int32(4))*_cgos_pio2_1t___rem_pio2
		return -4
	}
_cgol_4:
	;
_cgol_3:
	if !(ix < uint32(1094263291)) {
		goto _cgol_8
	}
medium:
	fn = float64(x)*_cgos_invpio2___rem_pio2 + _cgos_toint___rem_pio2 - _cgos_toint___rem_pio2
	n = int32(fn)
	r = x - fn*_cgos_pio2_1___rem_pio2
	w = fn * _cgos_pio2_1t___rem_pio2
	if func() int64 {
		if r-w < -_cgos_pio4___rem_pio2 {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		n--
		fn--
		r = x - fn*_cgos_pio2_1___rem_pio2
		w = fn * _cgos_pio2_1t___rem_pio2
	} else if func() int64 {
		if r-w > _cgos_pio4___rem_pio2 {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		n++
		fn++
		r = x - fn*_cgos_pio2_1___rem_pio2
		w = fn * _cgos_pio2_1t___rem_pio2
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = r - w
	u.f = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8))
	ey = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	ex = int32(ix >> int32(20))
	if ex-ey > int32(16) {
		t = r
		w = fn * _cgos_pio2_2___rem_pio2
		r = t - w
		w = fn*_cgos_pio2_2t___rem_pio2 - (t - r - w)
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = r - w
		u.f = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8))
		ey = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
		if ex-ey > int32(49) {
			t = r
			w = fn * _cgos_pio2_3___rem_pio2
			r = t - w
			w = fn*_cgos_pio2_3t___rem_pio2 - (t - r - w)
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = r - w
		}
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = r - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) - w
	return n
_cgol_8:
	if ix >= uint32(2146435072) {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = func() (_cgo_ret float64) {
			_cgo_addr := &*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8))
			*_cgo_addr = x - x
			return *_cgo_addr
		}()
		return int32(0)
	}
	u.f = x
	*(*uint64)(unsafe.Pointer(&u)) &= 4503599627370495
	*(*uint64)(unsafe.Pointer(&u)) |= 4710765210229538816
	z = u.f
	for i = int32(0); i < int32(2); i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&tx)))) + uintptr(i)*8)) = float64(int32(z))
		z = (z - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&tx)))) + uintptr(i)*8))) * 16777216.0
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&tx)))) + uintptr(i)*8)) = z
	for *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&tx)))) + uintptr(i)*8)) == 0.0 {
		i--
	}
	n = __rem_pio2_large((*float64)(unsafe.Pointer(&tx)), (*float64)(unsafe.Pointer(&ty)), int32(ix>>int32(20))-1046, i+int32(1), int32(1))
	if sign != 0 {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = -*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&ty)))) + uintptr(int32(0))*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = -*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&ty)))) + uintptr(int32(1))*8))
		return -n
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&ty)))) + uintptr(int32(0))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&ty)))) + uintptr(int32(1))*8))
	return n
}
