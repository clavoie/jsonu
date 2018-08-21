package jsonu

import "encoding/json"

// Int is a type that can be used for fields in json serialization
// structs. It can deserialize from both an integer number and a
// quoted integer number. For example if the json value to be deserialized
// was "9" the value of a field of type Int after deserialization would be
// 9
type Int int

// MarshalJSON serializes Int the same way as encoding/json.Marshal(int)
func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

// UnmarshalJSON deserializes Int the same way as encoding/json.Unmarshal([]byte, &int)
// with the exception that it can deserialize quoted values
func (i *Int) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	var tmp int
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	*i = Int(tmp)
	return nil
}

// Float32 is a type that can be used for fields in json serialization
// structs. It can deserialize from both a floating point number and a
// quoted floating point number. For example if the json value to be deserialized
// was "9.5" the value of a field of type Float32 after deserialization would be
// 9.5
type Float32 float32

// MarshalJSON serializes Int the same way as encoding.json.Marshal(float32)
func (f Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(f))
}

// UnmarshalJSON deserializes Int the same way as encoding/json.Unmarshal([]byte, &float32)
// with the exception that it can deserialize quoted values
func (f *Float32) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	var tmp float32
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	*f = Float32(tmp)
	return nil
}

// Float64 is a type that can be used for fields in json serialization
// structs. It can deserialize from both a floating point number and a
// quoted floating point number. For example if the json value to be deserialized
// was "9.5" the value of a field of type Float64 after deserialization would be
// 9.5
type Float64 float64

// MarshalJSON serializes Int the same way as encoding.json.Marshal(float64)
func (f Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

// UnmarshalJSON deserializes Int the same way as encoding/json.Unmarshal([]byte, &float64)
// with the exception that it can deserialize quoted values
func (f *Float64) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	var tmp float64
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	*f = Float64(tmp)
	return nil
}
