package libc

func sync() {
	__syscall0(int64(36))
}
