/*
Package kdm privoides Keyboard Design Model
*/
package kdm

type Keyboard struct {
	Name string

	VendorID  uint16
	ProductID uint16

	Keys Keys
}

func (kb *Keyboard) AddKeys(keys ...Key) {
	if kb.Keys == nil {
		kb.Keys = make([]Key, 0, 100)
	}
	kb.Keys = append(kb.Keys, keys...)
}
