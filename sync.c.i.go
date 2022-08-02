package libc

func Sync() {
	__syscall0(int64(36))
}
