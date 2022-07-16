package libc

import unsafe "unsafe"

func wcsncat(d *uint32, s *uint32, n uint64) *uint32 {
	var a *uint32 = d
	*(*uintptr)(unsafe.Pointer(&d)) += uintptr(wcslen(d)) * 4
	for n != 0 && *s != 0 {
		func() uint32 {
			n--
			return func() (_cgo_ret uint32) {
				_cgo_addr := &*func() (_cgo_ret *uint32) {
					_cgo_addr := &d
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
					return
				}()
				*_cgo_addr = *func() (_cgo_ret *uint32) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
					return
				}()
				return *_cgo_addr
			}()
		}()
	}
	*func() (_cgo_ret *uint32) {
		_cgo_addr := &d
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
		return
	}() = uint32(0)
	return a
}
