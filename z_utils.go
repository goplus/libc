package libc

import (
	"unsafe"
)

func toBytes(buf *int8, size uint64) []byte {
	return (*[1 << 30]byte)(unsafe.Pointer(buf))[:size:size]
}

func setErrno(err error) {
	panic("todo")
}
