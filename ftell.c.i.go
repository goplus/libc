package libc

import unsafe "unsafe"

func __ftello_unlocked(f *Struct__IO_FILE) int64 {
	var pos int64 = f.Seek(f, int64(0), func() int32 {
		if f.Flags&uint32(128) != 0 && uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wbase)) {
			return int32(2)
		} else {
			return int32(1)
		}
	}())
	if pos < int64(0) {
		return pos
	}
	if f.Rend != nil {
		pos += int64(uintptr(unsafe.Pointer(f.Rpos)) - uintptr(unsafe.Pointer(f.Rend)))
	} else if f.Wbase != nil {
		pos += int64(uintptr(unsafe.Pointer(f.Wpos)) - uintptr(unsafe.Pointer(f.Wbase)))
	}
	return pos
}
func __ftello(f *Struct__IO_FILE) int64 {
	var pos int64
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	pos = __ftello_unlocked(f)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return pos
}
func Ftell(f *Struct__IO_FILE) int64 {
	var pos int64 = __ftello(f)
	if pos > int64(2147483647) {
		*X__errno_location() = int32(75)
		return int64(-1)
	}
	return int64(pos)
}
