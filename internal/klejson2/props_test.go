package klejson2_test

import (
	"encoding/json"
	"testing"

	"github.com/koron-go/kcalign/internal/klejson2"
)

func json2props(t *testing.T, s string) klejson2.Props {
	t.Helper()
	var p klejson2.Props
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		t.Fatalf("failed to JSON unmarshal: %s", err)
	}
	return p
}

func props2json(t *testing.T, p klejson2.Props) string {
	t.Helper()
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("failed to JSON marshal: %s", err)
	}
	return string(b)
}

func TestPropsMerge(t *testing.T) {
	for i, c := range []struct {
		b   string
		n   string
		exp string
	}{
		{`{}`, `{}`, `{}`},
		{`{"x":123.4}`, `{}`, `{"x":123.4}`},
		{`{"x":123.4}`, `{"x":0}`, `{"x":0}`},
		{`{"x":123.4}`, `{"y":456.7}`, `{"x":123.4,"y":456.7}`},
		{`{"x":123.4}`, `{"w":1.5}`, `{"x":123.4,"w":1.5}`},
		{`{"x":123.4}`, `{"w":1.5,"r":90}`, `{"x":123.4,"r":90,"w":1.5}`},
	} {
		bp := json2props(t, c.b)
		np := json2props(t, c.n)
		bp.Merge(np)
		got := props2json(t, bp)
		if got != c.exp {
			t.Errorf("unmatch props: i=%d c=%+v got=%s", i, c, got)
		}
	}
}
