package libc

func iswgraph(wc uint32) int32 {
	return func() int32 {
		if !(iswspace(wc) != 0) && iswprint(wc) != 0 {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswgraph_l(c uint32, l *struct___locale_struct) int32 {
	return iswgraph(c)
}
