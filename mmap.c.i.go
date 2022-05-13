package libc

import unsafe "unsafe"

func dummy_cgo501() {
}
func __mmap(start unsafe.Pointer, len uint64, prot int32, flags int32, fd int32, off int64) unsafe.Pointer {
	var ret int64
	if uint64(off)&(0|(4096-uint64(1))) != 0 {
		*__errno_location() = int32(22)
		return unsafe.Pointer(uintptr(18446744073709551615))
	}
	if len >= uint64(2147483647) {
		*__errno_location() = int32(12)
		return unsafe.Pointer(uintptr(18446744073709551615))
	}
	if flags&16 != 0 {
		__vm_wait()
	}
	ret = __syscall6(int64(197), int64(uintptr(start)), int64(len), int64(prot), int64(flags), int64(fd), int64(off))
	if ret == int64(-1) && !(start != nil) && flags&32 != 0 && !(flags&16 != 0) {
		ret = int64(-12)
	}
	return unsafe.Pointer(uintptr(__syscall_ret(uint64(ret))))
}
