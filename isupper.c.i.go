package libc

func Isupper(c int32) int32 {
	return func() int32 {
		if uint32(c)-uint32('A') < uint32(26) {
			return 1
		} else {
			return 0
		}
	}()
}
func __isupper_l(c int32, l *Struct___locale_struct) int32 {
	return Isupper(c)
}
