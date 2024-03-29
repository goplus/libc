package libc

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func Getcwd(buf *int8, size uint64) *int8 {
	const max = 300
	b := make([]uint16, max)
	n, e := syscall.GetCurrentDirectory(max, &b[0])
	if e != nil {
		setErrno(e)
		return nil
	}
	return fromUtf8(buf, size, b[0:n])
}

// TODO: can be faster
func fromUtf8(buf *int8, size uint64, src []uint16) *int8 {
	s := string(utf16.Decode(src))
	if uint64(len(s)) >= size {
		setErrno(syscall.ERANGE)
		return nil
	}
	psz := rawBytes(unsafe.Pointer(buf))
	psz[copy(psz[:size:size], s)] = 0
	return buf
}
