package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/koron-go/kcalign"
)

var formats = map[string]func(string) *kcalign.Formatter{
	"@crkbd": newFormatterCrkbd,
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
