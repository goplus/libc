package libc

import unsafe "unsafe"

func Vdprintf(fd int32, fmt *int8, ap []interface {
}) int32 {
	var f struct__IO_FILE = struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, nil, __stdio_write, nil, (*uint8)(unsafe.Pointer(fmt)), uint64(0), nil, nil, fd, 0, 0, 0, -1, -1, nil, 0, nil, nil, nil, 0, 0, nil, nil, nil}
	return Vfprintf(&f, fmt, ap)
}
