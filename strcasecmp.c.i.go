package libc

import unsafe "unsafe"

func strcasecmp(_l *int8, _r *int8) int32 {
	var l *uint8 = (*uint8)(unsafe.Pointer(_l))
	var r *uint8 = (*uint8)(unsafe.Pointer(_r))
	for ; int32(*l) != 0 && int32(*r) != 0 && (int32(*l) == int32(*r) || tolower(int32(*l)) == tolower(int32(*r))); func() *uint8 {
		*(*uintptr)(unsafe.Pointer(&l))++
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
	}
	return tolower(int32(*l)) - tolower(int32(*r))
}
func __strcasecmp_l(l *int8, r *int8, loc *struct___locale_struct) int32 {
	return strcasecmp(l, r)
}
