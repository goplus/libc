package libc

import unsafe "unsafe"

func wmemmove(d *uint32, s *uint32, n uint64) *uint32 {
	var d0 *uint32 = d
	if uintptr(unsafe.Pointer(d)) == uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(s)))) {
		return d
	}
	if uint64(uintptr(unsafe.Pointer(d)))-uint64(uintptr(unsafe.Pointer(s))) < n*4 {
		for func() (_cgo_ret uint64) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(n)*4)) = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n)*4))
		}
	} else {
		for func() (_cgo_ret uint64) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
		}
	}
	return d0
}
