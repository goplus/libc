package libc

import unsafe "unsafe"

func wmemchr(s *uint32, c uint32, n uint64) *uint32 {
	for ; n != 0 && *s != c; func() *uint32 {
		n--
		return func() (_cgo_ret *uint32) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return func() *uint32 {
		if n != 0 {
			return (*uint32)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
