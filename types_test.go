package jsonu_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/clavoie/jsonu"
)

func TestTypes(t *testing.T) {
	type Struct struct {
		Integer    jsonu.Int
		Floating32 jsonu.Float32
		Floating64 jsonu.Float64
	}

	testStruct := &Struct{
		Integer:    10,
		Floating32: 88.8,
		Floating64: 99.1,
	}

	expectedJson := `{"Integer":10,"Floating32":88.8,"Floating64":99.1}`
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(testStruct)

	if err != nil {
		t.Fatal(err)
	}

	actualStr := strings.TrimSpace(string(buffer.Bytes()))
	if actualStr != expectedJson {
		t.Fatal("\n", expectedJson, "\n", actualStr)
	}

	decodeStruct := new(Struct)
	decoder := json.NewDecoder(buffer)
	err = decoder.Decode(decodeStruct)

	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(testStruct, decodeStruct) == false {
		t.Fatal(testStruct, decodeStruct)
	}

	quotedJson := `{"Integer":"10","Floating32":"88.8","Floating64":"99.1"}`
	decodeStruct = new(Struct)
	decoder = json.NewDecoder(strings.NewReader(quotedJson))
	err = decoder.Decode(decodeStruct)

	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(testStruct, decodeStruct) == false {
		t.Fatal(testStruct, decodeStruct)
	}

	errJson := `{"Integer":"x","Floating32":"88.8","Floating64":"99.1"}`
	decoder = json.NewDecoder(strings.NewReader(errJson))
	err = decoder.Decode(decodeStruct)

	if err == nil {
		t.Fatal("was expecting err but never got one")
	}

	errJson = `{"Integer":"10","Floating32":"x","Floating64":"99.1"}`
	decoder = json.NewDecoder(strings.NewReader(errJson))
	err = decoder.Decode(decodeStruct)

	if err == nil {
		t.Fatal("was expecting err but never got one")
	}

	errJson = `{"Integer":"10","Floating32":"88.8","Floating64":"x"}`
	decoder = json.NewDecoder(strings.NewReader(errJson))
	err = decoder.Decode(decodeStruct)

	if err == nil {
		t.Fatal("was expecting err but never got one")
	}
}
