package kcalign

// RowAlign defines alignment property for each rows.
type RowAlign struct {
	// Number of columns.
	Num int `json:"num"`

	// Indent is left indent.
	Indent int `json:"indent,omitempty"`

	// TextAlign is default alignment of coloumn.
	TextAlign TextAlign `json:"text_align,omitempty"`

	// ExMargins extends merge between columns.
	ExMargins map[int]int `json:"ex_margins,omitempty"`

	// ExWidths extends column width by position.
	ExWidths map[int]int `json:"ex_widths,omitempty"`

	// ExTextAligns extedns text alignment of columns.
	ExTextAligns map[int]TextAlign `json:"ex_text_aligns,omitempty"`
}

func (a RowAlign) textAlign(n int) TextAlign {
	ta, ok := a.ExTextAligns[n]
	if ok {
		return ta
	}
	return a.TextAlign
}
