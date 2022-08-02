package libc

import unsafe "unsafe"

func preadv(fd int32, iov *Struct_iovec, count int32, ofs int64) int64 {
	return __syscall_ret(uint64(__syscall_cp(int64(540), int64(fd), int64(uintptr(unsafe.Pointer(iov))), int64(count), int64(int64(ofs)), int64(int64(ofs>>int32(32))), int64(0))))
}
