package libc

import unsafe "unsafe"

func _cgos_dummy_clearenv(old *int8, new *int8) {
}
func Clearenv() int32 {
	var e **int8 = X__environ
	X__environ = (**int8)(nil)
	if e != nil {
		for *e != nil {
			__env_rm_add(*func() (_cgo_ret **int8) {
				_cgo_addr := &e
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
				return
			}(), nil)
		}
	}
	return int32(0)
}
