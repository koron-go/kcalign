package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/koron-go/kcalign"
)

func check(t *testing.T, in, out, layout string) {
	t.Helper()
	fin, err := os.Open(in)
	if err != nil {
		t.Errorf("failed to open/read in-file: %s", err)
		return
	}
	bb := &bytes.Buffer{}
	err = formatKeymap(bb, fin, layout)
	fin.Close()
	if err != nil {
		t.Errorf("failed to format: %s", err)
		return
	}
	bout, err := ioutil.ReadFile(out)
	if err != nil {
		t.Errorf("failed to open/read out-file: %s", err)
		return
	}
	want := string(bout)
	got := bb.String()
	if d := cmp.Diff(want, got); d != "" {
		t.Errorf("unmatch result: -want +got\n%s", d)
		return
	}
}

func TestCrkbd(t *testing.T) {
	defaultTextAlign = kcalign.Right
	check(t, "testdata/crkbd_in.json", "testdata/crkbd_out.json", "@crkbd")
}
