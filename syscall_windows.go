package libc

import (
	"syscall"
)

func __syscall0(n int64) int64 {
	r1, _, ret := syscall.Syscall(uintptr(n), 0, 0, 0, 0)
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall0_r1(n int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), 0, 0, 0, 0)
	return int64(r1)
}

func __syscall1(n int64, a int64) int64 {
	r1, _, ret := syscall.Syscall(uintptr(n), 1, uintptr(a), 0, 0)
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall1_r1(n int64, a int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), 1, uintptr(a), 0, 0)
	return int64(r1)
}

func __syscall2(n int64, a, b int64) int64 {
	r1, _, ret := syscall.Syscall(uintptr(n), 2, uintptr(a), uintptr(b), 0)
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall2_r1(n int64, a, b int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), 2, uintptr(a), uintptr(b), 0)
	return int64(r1)
}

func __syscall3(n int64, a int64, b int64, c int64) int64 {
	r1, _, ret := syscall.Syscall(uintptr(n), 3, uintptr(a), uintptr(b), uintptr(c))
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall3_r1(n int64, a int64, b int64, c int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), 3, uintptr(a), uintptr(b), uintptr(c))
	return int64(r1)
}

func __syscall4(n int64, a int64, b int64, c int64, d int64) int64 {
	r1, _, ret := syscall.Syscall6(uintptr(n), 4, uintptr(a), uintptr(b), uintptr(c), uintptr(d), 0, 0)
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall4_r1(n int64, a int64, b int64, c int64, d int64) int64 {
	r1, _, _ := syscall.Syscall6(uintptr(n), 4, uintptr(a), uintptr(b), uintptr(c), uintptr(d), 0, 0)
	return int64(r1)
}

func __syscall5(n int64, a int64, b int64, c int64, d int64, e int64) int64 {
	r1, _, ret := syscall.Syscall6(uintptr(n), 5, uintptr(a), uintptr(b), uintptr(c), uintptr(d), uintptr(e), 0)
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall5_r1(n int64, a int64, b int64, c int64, d int64, e int64) int64 {
	r1, _, _ := syscall.Syscall6(uintptr(n), 5, uintptr(a), uintptr(b), uintptr(c), uintptr(d), uintptr(e), 0)
	return int64(r1)
}

func __syscall6(n int64, a int64, b int64, c int64, d int64, e int64, f int64) int64 {
	r1, _, ret := syscall.Syscall6(uintptr(n), 6, uintptr(a), uintptr(b), uintptr(c), uintptr(d), uintptr(e), uintptr(f))
	__pthread_self().sys_r1 = int64(r1)
	return int64(ret)
}

func __syscall6_r1(n int64, a int64, b int64, c int64, d int64, e int64, f int64) int64 {
	r1, _, _ := syscall.Syscall6(uintptr(n), 6, uintptr(a), uintptr(b), uintptr(c), uintptr(d), uintptr(e), uintptr(f))
	return int64(r1)
}

func __syscall_cp(n int64, a int64, b int64, c int64, d int64, e int64, f int64) int64 {
	return __syscall6(n, a, b, c, d, e, f)
}

func __syscall_ret(uint64) int64 {
	return __pthread_self().sys_r1
}
