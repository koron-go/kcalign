package kcalign

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func loadItems(t *testing.T, name string) []string {
	t.Helper()
	b, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	var data []string
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func loadString(t *testing.T, name string) string {
	t.Helper()
	b, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func format1(ta TextAlign) *Formatter {
	return &Formatter{
		Width: 10,
		Span:  0,
		//Quote: None,
		Align: Align{
			Num:       12,
			TextAlign: ta,
			ExMargin: map[int]int{
				6: 36,
			},
		},
		ExAligns: map[int]Align{
			3: {
				Num:       6,
				Indent:    44,
				TextAlign: ta,
				ExWidth: map[int]int{
					2: 15,
					3: 15,
				},
				ExMargin: map[int]int{
					3: 4,
				},
			},
		},
	}
}

func formatCheck(t *testing.T, data []string, f *Formatter, want string) {
	t.Helper()
	bb := &bytes.Buffer{}
	err := f.Format(bb, data)
	if err != nil {
		t.Fatal(err)
	}
	got := bb.String()
	if got != want {
		t.Errorf("format unmatch:\nwant=%q\ngot=%q", got, want)
	}
}

func Test_crkbd(t *testing.T) {
	data := loadItems(t, "testdata/crkbd.json")

	f1 := format1(Left)
	want1 := loadString(t, "testdata/crkbd_left.want")
	formatCheck(t, data, f1, want1)

	f2 := format1(Right)
	want2 := loadString(t, "testdata/crkbd_right.want")
	formatCheck(t, data, f2, want2)

	f3 := format1(Center)
	want3 := loadString(t, "testdata/crkbd_center.want")
	formatCheck(t, data, f3, want3)
}
