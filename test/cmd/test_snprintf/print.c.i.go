package main

import (
	libc "github.com/goplus/libc"
	unsafe "unsafe"
)

var t_status int32 = int32(0)

func t_printf(s *int8, __cgo_args ...interface {
}) int32 {
	var ap []interface {
	}
	var buf [512]int8
	var n int32
	t_status = int32(1)
	ap = __cgo_args
	n = libc.Vsnprintf((*int8)(unsafe.Pointer(&buf)), 512, s, ap)
	if n < int32(0) {
		n = int32(0)
	} else if uint64(n) >= 512 {
		n = int32(512)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-int32(1)))) = int8('\n')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-int32(2)))) = int8('.')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-int32(3)))) = int8('.')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(n-int32(4)))) = int8('.')
	}
	return int32(libc.Write(int32(1), unsafe.Pointer((*int8)(unsafe.Pointer(&buf))), uint64(n)))
}
