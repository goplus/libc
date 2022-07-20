package libc

func Isprint(c int32) int32 {
	return func() int32 {
		if uint32(c)-uint32(32) < uint32(95) {
			return 1
		} else {
			return 0
		}
	}()
}
func __isprint_l(c int32, l *struct___locale_struct) int32 {
	return Isprint(c)
}
