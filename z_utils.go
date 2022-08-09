package libc

import (
	"unsafe"
)

type rawBytes = *[1 << 30]byte

func toBytes(buf *int8, size uint64) []byte {
	return rawBytes(unsafe.Pointer(buf))[:size:size]
}

func setErrno(err error) {
	panic("todo")
}
