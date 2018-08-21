package jsonu

// FromBytes takes a byte slice and a pointer to a data
// structure and attempts to decode the slice into the
// data structure. See encoding/json for details.
func FromBytes(data []byte, dst interface{}) error {
	return NewSerializer().FromBytes(data, dst)
}

// ToBytes takes a pointer to a data structure and encodes
// it as a json byte slice. See encoding/json for details
func ToBytes(data interface{}) ([]byte, error) {
	return NewSerializer().ToBytes(data)
}
