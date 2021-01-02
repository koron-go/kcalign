package kcalign

import "fmt"

// QuoteType specifies runes to quote each items.
type QuoteType int

const (
	// Double uses dobule quote (`"`).
	Double QuoteType = iota + 1
	// NoQuote doesn't use quote runes, write each items as is.
	NoQuote
)

// MarshalJSON provides json.Marshaler.
func (qt QuoteType) MarshalJSON() ([]byte, error) {
	switch qt {
	case Double:
		return []byte(`"double"`), nil
	case NoQuote:
		return []byte(`"none"`), nil
	default:
		return nil, nil
	}
}

// UnmarshalJSON provides json.Unmarshaler.
func (qt *QuoteType) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"double"`:
		*qt = Double
	case `"none"`:
		*qt = NoQuote
	default:
		return fmt.Errorf("unknown value for QuoteType: %s", string(b))
	}
	return nil
}
