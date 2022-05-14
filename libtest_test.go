package libc_test

import (
	"testing"

	"github.com/goplus/libc/test/cmd/test/string_memcpy"
	"github.com/goplus/libc/test/cmd/test/string_memset"
)

func TestFromTestdata(t *testing.T) {
	string_memcpy.TestMain(t)
	string_memset.TestMain(t)
}
