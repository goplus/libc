package libc

func iscntrl(c int32) int32 {
	return func() int32 {
		if uint32(c) < uint32(32) || c == int32(127) {
			return 1
		} else {
			return 0
		}
	}()
}
func __iscntrl_l(c int32, l *struct___locale_struct) int32 {
	return iscntrl(c)
}
