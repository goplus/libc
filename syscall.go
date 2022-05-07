package libc

import "syscall"

var g_r1 uintptr // TODO: goroutine safe

func __syscall1(n int64, a int64) int64 {
	var ret syscall.Errno
	g_r1, _, ret = syscall.Syscall(uintptr(n), uintptr(a), 0, 0)
	return int64(ret)
}

func __syscall3(n int64, a int64, b int64, c int64) int64 {
	var ret syscall.Errno
	g_r1, _, ret = syscall.Syscall(uintptr(n), uintptr(a), uintptr(b), uintptr(c))
	return int64(ret)
}

func __syscall4(n int64, a int64, b int64, c int64, d int64) int64 {
	var ret syscall.Errno
	g_r1, _, ret = syscall.Syscall6(uintptr(n), uintptr(a), uintptr(b), uintptr(c), uintptr(d), 0, 0)
	return int64(ret)
}

func __syscall6(n int64, a int64, b int64, c int64, d int64, e int64, f int64) int64 {
	var ret syscall.Errno
	g_r1, _, ret = syscall.Syscall6(uintptr(n), uintptr(a), uintptr(b), uintptr(c), uintptr(d), uintptr(e), uintptr(f))
	return int64(ret)
}

func __syscall_ret(uint64) int64 {
	r := g_r1
	if r > 18446744073709547520 {
		*__errno_location() = int32(-r)
		return int64(-1)
	}
	return int64(r)
}
