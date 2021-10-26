package klejson2

type Props struct {
	X  *float64 `json:"x,omitempty"`
	Y  *float64 `json:"y,omitempty"`
	X2 *float64 `json:"x2,omitempty"`
	Y2 *float64 `json:"y2,omitempty"`

	RotationAngle *float64 `json:"r,omitempty"`
	RotationX     *float64 `json:"rx,omitempty"`
	RotationY     *float64 `json:"ry,omitempty"`

	Width   *float64 `json:"w,omitempty"`
	Height  *float64 `json:"h,omitempty"`
	Width2  *float64 `json:"w2,omitempty"`
	Height2 *float64 `json:"h2,omitempty"`

	Align *int `json:"a,omitempty"`

	TextSize      *float64  `json:"f,omitempty"`
	OptimizedSize *float64  `json:"f2,omitempty"`
	OrderedSize   []float64 `json:"fa,omitempty"`

	Color     *Color `json:"c,omitempty"`
	TextColor *Color `json:"t,omitempty"`

	Ghost   *string `json:"g,omitempty"`
	Profile *string `json:"p,omitempty"`

	Nub     *bool `json:"n,omitempty"`
	Stepped *bool `json:"l,omitempty"`
	Decal   *bool `json:"d,omitempty"`

	SwitchMount *string `json:"sm,omitempty"`
	SwitchBrand *string `json:"sb,omitempty"`
	SwitchType  *string `json:"st,omitempty"`
}

func (p *Props) Merge(v Props) {
	if v.X != nil {
		p.X = v.X
	}
	if v.Y != nil {
		p.Y = v.Y
	}
	if v.X2 != nil {
		p.X2 = v.X2
	}
	if v.Y2 != nil {
		p.Y2 = v.Y2
	}
	if v.RotationAngle != nil {
		p.RotationAngle = v.RotationAngle
	}
	if v.RotationX != nil {
		p.RotationX = v.RotationX
	}
	if v.RotationY != nil {
		p.RotationY = v.RotationY
	}
	if v.Width != nil {
		p.Width = v.Width
	}
	if v.Height != nil {
		p.Height = v.Height
	}
	if v.Width2 != nil {
		p.Width2 = v.Width2
	}
	if v.Height2 != nil {
		p.Height2 = v.Height2
	}
	if v.Align != nil {
		p.Align = v.Align
	}
	if v.TextSize != nil {
		p.TextSize = v.TextSize
	}
	if v.OptimizedSize != nil {
		p.OptimizedSize = v.OptimizedSize
	}
	if v.OrderedSize != nil {
		p.OrderedSize = v.OrderedSize
	}
	if v.Color != nil {
		p.Color = v.Color
	}
	if v.TextColor != nil {
		p.TextColor = v.TextColor
	}
	if v.Ghost != nil {
		p.Ghost = v.Ghost
	}
	if v.Profile != nil {
		p.Profile = v.Profile
	}
	if v.Nub != nil {
		p.Nub = v.Nub
	}
	if v.Stepped != nil {
		p.Stepped = v.Stepped
	}
	if v.Decal != nil {
		p.Decal = v.Decal
	}
	if v.SwitchMount != nil {
		p.SwitchMount = v.SwitchMount
	}
	if v.SwitchBrand != nil {
		p.SwitchBrand = v.SwitchBrand
	}
	if v.SwitchType != nil {
		p.SwitchType = v.SwitchType
	}
}

func float2ptr(v float64) *float64 {
	return &v
}

var DefaultProps = Props{
	Width:   float2ptr(1),
	Height:  float2ptr(1),
	Width2:  float2ptr(1),
	Height2: float2ptr(1),
	Color:   mustParseColor("#cccccc"),
}
