package libc

import unsafe "unsafe"

func Strncat(d *int8, s *int8, n uint64) *int8 {
	var a *int8 = d
	*(*uintptr)(unsafe.Pointer(&d)) += uintptr(Strlen(d))
	for n != 0 && int32(*s) != 0 {
		func() int8 {
			n--
			return func() (_cgo_ret int8) {
				_cgo_addr := &*func() (_cgo_ret *int8) {
					_cgo_addr := &d
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
				*_cgo_addr = *func() (_cgo_ret *int8) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
				return *_cgo_addr
			}()
		}()
	}
	*func() (_cgo_ret *int8) {
		_cgo_addr := &d
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = int8(0)
	return a
}
