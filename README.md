libc - Porting libc from C to Go
======

[![Build Status](https://github.com/goplus/libc/actions/workflows/go.yml/badge.svg)](https://github.com/goplus/libc/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/goplus/libc)](https://goreportcard.com/report/github.com/goplus/libc)
[![GitHub release](https://img.shields.io/github/v/tag/goplus/libc.svg?label=release)](https://github.com/goplus/libc/releases)
[![Coverage Status](https://codecov.io/gh/goplus/libc/branch/musl-go/graph/badge.svg)](https://codecov.io/gh/goplus/libc)
[![GoDoc](https://pkg.go.dev/badge/github.com/goplus/libc.svg)](https://pkg.go.dev/mod/github.com/goplus/libc)

This is a subproject of [the c2go project](https://github.com/goplus/c2go).

## What's our plan?

- First, port `printf` and its dependent C functions to Go. Stage: `Almost Done`, see [supported C standard libary functions](https://github.com/goplus/libc/blob/musl-go/c2go.pub).
- Second, support libc in [the Go+ language](https://github.com/goplus/gop). Stage: `Doing`.
- Third, port all `sqlite3` dependent C functions. Stage: `Planning`, see [sqlite3 dependent fuctions](https://github.com/goplus/sqlite/blob/main/c2go_autogen.go).
- Last, support most of C standard libaries and can import them by Go+. Stage: `Planning`.
