package libc

import (
	"syscall"
	"unicode/utf16"
)

func Getcwd(buf *int8, size uint64) *int8 {
	const max = 300
	b := make([]uint16, max)
	n, e := syscall.GetCurrentDirectory(max, &b[0])
	if e != nil {
		setErrno(e)
		return nil
	}
	return string(utf16.Decode(b[0:n])), nil
	return buf
}

// TODO: can be faster
func fromUtf8(buf *int8, size uint64, src []uint16) *int8 {
	s := string(utf16.Decode(src))
	if uint64(len(s)) >= size {
		setErrno(syscall.ERANGE)
		return nil
	}
	psz := (*[1 << 30]byte)(unsafe.Pointer(buf))
	psz[copy(psz[:size], s)] = 0
	return buf
}
