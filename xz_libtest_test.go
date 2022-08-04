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

	"github.com/goplus/libc/test/cmd/test/expm1"
	"github.com/goplus/libc/test/cmd/test/expm1f"
	"github.com/goplus/libc/test/cmd/test/expm1l"

	"github.com/goplus/libc/test/cmd/test/fabs"
	"github.com/goplus/libc/test/cmd/test/fabsf"
	"github.com/goplus/libc/test/cmd/test/fabsl"

	"github.com/goplus/libc/test/cmd/test/fdim"
	"github.com/goplus/libc/test/cmd/test/fdiml"

	"github.com/goplus/libc/test/cmd/test/floor"
	"github.com/goplus/libc/test/cmd/test/floorf"
	"github.com/goplus/libc/test/cmd/test/floorl"

	"github.com/goplus/libc/test/cmd/test/fma"
	"github.com/goplus/libc/test/cmd/test/fmal"

	"github.com/goplus/libc/test/cmd/test/fmax"
	"github.com/goplus/libc/test/cmd/test/fmaxl"

	"github.com/goplus/libc/test/cmd/test/fmin"
	"github.com/goplus/libc/test/cmd/test/fminl"

	"github.com/goplus/libc/test/cmd/test/hypot"
	"github.com/goplus/libc/test/cmd/test/hypotl"

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

	expm1.TestMain(t)
	expm1f.TestMain(t)
	expm1l.TestMain(t)

	fabs.TestMain(t)
	fabsf.TestMain(t)
	fabsl.TestMain(t)

	fdim.TestMain(t)
	fdiml.TestMain(t)

	floor.TestMain(t)
	floorf.TestMain(t)
	floorl.TestMain(t)

	fma.TestMain(t)
	fmal.TestMain(t)

	fmax.TestMain(t)
	fmaxl.TestMain(t)

	fmin.TestMain(t)
	fminl.TestMain(t)

	hypot.TestMain(t)
	hypotl.TestMain(t)

	qsort.TestMain(t)
	random.TestMain(t)
	string_memcpy.TestMain(t)
	string_memset.TestMain(t)
	teststring.TestMain(t)
	strtod.TestMain(t)
}
