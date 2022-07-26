package libc

import unsafe "unsafe"

func __shlim(f *Struct__IO_FILE, lim int64) {
	f.Shlim = lim
	f.Shcnt = int64(uintptr(unsafe.Pointer(f.Buf)) - uintptr(unsafe.Pointer(f.Rpos)))
	if lim != 0 && int64(uintptr(unsafe.Pointer(f.Rend))-uintptr(unsafe.Pointer(f.Rpos))) > lim {
		f.Shend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Rpos)) + uintptr(lim)))
	} else {
		f.Shend = f.Rend
	}
}
func __shgetc(f *Struct__IO_FILE) int32 {
	var c int32
	var cnt int64 = f.Shcnt + int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf)))
	if f.Shlim != 0 && cnt >= f.Shlim || func() (_cgo_ret int32) {
		_cgo_addr := &c
		*_cgo_addr = __uflow(f)
		return *_cgo_addr
	}() < int32(0) {
		f.Shcnt = int64(uintptr(unsafe.Pointer(f.Buf))-uintptr(unsafe.Pointer(f.Rpos))) + cnt
		f.Shend = f.Rpos
		f.Shlim = int64(-1)
		return -1
	}
	cnt++
	if f.Shlim != 0 && int64(uintptr(unsafe.Pointer(f.Rend))-uintptr(unsafe.Pointer(f.Rpos))) > f.Shlim-cnt {
		f.Shend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Rpos)) + uintptr(f.Shlim-cnt)))
	} else {
		f.Shend = f.Rend
	}
	f.Shcnt = int64(uintptr(unsafe.Pointer(f.Buf))-uintptr(unsafe.Pointer(f.Rpos))) + cnt
	if uintptr(unsafe.Pointer(f.Rpos)) <= uintptr(unsafe.Pointer(f.Buf)) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Rpos)) - uintptr(1))) = uint8(c)
	}
	return c
}
