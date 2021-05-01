package main

import "github.com/koron-go/kcalign"

func newFormatterUzu42(param string) *kcalign.Formatter {
	ta := defaultTextAlign
	return &kcalign.Formatter{
		Desc:  "uzu42 layout",
		Width: 11,
		Span:  0,
		//Quote: kcalign.Double,
		Align: kcalign.RowAlign{
			Num:       10,
			TextAlign: ta,
			ExMargins: map[int]int{
				5: 24,
			},
		},
		ExAligns: map[int]kcalign.RowAlign{
			3: {
				Num:       12,
				TextAlign: ta,
			},
		},
	}
}

func init() {
	registerFormatter("@uzu42", newFormatterUzu42)
}
