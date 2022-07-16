package libc

func islower(c int32) int32 {
	return func() int32 {
		if uint32(c)-uint32('a') < uint32(26) {
			return 1
		} else {
			return 0
		}
	}()
}
func __islower_l(c int32, l *struct___locale_struct) int32 {
	return islower(c)
}
