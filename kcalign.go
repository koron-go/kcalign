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

func (f *Formatter) align(n int) Align {
	a, ok := f.ExAligns[n]
	if ok {
		return a
	}
	return f.Align
}

func (f *Formatter) width() int {
	if f.Width < 3 {
		return 10
	}
	return f.Width
}

const padding = "                                                  "

func (f *Formatter) padding(n int) string {
	if n == 0 {
		return ""
	}
	if n <= len(padding) {
		return padding[:n]
	}
	return strings.Repeat(" ", n)
}

func (f *Formatter) writePadding(w io.Writer, n int) error {
	if n == 0 {
		return nil
	}
	_, err := io.WriteString(w, f.padding(n))
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


func (f *Formatter) format(w io.Writer, a Align, data []string, last bool) error {
	cwdefault := f.width()
	if a.Indent > 0 {
		err := f.writePadding(w, a.Indent)
		if err != nil {
			return err
		}
	}
	for i, d := range data {
		m, ok := a.ExMargin[i]
		if ok && m > 0 {
			err := f.writePadding(w, m)
			if err != nil {
				return err
			}
		}
		cw, ok := a.ExWidth[i]
		if !ok {
			cw = cwdefault
		}
		s := f.quoteString(d)
		// calculate indents on left and right.
		var pl, pr int
		switch a.textAlign(i) {
		case Left:
			pr = cw - len(s)
			if pr < 0 {
				pr = 0
			}
		case Right:
			pl = cw - len(s)
			if pl < 0 {
				pl = 0
			}
		case Center:
			pl = (cw - len(s)) / 2
			if pl < 0 {
				pl = 0
			}
			pr = cw - len(s) - pl
			if pr < 0 {
				pr = 0
			}
		}
		if last && i+1 >= len(data) {
			pr = 0
		}
		// write left indent.
		if pl > 0 {
			_, err := io.WriteString(w, f.padding(pl))
			if err != nil {
				return err
			}
		}
		// write data.
		_, err := io.WriteString(w, s)
		if err != nil {
			return err
		}
		// write right indent.
		if pr > 0 {
			_, err := io.WriteString(w, f.padding(pr))
			if err != nil {
				return err
			}
		}
		// tail comma.
		if !last || i+1 < len(data) {
			const comma = ","
			_, err := io.WriteString(w, comma)
			if err != nil {
				return err
			}
		}
	}
	_, err := io.WriteString(w, "\n")
	return err
}
