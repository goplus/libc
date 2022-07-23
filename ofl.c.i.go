package libc

import unsafe "unsafe"

var ofl_head_cgos__ofl *struct__IO_FILE
var ofl_lock_cgos__ofl [1]int32
var __stdio_ofl_lockptr *int32 = (*int32)(unsafe.Pointer(&ofl_lock_cgos__ofl))

func __ofl_lock() **struct__IO_FILE {
	__lock((*int32)(unsafe.Pointer(&ofl_lock_cgos__ofl)))
	return &ofl_head_cgos__ofl
}
func __ofl_unlock() {
	__unlock((*int32)(unsafe.Pointer(&ofl_lock_cgos__ofl)))
}
