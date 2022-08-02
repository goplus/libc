package libc

func dup(fd int32) int32 {
	return int32(__syscall1_r1(int64(41), int64(fd)))
}
