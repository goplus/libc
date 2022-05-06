package libc

import (
	"testing"
	"unsafe"
)

func gostring(s *int8) string {
	n, arr := 0, (*[1 << 20]byte)(unsafe.Pointer(s))
	for arr[n] != 0 {
		n++
	}
	return string(arr[:n])
}

func C(s string) *int8 {
	n := len(s)
	ret := make([]byte, n+1)
	copy(ret, s)
	ret[n] = '\x00'
	return (*int8)(unsafe.Pointer(&ret[0]))
}

func goSprintf(format *int8, args ...interface{}) string {
	var buf [1 << 12]int8
	Sprintf(&buf[0], format, args...)
	return gostring(&buf[0])
}

func TestPrintf(t *testing.T) {
	n := Printf(C("Hello, world\n"))
	if n != 13 {
		t.Fatal("TestPrintf:", n)
	}
}

func TestSprintf(t *testing.T) {
	if s := goSprintf(C("Hi %7.1f\n"), 3.1415926); s != "Hi     3.1\n" {
		t.Fatal("TestSprintf:", s)
	}
}
