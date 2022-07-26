package libc

import unsafe "unsafe"

func Fread(destv unsafe.Pointer, size uint64, nmemb uint64, f *Struct__IO_FILE) uint64 {
	var dest *uint8 = (*uint8)(destv)
	var len uint64 = size * nmemb
	var l uint64 = len
	var k uint64
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
	f.Mode |= f.Mode - int32(1)
	if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
		k = func() uint64 {
			if uint64(uintptr(unsafe.Pointer(f.Rend))-uintptr(unsafe.Pointer(f.Rpos))) < l {
				return uint64(uintptr(unsafe.Pointer(f.Rend)) - uintptr(unsafe.Pointer(f.Rpos)))
			} else {
				return l
			}
		}()
		Memcpy(unsafe.Pointer(dest), unsafe.Pointer(f.Rpos), k)
		*(*uintptr)(unsafe.Pointer(&f.Rpos)) += uintptr(k)
		*(*uintptr)(unsafe.Pointer(&dest)) += uintptr(k)
		l -= k
	}
	for ; l != 0; func() *uint8 {
		l -= k
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &dest
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(k)
			return *_cgo_addr
		}()
	}() {
		k = func() uint64 {
			if __toread(f) != 0 {
				return uint64(0)
			} else {
				return f.Read(f, dest, l)
			}
		}()
		if !(k != 0) {
			for {
				if __need_unlock != 0 {
					__unlockfile(f)
				}
				if true {
					break
				}
			}
			return (len - l) / size
		}
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return nmemb
}
