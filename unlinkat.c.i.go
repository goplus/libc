package libc

import unsafe "unsafe"

func unlinkat(fd int32, path *int8, flag int32) int32 {
	return int32(__syscall3_r1(int64(472), int64(fd), int64(uintptr(unsafe.Pointer(path))), int64(flag)))
}
