package libc

import unsafe "unsafe"

func wcsncmp(l *uint32, r *uint32, n uint64) int32 {
	for ; n != 0 && *l == *r && *l != 0 && *r != 0; func() *uint32 {
		func() *uint32 {
			n--
			return func() (_cgo_ret *uint32) {
				_cgo_addr := &l
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
		}()
		return func() (_cgo_ret *uint32) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return int32(func() uint32 {
		if n != 0 {
			return *l - *r
		} else {
			return uint32(0)
		}
	}())
}
