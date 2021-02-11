package klejson

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Layout struct {
	Name      string
	Author    string
	Notes     string
	Backcolor string
	Radii     string
	CSS       string

	Rows []Row
}

type Row []Key

func Parse(r io.Reader) (*Layout, error) {
	raw := make([]interface{}, 0, 10)
	d := json.NewDecoder(r)
	d.UseNumber()
	err := d.Decode(&raw)
	if err != nil {
		return nil, err
	}
	var l Layout
	kctx := Key{
		W: 1,
		H: 1,
	} // key context
	for _, v := range raw {
		if m, ok := v.(map[string]interface{}); ok {
			err := l.parseProps(m)
			if err != nil {
				return nil, err
			}
			continue
		}
		if a, ok := v.([]interface{}); ok {
			err := l.parseRow(kctx, a)
			if err != nil {
				return nil, err
			}
			kctx.X = 0
			kctx.Y++
		}
	}
	return &l, nil
}

func (l *Layout) parseProps(m map[string]interface{}) error {
	for k, v := range m {
		switch k {
		case "name":
			l.Name = fmt.Sprint(v)
		case "author":
			l.Author = fmt.Sprint(v)
		case "notes":
			l.Notes = fmt.Sprint(v)
		case "backcolor":
			l.Backcolor = fmt.Sprint(v)
		case "radii":
			l.Radii = fmt.Sprint(v)
		case "css":
			l.CSS = fmt.Sprint(v)
		default:
			log.Printf("ignore unkonwn property: %s", k)
		}
	}
	return nil
}

func (l *Layout) parseRow(kctx Key, row []interface{}) error {
	keyRow := make(Row, 0, len(row))
	for _, kr := range row {
		m, ok := kr.(map[string]interface{})
		if !ok {
			keyRow = append(keyRow, newKey(kctx.X, kctx.Y, kctx.W, kctx.H))
			kctx.X += kctx.W
			kctx.W = 1
			kctx.H = 1
			continue
		}
		k, err := l.parseKeyCtx(kctx, m)
		if err != nil {
			return err
		}
		kctx = k
	}
	l.Rows = append(l.Rows, keyRow)
	return nil
}

func toFloat64(v interface{}) (float64, error) {
	n, ok := v.(json.Number)
	if !ok {
		return 0, fmt.Errorf("not json.Number: %[1]v (%[1]T)", v)
	}
	return n.Float64()
}

func (l *Layout) parseKeyCtx(kctx Key, m map[string]interface{}) (Key, error) {
	kctx.W = 1
	kctx.H = 1
	for k, v := range m {
		switch k {
		case "x":
			x, err := toFloat64(v)
			if err != nil {
				return kctx, fmt.Errorf(`"x" with invalid value: %w`, err)
			}
			kctx.X += x
		case "y":
			y, err := toFloat64(v)
			if err != nil {
				return kctx, fmt.Errorf(`"y" with invalid value: %w`, err)
			}
			kctx.Y += y
		case "w":
			w, err := toFloat64(v)
			if err != nil {
				return kctx, fmt.Errorf(`"w" with invalid value: %w`, err)
			}
			kctx.W = w
		case "h":
			h, err := toFloat64(v)
			if err != nil {
				return kctx, fmt.Errorf(`"h" with invalid value: %w`, err)
			}
			kctx.H = h
		}
	}
	return kctx, nil
}

func (l *Layout) Keys() Keys {
	n := 0
	for _, row := range l.Rows {
		n += len(row)
	}
	if n == 0 {
		return nil
	}
	keys := make(Keys, 0, n)
	for _, row := range l.Rows {
		keys = append(keys, row...)
	}
	return keys
}
