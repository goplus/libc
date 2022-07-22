package libc

import unsafe "unsafe"

func dummy_cgo18_fclose(f *struct__IO_FILE) {
}
func Fclose(f *struct__IO_FILE) int32 {
	var r int32
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	r = Fflush(f)
	r |= f.close(f)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	if f.flags&uint32(1) != 0 {
		return r
	}
	__unlist_locked_file(f)
	var head **struct__IO_FILE = __ofl_lock()
	if f.prev != nil {
		f.prev.next = f.next
	}
	if f.next != nil {
		f.next.prev = f.prev
	}
	if uintptr(unsafe.Pointer(*head)) == uintptr(unsafe.Pointer(f)) {
		*head = f.next
	}
	__ofl_unlock()
	Free(unsafe.Pointer(f.getln_buf))
	Free(unsafe.Pointer(f))
	return r
}
