package libc

import unsafe "unsafe"

func __env_rm_add(old *int8, new *int8) {
	for i := uint64(uint64(0)); i < _cgos_env_alloced_n_setenv; i++ {
		if uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_env_alloced_setenv)) + uintptr(i)*8)))) == uintptr(unsafe.Pointer(old)) {
			*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_env_alloced_setenv)) + uintptr(i)*8)) = new
			Free(unsafe.Pointer(old))
			return
		} else if !(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_env_alloced_setenv)) + uintptr(i)*8)) != nil) && new != nil {
			*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_env_alloced_setenv)) + uintptr(i)*8)) = new
			new = (*int8)(nil)
		}
	}
	if !(new != nil) {
		return
	}
	var t **int8 = (**int8)(Realloc(unsafe.Pointer(_cgos_env_alloced_setenv), 8*(_cgos_env_alloced_n_setenv+uint64(1))))
	if !(t != nil) {
		return
	}
	*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(func() (_cgo_ret **int8) {
		_cgo_addr := &_cgos_env_alloced_setenv
		*_cgo_addr = t
		return *_cgo_addr
	}())) + uintptr(func() (_cgo_ret uint64) {
		_cgo_addr := &_cgos_env_alloced_n_setenv
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}())*8)) = new
}

var _cgos_env_alloced_setenv **int8
var _cgos_env_alloced_n_setenv uint64

func Setenv(var_ *int8, value *int8, overwrite int32) int32 {
	var s *int8
	var l1 uint64
	var l2 uint64
	if !(var_ != nil) || !(func() (_cgo_ret uint64) {
		_cgo_addr := &l1
		*_cgo_addr = uint64(uintptr(unsafe.Pointer(__strchrnul(var_, '='))) - uintptr(unsafe.Pointer(var_)))
		return *_cgo_addr
	}() != 0) || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(var_)) + uintptr(l1)))) != 0 {
		*__errno_location() = int32(22)
		return -1
	}
	if !(overwrite != 0) && Getenv(var_) != nil {
		return int32(0)
	}
	l2 = Strlen(value)
	s = (*int8)(Malloc(l1 + l2 + uint64(2)))
	if !(s != nil) {
		return -1
	}
	Memcpy(unsafe.Pointer(s), unsafe.Pointer(var_), l1)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(l1))) = int8('=')
	Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(l1)))))+uintptr(int32(1))))), unsafe.Pointer(value), l2+uint64(1))
	return __putenv(s, l1, s)
}
