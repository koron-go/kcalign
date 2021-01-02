package kcalign

import "fmt"

// TextAlign specifies text alignment in column.
type TextAlign int

const (
	// Left is left alignment.
	Left TextAlign = iota + 1
	// Right is right alignment.
	Right
	// Center is center alignment.
	Center
)

// MarshalJSON provides json.Marshaler.
func (ta TextAlign) MarshalJSON() ([]byte, error) {
	switch ta {
	case Left:
		return []byte(`"left"`), nil
	case Right:
		return []byte(`"right"`), nil
	case Center:
		return []byte(`"center"`), nil
	default:
		return nil, nil
	}
}

// UnmarshalJSON provides json.Unmarshaler.
func (ta *TextAlign) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"left"`:
		*ta = Left
		return nil
	case `"right"`:
		*ta = Right
		return nil
	case `"center"`:
		*ta = Center
		return nil
	default:
		return fmt.Errorf("unknown value for TextAlign: %q", string(b))
	}
}
