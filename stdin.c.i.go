package libc

import unsafe "unsafe"

var _cgos_buf_stdin [1032]uint8
var __stdin_FILE Struct__IO_FILE

func init() {
	__stdin_FILE = Struct__IO_FILE{uint32(9), nil, nil, __stdio_close, nil, nil, nil, nil, __stdio_read, nil, __stdio_seek, (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_buf_stdin)))) + uintptr(int32(8)))), 1024, nil, nil, int32(0), 0, 0, 0, -1, 0, nil, 0, nil, nil, nil, 0, 0, nil, nil, nil}
}

var Stdin *Struct__IO_FILE = &__stdin_FILE
var __stdin_used *Struct__IO_FILE = &__stdin_FILE
