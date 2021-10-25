package klejson2

type Props struct {
	X  float64 `json:"x,omitempty"`
	Y  float64 `json:"y,omitempty"`
	X2 float64 `json:"x2,omitempty"`
	Y2 float64 `json:"y2,omitempty"`

	RotationAngle float64 `json:"r,omitempty"`
	RotationX     float64 `json:"rx,omitempty"`
	RotationY     float64 `json:"ry,omitempty"`

	Width   float64 `json:"w,omitempty"`
	Height  float64 `json:"h,omitempty"`
	Width2  float64 `json:"w2,omitempty"`
	Height2 float64 `json:"h2,omitempty"`

	Align int `json:"a,omitempty"`

	TextSize      float64   `json:"f,omitempty"`
	OptimizedSize float64   `json:"f2,omitempty"`
	OrderedSize   []float64 `json:"fa,omitempty"`

	Color     Color `json:"c,omitempty"`
	TextColor Color `json:"t,omitempty"`

	Ghost   string `json:"g,omitempty"`
	Profile string `json:"p,omitempty"`

	Nub     bool `json:"n,omitempty"`
	Stepped bool `json:"l,omitempty"`
	Decal   bool `json:"d,omitempty"`

	SwitchMount string `json:"sm,omitempty"`
	SwitchBrand string `json:"sb,omitempty"`
	SwitchType  string `json:"st,omitempty"`
}

var DefaultProps = Props{
	Width:   1,
	Height:  1,
	Width2:  1,
	Height2: 1,
	Color:   mustParseColor("#cccccc"),
}
