package libc

import unsafe "unsafe"

var _cgos_toint___rem_pio2f float64 = 1.5 / 2.2204460492503131e-16
var _cgos_pio4___rem_pio2f float64 = 0.78539818525314331
var _cgos_invpio2___rem_pio2f float64 = 0.63661977236758138
var _cgos_pio2_1___rem_pio2f float64 = 1.5707963109016418
var _cgos_pio2_1t___rem_pio2f float64 = 1.5893254773528196e-8

func __rem_pio2f(x float32, y *float64) int32 {
	type _cgoa_18___rem_pio2f struct {
		f float32
	}
	var u _cgoa_18___rem_pio2f
	u.f = x
	var tx [1]float64
	var ty [1]float64
	var fn float64
	var ix uint32
	var n int32
	var sign int32
	var e0 int32
	ix = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
	if ix < uint32(1305022427) {
		fn = float64(x)*_cgos_invpio2___rem_pio2f + _cgos_toint___rem_pio2f - _cgos_toint___rem_pio2f
		n = int32(fn)
		*y = float64(x) - fn*_cgos_pio2_1___rem_pio2f - fn*_cgos_pio2_1t___rem_pio2f
		if func() int64 {
			if *y < -_cgos_pio4___rem_pio2f {
				return 1
			} else {
				return 0
			}
		}() == int64(0) {
			n--
			fn--
			*y = float64(x) - fn*_cgos_pio2_1___rem_pio2f - fn*_cgos_pio2_1t___rem_pio2f
		} else if func() int64 {
			if *y > _cgos_pio4___rem_pio2f {
				return 1
			} else {
				return 0
			}
		}() == int64(0) {
			n++
			fn++
			*y = float64(x) - fn*_cgos_pio2_1___rem_pio2f - fn*_cgos_pio2_1t___rem_pio2f
		}
		return n
	}
	if ix >= uint32(2139095040) {
		*y = float64(x - x)
		return int32(0)
	}
	sign = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(31))
	e0 = int32(ix>>int32(23) - uint32(150))
	*(*uint32)(unsafe.Pointer(&u)) = ix - uint32(e0<<int32(23))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&tx)))) + uintptr(int32(0))*8)) = float64(u.f)
	n = __rem_pio2_large((*float64)(unsafe.Pointer(&tx)), (*float64)(unsafe.Pointer(&ty)), e0, int32(1), int32(0))
	if sign != 0 {
		*y = -*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&ty)))) + uintptr(int32(0))*8))
		return -n
	}
	*y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&ty)))) + uintptr(int32(0))*8))
	return n
}
