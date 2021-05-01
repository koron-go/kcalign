package main

import "github.com/koron-go/kcalign"

func newFormatterOneitem(param string) *kcalign.Formatter {
	ta := defaultTextAlign
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

func init() {
	registerFormatter("@oneitem", newFormatterOneitem)
}
