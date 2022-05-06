package libc

func Isdigit(c int32) int32 {
	return func() int32 {
		if uint32(c)-uint32('0') < uint32(10) {
			return 1
		} else {
			return 0
		}
	}()
}
func __isdigit_l(c int32, l *struct___locale_struct) int32 {
	return Isdigit(c)
}
