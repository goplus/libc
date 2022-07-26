package libc

func iswalnum(wc uint32) int32 {
	return func() int32 {
		if func() int32 {
			if int32(0) != 0 {
				return iswdigit(wc)
			} else {
				return func() int32 {
					if uint32(wc)-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 || iswalpha(wc) != 0 {
			return 1
		} else {
			return 0
		}
	}()
}
func __iswalnum_l(c uint32, l *Struct___locale_struct) int32 {
	return iswalnum(c)
}
