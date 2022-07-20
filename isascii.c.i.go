package libc

func Isascii(c int32) int32 {
	return func() int32 {
		if !(c&-128 != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
