package common

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
)

var seed_cgo1_rand uint64 = uint64(18446744073709551615)

func rand32_cgo2_rand() uint32 {
	seed_cgo1_rand = uint64(6364136223846793005)*seed_cgo1_rand + uint64(1)
	return uint32(seed_cgo1_rand >> int32(32))
}
func rand64_cgo3_rand() uint64 {
	var u uint64 = uint64(rand32_cgo2_rand())
	return u<<int32(32) | uint64(rand32_cgo2_rand())
}
func frand_cgo4_rand() float64 {
	return float64(rand64_cgo3_rand()) * 5.4210108624275222e-20
}
func frandf_cgo5_rand() float32 {
	return float32(rand32_cgo2_rand()) * 2.32830644e-10
}
func frandl_cgo6_rand() float64 {
	return float64(rand64_cgo3_rand()) * 5.42101086242752217004e-20
}
func t_randseed(s uint64) {
	seed_cgo1_rand = s
}
func t_randn(n uint64) uint64 {
	var r uint64
	var m uint64
	m = uint64(18446744073709551615)
	m -= m % n
	for func() (_cgo_ret uint64) {
		_cgo_addr := &r
		*_cgo_addr = rand64_cgo3_rand()
		return *_cgo_addr
	}() >= m {
	}
	return r % n
}
func t_randint(a uint64, b uint64) uint64 {
	var n uint64 = b - a + uint64(1)
	if n != 0 {
		return a + t_randn(n)
	}
	return rand64_cgo3_rand()
}
func shuffle2_cgo7_rand(p *uint64, q *uint64, np uint64, nq uint64) {
	var r uint64
	var t uint64
	for np != 0 {
		r = uint64(t_randn(uint64(nq + func() (_cgo_ret uint64) {
			_cgo_addr := &np
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}())))
		t = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(np)*8))
		if r < nq {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(np)*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(r)*8))
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(r)*8)) = t
		} else {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(np)*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(r-nq)*8))
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(r-nq)*8)) = t
		}
	}
}
func t_shuffle(p *uint64, n uint64) {
	shuffle2_cgo7_rand(p, nil, n, uint64(0))
}
func t_randrange(p *uint64, n uint64) {
	var i uint64
	for i = uint64(0); i < n; i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)) = uint64(i)
	}
	t_shuffle(p, n)
}
func insert_cgo8_rand(tab *uint64, len uint64, v uint64) int32 {
	var i uint64 = uint64(v & uint64(len-uint64(1)))
	var j uint64 = uint64(1)
	for *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(tab)) + uintptr(i)*8)) != 0 {
		if *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(tab)) + uintptr(i)*8)) == v {
			return -1
		}
		i += func() (_cgo_ret uint64) {
			_cgo_addr := &j
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()
		i &= len - uint64(1)
	}
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(tab)) + uintptr(i)*8)) = v
	return int32(0)
}
func t_choose(n uint64, k uint64, p *uint64) int32 {
	var tab *uint64
	var i uint64
	var j uint64
	var len uint64
	if n < uint64(k) {
		return -1
	}
	if n < uint64(16) {
		for k != 0 {
			if t_randn(func() (_cgo_ret uint64) {
				_cgo_addr := &n
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()) < uint64(k) {
				*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(func() (_cgo_ret uint64) {
					_cgo_addr := &k
					*_cgo_addr--
					return *_cgo_addr
				}())*8)) = n
			}
		}
		return int32(0)
	}
	if k < uint64(8) {
		for i = uint64(0); i < k; {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)) = t_randn(n)
			for j = uint64(0); *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(j)*8)) != *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)); j++ {
			}
			if j == i {
				i++
			}
		}
		return int32(0)
	}
	if n < uint64(uint64(5)*k) && (n-uint64(k))*uint64(8) < uint64(18446744073709551615) {
		tab = (*uint64)(libc.Malloc(uint64((n - uint64(k)) * uint64(8))))
		if !(tab != nil) {
			return -1
		}
		for i = uint64(0); i < k; i++ {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)) = uint64(i)
		}
		for ; uint64(i) < n; i++ {
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(tab)) + uintptr(i-k)*8)) = uint64(i)
		}
		if uint64(k) < n-uint64(k) {
			shuffle2_cgo7_rand(p, tab, k, uint64(n-uint64(k)))
		} else {
			shuffle2_cgo7_rand(tab, p, uint64(n-uint64(k)), k)
		}
		libc.Free(unsafe.Pointer(tab))
		return int32(0)
	}
	for len = uint64(16); len < uint64(2)*k; len *= uint64(2) {
	}
	tab = (*uint64)(libc.Calloc(len, 8))
	if !(tab != nil) {
		return -1
	}
	for i = uint64(0); i < k; i++ {
		for insert_cgo8_rand(tab, len, t_randn(n)+uint64(1)) != 0 {
		}
	}
	for i = uint64(0); i < len; i++ {
		if *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(tab)) + uintptr(i)*8)) != 0 {
			*func() (_cgo_ret *uint64) {
				_cgo_addr := &p
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
				return
			}() = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(tab)) + uintptr(i)*8)) - uint64(1)
		}
	}
	libc.Free(unsafe.Pointer(tab))
	return int32(0)
}
