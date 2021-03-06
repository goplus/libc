package libc

import unsafe "unsafe"

func __fwritex(s *uint8, l uint64, f *Struct__IO_FILE) uint64 {
	var i uint64 = uint64(0)
	if !(f.Wend != nil) && __towrite(f) != 0 {
		return uint64(0)
	}
	if l > uint64(uintptr(unsafe.Pointer(f.Wend))-uintptr(unsafe.Pointer(f.Wpos))) {
		return f.Write(f, s, l)
	}
	if f.Lbf >= int32(0) {
		for i = l; i != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i-uint64(1))))) != '\n'; i-- {
		}
		if i != 0 {
			var n uint64 = f.Write(f, s, i)
			if n < i {
				return n
			}
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(i)
			l -= i
		}
	}
	Memcpy(unsafe.Pointer(f.Wpos), unsafe.Pointer(s), l)
	*(*uintptr)(unsafe.Pointer(&f.Wpos)) += uintptr(l)
	return l + i
}
func Fwrite(src unsafe.Pointer, size uint64, nmemb uint64, f *Struct__IO_FILE) uint64 {
	var k uint64
	var l uint64 = size * nmemb
	if !(size != 0) {
		nmemb = uint64(0)
	}
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	k = __fwritex((*uint8)(src), l, f)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return func() uint64 {
		if k == l {
			return nmemb
		} else {
			return k / size
		}
	}()
}
