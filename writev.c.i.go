package libc

import unsafe "unsafe"

func writev(fd int32, iov *Struct_iovec, count int32) int64 {
	return __syscall_ret(uint64(__syscall_cp(int64(121), int64(fd), int64(uintptr(unsafe.Pointer(iov))), int64(count), int64(0), int64(0), int64(0))))
}
