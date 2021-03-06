//go:build darwin
// +build darwin

package libc_test

import (
	"testing"

	teststring "github.com/goplus/libc/test/cmd/test/string"

	"github.com/goplus/libc/test/cmd/test/acos"
	"github.com/goplus/libc/test/cmd/test/qsort"
	"github.com/goplus/libc/test/cmd/test/random"
	"github.com/goplus/libc/test/cmd/test/string_memcpy"
	"github.com/goplus/libc/test/cmd/test/string_memset"
	"github.com/goplus/libc/test/cmd/test/strtod"
)

func TestFromTestdata(t *testing.T) {
	acos.TestMain(t)
	qsort.TestMain(t)
	random.TestMain(t)
	string_memcpy.TestMain(t)
	string_memset.TestMain(t)
	teststring.TestMain(t)
	strtod.TestMain(t)
}
