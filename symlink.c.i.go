package libc

import unsafe "unsafe"

func symlink(existing *int8, new *int8) int32 {
	return int32(__syscall2_r1(int64(57), int64(uintptr(unsafe.Pointer(existing))), int64(uintptr(unsafe.Pointer(new)))))
}
