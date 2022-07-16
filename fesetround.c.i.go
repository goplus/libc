package libc

func Fesetround(r int32) int32 {
	if r != int32(0) {
		return -1
	}
	return __fesetround(r)
}
