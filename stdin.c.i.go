package libc

import unsafe "unsafe"

var buf_cgo407 [1032]uint8
var __stdin_FILE struct__IO_FILE

func init() {
	__stdin_FILE = struct__IO_FILE{uint32(1 | 8), nil, nil, __stdio_close, nil, nil, nil, nil, __stdio_read, nil, __stdio_seek, (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&buf_cgo407)))) + uintptr(8))), 1032 - uint64(8), nil, nil, 0, 0, 0, 0, -1, 0, nil, 0, nil, nil, nil, 0, 0, nil, nil, nil}
}

var Stdin *struct__IO_FILE = &__stdin_FILE
var __stdin_used *struct__IO_FILE = &__stdin_FILE
