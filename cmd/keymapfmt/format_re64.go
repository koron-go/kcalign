package main

import "github.com/koron-go/kcalign"

func newFormatterRe64(param string) *kcalign.Formatter {
	return &kcalign.Formatter{
		Desc:  "Re64 formatter",
		Width: 10,

		Align: kcalign.RowAlign{
			Num:       8,
			TextAlign: kcalign.Center,
		},

		ExAligns: map[int]kcalign.RowAlign{

			0: {
				Num:       15,
				TextAlign: kcalign.Center,
				ExMargins: map[int]int{
					7: 44,
				},
			},

			1: {
				Num:       14,
				TextAlign: kcalign.Center,
				ExMargins: map[int]int{
					6: 42,
				},
				ExWidths: map[int]int{
					0:  17,
					13: 16,
				},
			},

			2: {
				Num:       13,
				TextAlign: kcalign.Center,
				ExMargins: map[int]int{
					6: 41,
				},
				ExWidths: map[int]int{
					0:  20,
					12: 25,
				},
			},

			3: {
				Num:       13,
				TextAlign: kcalign.Center,
				ExMargins: map[int]int{
					6: 42,
				},
				ExWidths: map[int]int{
					0:  24,
					11: 20,
				},
			},

			4: {
				Num:       11,
				TextAlign: kcalign.Center,
				ExMargins: map[int]int{
					3: 10,
					6: 2,
					9: 10,
				},
				ExWidths: map[int]int{
					0: 15,
					2: 15,
					3: 12,
					4: 13,
					5: 16,
					6: 16,
					7: 13,
					8: 12,
					9: 15,
				},
			},
		},
	}
}
