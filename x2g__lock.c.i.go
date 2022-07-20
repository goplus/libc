package libc

import unsafe "unsafe"

func __lock(l *int32) {
	var need_locks int32 = int32(__libc.need_locks)
	if !(need_locks != 0) {
		return
	}
	var current int32 = a_cas(l, int32(0), -2147483647)
	if need_locks < int32(0) {
		__libc.need_locks = int8(0)
	}
	if !(current != 0) {
		return
	}
	for i := uint32(uint32(0)); i < uint32(10); i++ {
		if current < int32(0) {
			current -= -2147483647
		}
		var val int32 = a_cas(l, current, -2147483648+(current+int32(1)))
		if val == current {
			return
		}
		current = val
	}
	current = a_fetch_add(l, int32(1)) + int32(1)
	for {
		if current < int32(0) {
			__futexwait(unsafe.Pointer(l), current, int32(1))
			current -= -2147483647
		}
		var val int32 = a_cas(l, current, -2147483648+current)
		if val == current {
			return
		}
		current = val
	}
}
func __unlock(l *int32) {
	if *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(int32(0))*4)) < int32(0) {
		if a_fetch_add(l, 2147483647) != -2147483647 {
			__wake(unsafe.Pointer(l), int32(1), int32(1))
		}
	}
}
