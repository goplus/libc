package libc

import unsafe "unsafe"

func linkat(fd1 int32, existing *int8, fd2 int32, new *int8, flag int32) int32 {
	return int32(__syscall5_r1(int64(471), int64(fd1), int64(uintptr(unsafe.Pointer(existing))), int64(fd2), int64(uintptr(unsafe.Pointer(new))), int64(flag)))
}
