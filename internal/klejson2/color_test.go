package klejson2_test

import (
	"encoding/json"
	"testing"

	"github.com/koron-go/kcalign/internal/klejson2"
)

func TestColorFormat(t *testing.T) {
	for i, c := range []struct {
		color  klejson2.Color
		expect string
	}{
		{klejson2.Color{R: 0x00, G: 0x00, B: 0x00}, "#000000"},
		{klejson2.Color{R: 0xff, G: 0x00, B: 0x00}, "#ff0000"},
		{klejson2.Color{R: 0x00, G: 0xff, B: 0x00}, "#00ff00"},
		{klejson2.Color{R: 0x00, G: 0x00, B: 0xff}, "#0000ff"},
		{klejson2.Color{R: 0x12, G: 0x34, B: 0x56}, "#123456"},
		{klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}, "#abcdef"},
	} {
		got := c.color.Format()
		if got != c.expect {
			t.Errorf("unexpected: i=%d c=%+v got=%s", i, c, got)
		}
	}
}

func TestParseColor(t *testing.T) {
	for i, c := range []struct {
		s      string
		expect klejson2.Color
	}{
		{"#000000", klejson2.Color{R: 0x00, G: 0x00, B: 0x00}},
		{"#ff0000", klejson2.Color{R: 0xff, G: 0x00, B: 0x00}},
		{"#00ff00", klejson2.Color{R: 0x00, G: 0xff, B: 0x00}},
		{"#0000ff", klejson2.Color{R: 0x00, G: 0x00, B: 0xff}},
		{"#123456", klejson2.Color{R: 0x12, G: 0x34, B: 0x56}},
		{"#abcdef", klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}},
		{"#ABCDEF", klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}},
		{"#AbCdEf", klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}},

		{"#000", klejson2.Color{R: 0x00, G: 0x00, B: 0x00}},
		{"#f00", klejson2.Color{R: 0xff, G: 0x00, B: 0x00}},
		{"#0f0", klejson2.Color{R: 0x00, G: 0xff, B: 0x00}},
		{"#00f", klejson2.Color{R: 0x00, G: 0x00, B: 0xff}},
		{"#333", klejson2.Color{R: 0x33, G: 0x33, B: 0x33}},
		{"#fff", klejson2.Color{R: 0xff, G: 0xff, B: 0xff}},
	} {
		got, err := klejson2.ParseColor(c.s)
		if err != nil {
			t.Errorf("parse failure: i=%d c=%+v: %s", i, c, err)
			continue
		}
		if got != c.expect {
			t.Errorf("unexpected: i=%d c=%+v got=%+v", i, c, got)
			continue
		}
	}
}

func TestParseColorError(t *testing.T) {
	for i, c := range []struct {
		s      string
		expect string
	}{
		{"", "color string should start with a hash (#): "},
		{"abc", "color string should start with a hash (#): abc"},
		{"#", "invalid length of color string: #"},
		{"#1", "invalid length of color string: #1"},
		{"#12", "invalid length of color string: #12"},
		{"#1234", "invalid length of color string: #1234"},
		{"#12345", "invalid length of color string: #12345"},
		{"#1234567", "invalid length of color string: #1234567"},
		{"#Z00", `syntax error at RED (first): strconv.ParseUint: parsing "Z": invalid syntax`},
		{"#0Z0", `syntax error at GREEN (second): strconv.ParseUint: parsing "Z": invalid syntax`},
		{"#00Z", `syntax error at BLUE (third): strconv.ParseUint: parsing "Z": invalid syntax`},
		{"#ZZ0000", `syntax error at RED (first): strconv.ParseUint: parsing "ZZ": invalid syntax`},
		{"#00ZZ00", `syntax error at GREEN (second): strconv.ParseUint: parsing "ZZ": invalid syntax`},
		{"#0000ZZ", `syntax error at BLUE (third): strconv.ParseUint: parsing "ZZ": invalid syntax`},
	} {
		got, err := klejson2.ParseColor(c.s)
		if err == nil {
			t.Errorf("unexpected succeed: i=%d c=%+v got=%+v", i, c, got)
			continue
		}
		if err.Error() != c.expect {
			t.Errorf("unexpected error: i=%d c=%+v got=%+v", i, c, err.Error())
			continue
		}
	}
}

func TestColorMarshalJSON(t *testing.T) {
	for i, c := range []struct {
		color  klejson2.Color
		expect string
	}{
		{klejson2.Color{R: 0x00, G: 0x00, B: 0x00}, `"#000000"`},
		{klejson2.Color{R: 0xff, G: 0x00, B: 0x00}, `"#ff0000"`},
		{klejson2.Color{R: 0x00, G: 0xff, B: 0x00}, `"#00ff00"`},
		{klejson2.Color{R: 0x00, G: 0x00, B: 0xff}, `"#0000ff"`},
		{klejson2.Color{R: 0x12, G: 0x34, B: 0x56}, `"#123456"`},
		{klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}, `"#abcdef"`},
	} {
		b, err := json.Marshal(c.color)
		if err != nil {
			t.Errorf("unexpected failure: i=%d c=%+v: %s", i, c, err)
			continue
		}
		got := string(b)
		if got != c.expect {
			t.Errorf("unexpected JSON: i=%d c=%+v got=%s", i, c, got)
		}
	}
}

func TestColorUnmarshalJSON(t *testing.T) {
	for i, c := range []struct {
		s      string
		expect klejson2.Color
	}{
		{`"#000000"`, klejson2.Color{R: 0x00, G: 0x00, B: 0x00}},
		{`"#ff0000"`, klejson2.Color{R: 0xff, G: 0x00, B: 0x00}},
		{`"#00ff00"`, klejson2.Color{R: 0x00, G: 0xff, B: 0x00}},
		{`"#0000ff"`, klejson2.Color{R: 0x00, G: 0x00, B: 0xff}},
		{`"#123456"`, klejson2.Color{R: 0x12, G: 0x34, B: 0x56}},
		{`"#abcdef"`, klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}},
		{`"#ABCDEF"`, klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}},
		{`"#AbCdEf"`, klejson2.Color{R: 0xab, G: 0xcd, B: 0xef}},

		{`"#000"`, klejson2.Color{R: 0x00, G: 0x00, B: 0x00}},
		{`"#f00"`, klejson2.Color{R: 0xff, G: 0x00, B: 0x00}},
		{`"#0f0"`, klejson2.Color{R: 0x00, G: 0xff, B: 0x00}},
		{`"#00f"`, klejson2.Color{R: 0x00, G: 0x00, B: 0xff}},
		{`"#333"`, klejson2.Color{R: 0x33, G: 0x33, B: 0x33}},
		{`"#fff"`, klejson2.Color{R: 0xff, G: 0xff, B: 0xff}},
	} {
		var got klejson2.Color
		err := json.Unmarshal([]byte(c.s), &got)
		if err != nil {
			t.Errorf("unexpected failure: i=%d c=%+v: %s", i, c, err)
			continue
		}
		if got != c.expect {
			t.Errorf("unexpected Color: i=%d c=%+v got=%+v", i, c, got)
			continue
		}
	}
}
