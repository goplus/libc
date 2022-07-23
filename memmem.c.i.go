package libc

import unsafe "unsafe"

func _cgos_twobyte_memmem__memmem(h *uint8, k uint64, n *uint8) *int8 {
	var nw uint16 = uint16(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))))<<int32(8) | int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1))))))
	var hw uint16 = uint16(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0)))))<<int32(8) | int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1))))))
	for func() uint64 {
		*(*uintptr)(unsafe.Pointer(&h)) += uintptr(int32(2))
		return func() (_cgo_ret uint64) {
			_cgo_addr := &k
			*_cgo_addr -= uint64(2)
			return *_cgo_addr
		}()
	}(); k != 0; func() uint16 {
		k--
		return func() (_cgo_ret uint16) {
			_cgo_addr := &hw
			*_cgo_addr = uint16(int32(hw)<<int32(8) | int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &h
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()))
			return *_cgo_addr
		}()
	}() {
		if int32(hw) == int32(nw) {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(2))))
		}
	}
	return func() *int8 {
		if int32(hw) == int32(nw) {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(2))))
		} else {
			return nil
		}
	}()
}
func _cgos_threebyte_memmem__memmem(h *uint8, k uint64, n *uint8) *int8 {
	var nw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(2)))))<<int32(8))
	var hw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(2)))))<<int32(8))
	for func() uint64 {
		*(*uintptr)(unsafe.Pointer(&h)) += uintptr(int32(3))
		return func() (_cgo_ret uint64) {
			_cgo_addr := &k
			*_cgo_addr -= uint64(3)
			return *_cgo_addr
		}()
	}(); k != 0; func() uint32 {
		k--
		return func() (_cgo_ret uint32) {
			_cgo_addr := &hw
			*_cgo_addr = (hw | uint32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &h
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())) << int32(8)
			return *_cgo_addr
		}()
	}() {
		if hw == nw {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(3))))
		}
	}
	return func() *int8 {
		if hw == nw {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(3))))
		} else {
			return nil
		}
	}()
}
func _cgos_fourbyte_memmem__memmem(h *uint8, k uint64, n *uint8) *int8 {
	var nw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(2)))))<<int32(8)) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(3)))))
	var hw uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0)))))<<int32(24) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1)))))<<int32(16)) | uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(2)))))<<int32(8)) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(3)))))
	for func() uint64 {
		*(*uintptr)(unsafe.Pointer(&h)) += uintptr(int32(4))
		return func() (_cgo_ret uint64) {
			_cgo_addr := &k
			*_cgo_addr -= uint64(4)
			return *_cgo_addr
		}()
	}(); k != 0; func() uint32 {
		k--
		return func() (_cgo_ret uint32) {
			_cgo_addr := &hw
			*_cgo_addr = hw<<int32(8) | uint32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &h
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
			return *_cgo_addr
		}()
	}() {
		if hw == nw {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(4))))
		}
	}
	return func() *int8 {
		if hw == nw {
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(h)))) - uintptr(int32(4))))
		} else {
			return nil
		}
	}()
}
func _cgos_twoway_memmem__memmem(h *uint8, z *uint8, n *uint8, l uint64) *int8 {
	var i uint64
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
	for i = uint64(0); i < l; i++ {
		func() uint64 {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&byteset)))) + uintptr(uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(i))))/64)*8)) |= uint64(1) << (uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(i)))) % 64)
			return func() (_cgo_ret uint64) {
				_cgo_addr := &*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&shift)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(i))))*8))
				*_cgo_addr = i + uint64(1)
				return *_cgo_addr
			}()
		}()
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
	for {
		if uint64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(h))) < l {
			return (*int8)(nil)
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
		}(); k < l && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(k)))); k++ {
		}
		if k < l {
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
func Memmem(h0 unsafe.Pointer, k uint64, n0 unsafe.Pointer, l uint64) unsafe.Pointer {
	var h *uint8 = (*uint8)(h0)
	var n *uint8 = (*uint8)(n0)
	if !(l != 0) {
		return unsafe.Pointer(h)
	}
	if k < l {
		return unsafe.Pointer(nil)
	}
	h = (*uint8)(Memchr(h0, int32(*n), k))
	if !(h != nil) || l == uint64(1) {
		return unsafe.Pointer(h)
	}
	k -= uint64(uintptr(unsafe.Pointer(h)) - uintptr(unsafe.Pointer((*uint8)(h0))))
	if k < l {
		return unsafe.Pointer(nil)
	}
	if l == uint64(2) {
		return unsafe.Pointer(_cgos_twobyte_memmem__memmem(h, k, n))
	}
	if l == uint64(3) {
		return unsafe.Pointer(_cgos_threebyte_memmem__memmem(h, k, n))
	}
	if l == uint64(4) {
		return unsafe.Pointer(_cgos_fourbyte_memmem__memmem(h, k, n))
	}
	return unsafe.Pointer(_cgos_twoway_memmem__memmem(h, (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(h))+uintptr(k))), n, l))
}
