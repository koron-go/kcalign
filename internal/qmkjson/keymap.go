package qmkjson

// Keymap provides QMK Configurator's keymap.json
// https://beta.docs.qmk.fm/configurator/qmk-api/configurator_default_keymaps#technical-information-id-technical-information
type Keymap struct {
	Keyboard string `json:"keyboard"`

	Layout string `json:"layout"`

	Keymap string `json:"keymap"`

	Version int `json:"version"`

	Author string `json:"author"`

	Notes string `json:"notes"`

	Documentation string `json:"documentation"`

	Layers []Layer `json:"layers"`
}

// Layer represents a layer in keymap.
type Layer []string

// MarshalJSON provides json.Marshaler.
func (l Layer) MarshalJSON() ([]byte, error) {
	return []byte(`"QMKJSON_LAYER"`), nil
}
