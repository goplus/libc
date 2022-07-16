package libc

func isblank(c int32) int32 {
	return func() int32 {
		if c == ' ' || c == '\t' {
			return 1
		} else {
			return 0
		}
	}()
}
func __isblank_l(c int32, l *struct___locale_struct) int32 {
	return isblank(c)
}
