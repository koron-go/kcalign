package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/koron-go/kcalign"
	"github.com/koron-go/kcalign/internal/qmkjson"
)

const defaultLayerFormat = "@oneitem"

var defaultTextAlign = kcalign.Center

var formatters = map[string]func(string) *kcalign.Formatter{}

func registerFormatter(name string, factory func(string) *kcalign.Formatter) {
	_, has := formatters[name]
	if has {
		panic(fmt.Sprintf("formatter for %q have been registered alreadly", name))
	}
	formatters[name] = factory
}

// detect layer format from qmkjson.Keymap.
func detectLayerFormat(km *qmkjson.Keymap) string {
	// FIXME: add custom detection algorithms at here.
	for k := range formatters {
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
	fn, ok := formatters[name]
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
