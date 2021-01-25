package klejson

import "sort"

type Key struct {
	X float64
	Y float64
	W float64
	H float64

	CX float64
	CY float64
}

func newKey(x, y, w, h float64) Key {
	return Key{
		X:  x,
		Y:  y,
		W:  w,
		H:  h,
		CX: x + w/2,
		CY: y + h/2,
	}
}

type Keys []Key

func (keys Keys) SortByCenter() Keys {
	sort.SliceStable(keys, func(i, j int) bool {
		if keys[i].X < keys[j].X {
			return true
		}
		if keys[i].X > keys[j].X {
			return false
		}
		if keys[i].Y < keys[j].Y {
			return true
		}
		return i < j
	})

	// TODO:
	return keys
}
