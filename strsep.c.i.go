package libc

import unsafe "unsafe"

func Strsep(str **int8, sep *int8) *int8 {
	var s *int8 = *str
	var end *int8
	if !(s != nil) {
		return (*int8)(nil)
	}
	end = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Strcspn(s, sep))))
	if *end != 0 {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &end
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(0)
	} else {
		end = (*int8)(nil)
	}
	*str = end
	return s
}
