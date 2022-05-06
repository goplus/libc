package libc

import (
	"testing"
	"unsafe"
)

func C(s string) *int8 {
	n := len(s)
	ret := make([]byte, n+1)
	copy(ret, s)
	ret[n] = '\x00'
	return (*int8)(unsafe.Pointer(&ret[0]))
}

func TestPrintf(t *testing.T) {
	n := Printf(C("Hello, world\n"))
	if n != 13 {
		t.Fatal("TestPrintf:", n)
	}
}
