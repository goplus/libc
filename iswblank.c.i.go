package libc

func iswblank(wc uint32) int32 {
	return Isblank(int32(wc))
}
func __iswblank_l(c uint32, l *Struct___locale_struct) int32 {
	return iswblank(c)
}
