package libc

import unsafe "unsafe"

var buf_cgo925 [8]uint8
var __stderr_FILE struct__IO_FILE

func init() {
	__stderr_FILE = struct__IO_FILE{uint32(5), nil, nil, __stdio_close, nil, nil, nil, nil, nil, __stdio_write, __stdio_seek, (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&buf_cgo925)))) + uintptr(int32(8)))), uint64(0), nil, nil, int32(2), 0, 0, 0, -1, -1, nil, 0, nil, nil, nil, 0, 0, nil, nil, nil}
}

var Stderr *struct__IO_FILE = &__stderr_FILE
var __stderr_used *struct__IO_FILE = &__stderr_FILE
