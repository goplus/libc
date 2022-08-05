package libc

import unsafe "unsafe"

func _cgos_dummy_unsetenv(old *int8, new *int8) {
}
func Unsetenv(name *int8) int32 {
	var l uint64 = uint64(uintptr(unsafe.Pointer(__strchrnul(name, '='))) - uintptr(unsafe.Pointer(name)))
	if !(l != 0) || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(name)) + uintptr(l)))) != 0 {
		*X__errno_location() = int32(22)
		return -1
	}
	if X__environ != nil {
		var e **int8 = X__environ
		var eo **int8 = e
		for ; *e != nil; *(*uintptr)(unsafe.Pointer(&e)) += 8 {
			if !(Strncmp(name, *e, l) != 0) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*e)) + uintptr(l)))) == '=' {
				__env_rm_add(*e, nil)
			} else if uintptr(unsafe.Pointer(eo)) != uintptr(unsafe.Pointer(e)) {
				*func() (_cgo_ret **int8) {
					_cgo_addr := &eo
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
					return
				}() = *e
			} else {
				*(*uintptr)(unsafe.Pointer(&eo)) += 8
			}
		}
		if uintptr(unsafe.Pointer(eo)) != uintptr(unsafe.Pointer(e)) {
			*eo = (*int8)(nil)
		}
	}
	return int32(0)
}
