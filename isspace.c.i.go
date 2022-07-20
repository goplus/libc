package libc

func Isspace(c int32) int32 {
	return func() int32 {
		if c == ' ' || uint32(c)-uint32('\t') < uint32(5) {
			return 1
		} else {
			return 0
		}
	}()
}
func __isspace_l(c int32, l *struct___locale_struct) int32 {
	return Isspace(c)
}
