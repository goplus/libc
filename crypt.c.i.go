package libc

import unsafe "unsafe"

func crypt(key *int8, salt *int8) *int8 {
	return __crypt_r(key, salt, (*struct_crypt_data)(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_buf_crypt)))))
}

var _cgos_buf_crypt [128]int8
