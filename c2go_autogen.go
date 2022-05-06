package libc

import unsafe "unsafe"

type struct___locale_struct struct {
}
type struct_sigevent struct {
}

func ___errno_location() *int32 {
	panic("notimpl")
}
func __aio_close(int32) int32 {
	panic("notimpl")
}
func __futexwait(addr unsafe.Pointer, val int32, priv int32) {
	panic("notimpl")
}
func __lock(*int32) {
	panic("notimpl")
}
func __pthread_self() *struct___pthread {
	panic("notimpl")
}
func __stdio_seek(*struct__IO_FILE, int64, int32) int64 {
	panic("notimpl")
}
func __syscall_cp(int64, int64, int64, int64, int64, int64, int64) int64 {
	panic("notimpl")
}
func __unlock(*int32) {
	panic("notimpl")
}
func __wake(addr unsafe.Pointer, cnt int32, priv int32) {
	panic("notimpl")
}
func strerror(int32) *int8 {
	panic("notimpl")
}
func wctomb(*int8, uint32) int32 {
	panic("notimpl")
}
