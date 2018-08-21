package jsonu_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/clavoie/jsonu"
)

func TestSerializer(t *testing.T) {
	type Struct struct {
		Name  string
		Count int
	}

	testStruct := &Struct{"name", 101}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(testStruct)

	if err != nil {
		t.Fatal(err)
	}

	serializeBytes := buffer.Bytes()
	serializeStr := strings.TrimSpace(string(serializeBytes))
	serializeBytes = []byte(serializeStr)
	serializer := jsonu.NewSerializer()

	t.Run("FromBytes", func(t *testing.T) {
		actualStruct := new(Struct)
		err := serializer.FromBytes(serializeBytes, actualStruct)

		if err != nil {
			t.Fatal(err)
		}

		if reflect.DeepEqual(testStruct, actualStruct) == false {
			t.Fatal(testStruct, actualStruct)
		}
	})

	t.Run("ToBytes", func(t *testing.T) {
		actualBytes, err := serializer.ToBytes(testStruct)

		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(serializeBytes, actualBytes) == false {
			t.Fatal(serializeStr, string(actualBytes))
		}
	})
}
