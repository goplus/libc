package libc

func iswxdigit(wc uint32) int32 {
	return func() int32 {
		if uint32(wc-uint32('0')) < uint32(10) || uint32(wc|uint32(32)-uint32('a')) < uint32(6) {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswxdigit_l(c uint32, l *Struct___locale_struct) int32 {
	return iswxdigit(c)
}
