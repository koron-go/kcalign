# koron-go/kcalign

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/kcalign)](https://pkg.go.dev/github.com/koron-go/kcalign)
[![GoDoc](https://godoc.org/github.com/koron-go/kcalign?status.svg)](https://godoc.org/github.com/koron-go/kcalign)
[![Actions/Go](https://github.com/koron-go/kcalign/workflows/Go/badge.svg)](https://github.com/koron-go/kcalign/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/kcalign)](https://goreportcard.com/report/github.com/koron-go/kcalign)

Key code alignment library.  Designed to use with QMK keymap.json.

## keymapfmt

kcalign includes a tool called keymapfmt. it formats `"layers"` property in QMK
keymap.json with layout which affected its phisical key alignments.

Currently it supports crkbd's layout only.

### How to get keymapfmt

```console
$ go install github.com/koron-go/kcalign/cmd/keymapfmt
```

TO BE WRITTEN...
