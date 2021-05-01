package main

import "github.com/koron-go/kcalign"

func oneitemFormat(ta kcalign.TextAlign) *kcalign.Formatter {
	return &kcalign.Formatter{
		Desc:  "one item, one line layout",
		Width: 10,
		Span:  0,
		Align: kcalign.RowAlign{
			Num:       1,
			TextAlign: ta,
		},
	}
}

func newFormatterOneitem(param string) *kcalign.Formatter {
	return oneitemFormat(kcalign.Right)
}

func init() {
	registerFormatter("@oneitem", newFormatterOneitem)
}
