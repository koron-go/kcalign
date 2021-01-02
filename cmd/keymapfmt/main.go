package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/koron-go/kcalign"
	"github.com/koron-go/kcalign/internal/qmkjson"
)

func main() {
	var layerFormat string
	flag.StringVar(&layerFormat, "format", "@crkbd", `layer format`)
	flag.Parse()
	err := formatKeymap(os.Stdout, os.Stdin, layerFormat)
	if err != nil {
		log.Fatal(err)
	}
}

func formatKeymap(w io.Writer, r io.Reader, layerFormat string) error {
	f, err := loadFormat(layerFormat)
	if err != nil {
		return err
	}
	var km *qmkjson.Keymap
	err = json.NewDecoder(r).Decode(&km)
	if err != nil {
		return err
	}
	err = prettyKeymapJSON(w, f, km)
	if err != nil {
		return err
	}
	return nil
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
