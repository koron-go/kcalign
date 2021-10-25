package klejson2

type LabelDefault struct {
	TextColor string
	TextSize  float64
}

type KeyProps struct {
	// position
	X  float64 `json:"x,omitempty"`
	Y  float64 `json:"y,omitempty"`
	X2 float64 `json:"x2,omitempty"`
	Y2 float64 `json:"y2,omitempty"`

	// size
	Width   float64 `json:"w,omitempty"`
	Height  float64 `json:"h,omitempty"`
	Width2  float64 `json:"w2,omitempty"`
	Height2 float64 `json:"h2,omitempty"`

	// rotation
	RotationAngle float64 `json:"r,omitempty"`
	RotationX     float64 `json:"rx,omitempty"`
	RotationY     float64 `json:"ry,omitempty"`

	// label properties
	Labels    []string
	TextColor []string `json:"t,omitempty"`
	Textsize  []float64

	// label defaults
	Default LabelDefault

	// cap appearance
	Color   string `json:"c,omitempty"`
	Profile string `json:"p,omitempty"`
	Nub     bool   `json:"n,omitempty"`

	// miscellaneous options
	Ghost   bool `json:"g,omitempty"`
	Stepped bool `json:"l,omitempty"`
	Decal   bool `json:"d,omitempty"`

	// switch
	SwitchMount string `json:"sm,omitempty"`
	SwitchBrand string `json:"sb,omitempty"`
	SwitchType  string `json:"st,omitempty"`
}

var DefaultKeyProps = KeyProps{
	Width:   1,
	Height:  1,
	Width2:  1,
	Height2: 1,
	Default: LabelDefault{
		TextColor: "#000000",
		TextSize:  3,
	},
	Color: "#cccccc",
}
