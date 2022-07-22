package libc

func Llrintf(x float32) int64 {
	return int64(Rintf(x))
}
