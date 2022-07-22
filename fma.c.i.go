package libc

import unsafe "unsafe"

type struct_num struct {
	m    uint64
	e    int32
	sign int32
}

func normalize_cgo22_fma(x float64) struct_num {
	var ix uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_23_fma{x}))
	var e int32 = int32(ix >> int32(52))
	var sign int32 = e & int32(2048)
	e &= int32(2047)
	if !(e != 0) {
		ix = *(*uint64)(unsafe.Pointer(&_cgoz_24_fma{x * 9.2233720368547758e+18}))
		e = int32(ix >> int32(52) & uint64(2047))
		e = func() int32 {
			if e != 0 {
				return e - int32(63)
			} else {
				return int32(2048)
			}
		}()
	}
	ix &= 4503599627370495
	ix |= 4503599627370496
	ix <<= int32(1)
	e -= 1076
	return struct_num{ix, e, sign}
}

type _cgoz_23_fma struct {
	f float64
}
type _cgoz_24_fma struct {
	f float64
}

func mul_cgo25_fma(hi *uint64, lo *uint64, x uint64, y uint64) {
	var t1 uint64
	var t2 uint64
	var t3 uint64
	var xlo uint64 = uint64(uint32(x))
	var xhi uint64 = x >> int32(32)
	var ylo uint64 = uint64(uint32(y))
	var yhi uint64 = y >> int32(32)
	t1 = xlo * ylo
	t2 = xlo*yhi + xhi*ylo
	t3 = xhi * yhi
	*lo = t1 + t2<<int32(32)
	*hi = t3 + t2>>int32(32) + func() uint64 {
		if t1 > *lo {
			return 1
		} else {
			return 0
		}
	}()
}
func Fma(x float64, y float64, z float64) float64 {
	var nx struct_num
	var ny struct_num
	var nz struct_num
	nx = normalize_cgo22_fma(x)
	ny = normalize_cgo22_fma(y)
	nz = normalize_cgo22_fma(z)
	if nx.e >= 971 || ny.e >= 971 {
		return x*y + z
	}
	if nz.e >= 971 {
		if nz.e > 971 {
			return x*y + z
		}
		return z
	}
	var rhi uint64
	var rlo uint64
	var zhi uint64
	var zlo uint64
	mul_cgo25_fma(&rhi, &rlo, nx.m, ny.m)
	var e int32 = nx.e + ny.e
	var d int32 = nz.e - e
	if d > int32(0) {
		if d < int32(64) {
			zlo = nz.m << d
			zhi = nz.m >> (int32(64) - d)
		} else {
			zlo = uint64(0)
			zhi = nz.m
			e = nz.e - int32(64)
			d -= int32(64)
			if d == int32(0) {
			} else if d < int32(64) {
				rlo = rhi<<(int32(64)-d) | rlo>>d | func() uint64 {
					if !!(rlo<<(int32(64)-d) != 0) {
						return 1
					} else {
						return 0
					}
				}()
				rhi = rhi >> d
			} else {
				rlo = uint64(1)
				rhi = uint64(0)
			}
		}
	} else {
		zhi = uint64(0)
		d = -d
		if d == int32(0) {
			zlo = nz.m
		} else if d < int32(64) {
			zlo = nz.m>>d | func() uint64 {
				if !!(nz.m<<(int32(64)-d) != 0) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			zlo = uint64(1)
		}
	}
	var sign int32 = nx.sign ^ ny.sign
	var samesign int32 = func() int32 {
		if !(sign^nz.sign != 0) {
			return 1
		} else {
			return 0
		}
	}()
	var nonzero int32 = int32(1)
	if samesign != 0 {
		rlo += zlo
		rhi += zhi + func() uint64 {
			if rlo < zlo {
				return 1
			} else {
				return 0
			}
		}()
	} else {
		var t uint64 = rlo
		rlo -= zlo
		rhi = rhi - zhi - func() uint64 {
			if t < rlo {
				return 1
			} else {
				return 0
			}
		}()
		if rhi>>int32(63) != 0 {
			rlo = -rlo
			rhi = -rhi - func() uint64 {
				if !!(rlo != 0) {
					return 1
				} else {
					return 0
				}
			}()
			sign = func() int32 {
				if !(sign != 0) {
					return 1
				} else {
					return 0
				}
			}()
		}
		nonzero = func() int32 {
			if !!(rhi != 0) {
				return 1
			} else {
				return 0
			}
		}()
	}
	if nonzero != 0 {
		e += int32(64)
		d = a_clz_64(rhi) - int32(1)
		rhi = rhi<<d | rlo>>(int32(64)-d) | func() uint64 {
			if !!(rlo<<d != 0) {
				return 1
			} else {
				return 0
			}
		}()
	} else if rlo != 0 {
		d = a_clz_64(rlo) - int32(1)
		if d < int32(0) {
			rhi = rlo>>int32(1) | rlo&uint64(1)
		} else {
			rhi = rlo << d
		}
	} else {
		return x*y + z
	}
	e -= d
	var i int64 = int64(rhi)
	if sign != 0 {
		i = -i
	}
	var r float64 = float64(i)
	if e < -1084 {
		if e == -1085 {
			var c float64 = 9.2233720368547758e+18
			if sign != 0 {
				c = -c
			}
			if r == c {
				var fltmin float32 = float32(1.0842021401737618e-19 * float64(1.17549435e-38) * r)
				return 2.2250738585072014e-308 / float64(1.17549435e-38) * float64(fltmin)
			}
			if rhi<<int32(53) != 0 {
				i = int64(rhi>>int32(1) | rhi&uint64(1) | 4611686018427387904)
				if sign != 0 {
					i = -i
				}
				r = float64(i)
				r = float64(int32(2))*r - c
				{
					var tiny float64 = 2.2250738585072014e-308 / float64(1.17549435e-38) * r
					r += float64(tiny*tiny) * (r - r)
				}
			}
		} else {
			d = int32(10)
			i = int64((rhi>>d | func() uint64 {
				if !!(rhi<<(int32(64)-d) != 0) {
					return 1
				} else {
					return 0
				}
			}()) << d)
			if sign != 0 {
				i = -i
			}
			r = float64(i)
		}
	}
	return Scalbn(r, e)
}
