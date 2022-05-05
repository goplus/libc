package libc

import unsafe "unsafe"

func __lockfile(f *struct__IO_FILE) int32 {
	var owner int32 = f.lock
	var tid int32 = (*struct___pthread)(unsafe.Pointer(uintptr(__get_tp() - 200 - uint64(0)))).tid
	if owner & ^1073741824 == tid {
		return int32(0)
	}
	owner = a_cas(&f.lock, 0, tid)
	if !(owner != 0) {
		return int32(1)
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &owner
		*_cgo_addr = a_cas(&f.lock, 0, tid|1073741824)
		return *_cgo_addr
	}() != 0 {
		if owner&1073741824 != 0 || a_cas(&f.lock, owner, owner|1073741824) == owner {
			__futexwait(unsafe.Pointer(&f.lock), owner|1073741824, 1)
		}
	}
	return int32(1)
}
func __unlockfile(f *struct__IO_FILE) {
	if a_swap(&f.lock, 0)&1073741824 != 0 {
		__wake(unsafe.Pointer(&f.lock), 1, 1)
	}
}
