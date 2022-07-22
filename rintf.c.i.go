package libc

import unsafe "unsafe"

var toint_cgo18_rintf float32 = float32(int32(1)) / 1.1920929e-7

func Rintf(x float32) float32 {
	type _cgoa_19_rintf struct {
		f float32
	}
	var u _cgoa_19_rintf
	u.f = x
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(23) & uint32(255))
	var s int32 = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(31))
	var y float32
	if e >= 150 {
		return x
	}
	if s != 0 {
		y = x - toint_cgo18_rintf + toint_cgo18_rintf
	} else {
		y = x + toint_cgo18_rintf - toint_cgo18_rintf
	}
	if y == float32(int32(0)) {
		return func() float32 {
			if s != 0 {
				return -0
			} else {
				return 0
			}
		}()
	}
	return y
}
