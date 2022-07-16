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
func __lock(*int32) {
	panic("notimpl")
}
func __uflow(*struct__IO_FILE) int32 {
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
func copysignl(float64, float64) float64 {
	panic("notimpl")
}
func fabsl(float64) float64 {
	panic("notimpl")
}
func fmodl(float64, float64) float64 {
	panic("notimpl")
}
func scalbn(float64, int32) float64 {
	panic("notimpl")
}
func scalbnl(float64, int32) float64 {
	panic("notimpl")
}
func wctomb(*int8, uint32) int32 {
	panic("notimpl")
}
