package libc

import unsafe "unsafe"

func fputws(ws *uint32, f *Struct__IO_FILE) int32 {
	var buf [1024]uint8
	var l uint64 = uint64(0)
	var ploc **Struct___locale_struct = &__pthread_self().Locale
	var loc *Struct___locale_struct = *ploc
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	fwide(f, int32(1))
	*ploc = f.Locale
	for ws != nil && func() (_cgo_ret uint64) {
		_cgo_addr := &l
		*_cgo_addr = wcsrtombs((*int8)(unsafe.Pointer((*uint8)(unsafe.Pointer(&buf)))), (**uint32)(unsafe.Pointer(&ws)), 1024, nil)
		return *_cgo_addr
	}()+uint64(1) > uint64(1) {
		if __fwritex((*uint8)(unsafe.Pointer(&buf)), l, f) < l {
			for {
				if __need_unlock != 0 {
					__unlockfile(f)
				}
				if true {
					break
				}
			}
			*ploc = loc
			return -1
		}
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	*ploc = loc
	return int32(l)
}
