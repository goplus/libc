package libc

import unsafe "unsafe"

func Ptsname(fd int32) *int8 {
	var err int32 = __ptsname_r(fd, (*int8)(unsafe.Pointer(&_cgos_buf_ptsname)), 22)
	if err != 0 {
		*X__errno_location() = err
		return (*int8)(nil)
	}
	return (*int8)(unsafe.Pointer(&_cgos_buf_ptsname))
}

var _cgos_buf_ptsname [22]int8
