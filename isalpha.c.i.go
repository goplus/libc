package libc

func Isalpha(c int32) int32 {
	return func() int32 {
		if uint32(c)|uint32(32)-uint32('a') < uint32(26) {
			return 1
		} else {
			return 0
		}
	}()
}
func __isalpha_l(c int32, l *Struct___locale_struct) int32 {
	return Isalpha(c)
}
