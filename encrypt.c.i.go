package libc

import unsafe "unsafe"

var _cgos___encrypt_key_encrypt struct_expanded_key

func Setkey(key *int8) {
	var bkey [8]uint8
	var i int32
	var j int32
	for i = int32(0); i < int32(8); i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&bkey)))) + uintptr(i))) = uint8(0)
		for j = int32(7); j >= int32(0); func() *int8 {
			j--
			return func() (_cgo_ret *int8) {
				_cgo_addr := &key
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}() {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&bkey)))) + uintptr(i))) |= uint8(uint32(int32(*key)&int32(1)) << j)
		}
	}
	__des_setkey((*uint8)(unsafe.Pointer(&bkey)), &_cgos___encrypt_key_encrypt)
}
func encrypt(block *int8, edflag int32) {
	var decrypt_key struct_expanded_key
	var key *struct_expanded_key
	var b [2]uint32
	var i int32
	var j int32
	var p *int8
	p = block
	for i = int32(0); i < int32(2); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&b)))) + uintptr(i)*4)) = uint32(0)
		for j = int32(31); j >= int32(0); func() *int8 {
			j--
			return func() (_cgo_ret *int8) {
				_cgo_addr := &p
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}() {
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&b)))) + uintptr(i)*4)) |= uint32(int32(*p)&int32(1)) << j
		}
	}
	key = &_cgos___encrypt_key_encrypt
	if edflag != 0 {
		key = &decrypt_key
		for i = int32(0); i < int32(16); i++ {
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&decrypt_key.l)))) + uintptr(i)*4)) = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos___encrypt_key_encrypt.l)))) + uintptr(int32(15)-i)*4))
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&decrypt_key.r)))) + uintptr(i)*4)) = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos___encrypt_key_encrypt.r)))) + uintptr(int32(15)-i)*4))
		}
	}
	__do_des(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&b)))) + uintptr(int32(0))*4)), *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&b)))) + uintptr(int32(1))*4)), (*uint32)(unsafe.Pointer(&b)), (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&b))))+uintptr(int32(1))*4)), uint32(1), uint32(0), key)
	p = block
	for i = int32(0); i < int32(2); i++ {
		for j = int32(31); j >= int32(0); j-- {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &p
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = int8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&b)))) + uintptr(i)*4)) >> j & uint32(1))
		}
	}
}
