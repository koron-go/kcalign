package klejson2

type Metadata struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Notes  string `json:"notes,omitempty"`

	Backcolor *Color `json:"backcolor,omitempty"`

	Radii        *string `json:"radii,omitempty"`
	SwitchMount  *string `json:"switchMount,omitempty"`
	SwitchBrand *string `json:"switchBrand,omitempty"`
	SwitchType   *string `json:"switchType,omitempty"`

	Background *interface{} `json:"background,omitempty"`
}

var DefaultMetadata = Metadata{
	Backcolor: MustParseColor("#eeeeee"),
}

func defaultMetadata() Metadata {
	m := DefaultMetadata
	c := *m.Backcolor
	m.Backcolor = &c
	return m
}
