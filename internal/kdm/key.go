package kdm

import (
	"math"
	"sort"
)

type Pos struct {
	X float64
	Y float64
}

type Box [4]Pos

func (b Box) Rotate(theta float64) Box {
	s, c := math.Sincos(theta)
	for i, p := range b {
		b[i].X = s*p.X + c*p.Y
		b[i].Y = -c*p.X + s*p.Y
	}
	return b
}

func (b Box) MinMax() (min, max Pos) {
	for _, p := range b {
		if p.X < min.X {
			min.X = p.X
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return min, max
}

type Dim struct {
	W float64
	H float64
}

type Key struct {
	ID int

	Row    int
	Column int

	Center   Pos
	Size     Dim
	Rotation float64 // degree

	Label string
}

func (k Key) Box() Box {
	return Box{
		{
			X: k.Center.X - k.Size.W/2,
			Y: k.Center.Y - k.Size.H/2,
		},
		{
			X: k.Center.X + k.Size.W/2,
			Y: k.Center.Y - k.Size.H/2,
		},
		{
			X: k.Center.X - k.Size.W/2,
			Y: k.Center.Y + k.Size.H/2,
		},
		{
			X: k.Center.X + k.Size.W/2,
			Y: k.Center.Y + k.Size.H/2,
		},
	}
}

func (k Key) BoundingBox() (min, max Pos) {
	box := k.Box()
	// consider rotate
	if k.Rotation != 0 {
		box.Rotate(k.Rotation * math.Pi / 180)
	}
	return box.MinMax()
}

type Keys []Key

func (kk Keys) ByRow(row int) Keys {
	r := make(Keys, 0, 10)
	for _, k := range kk {
		if k.Row == row {
			r = append(r, k)
		}
	}
	return r
}

func (kk Keys) ByColumn(col int) Keys {
	r := make(Keys, 0, 10)
	for _, k := range kk {
		if k.Column == col {
			r = append(r, k)
		}
	}
	return r
}

// Sort sorts keys order by row and column
func (kk Keys) Sort() {
	sort.SliceStable(kk, func(i, j int) bool {
		if kk[i].Column != kk[j].Column {
			return kk[i].Column < kk[j].Column
		}
		return kk[i].Row < kk[j].Row
	})
}
