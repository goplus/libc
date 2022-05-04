libc - Porting libc from C to Go
======

This is a subproject of the [c2go](https://github.com/goplus/c2go) project.

## What's our plan?

- First, port `printf` and its dependent C functions to Go. Stage: `Doing`, see [list of unfinished functions](https://github.com/goplus/libc/blob/musl-go/c2go_autogen.go).
- Second, port all `sqlite3` dependent C functions. Stage: `Planning`, see [dependent fuctions of sqlite3](https://github.com/goplus/sqlite/blob/main/c2go_autogen.go).
- Last, support most of C standard libaries and can import them by Go+. Stage: `Planning`.
