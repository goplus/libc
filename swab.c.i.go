package libc

import unsafe "unsafe"

func swab(_src unsafe.Pointer, _dest unsafe.Pointer, n int64) {
	var src *int8 = (*int8)(_src)
	var dest *int8 = (*int8)(_dest)
	for ; n > int64(1); n -= int64(2) {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest)) + uintptr(int32(0)))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(int32(1))))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest)) + uintptr(int32(1)))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(int32(0))))
		*(*uintptr)(unsafe.Pointer(&dest)) += uintptr(int32(2))
		*(*uintptr)(unsafe.Pointer(&src)) += uintptr(int32(2))
	}
}
