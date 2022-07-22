package libc

import unsafe "unsafe"

func Tempnam(dir *int8, pfx *int8) *int8 {
	var s [4096]int8
	var l uint64
	var dl uint64
	var pl uint64
	var try int32
	var r int32
	if !(dir != nil) {
		dir = (*int8)(unsafe.Pointer(&[5]int8{'/', 't', 'm', 'p', '\x00'}))
	}
	if !(pfx != nil) {
		pfx = (*int8)(unsafe.Pointer(&[5]int8{'t', 'e', 'm', 'p', '\x00'}))
	}
	dl = Strlen(dir)
	pl = Strlen(pfx)
	l = dl + uint64(1) + pl + uint64(1) + uint64(6)
	if l >= uint64(4096) {
		*__errno_location() = int32(36)
		return (*int8)(nil)
	}
	Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), unsafe.Pointer(dir), dl)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s)))) + uintptr(dl))) = int8('/')
	Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))))+uintptr(dl)))))+uintptr(int32(1))))), unsafe.Pointer(pfx), pl)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s)))) + uintptr(dl+uint64(1)+pl))) = int8('_')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s)))) + uintptr(l))) = int8(0)
	for try = int32(0); try < int32(100); try++ {
		__randname((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))))+uintptr(l))))) - uintptr(int32(6)))))
		r = int32(__syscall2(int64(340), int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))))), int64(uintptr(unsafe.Pointer(&struct_kstat{uint64(0), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})))))
		if r == -2 {
			return Strdup((*int8)(unsafe.Pointer(&s)))
		}
	}
	return (*int8)(nil)
}
