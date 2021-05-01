package main

import "github.com/koron-go/kcalign"

// newFormatterDz60Rgb creates a fromatter for DZ60 RGB  V1/V2 (supported both)
func newFormatterDz60Rgb(param string) *kcalign.Formatter {
	return &kcalign.Formatter{
		Desc:  "DZ60RGB formatter",
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

func init() {
	registerFormatter("@dz60rgb", newFormatterDz60Rgb)
}
