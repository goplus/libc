package libc

import unsafe "unsafe"

type cmpfun = func(unsafe.Pointer, unsafe.Pointer) int32

func pntz_cgo18_qsort(p *uint64) int32 {
	var r int32 = a_ctz_l(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) - uint64(1))
	if r != int32(0) || uint64(func() (_cgo_ret int32) {
		_cgo_addr := &r
		*_cgo_addr = int32(64 + uint64(a_ctz_l(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)))))
		return *_cgo_addr
	}()) != 64 {
		return r
	}
	return int32(0)
}
func cycle_cgo19_qsort(width uint64, ar **uint8, n int32) {
	var tmp [256]uint8
	var l uint64
	var i int32
	if n < int32(2) {
		return
	}
	*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ar)) + uintptr(n)*8)) = (*uint8)(unsafe.Pointer(&tmp))
	for width != 0 {
		l = func() uint64 {
			if 256 < width {
				return 256
			} else {
				return width
			}
		}()
		Memcpy(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ar)) + uintptr(n)*8))), unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ar)) + uintptr(int32(0))*8))), l)
		for i = int32(0); i < n; i++ {
			Memcpy(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ar)) + uintptr(i)*8))), unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ar)) + uintptr(i+int32(1))*8))), l)
			*(*uintptr)(unsafe.Pointer(&*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(ar)) + uintptr(i)*8)))) += uintptr(l)
		}
		width -= l
	}
}
func shl_cgo20_qsort(p *uint64, n int32) {
	if uint64(n) >= 64 {
		n -= int32(64)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8))
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) = uint64(0)
	}
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)) <<= n
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)) |= *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) >> (64 - uint64(n))
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) <<= n
}
func shr_cgo21_qsort(p *uint64, n int32) {
	if uint64(n) >= 64 {
		n -= int32(64)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8))
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)) = uint64(0)
	}
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) >>= n
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) |= *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)) << (64 - uint64(n))
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8)) >>= n
}
func sift_cgo22_qsort(head *uint8, width uint64, cmp func(unsafe.Pointer, unsafe.Pointer) int32, pshift int32, lp *uint64) {
	var rt *uint8
	var lf *uint8
	var ar [113]*uint8
	var i int32 = int32(1)
	*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(int32(0))*8)) = head
	for pshift > int32(1) {
		rt = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head)) - uintptr(width)))
		lf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head))-uintptr(width))))) - uintptr(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(lp)) + uintptr(pshift-int32(2))*8)))))
		if cmp(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(int32(0))*8))), unsafe.Pointer(lf)) >= int32(0) && cmp(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(int32(0))*8))), unsafe.Pointer(rt)) >= int32(0) {
			break
		}
		if cmp(unsafe.Pointer(lf), unsafe.Pointer(rt)) >= int32(0) {
			*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())*8)) = lf
			head = lf
			pshift -= int32(1)
		} else {
			*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())*8)) = rt
			head = rt
			pshift -= int32(2)
		}
	}
	cycle_cgo19_qsort(width, (**uint8)(unsafe.Pointer(&ar)), i)
}
func trinkle_cgo23_qsort(head *uint8, width uint64, cmp func(unsafe.Pointer, unsafe.Pointer) int32, pp *uint64, pshift int32, trusty int32, lp *uint64) {
	var stepson *uint8
	var rt *uint8
	var lf *uint8
	var p [2]uint64
	var ar [113]*uint8
	var i int32 = int32(1)
	var trail int32
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + uintptr(int32(0))*8))
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(1))*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(pp)) + uintptr(int32(1))*8))
	*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(int32(0))*8)) = head
	for *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8)) != uint64(1) || *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(1))*8)) != uint64(0) {
		stepson = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head)) - uintptr(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(lp)) + uintptr(pshift)*8)))))
		if cmp(unsafe.Pointer(stepson), unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(int32(0))*8)))) <= int32(0) {
			break
		}
		if !(trusty != 0) && pshift > int32(1) {
			rt = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head)) - uintptr(width)))
			lf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head))-uintptr(width))))) - uintptr(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(lp)) + uintptr(pshift-int32(2))*8)))))
			if cmp(unsafe.Pointer(rt), unsafe.Pointer(stepson)) >= int32(0) || cmp(unsafe.Pointer(lf), unsafe.Pointer(stepson)) >= int32(0) {
				break
			}
		}
		*(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((**uint8)(unsafe.Pointer(&ar)))) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}())*8)) = stepson
		head = stepson
		trail = pntz_cgo18_qsort((*uint64)(unsafe.Pointer(&p)))
		shr_cgo21_qsort((*uint64)(unsafe.Pointer(&p)), trail)
		pshift += trail
		trusty = int32(0)
	}
	if !(trusty != 0) {
		cycle_cgo19_qsort(width, (**uint8)(unsafe.Pointer(&ar)), i)
		sift_cgo22_qsort(head, width, cmp, pshift, lp)
	}
}
func Qsort(base unsafe.Pointer, nel uint64, width uint64, cmp func(unsafe.Pointer, unsafe.Pointer) int32) {
	var lp [96]uint64
	var i uint64
	var size uint64 = width * nel
	var head *uint8
	var high *uint8
	var p [2]uint64 = [2]uint64{uint64(1), uint64(0)}
	var pshift int32 = int32(1)
	var trail int32
	if !(size != 0) {
		return
	}
	head = (*uint8)(base)
	high = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head))+uintptr(size))))) - uintptr(width)))
	for func() uint64 {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(int32(0))*8)) = func() (_cgo_ret uint64) {
			_cgo_addr := &*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(int32(1))*8))
			*_cgo_addr = width
			return *_cgo_addr
		}()
		return func() (_cgo_ret uint64) {
			_cgo_addr := &i
			*_cgo_addr = uint64(2)
			return *_cgo_addr
		}()
	}(); func() (_cgo_ret uint64) {
		_cgo_addr := &*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(i)*8))
		*_cgo_addr = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(i-uint64(2))*8)) + *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(i-uint64(1))*8)) + width
		return *_cgo_addr
	}() < size; i++ {
	}
	for uintptr(unsafe.Pointer(head)) < uintptr(unsafe.Pointer(high)) {
		if *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8))&uint64(3) == uint64(3) {
			sift_cgo22_qsort(head, width, cmp, pshift, (*uint64)(unsafe.Pointer(&lp)))
			shr_cgo21_qsort((*uint64)(unsafe.Pointer(&p)), int32(2))
			pshift += int32(2)
		} else {
			if *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(pshift-int32(1))*8)) >= uint64(uintptr(unsafe.Pointer(high))-uintptr(unsafe.Pointer(head))) {
				trinkle_cgo23_qsort(head, width, cmp, (*uint64)(unsafe.Pointer(&p)), pshift, int32(0), (*uint64)(unsafe.Pointer(&lp)))
			} else {
				sift_cgo22_qsort(head, width, cmp, pshift, (*uint64)(unsafe.Pointer(&lp)))
			}
			if pshift == int32(1) {
				shl_cgo20_qsort((*uint64)(unsafe.Pointer(&p)), int32(1))
				pshift = int32(0)
			} else {
				shl_cgo20_qsort((*uint64)(unsafe.Pointer(&p)), pshift-int32(1))
				pshift = int32(1)
			}
		}
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8)) |= uint64(1)
		*(*uintptr)(unsafe.Pointer(&head)) += uintptr(width)
	}
	trinkle_cgo23_qsort(head, width, cmp, (*uint64)(unsafe.Pointer(&p)), pshift, int32(0), (*uint64)(unsafe.Pointer(&lp)))
	for pshift != int32(1) || *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8)) != uint64(1) || *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(1))*8)) != uint64(0) {
		if pshift <= int32(1) {
			trail = pntz_cgo18_qsort((*uint64)(unsafe.Pointer(&p)))
			shr_cgo21_qsort((*uint64)(unsafe.Pointer(&p)), trail)
			pshift += trail
		} else {
			shl_cgo20_qsort((*uint64)(unsafe.Pointer(&p)), int32(2))
			pshift -= int32(2)
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8)) ^= uint64(7)
			shr_cgo21_qsort((*uint64)(unsafe.Pointer(&p)), int32(1))
			trinkle_cgo23_qsort((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head))-uintptr(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&lp)))) + uintptr(pshift)*8)))))))-uintptr(width))), width, cmp, (*uint64)(unsafe.Pointer(&p)), pshift+int32(1), int32(1), (*uint64)(unsafe.Pointer(&lp)))
			shl_cgo20_qsort((*uint64)(unsafe.Pointer(&p)), int32(1))
			*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&p)))) + uintptr(int32(0))*8)) |= uint64(1)
			trinkle_cgo23_qsort((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(head))-uintptr(width))), width, cmp, (*uint64)(unsafe.Pointer(&p)), pshift, int32(1), (*uint64)(unsafe.Pointer(&lp)))
		}
		*(*uintptr)(unsafe.Pointer(&head)) -= uintptr(width)
	}
}
