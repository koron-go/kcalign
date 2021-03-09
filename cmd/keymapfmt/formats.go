package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/koron-go/kcalign"
	"github.com/koron-go/kcalign/internal/qmkjson"
)

const defaultLayerFormat = "@oneitem"

var formats = map[string]func(string) *kcalign.Formatter{
	"@crkbd":   newFormatterCrkbd,
	"@oneitem": newFormatterOneitem,
	"@re64":    newFormatterRe64,
	"@uzu42":   newFormatterUzu42,

	// DZ60 RGB V1/V2, both supported
	"@dztech/dz60rgb": newFormatterDz60Rgb,
}

// detect layer format from qmkjson.Keymap.
func detectLayerFormat(km *qmkjson.Keymap) string {
	// FIXME: add custom detection algorithms at here.
	for k := range formats {
		if strings.HasPrefix(km.Keyboard, k[1:]) {
			return k
		}
	}
	return defaultLayerFormat
}

func loadFormat(name string) (*kcalign.Formatter, error) {
	if name == "" {
		return nil, errors.New("no format supplied")
	}
	var param string
	if n := strings.Index(name, ":"); n > 0 {
		name, param = name[:n], name[n+1:]
	}
	fn, ok := formats[name]
	if ok {
		return fn(param), nil
	}

	// load kcalign.Formatter from a file.
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var f *kcalign.Formatter
	err = json.Unmarshal(b, &f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
