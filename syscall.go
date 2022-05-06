package libc

import "syscall"

func __syscall3(n int64, a int64, b int64, c int64) int64 {
	_, _, ret := syscall.Syscall(uintptr(n), uintptr(a), uintptr(b), uintptr(c))
	return int64(ret)
}
