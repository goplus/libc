package common

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
)

var _cgos_seed_rand uint64 = uint64(18446744073709551615)

func _cgos_rand32_rand() uint32 {
	_cgos_seed_rand = uint64(6364136223846793005)*_cgos_seed_rand + uint64(1)
	return uint32(_cgos_seed_rand >> int32(32))
}
func _cgos_rand64_rand() uint64 {
	var u uint64 = uint64(_cgos_rand32_rand())
	return u<<int32(32) | uint64(_cgos_rand32_rand())
}
func _cgos_frand_rand() float64 {
	return float64(_cgos_rand64_rand()) * 5.4210108624275222e-20
}
func _cgos_frandf_rand() float32 {
	return float32(_cgos_rand32_rand()) * 2.32830644e-10
}
func _cgos_frandl_rand() float64 {
	return float64(_cgos_rand64_rand()) * 5.42101086242752217004e-20
}
func T_randseed(s uint64) {
	_cgos_seed_rand = s
}
func T_randn(n uint64) uint64 {
	var r uint64
	var m uint64
	m = uint64(18446744073709551615)
	m -= m % n
	for func() (_cgo_ret uint64) {
		_cgo_addr := &r
		*_cgo_addr = _cgos_rand64_rand()
		return *_cgo_addr
	}() >= m {
	}
	return r % n
}
func T_randint(a uint64, b uint64) uint64 {
	var n uint64 = b - a + uint64(1)
	if n != 0 {
		return a + T_randn(n)
	}
	return _cgos_rand64_rand()
}
func _cgos_shuffle2_rand(p *uint64, q *uint64, np uint64, nq uint64) {
	var r uint64
	var t uint64
	for np != 0 {
		r = uint64(T_randn(uint64(nq + func() (_cgo_ret uint64) {
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
func T_shuffle(p *uint64, n uint64) {
	_cgos_shuffle2_rand(p, nil, n, uint64(0))
}
func T_randrange(p *uint64, n uint64) {
	var i uint64
	for i = uint64(0); i < n; i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)) = uint64(i)
	}
	T_shuffle(p, n)
}
func _cgos_insert_rand(tab *uint64, len uint64, v uint64) int32 {
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
func T_choose(n uint64, k uint64, p *uint64) int32 {
	var tab *uint64
	var i uint64
	var j uint64
	var len uint64
	if n < uint64(k) {
		return -1
	}
	if n < uint64(16) {
		for k != 0 {
			if T_randn(func() (_cgo_ret uint64) {
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
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)) = T_randn(n)
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
			_cgos_shuffle2_rand(p, tab, k, uint64(n-uint64(k)))
		} else {
			_cgos_shuffle2_rand(tab, p, uint64(n-uint64(k)), k)
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
		for _cgos_insert_rand(tab, len, T_randn(n)+uint64(1)) != 0 {
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
