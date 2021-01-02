package kcalign

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// TextAlign specifies text alignment in column.
type TextAlign int

const (
	// Left is left alignment.
	Left TextAlign = iota
	// Right is right alignment.
	Right
	// Center is center alignment.
	Center
)

// QuoteType specifies runes to quote each items.
type QuoteType int

const (
	// Double uses dobule quote (`"`).
	Double QuoteType = iota
	// NoQuote doesn't use quote runes, write each items as is.
	NoQuote
)

// Align defines alignment property for each rows.
type Align struct {
	// Number of columns.
	Num int

	// Indent is left indent.
	Indent int

	// TextAlign is default alignment of coloumn.
	TextAlign TextAlign

	// ExMargin extends merge between columns.
	ExMargin map[int]int

	// ExWidth extends column width by position.
	ExWidth map[int]int

	// ExTextAligns extedns text alignment of columns.
	ExTextAligns map[int]TextAlign
}

func (a Align) textAlign(n int) TextAlign {
	ta, ok := a.ExTextAligns[n]
	if ok {
		return ta
	}
	return a.TextAlign
}

// Formatter provides format strign array, with alignment like key board.
type Formatter struct {
	// Width is width of a column. Two double quotes are included in count.
	// Minimal is 3, default is 10.
	Width int

	// Span is number of white spaces between columns. Minimal is 0.
	Span int

	// Quote is quote type for each data.
	Quote QuoteType

	// Align is default alignment.
	Align Align

	// ExAligns is exceptional alignments per lines.
	ExAligns map[int]Align
}

// Format writes formatted data by defined alignment.
func (f *Formatter) Format(w io.Writer, data []string) error {
	var lnum = 0
	for len(data) > 0 {
		a := f.align(lnum)
		if len(data) < a.Num {
			return fmt.Errorf("less data, want %d got %d at line %d", a.Num, len(data), lnum+1)
		}
		last := len(data) == a.Num
		err := f.format(w, a, data[0:a.Num], last)
		if err != nil {
			return fmt.Errorf("format failed at line %d: %w", lnum+1, err)
		}
		data = data[a.Num:]
		lnum++
	}
	return nil
}

// align gets an Align for a N'th row.
func (f *Formatter) align(n int) Align {
	a, ok := f.ExAligns[n]
	if ok {
		return a
	}
	return f.Align
}

const paddingStr = "                                                  "

// writePadding writes white spaces as padding.
func (f *Formatter) writePadding(w io.Writer, n int) error {
	if n == 0 {
		return nil
	}
	var s string
	if n <= len(paddingStr) {
		s = paddingStr[:n]
	} else {
		s = strings.Repeat(" ", n)
	}
	_, err := io.WriteString(w, s)
	return err
}

func (f *Formatter) quoteString(s string) string {
	switch f.Quote {
	default:
		fallthrough
	case Double:
		return strconv.Quote(s)
	case NoQuote:
		return s
	}
}

// columnWidth calculate column width requirement.
func (f *Formatter) columnWidth(a Align, n int) int {
	w, ok := a.ExWidth[n]
	if ok {
		return w
	}
	if f.Width < 3 {
		return 10
	}
	return f.Width
}

// columnPadding calculates paddings for a column.
func (f *Formatter) columnPadding(ta TextAlign, columnWidth int, s string) (left, right int) {
	switch ta {
	case Left:
		right = columnWidth - len(s)
		if right < 0 {
			right = 0
		}
	case Right:
		left = columnWidth - len(s)
		if left < 0 {
			left = 0
		}
	case Center:
		left = (columnWidth - len(s)) / 2
		if left < 0 {
			left = 0
		}
		right = columnWidth - len(s) - left
		if right < 0 {
			right = 0
		}
	}
	return left, right
}

// formatColumn format/output contents, paddings, and separator of a column.
func (f *Formatter) formatColumn(w io.Writer, s string, left, right int, lastCol bool) error {
	// write left padding.
	if err := f.writePadding(w, left); err != nil {
		return err
	}
	// write data.
	_, err := io.WriteString(w, s)
	if err != nil {
		return err
	}
	// write right padding.
	if err := f.writePadding(w, right); err != nil {
		return err
	}
	// tail comma.
	if !lastCol {
		const comma = ","
		if _, err := io.WriteString(w, comma); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) format(w io.Writer, a Align, data []string, lastRow bool) error {
	if a.Indent > 0 {
		err := f.writePadding(w, a.Indent)
		if err != nil {
			return err
		}
	}
	for i, d := range data {
		lastCol := lastRow && i+1 >= len(data)
		m, ok := a.ExMargin[i]
		if ok && m > 0 {
			if err := f.writePadding(w, m); err != nil {
				return err
			}
		}
		cw := f.columnWidth(a, i)
		s := f.quoteString(d)
		// calculate indents on left and right.
		padL, padR := f.columnPadding(a.textAlign(i), cw, s)
		if lastCol {
			padR = 0
		}
		// write left padding.
		if err := f.formatColumn(w, s, padL, padR, lastCol); err != nil {
			return err
		}
	}
	_, err := io.WriteString(w, "\n")
	return err
}
