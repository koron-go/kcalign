package klejson2_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/koron-go/kcalign/internal/klejson2"
)

func TestReadMetadata(t *testing.T) {
	for i, c := range []struct {
		json string
		want klejson2.Metadata
	}{
		{`[{}]`, klejson2.Metadata{
			Backcolor: klejson2.MustParseColor("#eeeeee"),
		}},
		{`[{"backcolor":"#aaaaaa"}]`, klejson2.Metadata{
			Backcolor: klejson2.MustParseColor("#aaaaaa"),
		}},
		{`[{"backcolor":"#ffffff"}]`, klejson2.Metadata{
			Backcolor: klejson2.MustParseColor("#ffffff"),
		}},
		{`[{}]`, klejson2.Metadata{
			Backcolor: klejson2.MustParseColor("#eeeeee"),
		}},
	} {
		l, err := klejson2.Read(strings.NewReader(c.json))
		if err != nil {
			t.Errorf("failed to read: i=%d c=%+v: %s", i, c, err)
			continue
		}
		got := l.Metadata
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("unmatch: i=%d c=%+v got=%+v", i, c, got)
			continue
		}
	}
}

func keyMerge(base, v klejson2.Key) klejson2.Key {
	k := base
	for i, l := range v.Legends {
		if reflect.ValueOf(l).IsZero() {
			continue
		}
		k.Legends[i] = l
	}
	if v.TextSize != 0 {
		k.TextSize = v.TextSize
	}
	if !reflect.ValueOf(v.TextColor).IsZero() {
		k.TextColor = v.TextColor
	}
	if !reflect.ValueOf(v.Color).IsZero() {
		k.Color = v.Color
	}
	if v.Width != 0 {
		k.Width = v.Width
	}
	if v.Width2 != 0 {
		k.Width2 = v.Width2
	}
	if v.Height != 0 {
		k.Height = v.Height
	}
	if v.Height2 != 0 {
		k.Height2 = v.Height2
	}
	if v.X != 0 {
		k.X = v.X
	}
	if v.X2 != 0 {
		k.X2 = v.X2
	}
	if v.Y != 0 {
		k.Y = v.Y
	}
	if v.Y2 != 0 {
		k.Y2 = v.Y2
	}
	if v.RotationAngle != 0 {
		k.RotationAngle = v.RotationAngle
	}
	if v.RotationX != 0 {
		k.RotationX = v.RotationX
	}
	if v.RotationY != 0 {
		k.RotationY = v.RotationY
	}
	if v.Profile != "" {
		k.Profile = v.Profile
	}
	if v.SwitchMount != "" {
		k.SwitchMount = v.SwitchMount
	}
	if v.SwitchBrand != "" {
		k.SwitchBrand = v.SwitchBrand
	}
	if v.SwitchType != "" {
		k.SwitchType = v.SwitchType
	}
	v.Ghost = k.Ghost
	v.Stepped = k.Stepped
	v.Nub = k.Nub
	v.Decal = k.Decal
	return k
}

func testKey(v klejson2.Key) klejson2.Key {
	return keyMerge(klejson2.DefaultKey, v)
}

func TestReadRows(t *testing.T) {
	for i, c := range []struct {
		json string
		want []klejson2.Row
	}{
		{`[{},[]]`, nil},
		{`[{},["foo"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "foo"}},
				}),
			},
		}},

		// multiple keys
		{`[{},["foo","bar"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "foo"}},
				}),
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "bar"}},
					X:       1,
				}),
			},
		}},
		{`[{},["foo","bar"],["baz","qux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "foo"}},
				}),
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "bar"}},
					X:       1,
				}),
			},
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "baz"}},
					Y:       1,
				}),
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "qux"}},
					X:       1,
					Y:       1,
				}),
			},
		}},

		// postion and width: "x", "w"
		{`[{},[{"x":0.25,"w":1.5},"foo","bar"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "foo"}},
					Width:   1.5,
					X:       0.25,
				}),
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "bar"}},
					X:       1.75,
				}),
			},
		}},

		// alignment: "a"
		{`[{},["foo\nbar\nbaz\nqux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{
						0: {Label: "foo"},
						2: {Label: "baz"},
						6: {Label: "bar"},
						8: {Label: "qux"},
					},
				}),
			},
		}},
		{`[{},[{"a":7},"foo\nbar\nbaz\nqux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{
						4: {Label: "foo"},
					},
				}),
			},
		}},

		// color: "c", "t"
		{`[{},[{"c":"#aaaaaa"},"foo"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{{Label: "foo"}},
					Color:   *klejson2.MustParseColor("#aaaaaa"),
				}),
			},
		}},
		{`[{},[{"t":"#333333,#666666,#999999,#cccccc"},"foo\nbar\nbaz\nqux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{
						0: {
							Label: "foo",
							Color: *klejson2.MustParseColor("#333333"),
						},
						2: {
							Label: "baz",
							Color: *klejson2.MustParseColor("#999999"),
						},
						6: {
							Label: "bar",
							Color: *klejson2.MustParseColor("#666666"),
						},
						8: {
							Label: "qux",
							Color: *klejson2.MustParseColor("#cccccc"),
						},
					},
					TextColor: *klejson2.MustParseColor("#333333"),
				}),
			},
		}},

		// text size
		{`[{},[{"fa":[1,2,0.5,1.5]},"foo\nbar\nbaz\nqux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{
						0: {
							Label: "foo",
							Size:  1,
						},
						2: {
							Label: "baz",
							Size:  0.5,
						},
						6: {
							Label: "bar",
							Size:  2,
						},
						8: {
							Label: "qux",
							Size:  1.5,
						},
					},
				}),
			},
		}},
		{`[{},[{"f2":1.5},"foo\nbar\nbaz\nqux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{
						0: {Label: "foo"},
						2: {
							Label: "baz",
							Size:  1.5,
						},
						6: {
							Label: "bar",
							Size:  1.5,
						},
						8: {
							Label: "qux",
							Size:  1.5,
						},
					},
				}),
			},
		}},
		{`[{},[{"f":1.5},"foo\nbar\nbaz\nqux"]]`, []klejson2.Row{
			{
				testKey(klejson2.Key{
					Legends: [12]klejson2.Legend{
						0: {Label: "foo"},
						2: {Label: "baz"},
						6: {Label: "bar"},
						8: {Label: "qux"},
					},
					TextSize: 1.5,
				}),
			},
		}},
	} {
		l, err := klejson2.Read(strings.NewReader(c.json))
		if err != nil {
			t.Errorf("failed to read: i=%d c=%+v: %s", i, c, err)
			continue
		}
		got := l.Rows
		if d := cmp.Diff(c.want, got); d != "" {
			t.Errorf("unmatch: i=%d c.json=%s -want +got\n%s", i, c.json, d)
			continue
		}
	}
}
