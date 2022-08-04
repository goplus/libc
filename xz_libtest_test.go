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

	"github.com/goplus/libc/test/cmd/test/atan"
	"github.com/goplus/libc/test/cmd/test/atanf"
	"github.com/goplus/libc/test/cmd/test/atanl"

	"github.com/goplus/libc/test/cmd/test/atanh"
	"github.com/goplus/libc/test/cmd/test/atanhf"
	"github.com/goplus/libc/test/cmd/test/atanhl"

	"github.com/goplus/libc/test/cmd/test/cbrt"
	"github.com/goplus/libc/test/cmd/test/cbrtf"
	"github.com/goplus/libc/test/cmd/test/cbrtl"

	"github.com/goplus/libc/test/cmd/test/ceil"
	"github.com/goplus/libc/test/cmd/test/ceilf"
	"github.com/goplus/libc/test/cmd/test/ceill"

	"github.com/goplus/libc/test/cmd/test/copysign"
	"github.com/goplus/libc/test/cmd/test/copysignl"

	"github.com/goplus/libc/test/cmd/test/erf"
	"github.com/goplus/libc/test/cmd/test/erff"
	"github.com/goplus/libc/test/cmd/test/erfl"

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

	atan.TestMain(t)
	atanf.TestMain(t)
	atanl.TestMain(t)

	atanh.TestMain(t)
	atanhf.TestMain(t)
	atanhl.TestMain(t)

	cbrt.TestMain(t)
	cbrtf.TestMain(t)
	cbrtl.TestMain(t)

	ceil.TestMain(t)
	ceilf.TestMain(t)
	ceill.TestMain(t)

	copysign.TestMain(t)
	copysignl.TestMain(t)

	erf.TestMain(t)
	erff.TestMain(t)
	erfl.TestMain(t)

	qsort.TestMain(t)
	random.TestMain(t)
	string_memcpy.TestMain(t)
	string_memset.TestMain(t)
	teststring.TestMain(t)
	strtod.TestMain(t)
}
