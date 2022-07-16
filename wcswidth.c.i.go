package libc

import unsafe "unsafe"

func wcswidth(wcs *uint32, n uint64) int32 {
	var l int32 = int32(0)
	var k int32 = int32(0)
	for ; func() (_cgo_ret uint64) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 && *wcs != 0 && func() (_cgo_ret int32) {
		_cgo_addr := &k
		*_cgo_addr = wcwidth(*wcs)
		return *_cgo_addr
	}() >= int32(0); func() *uint32 {
		l += k
		return func() (_cgo_ret *uint32) {
			_cgo_addr := &wcs
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return func() int32 {
		if k < int32(0) {
			return k
		} else {
			return l
		}
	}()
}
