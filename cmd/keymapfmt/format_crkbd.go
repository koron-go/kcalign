package main

import "github.com/koron-go/kcalign"

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

func newFormatterCrkbd(param string) *kcalign.Formatter {
	return crkbdFormat(kcalign.Right)
}

func init() {
	registerFormatter("@crkbd", newFormatterCrkbd)
}
