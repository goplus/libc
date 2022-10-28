package libc

import unsafe "unsafe"

var _cgos_invsqrtpi_jn float64 = 0.56418958354775628

func Jn(n int32, x float64) float64 {
	var ix uint32
	var lx uint32
	var nm1 int32
	var i int32
	var sign int32
	var a float64
	var b float64
	var temp float64
	for {
		var __u uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_18_jn{x}))
		ix = uint32(__u >> int32(32))
		lx = uint32(__u)
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix|(lx|-lx)>>int32(31) > uint32(2146435072) {
		return x
	}
	if n == int32(0) {
		return J0(x)
	}
	if n < int32(0) {
		nm1 = -(n + int32(1))
		x = -x
		sign ^= int32(1)
	} else {
		nm1 = n - int32(1)
	}
	if nm1 == int32(0) {
		return J1(x)
	}
	sign &= n
	x = Fabs(x)
	if ix|lx == uint32(0) || ix == uint32(2146435072) {
		b = float64(0.0)
	} else if float64(nm1) < x {
		if ix >= uint32(1389363200) {
			switch nm1 & int32(3) {
			case int32(0):
				temp = -Cos(x) + Sin(x)
				break
			case int32(1):
				temp = -Cos(x) - Sin(x)
				break
			case int32(2):
				temp = Cos(x) - Sin(x)
				break
			default:
				fallthrough
			case int32(3):
				temp = Cos(x) + Sin(x)
				break
			}
			b = _cgos_invsqrtpi_jn * temp / Sqrt(x)
		} else {
			a = J0(x)
			b = J1(x)
			for i = int32(0); i < nm1; {
				i++
				temp = b
				b = b*(2.0*float64(i)/x) - a
				a = temp
			}
		}
	} else if ix < uint32(1041235968) {
		if nm1 > int32(32) {
			b = float64(0.0)
		} else {
			temp = x * 0.5
			b = temp
			a = float64(1.0)
			for i = int32(2); i <= nm1+int32(1); i++ {
				a *= float64(i)
				b *= temp
			}
			b = b / a
		}
	} else {
		var t float64
		var q0 float64
		var q1 float64
		var w float64
		var h float64
		var z float64
		var tmp float64
		var nf float64
		var k int32
		nf = float64(nm1) + 1.0
		w = float64(int32(2)) * nf / x
		h = float64(int32(2)) / x
		z = w + h
		q0 = w
		q1 = w*z - 1.0
		k = int32(1)
		for q1 < 1.0e+9 {
			k += int32(1)
			z += h
			tmp = z*q1 - q0
			q0 = q1
			q1 = tmp
		}
		for func() int32 {
			t = float64(0.0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				*_cgo_addr = k
				return *_cgo_addr
			}()
		}(); i >= int32(0); i-- {
			t = float64(int32(1)) / (float64(int32(2))*(float64(i)+nf)/x - t)
		}
		a = t
		b = float64(1.0)
		tmp = nf * Log(Fabs(w))
		if tmp < 709.78271289338397 {
			for i = nm1; i > int32(0); i-- {
				temp = b
				b = b*(2.0*float64(i))/x - a
				a = temp
			}
		} else {
			for i = nm1; i > int32(0); i-- {
				temp = b
				b = b*(2.0*float64(i))/x - a
				a = temp
				if b > 3.2733906078961419e+150 {
					a /= b
					t /= b
					b = float64(1.0)
				}
			}
		}
		z = J0(x)
		w = J1(x)
		if Fabs(z) >= Fabs(w) {
			b = t * z / b
		} else {
			b = t * w / a
		}
	}
	return func() float64 {
		if sign != 0 {
			return -b
		} else {
			return b
		}
	}()
}

type _cgoz_18_jn struct {
	_f float64
}

func Yn(n int32, x float64) float64 {
	var ix uint32
	var lx uint32
	var ib uint32
	var nm1 int32
	var sign int32
	var i int32
	var a float64
	var b float64
	var temp float64
	for {
		var __u uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_19_jn{x}))
		ix = uint32(__u >> int32(32))
		lx = uint32(__u)
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix|(lx|-lx)>>int32(31) > uint32(2146435072) {
		return x
	}
	if sign != 0 && ix|lx != uint32(0) {
		return func() float64 {
			return float64(int32(0))
		}() / 0.0
	}
	if ix == uint32(2146435072) {
		return float64(0.0)
	}
	if n == int32(0) {
		return Y0(x)
	}
	if n < int32(0) {
		nm1 = -(n + int32(1))
		sign = n & int32(1)
	} else {
		nm1 = n - int32(1)
		sign = int32(0)
	}
	if nm1 == int32(0) {
		return func() float64 {
			if sign != 0 {
				return -Y1(x)
			} else {
				return Y1(x)
			}
		}()
	}
	if ix >= uint32(1389363200) {
		switch nm1 & int32(3) {
		case int32(0):
			temp = -Sin(x) - Cos(x)
			break
		case int32(1):
			temp = -Sin(x) + Cos(x)
			break
		case int32(2):
			temp = Sin(x) + Cos(x)
			break
		default:
			fallthrough
		case int32(3):
			temp = Sin(x) - Cos(x)
			break
		}
		b = _cgos_invsqrtpi_jn * temp / Sqrt(x)
	} else {
		a = Y0(x)
		b = Y1(x)
		for {
			ib = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_jn{b})) >> int32(32))
			if true {
				break
			}
		}
		for i = int32(0); i < nm1 && ib != uint32(4293918720); {
			i++
			temp = b
			b = 2.0*float64(i)/x*b - a
			for {
				ib = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_21_jn{b})) >> int32(32))
				if true {
					break
				}
			}
			a = temp
		}
	}
	return func() float64 {
		if sign != 0 {
			return -b
		} else {
			return b
		}
	}()
}

type _cgoz_19_jn struct {
	_f float64
}
type _cgoz_20_jn struct {
	_f float64
}
type _cgoz_21_jn struct {
	_f float64
}
