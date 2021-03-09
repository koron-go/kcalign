package main

import "github.com/koron-go/kcalign"

func uzu42Format(ta kcalign.TextAlign) *kcalign.Formatter {
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

func newFormatterUzu42(param string) *kcalign.Formatter {
	return uzu42Format(kcalign.Right)
}
