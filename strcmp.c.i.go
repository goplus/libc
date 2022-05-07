package libc

import unsafe "unsafe"

func Strcmp(l *int8, r *int8) int32 {
	for ; int32(*l) == int32(*r) && int32(*l) != 0; func() *int8 {
		*(*uintptr)(unsafe.Pointer(&l))++
		return func() (_cgo_ret *int8) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
	}
	return int32(*(*uint8)(unsafe.Pointer(l))) - int32(*(*uint8)(unsafe.Pointer(r)))
}
