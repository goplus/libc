package libc

import unsafe "unsafe"

func __crypt_r(key *int8, salt *int8, data *struct_crypt_data) *int8 {
	var output *int8 = (*int8)(unsafe.Pointer(data))
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(0))))) == '$' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(1))))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(2))))) != 0 {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(1))))) == '1' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(2))))) == '$' {
			return __crypt_md5(key, salt, output)
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(1))))) == '2' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(3))))) == '$' {
			return __crypt_blowfish(key, salt, output)
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(1))))) == '5' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(2))))) == '$' {
			return __crypt_sha256(key, salt, output)
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(1))))) == '6' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(int32(2))))) == '$' {
			return __crypt_sha512(key, salt, output)
		}
	}
	return __crypt_des(key, salt, output)
}
