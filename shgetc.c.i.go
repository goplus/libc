package libc

import unsafe "unsafe"

func __shlim(f *struct__IO_FILE, lim int64) {
	f.shlim = lim
	f.shcnt = int64(uintptr(unsafe.Pointer(f.buf)) - uintptr(unsafe.Pointer(f.rpos)))
	if lim != 0 && int64(uintptr(unsafe.Pointer(f.rend))-uintptr(unsafe.Pointer(f.rpos))) > lim {
		f.shend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.rpos)) + uintptr(lim)))
	} else {
		f.shend = f.rend
	}
}
func __shgetc(f *struct__IO_FILE) int32 {
	var c int32
	var cnt int64 = f.shcnt + int64(uintptr(unsafe.Pointer(f.rpos))-uintptr(unsafe.Pointer(f.buf)))
	if f.shlim != 0 && cnt >= f.shlim || func() (_cgo_ret int32) {
		_cgo_addr := &c
		*_cgo_addr = __uflow(f)
		return *_cgo_addr
	}() < int32(0) {
		f.shcnt = int64(uintptr(unsafe.Pointer(f.buf))-uintptr(unsafe.Pointer(f.rpos))) + cnt
		f.shend = f.rpos
		f.shlim = int64(-1)
		return -1
	}
	cnt++
	if f.shlim != 0 && int64(uintptr(unsafe.Pointer(f.rend))-uintptr(unsafe.Pointer(f.rpos))) > f.shlim-cnt {
		f.shend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.rpos)) + uintptr(f.shlim-cnt)))
	} else {
		f.shend = f.rend
	}
	f.shcnt = int64(uintptr(unsafe.Pointer(f.buf))-uintptr(unsafe.Pointer(f.rpos))) + cnt
	if uintptr(unsafe.Pointer(f.rpos)) <= uintptr(unsafe.Pointer(f.buf)) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.rpos)) - uintptr(1))) = uint8(c)
	}
	return c
}
