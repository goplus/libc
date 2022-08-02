package libc

import unsafe "unsafe"

func link(existing *int8, new *int8) int32 {
	return int32(__syscall2_r1(int64(9), int64(uintptr(unsafe.Pointer(existing))), int64(uintptr(unsafe.Pointer(new)))))
}
