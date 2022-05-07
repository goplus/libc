package libc

import unsafe "unsafe"

func dummy_cgo489() {
}
func __munmap(start unsafe.Pointer, len uint64) int32 {
	__vm_wait()
	return int32(__syscall_ret(uint64(__syscall2(int64(73), int64(uintptr(start)), int64(len)))))
}
