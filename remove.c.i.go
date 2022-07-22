package libc

import unsafe "unsafe"

func Remove(path *int8) int32 {
	var r int32 = int32(__syscall1(int64(10), int64(uintptr(unsafe.Pointer(path)))))
	if r == -21 {
		r = int32(__syscall1(int64(137), int64(uintptr(unsafe.Pointer(path)))))
	}
	return int32(__syscall_ret(uint64(r)))
}
