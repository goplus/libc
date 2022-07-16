package libc

import unsafe "unsafe"

func dummy_cgo18___lctrans(msg *int8, lm *struct___locale_map) *int8 {
	return msg
}
func __lctrans(msg *int8, lm *struct___locale_map) *int8 {
	return __lctrans_impl(msg, lm)
}
func __lctrans_cur(msg *int8) *int8 {
	return __lctrans_impl(msg, *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().locale.cat)))) + uintptr(int32(5))*8)))
}
