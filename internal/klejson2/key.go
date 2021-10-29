package klejson2

type Legend struct {
	Label string
	Size  float64
	Color Color
}

type Key struct {
	Legends [12]Legend

	TextSize  float64
	TextColor Color

	Color Color

	Width   float64
	Width2  float64
	Height  float64
	Height2 float64

	X  float64
	X2 float64
	Y  float64
	Y2 float64

	RotationAngle float64
	RotationX     float64
	RotationY     float64

	Profile string

	SwitchMount string
	SwitchBrand string
	SwitchType  string

	Ghost   bool
	Stepped bool
	Nub     bool
	Decal   bool
}

var DefaultKey = Key{
	Width:   1,
	Height:  1,
	Width2:  1,
	Height2: 1,

	TextSize:  3,
	TextColor: *mustParseColor("#000000"),

	Color: *mustParseColor("#cccccc"),
}
