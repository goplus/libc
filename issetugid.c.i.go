package libc

func issetugid() int32 {
	return int32(__libc.secure)
}
