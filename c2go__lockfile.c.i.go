package libc

import unsafe "unsafe"

func __lockfile(f *struct__IO_FILE) int32 {
	var owner int32 = f.lock
	var tid int32 = __pthread_self().tid
	if owner&-1073741825 == tid {
		return int32(0)
	}
	owner = a_cas(&f.lock, int32(0), tid)
	if !(owner != 0) {
		return int32(1)
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &owner
		*_cgo_addr = a_cas(&f.lock, int32(0), tid|int32(1073741824))
		return *_cgo_addr
	}() != 0 {
		if owner&int32(1073741824) != 0 || a_cas(&f.lock, owner, owner|int32(1073741824)) == owner {
			__futexwait(unsafe.Pointer(&f.lock), owner|int32(1073741824), int32(1))
		}
	}
	return int32(1)
}
func __unlockfile(f *struct__IO_FILE) {
	if a_swap(&f.lock, int32(0))&int32(1073741824) != 0 {
		__wake(unsafe.Pointer(&f.lock), int32(1), int32(1))
	}
}
