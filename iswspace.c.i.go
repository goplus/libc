package libc

import unsafe "unsafe"

func iswspace(wc uint32) int32 {
	return func() int32 {
		if wc != 0 && wcschr((*uint32)(unsafe.Pointer(&_cgos_spaces_iswspace)), wc) != nil {
			return 1
		} else {
			return 0
		}
	}()
}

var _cgos_spaces_iswspace [22]uint32 = [22]uint32{uint32(' '), uint32('\t'), uint32('\n'), uint32('\r'), uint32(11), uint32(12), uint32(133), uint32(8192), uint32(8193), uint32(8194), uint32(8195), uint32(8196), uint32(8197), uint32(8198), uint32(8200), uint32(8201), uint32(8202), uint32(8232), uint32(8233), uint32(8287), uint32(12288), uint32(0)}

func __iswspace_l(c uint32, l *struct___locale_struct) int32 {
	return iswspace(c)
}
