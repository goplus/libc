package libc

import unsafe "unsafe"

func wcscmp(l *uint32, r *uint32) int32 {
	for ; *l == *r && *l != 0 && *r != 0; func() *uint32 {
		*(*uintptr)(unsafe.Pointer(&l)) += 4
		return func() (_cgo_ret *uint32) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return int32(*l - *r)
}
