package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/koron-go/kcalign/internal/klejson"
)

type ledMapEntry struct {
	ID       string
	SwitchID string
	Rotate   float64
}

type ledMap map[string]ledMapEntry

func (m ledMap) put(records []string) error {
	if len(records) == 0 {
		return nil
	}
	if len(records) < 2 {
		return errors.New("require two or more records")
	}
	e := ledMapEntry{
		ID:       records[0],
		SwitchID: records[1],
	}
	if len(records) >= 3 {
		r, err := strconv.ParseFloat(records[2], 64)
		if err != nil {
			return fmt.Errorf("invalid rotate: %w", err)
		}
		e.Rotate = r
	}
	if _, ok := m[e.SwitchID]; ok {
		return fmt.Errorf("duplicated switch ID: %s", e.SwitchID)
	}
	m[e.SwitchID] = e
	return nil
}

func loadLEDMap(name string) (ledMap, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = '\t'
	r.Comment = '#'
	m := ledMap{}
	for {
		vv, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if len(vv) == 0 {
			continue
		}
		err = m.put(vv)
		if err != nil {
			return nil, fmt.Errorf("syntax error in LED map file %s: %w", name, err)
		}
	}
	return m, nil
}

func writeMappedLED(w io.Writer, l *klejson.Layout, keys klejson.Keys, mapFile string) error {
	m, err := loadLEDMap(mapFile)
	if err != nil {
		return err
	}
	for i, k := range keys {
		swid := fmt.Sprintf("%s%d", prefixSwitch, i+1)
		e, ok := m[swid]
		if !ok {
			continue
		}
		x := k.CX*unit.x + offLED.x + origin.x
		y := k.CY*unit.y + offLED.y + origin.y
		_, err := fmt.Fprintf(w, "%s\t%f\t%f\t%f\n", e.ID, x, y, e.Rotate)
		if err != nil {
			return err
		}
	}
	return nil
}
