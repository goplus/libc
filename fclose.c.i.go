package libc

import unsafe "unsafe"

func _cgos_dummy_fclose(f *Struct__IO_FILE) {
}
func Fclose(f *Struct__IO_FILE) int32 {
	var r int32
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	r = Fflush(f)
	r |= f.Close(f)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	if f.Flags&uint32(1) != 0 {
		return r
	}
	__unlist_locked_file(f)
	var head **Struct__IO_FILE = __ofl_lock()
	if f.Prev != nil {
		f.Prev.Next = f.Next
	}
	if f.Next != nil {
		f.Next.Prev = f.Prev
	}
	if uintptr(unsafe.Pointer(*head)) == uintptr(unsafe.Pointer(f)) {
		*head = f.Next
	}
	__ofl_unlock()
	Free(unsafe.Pointer(f.Getln_buf))
	Free(unsafe.Pointer(f))
	return r
}
