package libc

import unsafe "unsafe"

func _cgos_twobyte_strstr_strstr(h *uint8, n *uint8) *int8 {
	var nw uint16 = uint16(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))))<<int32(8) | int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1))))))
	var hw uint16 = uint16(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0)))))<<int32(8) | int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1))))))
	for *(*uintptr)(unsafe.Pointer(&h))++; int32(*h) != 0 && int32(hw) != int32(nw); hw = uint16(int32(hw)<<int32(8) | int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &h
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return *_cgo_addr
	}())) {
	}
	return func() *int8 {
		if int32(*h) != 0 {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(1))))
		} else {
			return nil
		}
	}()
}
func _cgos_threebyte_strstr_strstr(h *uint8, n *uint8) *int8 {
	var nw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(2)))))<<int32(8))
	var hw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(2)))))<<int32(8))
	for *(*uintptr)(unsafe.Pointer(&h)) += uintptr(int32(2)); int32(*h) != 0 && hw != nw; hw = (hw | uint32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &h
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return *_cgo_addr
	}())) << int32(8) {
	}
	return func() *int8 {
		if int32(*h) != 0 {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(2))))
		} else {
			return nil
		}
	}()
}
func _cgos_fourbyte_strstr_strstr(h *uint8, n *uint8) *int8 {
	var nw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(2)))))<<int32(8)) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(3)))))
	var hw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(2)))))<<int32(8)) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(3)))))
	for *(*uintptr)(unsafe.Pointer(&h)) += uintptr(int32(3)); int32(*h) != 0 && hw != nw; hw = hw<<int32(8) | uint32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &h
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return *_cgo_addr
	}()) {
	}
	return func() *int8 {
		if int32(*h) != 0 {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(3))))
		} else {
			return nil
		}
	}()
}
func _cgos_twoway_strstr_strstr(h *uint8, n *uint8) *int8 {
	var z *uint8
	var l uint64
	var ip uint64
	var jp uint64
	var k uint64
	var p uint64
	var ms uint64
	var p0 uint64
	var mem uint64
	var mem0 uint64
	var byteset [4]uint64 = [4]uint64{uint64(0)}
	var shift [256]uint64
	for l = uint64(0); int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l)))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(l)))) != 0; l++ {
		func() uint64 {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&byteset)))) + uintptr(uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l))))/64)*8)) |= uint64(1) << (uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l)))) % 64)
			return func() (_cgo_ret uint64) {
				_cgo_addr := &*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&shift)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l))))*8))
				*_cgo_addr = l + uint64(1)
				return *_cgo_addr
			}()
		}()
	}
	if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l))) != 0 {
		return (*int8)(nil)
	}
	ip = uint64(18446744073709551615)
	jp = uint64(0)
	k = func() (_cgo_ret uint64) {
		_cgo_addr := &p
		*_cgo_addr = uint64(1)
		return *_cgo_addr
	}()
	for jp+k < l {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)))) {
			if k == p {
				jp += p
				k = uint64(1)
			} else {
				k++
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)))) > int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)))) {
			jp += k
			k = uint64(1)
			p = jp - ip
		} else {
			ip = func() (_cgo_ret uint64) {
				_cgo_addr := &jp
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
			k = func() (_cgo_ret uint64) {
				_cgo_addr := &p
				*_cgo_addr = uint64(1)
				return *_cgo_addr
			}()
		}
	}
	ms = ip
	p0 = p
	ip = uint64(18446744073709551615)
	jp = uint64(0)
	k = func() (_cgo_ret uint64) {
		_cgo_addr := &p
		*_cgo_addr = uint64(1)
		return *_cgo_addr
	}()
	for jp+k < l {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)))) {
			if k == p {
				jp += p
				k = uint64(1)
			} else {
				k++
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)))) < int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)))) {
			jp += k
			k = uint64(1)
			p = jp - ip
		} else {
			ip = func() (_cgo_ret uint64) {
				_cgo_addr := &jp
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
			k = func() (_cgo_ret uint64) {
				_cgo_addr := &p
				*_cgo_addr = uint64(1)
				return *_cgo_addr
			}()
		}
	}
	if ip+uint64(1) > ms+uint64(1) {
		ms = ip
	} else {
		p = p0
	}
	if Memcmp(unsafe.Pointer(n), unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n))+uintptr(p)))), ms+uint64(1)) != 0 {
		mem0 = uint64(0)
		p = func() uint64 {
			if ms > l-ms-uint64(1) {
				return ms
			} else {
				return l - ms - uint64(1)
			}
		}() + uint64(1)
	} else {
		mem0 = l - p
	}
	mem = uint64(0)
	z = h
	for {
		if uint64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(h))) < l {
			var grow uint64 = l | uint64(63)
			var z2 *uint8 = (*uint8)(Memchr(unsafe.Pointer(z), int32(0), grow))
			if z2 != nil {
				z = z2
				if uint64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(h))) < l {
					return (*int8)(nil)
				}
			} else {
				*(*uintptr)(unsafe.Pointer(&z)) += uintptr(grow)
			}
		}
		if *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&byteset)))) + uintptr(uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(l-uint64(1)))))/64)*8))&(uint64(1)<<(uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(l-uint64(1)))))%64)) != 0 {
			k = l - *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&shift)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(l-uint64(1)))))*8))
			if k != 0 {
				if k < mem {
					k = mem
				}
				*(*uintptr)(unsafe.Pointer(&h)) += uintptr(k)
				mem = uint64(0)
				continue
			}
		} else {
			*(*uintptr)(unsafe.Pointer(&h)) += uintptr(l)
			mem = uint64(0)
			continue
		}
		for k = func() uint64 {
			if ms+uint64(1) > mem {
				return ms + uint64(1)
			} else {
				return mem
			}
		}(); int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k)))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(k)))); k++ {
		}
		if *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k))) != 0 {
			*(*uintptr)(unsafe.Pointer(&h)) += uintptr(k - ms)
			mem = uint64(0)
			continue
		}
		for k = ms + uint64(1); k > mem && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k-uint64(1))))) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(k-uint64(1))))); k-- {
		}
		if k <= mem {
			return (*int8)(unsafe.Pointer(h))
		}
		*(*uintptr)(unsafe.Pointer(&h)) += uintptr(p)
		mem = mem0
	}
	return nil
}
func Strstr(h *int8, n *int8) *int8 {
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))) != 0) {
		return (*int8)(unsafe.Pointer(h))
	}
	h = Strchr(h, int32(*n))
	if !(h != nil) || !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1)))) != 0) {
		return (*int8)(unsafe.Pointer(h))
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1)))) != 0) {
		return (*int8)(nil)
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(2)))) != 0) {
		return _cgos_twobyte_strstr_strstr((*uint8)(unsafe.Pointer(h)), (*uint8)(unsafe.Pointer(n)))
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(2)))) != 0) {
		return (*int8)(nil)
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(3)))) != 0) {
		return _cgos_threebyte_strstr_strstr((*uint8)(unsafe.Pointer(h)), (*uint8)(unsafe.Pointer(n)))
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(3)))) != 0) {
		return (*int8)(nil)
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(4)))) != 0) {
		return _cgos_fourbyte_strstr_strstr((*uint8)(unsafe.Pointer(h)), (*uint8)(unsafe.Pointer(n)))
	}
	return _cgos_twoway_strstr_strstr((*uint8)(unsafe.Pointer(h)), (*uint8)(unsafe.Pointer(n)))
}
