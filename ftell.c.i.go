package libc

import unsafe "unsafe"

func __ftello_unlocked(f *struct__IO_FILE) int64 {
	var pos int64 = f.seek(f, int64(0), func() int32 {
		if f.flags&uint32(128) != 0 && uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
			return int32(2)
		} else {
			return int32(1)
		}
	}())
	if pos < int64(0) {
		return pos
	}
	if f.rend != nil {
		pos += int64(uintptr(unsafe.Pointer(f.rpos)) - uintptr(unsafe.Pointer(f.rend)))
	} else if f.wbase != nil {
		pos += int64(uintptr(unsafe.Pointer(f.wpos)) - uintptr(unsafe.Pointer(f.wbase)))
	}
	return pos
}
func __ftello(f *struct__IO_FILE) int64 {
	var pos int64
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
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
func Ftell(f *struct__IO_FILE) int64 {
	var pos int64 = __ftello(f)
	if pos > int64(2147483647) {
		*__errno_location() = int32(75)
		return int64(-1)
	}
	return int64(pos)
}
