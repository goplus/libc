package libc

import unsafe "unsafe"

func strtok_r(s *int8, sep *int8, p **int8) *int8 {
	if !(s != nil) && !(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = *p
		return *_cgo_addr
	}() != nil) {
		return (*int8)(nil)
	}
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(Strspn(s, sep))
	if !(*s != 0) {
		return func() (_cgo_ret *int8) {
			_cgo_addr := &*p
			*_cgo_addr = (*int8)(nil)
			return *_cgo_addr
		}()
	}
	*p = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Strcspn(s, sep))))
	if **p != 0 {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &*p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(0)
	} else {
		*p = (*int8)(nil)
	}
	return s
}
