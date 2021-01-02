package qmkjson

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestRead_crkbd(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/keymap_crkbd.json")
	if err != nil {
		t.Fatal(err)
	}
	var km Keymap
	err = json.Unmarshal(b, &km)
	if err != nil {
		t.Fatal(err)
	}
	if len(km.Layers) != 7 {
		t.Errorf("unexpected number of layers: want=%d got=%d", 7, len(km.Layers))
	}
	for i, l := range km.Layers {
		if len(l) != 42 {
			t.Errorf("unexpected keycodes in layer#%d: want=%d got=%d", i, 42, len(l))
		}
	}
}
