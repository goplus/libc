package libc

import unsafe "unsafe"

func iswctype(wc uint32, type_ uint64) int32 {
	switch type_ {
	case uint64(1):
		return iswalnum(wc)
	case uint64(2):
		return iswalpha(wc)
	case uint64(3):
		return iswblank(wc)
	case uint64(4):
		return iswcntrl(wc)
	case uint64(5):
		return func() int32 {
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
		}()
	case uint64(6):
		return iswgraph(wc)
	case uint64(7):
		return iswlower(wc)
	case uint64(8):
		return iswprint(wc)
	case uint64(9):
		return iswpunct(wc)
	case uint64(10):
		return iswspace(wc)
	case uint64(11):
		return iswupper(wc)
	case uint64(12):
		return iswxdigit(wc)
	}
	return int32(0)
}
func wctype(s *int8) uint64 {
	var i int32
	var p *int8
	for func() *int8 {
		i = int32(1)
		return func() (_cgo_ret *int8) {
			_cgo_addr := &p
			*_cgo_addr = (*int8)(unsafe.Pointer(&names_cgo15_iswctype))
			return *_cgo_addr
		}()
	}(); *p != 0; func() *int8 {
		i++
		return func() (_cgo_ret *int8) {
			_cgo_addr := &p
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(6))
			return *_cgo_addr
		}()
	}() {
		if int32(*s) == int32(*p) && !(Strcmp(s, p) != 0) {
			return uint64(i)
		}
	}
	return uint64(0)
}

var names_cgo15_iswctype [73]int8 = [73]int8{'a', 'l', 'n', 'u', 'm', '\x00', 'a', 'l', 'p', 'h', 'a', '\x00', 'b', 'l', 'a', 'n', 'k', '\x00', 'c', 'n', 't', 'r', 'l', '\x00', 'd', 'i', 'g', 'i', 't', '\x00', 'g', 'r', 'a', 'p', 'h', '\x00', 'l', 'o', 'w', 'e', 'r', '\x00', 'p', 'r', 'i', 'n', 't', '\x00', 'p', 'u', 'n', 'c', 't', '\x00', 's', 'p', 'a', 'c', 'e', '\x00', 'u', 'p', 'p', 'e', 'r', '\x00', 'x', 'd', 'i', 'g', 'i', 't', '\x00'}

func __iswctype_l(c uint32, t uint64, l *struct___locale_struct) int32 {
	return iswctype(c, t)
}
func __wctype_l(s *int8, l *struct___locale_struct) uint64 {
	return wctype(s)
}
