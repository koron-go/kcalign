package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"

	"github.com/koron-go/kcalign"
	"github.com/koron-go/kcalign/internal/qmkjson"
)

func main() {
	err := formatKeymap(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

func crkbdFormat(ta kcalign.TextAlign) *kcalign.Formatter {
	return &kcalign.Formatter{
		Desc:  "crkbd layout",
		Width: 10,
		Span:  0,
		//Quote: kcalign.Double,
		Align: kcalign.RowAlign{
			Num:       12,
			TextAlign: ta,
			ExMargins: map[int]int{
				6: 34,
			},
		},
		ExAligns: map[int]kcalign.RowAlign{
			3: {
				Num:       6,
				Indent:    44,
				TextAlign: ta,
				ExWidths: map[int]int{
					2: 15,
					3: 15,
				},
				ExMargins: map[int]int{
					3: 2,
				},
			},
		},
	}
}

func formatKeymap(w io.Writer, r io.Reader) error {
	var km *qmkjson.Keymap
	err := json.NewDecoder(r).Decode(&km)
	if err != nil {
		return err
	}

	// FIXME: make replacable
	f := crkbdFormat(kcalign.Right)

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
