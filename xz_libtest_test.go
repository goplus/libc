//go:build darwin
// +build darwin

package libc_test

import (
	"testing"

	teststring "github.com/goplus/libc/test/cmd/test/string"

	"github.com/goplus/libc/test/cmd/test/acos"
	"github.com/goplus/libc/test/cmd/test/acosf"
	"github.com/goplus/libc/test/cmd/test/acosl"

	"github.com/goplus/libc/test/cmd/test/asin"
	"github.com/goplus/libc/test/cmd/test/asinf"
	"github.com/goplus/libc/test/cmd/test/asinl"

	"github.com/goplus/libc/test/cmd/test/qsort"
	"github.com/goplus/libc/test/cmd/test/random"
	"github.com/goplus/libc/test/cmd/test/string_memcpy"
	"github.com/goplus/libc/test/cmd/test/string_memset"
	"github.com/goplus/libc/test/cmd/test/strtod"
)

func TestFromTestdata(t *testing.T) {
	acos.TestMain(t)
	acosf.TestMain(t)
	acosl.TestMain(t)

	asin.TestMain(t)
	asinf.TestMain(t)
	asinl.TestMain(t)

	qsort.TestMain(t)
	random.TestMain(t)
	string_memcpy.TestMain(t)
	string_memset.TestMain(t)
	teststring.TestMain(t)
	strtod.TestMain(t)
}
