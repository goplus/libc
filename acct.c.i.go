package libc

import unsafe "unsafe"

func acct(filename *int8) int32 {
	return int32(__syscall1_r1(int64(51), int64(uintptr(unsafe.Pointer(filename)))))
}
