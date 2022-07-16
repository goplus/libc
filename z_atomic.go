package libc

import "sync/atomic"

func a_ll(p *int32) int32 {
	return atomic.LoadInt32(p)
}

func a_sc(p *int32, v int32) int32 {
	atomic.StoreInt32(p, v)
	return 1
}

func a_swap(p *int32, v int32) int32 {
	return atomic.SwapInt32(p, v)
}

func a_cas(p *int32, t int32, s int32) int32 {
	if atomic.CompareAndSwapInt32(p, t, s) {
		return t
	}
	return atomic.LoadInt32(p)
}
