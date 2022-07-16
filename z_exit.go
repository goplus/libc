package libc

import "os"

func Abort() {
	os.Exit(127)
}
