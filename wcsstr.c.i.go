package libc

import unsafe "unsafe"

func twoway_wcsstr_cgo15_wcsstr(h *uint32, n *uint32) *uint32 {
	var z *uint32
	var l uint64
	var ip uint64
	var jp uint64
	var k uint64
	var p uint64
	var ms uint64
	var p0 uint64
	var mem uint64
	var mem0 uint64
	for l = uint64(0); *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l)*4)) != 0 && *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(l)*4)) != 0; l++ {
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(l)*4)) != 0 {
		return (*uint32)(nil)
	}
	ip = uint64(18446744073709551615)
	jp = uint64(0)
	k = func() (_cgo_ret uint64) {
		_cgo_addr := &p
		*_cgo_addr = uint64(1)
		return *_cgo_addr
	}()
	for jp+k < l {
		if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)*4)) == *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)*4)) {
			if k == p {
				jp += p
				k = uint64(1)
			} else {
				k++
			}
		} else if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)*4)) > *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)*4)) {
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
		if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)*4)) == *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)*4)) {
			if k == p {
				jp += p
				k = uint64(1)
			} else {
				k++
			}
		} else if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(ip+k)*4)) < *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(jp+k)*4)) {
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
	if wmemcmp(n, (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n))+uintptr(p)*4)), ms+uint64(1)) != 0 {
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
		if uint64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(h)))*4) < l {
			var grow uint64 = l | uint64(63)
			var z2 *uint32 = wmemchr(z, uint32(0), grow)
			if z2 != nil {
				z = z2
				if uint64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(h)))*4) < l {
					return (*uint32)(nil)
				}
			} else {
				*(*uintptr)(unsafe.Pointer(&z)) += uintptr(grow) * 4
			}
		}
		for k = func() uint64 {
			if ms+uint64(1) > mem {
				return ms + uint64(1)
			} else {
				return mem
			}
		}(); *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k)*4)) != 0 && *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k)*4)) == *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(k)*4)); k++ {
		}
		if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k)*4)) != 0 {
			*(*uintptr)(unsafe.Pointer(&h)) += uintptr(k-ms) * 4
			mem = uint64(0)
			continue
		}
		for k = ms + uint64(1); k > mem && *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(k-uint64(1))*4)) == *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(k-uint64(1))*4)); k-- {
		}
		if k <= mem {
			return (*uint32)(unsafe.Pointer(h))
		}
		*(*uintptr)(unsafe.Pointer(&h)) += uintptr(p) * 4
		mem = mem0
	}
	return nil
}
func wcsstr(h *uint32, n *uint32) *uint32 {
	if !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(0))*4)) != 0) {
		return (*uint32)(unsafe.Pointer(h))
	}
	if !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(0))*4)) != 0) {
		return (*uint32)(nil)
	}
	h = wcschr(h, *n)
	if !(h != nil) || !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + uintptr(int32(1))*4)) != 0) {
		return (*uint32)(unsafe.Pointer(h))
	}
	if !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(int32(1))*4)) != 0) {
		return (*uint32)(nil)
	}
	return twoway_wcsstr_cgo15_wcsstr(h, n)
}
