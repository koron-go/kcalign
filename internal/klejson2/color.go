package klejson2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

var _ json.Marshaler = Color{}

var invalidColor = Color{}

func ParseColor(s string) (Color, error) {
	if !strings.HasPrefix(s, "#") {
		return invalidColor, fmt.Errorf("color string should start with a hash (#): %s", s)
	}
	var r, g, b uint8
	v := s[1:]
	switch len(v) {
	case 3:
		r64, err := strconv.ParseUint(v[0:1], 16, 8)
		if err != nil {
			return invalidColor, fmt.Errorf("syntax error at RED (first): %s", err)
		}
		r = uint8(r64) * 0x11
		g64, err := strconv.ParseUint(v[1:2], 16, 8)
		if err != nil {
			return invalidColor, fmt.Errorf("syntax error at GREEN (second): %s", err)
		}
		g = uint8(g64) * 0x11
		b64, err := strconv.ParseUint(v[2:3], 16, 8)
		if err != nil {
			return invalidColor, fmt.Errorf("syntax error at BLUE (third): %s", err)
		}
		b = uint8(b64) * 0x11

	case 6:
		r64, err := strconv.ParseUint(v[0:2], 16, 8)
		if err != nil {
			return invalidColor, fmt.Errorf("syntax error at RED (first): %s", err)
		}
		r = uint8(r64)
		g64, err := strconv.ParseUint(v[2:4], 16, 8)
		if err != nil {
			return invalidColor, fmt.Errorf("syntax error at GREEN (second): %s", err)
		}
		g = uint8(g64)
		b64, err := strconv.ParseUint(v[4:6], 16, 8)
		if err != nil {
			return invalidColor, fmt.Errorf("syntax error at BLUE (third): %s", err)
		}
		b = uint8(b64)

	default:
		return invalidColor, fmt.Errorf("invalid length of color string: %s", s)
	}
	return Color{R: r, G: g, B: b}, nil
}

func mustParseColor(s string) *Color {
	c, err := ParseColor(s)
	if err != nil {
		panic(err)
	}
	return &c
}

func (c Color) Format() string {
	bb := bytes.NewBuffer(make([]byte, 0, 7))
	c.write(bb)
	return bb.String()
}

func (c Color) write(w io.Writer) error {
	_, err := fmt.Fprintf(w, "#%02x%02x%02x", c.R, c.G, c.B)
	return err
}

func (c Color) MarshalJSON() ([]byte, error) {
	bb := bytes.NewBuffer(make([]byte, 0, 7+2))
	bb.WriteRune('"')
	c.write(bb)
	bb.WriteRune('"')
	return bb.Bytes(), nil
}

var _ json.Unmarshaler = (*Color)(nil)

func (c *Color) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) < 2 || !strings.HasPrefix(s, `"`) || !strings.HasSuffix(s, `"`) {
		return fmt.Errorf("invalid format, string required: %s", s)
	}
	v, err := ParseColor(s[1 : len(s)-1])
	if err != nil {
		return err
	}
	*c = v
	return nil
}

type ColorList []Color

func (l ColorList) MarshalJSON() ([]byte, error) {
	bb := bytes.NewBuffer(make([]byte, 0, len(l)*8+1))
	bb.WriteRune('"')
	for i, c := range l {
		if i > 0 {
			bb.WriteRune(',')
		}
		c.write(bb)
	}
	bb.WriteRune('"')
	return bb.Bytes(), nil
}

func (l *ColorList) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) < 2 || !strings.HasPrefix(s, `"`) || !strings.HasSuffix(s, `"`) {
		return fmt.Errorf("invalid format, string required: %s", s)
	}
	ss := strings.Split(s[1:len(s)-1], ",")
	cc := make(ColorList, 0, len(ss))
	for i, t := range ss {
		c, err := ParseColor(strings.TrimSpace(t))
		if err != nil {
			return fmt.Errorf("failed at #%d: %w", i, err)
		}
		cc = append(cc, c)
	}
	*l = append(*l, cc...)
	return nil
}
