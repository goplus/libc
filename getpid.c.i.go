package libc

func Getpid() int32 {
	return int32(__syscall0(int64(20)))
}
