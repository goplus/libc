package main

import unsafe "unsafe"

var t_status int32 = 0

func t_printf(s *int8, __cgo_args ...interface {
}) int32 {
	var ap []interface {
	}
	var buf [512]int8
	var n int32
	t_status = int32(1)
	ap = __cgo_args
	n = vsnprintf((*int8)(unsafe.Pointer(&buf)), 512, s, ap)
	if n < 0 {
		n = int32(0)
	} else if uint64(n) >= 512 {
		n = int32(512)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-1))) = int8('\n')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-2))) = int8('.')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-3))) = int8('.')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-4))) = int8('.')
	}
	return int32(write(1, unsafe.Pointer((*int8)(unsafe.Pointer(&buf))), uint64(n)))
}
