# koron-go/kcalign

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/kcalign)](https://pkg.go.dev/github.com/koron-go/kcalign)
[![GoDoc](https://godoc.org/github.com/koron-go/kcalign?status.svg)](https://godoc.org/github.com/koron-go/kcalign)
[![Actions/Go](https://github.com/koron-go/kcalign/workflows/Go/badge.svg)](https://github.com/koron-go/kcalign/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/kcalign)](https://goreportcard.com/report/github.com/koron-go/kcalign)

Key code alignment library.  Designed to use with QMK keymap.json.

## keymapfmt - A tool to format QMK's keymap.json

kcalign includes a tool called keymapfmt. it formats `"layers"` property in QMK
keymap.json with layout which affected its physical key alignments.

Samples:

* for crkbd
    * [Input JSON](./cmd/keymapfmt/testdata/crkbd_in.json)
    * [Result JSON](./cmd/keymapfmt/testdata/crkbd_out.json)

### keymapfmt: How to install

```console
$ go install github.com/koron-go/kcalign/cmd/keymapfmt
```

### keymapfmt: How to use

Most simple usage is reading from STDIN, writing to STDOUT.

```
keymapfmt [-format {format}] < in.json > out.json
```

Where `{format}` accepts two types. First type is file name, it read
formatter.json from the file.
See [formats/\_schema\_.json](formats/_schema_.json) and `kcalign` code for
details of its schema.

Second type is pre-defined formats. It is form of `@{name}` or
`@{name}:{param}`. Currently, keymapfmt includes these pre-defined formats:

* `@crkbd` - [crkbd](https://github.com/foostan/crkbd/). no `{param}` support.
* `@re64` - [Re64](). no `{param}` support.

Extended usage is reading from file.

```
keymapfmt [-format {format}] [-inplace] {files...}
```

In this form, keymapfmt read JSON from `{files...}` and write all to STDOUT.

When `-inplace` option given, keymapfmt replaces contents of each `{files...}`
in place.
