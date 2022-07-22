package libc

import unsafe "unsafe"

var toint_cgo18_rint float64 = float64(int32(1)) / 2.2204460492503131e-16

func Rint(x float64) float64 {
	type _cgoa_19_rint struct {
		f float64
	}
	var u _cgoa_19_rint
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	var s int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	var y float64
	if e >= 1075 {
		return x
	}
	if s != 0 {
		y = x - toint_cgo18_rint + toint_cgo18_rint
	} else {
		y = x + toint_cgo18_rint - toint_cgo18_rint
	}
	if y == float64(int32(0)) {
		return func() float64 {
			if s != 0 {
				return -0
			} else {
				return float64(int32(0))
			}
		}()
	}
	return y
}
