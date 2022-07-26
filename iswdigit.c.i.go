package libc

func iswdigit(wc uint32) int32 {
	return func() int32 {
		if uint32(wc)-uint32('0') < uint32(10) {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswdigit_l(c uint32, l *Struct___locale_struct) int32 {
	return iswdigit(c)
}
