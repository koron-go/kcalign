package main

import "github.com/koron-go/kcalign"

func newFormatterDz60RgbV2(param string) *kcalign.Formatter {
	return &kcalign.Formatter{
		Desc:  "DZ60 RGB V2 formatter",
		Width: 10,

		Align: kcalign.RowAlign{
			Num:       14,
			TextAlign: kcalign.Center,
		},

		ExAligns: map[int]kcalign.RowAlign{

			0: {
				Num:       14,
				TextAlign: kcalign.Center,
				ExWidths: map[int]int{
					13: 23,
				},
			},

			1: {
				Num:       14,
				TextAlign: kcalign.Center,
				ExWidths: map[int]int{
					0:  16,
					13: 17,
				},
			},
			2: {
				Num:       13,
				TextAlign: kcalign.Center,
				ExWidths: map[int]int{
					0:  19,
					12: 25,
				},
			},

			3: {
				Num:       13,
				TextAlign: kcalign.Center,
				ExWidths: map[int]int{
					0:  25,
					10: 19,
				},
			},

			4: {
				Num:       9,
				TextAlign: kcalign.Center,
				ExWidths: map[int]int{
					0:  13,
					1:  13,
					2:  13,
					3:  69,
				},
			},

		},
	}
}
