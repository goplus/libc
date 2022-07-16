package libc

func ispunct(c int32) int32 {
	return func() int32 {
		if func() int32 {
			if int32(0) != 0 {
				return isgraph(c)
			} else {
				return func() int32 {
					if uint32(c)-uint32(33) < uint32(94) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 && !(isalnum(c) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __ispunct_l(c int32, l *struct___locale_struct) int32 {
	return ispunct(c)
}
