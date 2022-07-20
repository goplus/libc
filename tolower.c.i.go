package libc

func Tolower(c int32) int32 {
	if func() int32 {
		if int32(0) != 0 {
			return Isupper(c)
		} else {
			return func() int32 {
				if uint32(c)-uint32('A') < uint32(26) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 {
		return c | int32(32)
	}
	return c
}
func __tolower_l(c int32, l *struct___locale_struct) int32 {
	return Tolower(c)
}
