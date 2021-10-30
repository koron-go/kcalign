package klejson2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Row []Key

type Layout struct {
	Metadata
	Rows []Row
}

type Mods struct {
	Align *int `json:"a,omitempty"`

	TextSize     *float64  `json:"f,omitempty"`
	TextSize2    *float64  `json:"f2,omitempty"`
	TextSizeList []float64 `json:"fa,omitempty"`

	Color      *Color    `json:"c,omitempty"`
	TextColors ColorList `json:"t,omitempty"`

	X       *float64 `json:"x,omitempty"`
	Y       *float64 `json:"y,omitempty"`
	Width   *float64 `json:"w,omitempty"`
	Height  *float64 `json:"h,omitempty"`
	X2      *float64 `json:"x2,omitempty"`
	Y2      *float64 `json:"y2,omitempty"`
	Width2  *float64 `json:"w2,omitempty"`
	Height2 *float64 `json:"h2,omitempty"`

	RotationAngle *float64 `json:"r,omitempty"`
	RotationX     *float64 `json:"rx,omitempty"`
	RotationY     *float64 `json:"ry,omitempty"`

	Profile *string `json:"p,omitempty"`

	Nub     *bool `json:"n,omitempty"`
	Stepped *bool `json:"l,omitempty"`
	Decal   *bool `json:"d,omitempty"`
	Ghost   *bool `json:"g,omitempty"`

	SwitchMount *string `json:"sm,omitempty"`
	SwitchBrand *string `json:"sb,omitempty"`
	SwitchType  *string `json:"st,omitempty"`
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
	var md Metadata = defaultMetadata()
	err = jsonReparse(raw0, &md)
	if err != nil {
		return nil, fmt.Errorf("invalid layout metadata: %w", err)
	}
	// read Rows
	var rows []Row
	var curr Key = DefaultKey
	var align int = 4
	for i, r := range raw[1:] {
		list, ok := r.([]interface{})
		if !ok {
			return nil, fmt.Errorf("non-row found at #%d: %T", i, r)
		}
		var row Row
		for j, el := range list {
			switch v := el.(type) {
			case map[string]interface{}:
				// parse as a context modification
				var mods Mods
				err := jsonReparse(v, &mods)
				if err != nil {
					return nil, fmt.Errorf("invalid key properties at #%d,%d: %w", i, j, err)
				}
				err = mergeKeyTemplate(&curr, &align, &mods)
				if err != nil {
					return nil, fmt.Errorf("invalid mods at #%d,%d: %w", i, j, err)
				}
			case string:
				// parse as a key
				k, err := newKey(curr, align, v)
				if err != nil {
					return nil, fmt.Errorf("invalid key at #%d,%d: %w", i, j, err)
				}
				row = append(row, *k)
				resetKeyTemplate(&curr)
			default:
				return nil, fmt.Errorf("detect neither properties/object nor key/string at #%d,%d: %T", i, j, v)
			}
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
		curr.Y++
		curr.X = curr.RotationX
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

func mergeKeyTemplate(key *Key, align *int, mods *Mods) error {
	if mods.Align != nil {
		if *mods.Align >= len(alignmentMap) {
			return fmt.Errorf("align, out of range: want<=%d got=%d", len(alignmentMap), *mods.Align)
		}
		*align = *mods.Align
	}
	if mods.TextSize != nil {
		key.TextSize = *mods.TextSize
		for i := range key.Legends {
			key.Legends[i].Size = 0
		}
	}
	if mods.TextSize2 != nil {
		for i := 1; i < len(key.Legends); i++ {
			key.Legends[i].Size = *mods.TextSize2
		}
	}
	if m := len(mods.TextSizeList); m > 0 {
		for i := range key.Legends {
			key.Legends[i].Size = 0
		}
		// set key.Legends[i].Size with re-order
		for i, x := range alignmentIndex(*align) {
			if x < 0 || i >= len(mods.TextSizeList) {
				continue
			}
			key.Legends[x].Size = mods.TextSizeList[i]
		}
	}
	if mods.Color != nil {
		key.Color = *mods.Color
	}
	if len(mods.TextColors) > 0 {
		key.TextColor = mods.TextColors[0]
		// set key.Legends[i].Color with re-order
		for i, x := range alignmentIndex(*align) {
			if x < 0 || i >= len(mods.TextColors) {
				continue
			}
			key.Legends[x].Color = mods.TextColors[i]
		}
	}
	if mods.X != nil {
		key.X = *mods.X
	}
	if mods.Y != nil {
		key.Y = *mods.Y
	}
	if mods.Width != nil {
		key.Width = *mods.Width
	}
	if mods.Height != nil {
		key.Height = *mods.Height
	}
	if mods.X2 != nil {
		key.X2 = *mods.X2
	}
	if mods.Y2 != nil {
		key.Y2 = *mods.Y2
	}
	if mods.Width2 != nil {
		key.Width2 = *mods.Width2
	}
	if mods.Height2 != nil {
		key.Height2 = *mods.Height2
	}
	if mods.RotationAngle != nil {
		key.RotationAngle = *mods.RotationAngle
		// FIXME: check this is at first column
	}
	if mods.RotationX != nil {
		key.RotationX = *mods.RotationX
		// FIXME: check this is at first column
	}
	if mods.RotationY != nil {
		key.RotationY = *mods.RotationY
		// FIXME: check this is at first column
	}
	if mods.Profile != nil {
		key.Profile = *mods.Profile
	}
	if mods.Nub != nil {
		key.Nub = *mods.Nub
	}
	if mods.Stepped != nil {
		key.Stepped = *mods.Stepped
	}
	if mods.Decal != nil {
		key.Decal = *mods.Decal
	}
	if mods.Ghost != nil {
		key.Ghost = *mods.Ghost
	}
	if mods.SwitchMount != nil {
		key.SwitchMount = *mods.SwitchMount
	}
	if mods.SwitchBrand != nil {
		key.SwitchBrand = *mods.SwitchBrand
	}
	if mods.SwitchType != nil {
		key.SwitchType = *mods.SwitchType
	}
	return nil
}

func resetKeyTemplate(key *Key) {
	key.X += key.Width
	key.Width = 1
	key.Height = 1
	key.X2 = 0
	key.Y2 = 0
	key.Width2 = 0
	key.Height2 = 0
	key.Nub = false
	key.Stepped = false
	key.Decal = false
}

func newKey(tmpl Key, align int, s string) (*Key, error) {
	key := tmpl
	if key.Width2 == 0 {
		key.Width2 = key.Width
	}
	if key.Height2 == 0 {
		key.Height2 = key.Height
	}
	// set key.Legends[i].Label with re-order
	labels := strings.SplitN(s, "\n", 13)
	for i, x := range alignmentIndex(align) {
		if x < 0 || i >= len(labels) {
			continue
		}
		key.Legends[x].Label = labels[i]
	}
	for i, l := range key.Legends {
		if l.Label == "" {
			key.Legends[i].Size = 0
			key.Legends[i].Color = Color{}
		}
	}
	return &key, nil
}

var alignmentMap = [][]int{
	{0, 6, 2, 8, 9, 11, 3, 5, 1, 4, 7, 10},          // 0 = no centering
	{1, 7, -1, -1, 9, 11, 4, -1, -1, -1, -1, 10},    // 1 = center x
	{3, -1, 5, -1, 9, 11, -1, -1, 4, -1, -1, 10},    // 2 = center y
	{4, -1, -1, -1, 9, 11, -1, -1, -1, -1, -1, 10},  // 3 = center x & y
	{0, 6, 2, 8, 10, -1, 3, 5, 1, 4, 7, -1},         // 4 = center front (default)
	{1, 7, -1, -1, 10, -1, 4, -1, -1, -1, -1, -1},   // 5 = center front & x
	{3, -1, 5, -1, 10, -1, -1, -1, 4, -1, -1, -1},   // 6 = center front & y
	{4, -1, -1, -1, 10, -1, -1, -1, -1, -1, -1, -1}, // 7 = center front & x & y
}

func alignmentIndex(align int) []int {
	return alignmentMap[align]
}
