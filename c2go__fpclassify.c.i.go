package libc

import unsafe "unsafe"

func __fpclassify(x float64) int32 {
	type _cgoa_51 struct {
		f float64
	}
	var u _cgoa_51
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> 52 & uint64(2047))
	if !(e != 0) {
		return func() int32 {
			if *(*uint64)(unsafe.Pointer(&u))<<1 != 0 {
				return 3
			} else {
				return 2
			}
		}()
	}
	if e == 2047 {
		return func() int32 {
			if *(*uint64)(unsafe.Pointer(&u))<<12 != 0 {
				return 0
			} else {
				return 1
			}
		}()
	}
	return int32(4)
}
