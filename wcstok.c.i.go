package libc

import unsafe "unsafe"

func wcstok(s *uint32, sep *uint32, p **uint32) *uint32 {
	if !(s != nil) && !(func() (_cgo_ret *uint32) {
		_cgo_addr := &s
		*_cgo_addr = *p
		return *_cgo_addr
	}() != nil) {
		return (*uint32)(nil)
	}
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(wcsspn(s, sep)) * 4
	if !(*s != 0) {
		return func() (_cgo_ret *uint32) {
			_cgo_addr := &*p
			*_cgo_addr = (*uint32)(nil)
			return *_cgo_addr
		}()
	}
	*p = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(wcscspn(s, sep))*4))
	if **p != 0 {
		*func() (_cgo_ret *uint32) {
			_cgo_addr := &*p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = uint32(0)
	} else {
		*p = (*uint32)(nil)
	}
	return s
}
