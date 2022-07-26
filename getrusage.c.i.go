package libc

import unsafe "unsafe"

func getrusage(who int32, ru *struct_rusage) int32 {
	var r int32
	var dest *int8 = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&ru.ru_maxrss)))) - uintptr(32)))
	r = int32(__syscall2(int64(117), int64(who), int64(uintptr(unsafe.Pointer(dest)))))
	if !(r != 0) && false {
		var kru [4]int64
		Memcpy(unsafe.Pointer((*int64)(unsafe.Pointer(&kru))), unsafe.Pointer(dest), 32)
		ru.ru_utime = Struct_timeval{int64(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&kru)))) + uintptr(int32(0))*8))), int64(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&kru)))) + uintptr(int32(1))*8)))}
		ru.ru_stime = Struct_timeval{int64(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&kru)))) + uintptr(int32(2))*8))), int64(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&kru)))) + uintptr(int32(3))*8)))}
	}
	return int32(__syscall_ret(uint64(r)))
}
