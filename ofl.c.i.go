package libc

import unsafe "unsafe"

var _cgos_ofl_head__ofl *struct__IO_FILE
var _cgos_ofl_lock__ofl [1]int32
var __stdio_ofl_lockptr *int32 = (*int32)(unsafe.Pointer(&_cgos_ofl_lock__ofl))

func __ofl_lock() **struct__IO_FILE {
	__lock((*int32)(unsafe.Pointer(&_cgos_ofl_lock__ofl)))
	return &_cgos_ofl_head__ofl
}
func __ofl_unlock() {
	__unlock((*int32)(unsafe.Pointer(&_cgos_ofl_lock__ofl)))
}
