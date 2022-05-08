libc - Porting libc from C to Go
======

This is a subproject of [the c2go project](https://github.com/goplus/c2go).

## What's our plan?

- First, port `printf` and its dependent C functions to Go. Stage: `Almost Done`, see [supported C standard libary functions](https://github.com/goplus/libc/blob/v0.2.1/c2go.pub).
- Second, support libc in [the Go+ language](https://github.com/goplus/gop). Stage: `Doing`.
- Third, port all `sqlite3` dependent C functions. Stage: `Planning`, see [sqlite3 dependent fuctions](https://github.com/goplus/sqlite/blob/main/c2go_autogen.go).
- Last, support most of C standard libaries and can import them by Go+. Stage: `Planning`.
