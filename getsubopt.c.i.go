package libc

import unsafe "unsafe"

func Getsubopt(opt **int8, keys **int8, val **int8) int32 {
	var s *int8 = *opt
	var i int32
	*val = (*int8)(nil)
	*opt = Strchr(s, ',')
	if *opt != nil {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &*opt
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(0)
	} else {
		*opt = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Strlen(s))))
	}
	for i = int32(0); *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(i)*8)) != nil; i++ {
		var l uint64 = Strlen(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(i)*8)))
		if Strncmp(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(i)*8)), s, l) != 0 {
			continue
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(l)))) == '=' {
			*val = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(l))))) + uintptr(int32(1))))
		} else if *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(l))) != 0 {
			continue
		}
		return i
	}
	return -1
}
