package libc

import unsafe "unsafe"

func _cgos_swapc____mo_lookup(x uint32, c int32) uint32 {
	return func() uint32 {
		if c != 0 {
			return x>>int32(24) | x>>int32(8)&uint32(65280) | x<<int32(8)&uint32(16711680) | x<<int32(24)
		} else {
			return x
		}
	}()
}
func __mo_lookup(p unsafe.Pointer, size uint64, s *int8) *int8 {
	var mo *uint32 = (*uint32)(p)
	var sw int32 = int32(*mo - uint32(2500072158))
	var b uint32 = uint32(0)
	var n uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(int32(2))*4)), sw)
	var o uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(int32(3))*4)), sw)
	var t uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(int32(4))*4)), sw)
	if uint64(n) >= size/uint64(4) || uint64(o) >= size-uint64(uint32(4)*n) || uint64(t) >= size-uint64(uint32(4)*n) || (o|t)%uint32(4) != 0 {
		return (*int8)(nil)
	}
	o /= uint32(4)
	t /= uint32(4)
	for {
		var ol uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(o+uint32(2)*(b+n/uint32(2)))*4)), sw)
		var os uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(o+uint32(2)*(b+n/uint32(2))+uint32(1))*4)), sw)
		if uint64(os) >= size || uint64(ol) >= size-uint64(os) || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(p))) + uintptr(os+ol)))) != 0 {
			return (*int8)(nil)
		}
		var sign int32 = Strcmp(s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(p)))+uintptr(os))))
		if !(sign != 0) {
			var tl uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(t+uint32(2)*(b+n/uint32(2)))*4)), sw)
			var ts uint32 = _cgos_swapc____mo_lookup(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(mo)) + uintptr(t+uint32(2)*(b+n/uint32(2))+uint32(1))*4)), sw)
			if uint64(ts) >= size || uint64(tl) >= size-uint64(ts) || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(p))) + uintptr(ts+tl)))) != 0 {
				return (*int8)(nil)
			}
			return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(p))) + uintptr(ts)))
		} else if n == uint32(1) {
			return (*int8)(nil)
		} else if sign < int32(0) {
			n /= uint32(2)
		} else {
			b += n / uint32(2)
			n -= n / uint32(2)
		}
	}
	return (*int8)(nil)
}
