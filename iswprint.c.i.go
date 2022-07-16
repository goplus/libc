package libc

func iswprint(wc uint32) int32 {
	if wc < uint32(255) {
		return func() int32 {
			if (wc+uint32(1))&uint32(127) >= uint32(33) {
				return 1
			} else {
				return 0
			}
		}()
	}
	if wc < uint32(8232) || wc-uint32(8234) < uint32(47062) || wc-uint32(57344) < uint32(8185) {
		return int32(1)
	}
	if wc-uint32(65532) > uint32(1048579) || wc&uint32(65534) == uint32(65534) {
		return int32(0)
	}
	return int32(1)
}
func __iswprint_l(c uint32, l *struct___locale_struct) int32 {
	return iswprint(c)
}
