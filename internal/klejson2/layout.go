package klejson2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Layout struct {
	Metadata
	Rows []Row
}

func Read(r io.Reader) (*Layout, error) {
	var raw []interface{}
	err := json.NewDecoder(r).Decode(&raw)
	if err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, errors.New("no elements in layout JSON")
	}
	// read Metadata
	raw0, ok := raw[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected error, 1st element should be object but got: %T", raw[0])
	}
	var md Metadata
	err = jsonReparse(raw0, &md)
	if err != nil {
		return nil, fmt.Errorf("invalid layout metadata: %w", err)
	}
	// read Rows
	var rows []Row
	var curr Props = DefaultProps
	for i, r := range raw[1:] {
		list, ok := r.([]interface{})
		if !ok {
			return nil, fmt.Errorf("non-row found at #%d: %T", i, r)
		}
		var row Row
		for j, el := range list {
			switch v := el.(type) {
			case map[string]interface{}:
				var p Props
				err := jsonReparse(v, &p)
				if err != nil {
					return nil, fmt.Errorf("invalid key properties at #%d,%d: %w", i, j, err)
				}
				curr = mergeProps(curr, p)
			case string:
				k, err := newKey(curr, v)
				if err != nil {
					return nil, fmt.Errorf("invalid key at #%d,%d: %w", i, j, err)
				}
				row = append(row, *k)
			default:
				return nil, fmt.Errorf("detect neither properties/object nor key/string at #%d,%d: %T", i, j, v)
			}
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
	}
	return &Layout{
		Metadata: md,
		Rows:     rows,
	}, nil
}

func jsonReparse(in, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func mergeProps(cv, nv Props) Props {
	// TODO: merge properties
	return nv
}

func newKey(p Props, s string) (*Key, error) {
	// TODO: create a new key
	return nil, nil
}
