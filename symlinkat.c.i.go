package libc

import unsafe "unsafe"

func symlinkat(existing *int8, fd int32, new *int8) int32 {
	return int32(__syscall3_r1(int64(474), int64(uintptr(unsafe.Pointer(existing))), int64(fd), int64(uintptr(unsafe.Pointer(new)))))
}
