package libc

import unsafe "unsafe"

func X__fpclassifyf(x float32) int32 {
	type _cgoa_18___fpclassifyf struct {
		f float32
	}
	var u _cgoa_18___fpclassifyf
	u.f = x
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(23) & uint32(255))
	if !(e != 0) {
		return func() int32 {
			if *(*uint32)(unsafe.Pointer(&u))<<int32(1) != 0 {
				return int32(3)
			} else {
				return int32(2)
			}
		}()
	}
	if e == int32(255) {
		return func() int32 {
			if *(*uint32)(unsafe.Pointer(&u))<<int32(9) != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}()
	}
	return int32(4)
}
