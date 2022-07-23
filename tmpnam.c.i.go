package libc

import unsafe "unsafe"

func Tmpnam(buf *int8) *int8 {
	var s [19]int8 = [19]int8{'/', 't', 'm', 'p', '/', 't', 'm', 'p', 'n', 'a', 'm', '_', 'X', 'X', 'X', 'X', 'X', 'X', '\x00'}
	var try int32
	var r int32
	for try = int32(0); try < int32(100); try++ {
		__randname((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s)))) + uintptr(int32(12)))))
		r = int32(__syscall2(int64(340), int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))))), int64(uintptr(unsafe.Pointer(&struct_kstat{uint64(0), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})))))
		if r == -2 {
			return Strcpy(func() *int8 {
				if buf != nil {
					return buf
				} else {
					return (*int8)(unsafe.Pointer(&_cgos_internal_tmpnam))
				}
			}(), (*int8)(unsafe.Pointer(&s)))
		}
	}
	return (*int8)(nil)
}

var _cgos_internal_tmpnam [20]int8
