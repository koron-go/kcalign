package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/koron-go/kcalign"
	"github.com/koron-go/kcalign/internal/qmkjson"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var layerFormat string
	var inPlace bool
	flag.StringVar(&layerFormat, "format", "", `layer format (default: auto detect)`)
	flag.BoolVar(&inPlace, "inplace", false, `rewrite JSON files in place`)
	flag.Parse()

	if flag.NArg() == 0 {
		if inPlace {
			log.Printf("ignore -inplace when read stdin")
		}
		return formatKeymap(os.Stdout, os.Stdin, layerFormat)
	}

	// mode: not in place
	if !inPlace {
		for _, in := range flag.Args() {
			f, err := os.Open(in)
			if err != nil {
				return err
			}
			err = formatKeymap(os.Stdout, f, layerFormat)
			f.Close()
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, inOut := range flag.Args() {
		f, err := os.Open(inOut)
		if err != nil {
			return err
		}
		bb := &bytes.Buffer{}
		err = formatKeymap(bb, f, layerFormat)
		f.Close()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(inOut, bb.Bytes(), 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

func formatKeymap(w io.Writer, r io.Reader, layerFormat string) error {
	var km *qmkjson.Keymap
	err := json.NewDecoder(r).Decode(&km)
	if err != nil {
		return err
	}
	// detect layer format.
	if layerFormat == "" {
		layerFormat = detectLayerFormat(km)
	}
	f, err := loadFormat(layerFormat)
	if err != nil {
		return err
	}
	err = prettyKeymapJSON(w, f, km)
	if err != nil {
		return err
	}
	return nil
}

// detect layer format from qmkjson.Keymap.
func detectLayerFormat(km *qmkjson.Keymap) string {
	if strings.HasPrefix(km.Keyboard, "crkbd") {
		return "@crkbd"
	}
	if strings.HasPrefix(km.Keyboard, "re64") {
		return "@re64"
	}
	return defaultLayerFormat
}

func prettyKeymapJSON(w io.Writer, f *kcalign.Formatter, km *qmkjson.Keymap) error {
	bb := &bytes.Buffer{}
	enc := json.NewEncoder(bb)
	enc.SetIndent("", "  ")
	err := enc.Encode(km)
	if err != nil {
		return err
	}

	var x int
	for {
		l, err := bb.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				_, err := io.WriteString(w, l)
				return err
			}
			return err
		}
		const marker = "    \"QMKJSON_LAYER\""
		if !strings.HasPrefix(l, marker) {
			_, err := io.WriteString(w, l)
			if err != nil {
				return err
			}
			continue
		}
		_, err = io.WriteString(w, "    [\n")
		if err != nil {
			return err
		}
		err = f.FormatIndent(w, 6, km.Layers[x])
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "    ]"+l[len(marker):])
		if err != nil {
			return err
		}
		x++
	}
}
