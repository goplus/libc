package libc

import unsafe "unsafe"

type _cgos_ctx_setxid struct {
	id  int32
	eid int32
	sid int32
	nr  int32
	ret int32
}

func _cgos_do_setxid_setxid(p unsafe.Pointer) {
	var c *_cgos_ctx_setxid = (*_cgos_ctx_setxid)(p)
	if c.ret < int32(0) {
		return
	}
	var ret int32 = int32(__syscall3(int64(c.nr), int64(c.id), int64(c.eid), int64(c.sid)))
	if ret != 0 && !(c.ret != 0) {
		__block_all_sigs(nil)
		__syscall2(int64(37), int64(__syscall0(int64(20))), int64(9))
	}
	c.ret = ret
}
func __setxid(nr int32, id int32, eid int32, sid int32) int32 {
	var c _cgos_ctx_setxid = _cgos_ctx_setxid{id, eid, sid, nr, int32(1)}
	__synccall(_cgos_do_setxid_setxid, unsafe.Pointer(&c))
	return int32(__syscall_ret(uint64(c.ret)))
}
