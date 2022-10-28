package libc

import unsafe "unsafe"

func Jnf(n int32, x float32) float32 {
	var ix uint32
	var nm1 int32
	var sign int32
	var i int32
	var a float32
	var b float32
	var temp float32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_jnf{x}))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix > uint32(2139095040) {
		return x
	}
	if n == int32(0) {
		return J0f(x)
	}
	if n < int32(0) {
		nm1 = -(n + int32(1))
		x = -x
		sign ^= int32(1)
	} else {
		nm1 = n - int32(1)
	}
	if nm1 == int32(0) {
		return J1f(x)
	}
	sign &= n
	x = Fabsf(x)
	if ix == uint32(0) || ix == uint32(2139095040) {
		b = float32(0.0)
	} else if float32(nm1) < x {
		a = J0f(x)
		b = J1f(x)
		for i = int32(0); i < nm1; {
			i++
			temp = b
			b = b*(2.0*float32(i)/x) - a
			a = temp
		}
	} else if ix < uint32(897581056) {
		if nm1 > int32(8) {
			nm1 = int32(8)
		}
		temp = 0.5 * x
		b = temp
		a = float32(1.0)
		for i = int32(2); i <= nm1+int32(1); i++ {
			a *= float32(i)
			b *= temp
		}
		b = b / a
	} else {
		var t float32
		var q0 float32
		var q1 float32
		var w float32
		var h float32
		var z float32
		var tmp float32
		var nf float32
		var k int32
		nf = float32(nm1) + 1.0
		w = float32(int32(2)) * nf / x
		h = float32(int32(2)) / x
		z = w + h
		q0 = w
		q1 = w*z - 1.0
		k = int32(1)
		for q1 < 1.0e+4 {
			k += int32(1)
			z += h
			tmp = z*q1 - q0
			q0 = q1
			q1 = tmp
		}
		for func() int32 {
			t = float32(0.0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				*_cgo_addr = k
				return *_cgo_addr
			}()
		}(); i >= int32(0); i-- {
			t = 1.0 / (float32(int32(2))*(float32(i)+nf)/x - t)
		}
		a = t
		b = float32(1.0)
		tmp = nf * Logf(Fabsf(w))
		if tmp < 88.7216796 {
			for i = nm1; i > int32(0); i-- {
				temp = b
				b = 2.0*float32(i)*b/x - a
				a = temp
			}
		} else {
			for i = nm1; i > int32(0); i-- {
				temp = b
				b = 2.0*float32(i)*b/x - a
				a = temp
				if b > 1.1529215e+18 {
					a /= b
					t /= b
					b = float32(1.0)
				}
			}
		}
		z = J0f(x)
		w = J1f(x)
		if Fabsf(z) >= Fabsf(w) {
			b = t * z / b
		} else {
			b = t * w / a
		}
	}
	return func() float32 {
		if sign != 0 {
			return -b
		} else {
			return b
		}
	}()
}

type _cgoz_18_jnf struct {
	_f float32
}

func Ynf(n int32, x float32) float32 {
	var ix uint32
	var ib uint32
	var nm1 int32
	var sign int32
	var i int32
	var a float32
	var b float32
	var temp float32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_19_jnf{x}))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix > uint32(2139095040) {
		return x
	}
	if sign != 0 && ix != uint32(0) {
		return func() float32 {
			return float32(int32(0))
		}() / 0.0
	}
	if ix == uint32(2139095040) {
		return float32(0.0)
	}
	if n == int32(0) {
		return Y0f(x)
	}
	if n < int32(0) {
		nm1 = -(n + int32(1))
		sign = n & int32(1)
	} else {
		nm1 = n - int32(1)
		sign = int32(0)
	}
	if nm1 == int32(0) {
		return func() float32 {
			if sign != 0 {
				return -Y1f(x)
			} else {
				return Y1f(x)
			}
		}()
	}
	a = Y0f(x)
	b = Y1f(x)
	for {
		ib = *(*uint32)(unsafe.Pointer(&_cgoz_20_jnf{b}))
		if true {
			break
		}
	}
	for i = int32(0); i < nm1 && ib != uint32(4286578688); {
		i++
		temp = b
		b = 2.0*float32(i)/x*b - a
		for {
			ib = *(*uint32)(unsafe.Pointer(&_cgoz_21_jnf{b}))
			if true {
				break
			}
		}
		a = temp
	}
	return func() float32 {
		if sign != 0 {
			return -b
		} else {
			return b
		}
	}()
}

type _cgoz_19_jnf struct {
	_f float32
}
type _cgoz_20_jnf struct {
	_f float32
}
type _cgoz_21_jnf struct {
	_f float32
}
