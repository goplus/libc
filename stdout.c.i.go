package libc

import unsafe "unsafe"

var buf_cgo46 [1032]uint8
var __stdout_FILE struct__IO_FILE

func init() {
	__stdout_FILE = struct__IO_FILE{uint32(1 | 4), nil, nil, __stdio_close, nil, nil, nil, nil, nil, __stdout_write, __stdio_seek, (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&buf_cgo46)))) + uintptr(8))), 1032 - uint64(8), nil, nil, 1, 0, 0, 0, -1, '\n', nil, 0, nil, nil, nil, 0, 0, nil, nil, nil}
}

var stdout *struct__IO_FILE = &__stdout_FILE
var __stdout_used *struct__IO_FILE = &__stdout_FILE
