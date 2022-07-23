package libc

import unsafe "unsafe"

func Strtok(s *int8, sep *int8) *int8 {
	if !(s != nil) && !(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = _cgos_p_strtok
		return *_cgo_addr
	}() != nil) {
		return (*int8)(nil)
	}
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(Strspn(s, sep))
	if !(*s != 0) {
		return func() (_cgo_ret *int8) {
			_cgo_addr := &_cgos_p_strtok
			*_cgo_addr = (*int8)(nil)
			return *_cgo_addr
		}()
	}
	_cgos_p_strtok = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Strcspn(s, sep))))
	if *_cgos_p_strtok != 0 {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &_cgos_p_strtok
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(0)
	} else {
		_cgos_p_strtok = (*int8)(nil)
	}
	return s
}

var _cgos_p_strtok *int8
