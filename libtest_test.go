//go:build darwin || linux
// +build darwin linux

package libc_test

import (
	"testing"

	"github.com/goplus/libc/test/cmd/test/qsort"
	"github.com/goplus/libc/test/cmd/test/string"
	"github.com/goplus/libc/test/cmd/test/string_memcpy"
	"github.com/goplus/libc/test/cmd/test/string_memset"
)

func TestFromTestdata(t *testing.T) {
	qsort.TestMain(t)
	string_memcpy.TestMain(t)
	string_memset.TestMain(t)
	string.TestMain(t)
}
