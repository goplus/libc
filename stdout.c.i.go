package libc

import unsafe "unsafe"

var _cgos_buf_stdout [1032]uint8
var __stdout_FILE Struct__IO_FILE

func init() {
	__stdout_FILE = Struct__IO_FILE{uint32(5), nil, nil, __stdio_close, nil, nil, nil, nil, nil, __stdout_write, __stdio_seek, (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_buf_stdout)))) + uintptr(int32(8)))), 1024, nil, nil, int32(1), 0, 0, 0, -1, '\n', nil, 0, nil, nil, nil, 0, 0, nil, nil, nil}
}

var Stdout *Struct__IO_FILE = &__stdout_FILE
var __stdout_used *Struct__IO_FILE = &__stdout_FILE
