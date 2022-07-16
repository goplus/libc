package libc

func iswlower(wc uint32) int32 {
	return func() int32 {
		if towupper(wc) != wc {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswlower_l(c uint32, l *struct___locale_struct) int32 {
	return iswlower(c)
}
