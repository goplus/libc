package libc

func Setsid() int32 {
	return int32(__syscall0_r1(int64(147)))
}
