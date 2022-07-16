package libc

func isalnum(c int32) int32 {
	return func() int32 {
		if func() int32 {
			if int32(0) != 0 {
				return isalpha(c)
			} else {
				return func() int32 {
					if uint32(c)|uint32(32)-uint32('a') < uint32(26) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 || func() int32 {
			if int32(0) != 0 {
				return Isdigit(c)
			} else {
				return func() int32 {
					if uint32(c)-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 {
			return 1
		} else {
			return 0
		}
	}()
}
func __isalnum_l(c int32, l *struct___locale_struct) int32 {
	return isalnum(c)
}
