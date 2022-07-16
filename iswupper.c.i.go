package libc

func iswupper(wc uint32) int32 {
	return func() int32 {
		if towlower(wc) != wc {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswupper_l(c uint32, l *struct___locale_struct) int32 {
	return iswupper(c)
}
