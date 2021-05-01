package main

import "github.com/koron-go/kcalign"

func newFormatterHHKB(param string) *kcalign.Formatter {
	ta := defaultTextAlign
	return &kcalign.Formatter{
		Desc:  "Happy Hacking Keybaord layout",
		Width: 11,

		Align: kcalign.RowAlign{
			Num:       15,
			TextAlign: ta,
		},

		ExAligns: map[int]kcalign.RowAlign{
			1: {
				Num:       14,
				TextAlign: ta,
				ExWidths: map[int]int{
					0:  17,
					13: 17,
				},
			},
			2: {
				Num:       13,
				TextAlign: ta,
				ExWidths: map[int]int{
					0:  20,
					12: 26,
				},
			},
			3: {
				Num:       13,
				TextAlign: ta,
				ExWidths: map[int]int{
					0:  26,
					11: 20,
				},
			},
			4: {
				Num:       5,
				TextAlign: ta,
				ExMargins: map[int]int{
					0: 17,
				},
				ExWidths: map[int]int{
					1: 17,
					2: 71,
					3: 17,
				},
			},
		},
	}
}

func init() {
	registerFormatter("@hhkb", newFormatterHHKB)
}
