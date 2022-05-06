package libc

import "syscall"

var g_r1 uintptr // TODO: goroutine safe

func __syscall3(n int64, a int64, b int64, c int64) int64 {
	var ret syscall.Errno
	g_r1, _, ret = syscall.Syscall(uintptr(n), uintptr(a), uintptr(b), uintptr(c))
	return int64(ret)
}

func __syscall_ret(uint64) int64 {
	return int64(g_r1)
}
