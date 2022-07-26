package libc

func iswcntrl(wc uint32) int32 {
	return func() int32 {
		if uint32(wc) < uint32(32) || uint32(wc-uint32(127)) < uint32(33) || uint32(wc-uint32(8232)) < uint32(2) || uint32(wc-uint32(65529)) < uint32(3) {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswcntrl_l(c uint32, l *Struct___locale_struct) int32 {
	return iswcntrl(c)
}
