package libc

func Lrintf(x float32) int64 {
	return int64(Rintf(x))
}
