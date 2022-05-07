package libc

import unsafe "unsafe"

func Getenv(name *int8) *int8 {
	var l uint64 = uint64(uintptr(unsafe.Pointer(__strchrnul(name, '='))) - uintptr(unsafe.Pointer(name)))
	if l != 0 && !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(name)) + uintptr(l))) != 0) && __environ != nil {
		for e := (**int8)(__environ); *e != nil; *(*uintptr)(unsafe.Pointer(&e)) += 8 {
			if !(Strncmp(name, *e, l) != 0) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*e)) + uintptr(l)))) == '=' {
				return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*e))+uintptr(l))))) + uintptr(1)))
			}
		}
	}
	return (*int8)(nil)
}
