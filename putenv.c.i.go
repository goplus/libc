package libc

import unsafe "unsafe"

func _cgos_dummy_putenv(old *int8, new *int8) {
}
func __putenv(s *int8, l uint64, r *int8) int32 {
	var i uint64 = uint64(0)
	if X__environ != nil {
		for e := (**int8)(X__environ); *e != nil; func() uint64 {
			*(*uintptr)(unsafe.Pointer(&e)) += 8
			return func() (_cgo_ret uint64) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
		}() {
			if !(Strncmp(s, *e, l+uint64(1)) != 0) {
				var tmp *int8 = *e
				*e = s
				__env_rm_add(tmp, r)
				return int32(0)
			}
		}
	}
	var newenv **int8
	if uintptr(unsafe.Pointer(X__environ)) == uintptr(unsafe.Pointer(_cgos_oldenv_putenv)) {
		newenv = (**int8)(Realloc(unsafe.Pointer(_cgos_oldenv_putenv), 8*(i+uint64(2))))
		if !(newenv != nil) {
			goto oom
		}
	} else {
		newenv = (**int8)(Malloc(8 * (i + uint64(2))))
		if !(newenv != nil) {
			goto oom
		}
		if i != 0 {
			Memcpy(unsafe.Pointer(newenv), unsafe.Pointer(X__environ), 8*i)
		}
		Free(unsafe.Pointer(_cgos_oldenv_putenv))
	}
	*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(newenv)) + uintptr(i)*8)) = s
	*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(newenv)) + uintptr(i+uint64(1))*8)) = (*int8)(nil)
	X__environ = func() (_cgo_ret **int8) {
		_cgo_addr := &_cgos_oldenv_putenv
		*_cgo_addr = newenv
		return *_cgo_addr
	}()
	if r != nil {
		__env_rm_add(nil, r)
	}
	return int32(0)
oom:
	Free(unsafe.Pointer(r))
	return -1
}

var _cgos_oldenv_putenv **int8

func Putenv(s *int8) int32 {
	var l uint64 = uint64(uintptr(unsafe.Pointer(__strchrnul(s, '='))) - uintptr(unsafe.Pointer(s)))
	if !(l != 0) || !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(l))) != 0) {
		return Unsetenv(s)
	}
	return __putenv(s, l, nil)
}
