package libc

import unsafe "unsafe"

func __ctype_get_mb_cur_max() uint64 {
	return uint64(func() int32 {
		if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().locale.cat)))) + uintptr(int32(0))*8)) != nil) {
			return int32(4)
		} else {
			return int32(1)
		}
	}())
}
