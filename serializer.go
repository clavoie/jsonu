package jsonu

import (
	"bytes"
	"encoding/json"
)

// Serializer encodes to and from []byte to json objects
type Serializer interface {
	// FromBytes takes a byte slice and a pointer to a data
	// structure and attempts to decode the slice into the
	// data structure. See encoding/json for details.
	FromBytes([]byte, interface{}) error

	// ToBytes takes a pointer to a data structure and encodes
	// it as a json byte slice. See encoding/json for details
	ToBytes(interface{}) ([]byte, error)
}

// serializer is an implementation of Serializer
type serializer struct{}

// NewSerializer returns a new instance of Serializer
func NewSerializer() Serializer {
	return new(serializer)
}

func (js *serializer) FromBytes(data []byte, dst interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)

	return decoder.Decode(dst)
}

func (js *serializer) ToBytes(data interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(data)

	if err == nil {
		return buffer.Bytes(), nil
	}

	return []byte{}, err
}
