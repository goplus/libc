package libc

import unsafe "unsafe"

type struct_sigevent struct {
}

func __aio_close(int32) int32 {
	panic("notimpl")
}
func __futexwait(addr unsafe.Pointer, val int32, priv int32) {
	panic("notimpl")
}
func __libc_malloc(uint64) unsafe.Pointer {
	panic("notimpl")
}
func __lock(*int32) {
	panic("notimpl")
}
func __syscall_cp(int64, int64, int64, int64, int64, int64, int64) int64 {
	panic("notimpl")
}
func __unlock(*int32) {
	panic("notimpl")
}
func __vm_wait() {
	panic("notimpl")
}
func __wake(addr unsafe.Pointer, cnt int32, priv int32) {
	panic("notimpl")
}
func wctomb(*int8, uint32) int32 {
	panic("notimpl")
}
