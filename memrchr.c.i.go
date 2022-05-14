package libc

import unsafe "unsafe"

func __memrchr(m unsafe.Pointer, c int32, n uint64) unsafe.Pointer {
	var s *uint8 = (*uint8)(m)
	c = int32(uint8(c))
	for func() (_cgo_ret uint64) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n)))) == c {
			return unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n))))
		}
	}
	return unsafe.Pointer(nil)
}
