package libc

import (
	"golang.org/x/sys/unix"
)

func Getcwd(buf *int8, size uint64) *int8 {
	if _, err := unix.Getcwd(toBytes(buf, size)); err != nil {
		setErrno(err)
		return nil
	}
	return buf
}
