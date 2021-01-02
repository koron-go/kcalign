package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/koron-go/kcalign"
)

func main() {
	err := run()
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
				6: 36,
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
					3: 4,
				},
			},
		},
	}
}

func run() error {
	var textAlign string
	flag.StringVar(&textAlign, "textalign", "", `text alignment: "left", "right", or "center"`)
	flag.Parse()

	var ta kcalign.TextAlign
	switch textAlign {
	case "left":
		ta = kcalign.Left
	case "right":
		ta = kcalign.Right
	case "center":
		ta = kcalign.Center
	default:
	}

	f := crkbdFormat(ta)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(f)
}
