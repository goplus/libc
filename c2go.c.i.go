package libc

import unsafe "unsafe"

func __stdio_exit_needed() {
	__stdio_exit()
}
func Memrchr(m unsafe.Pointer, c int32, n uint64) unsafe.Pointer {
	return __memrchr(m, c, n)
}
