package jsonu_test

import (
	"fmt"
	"log"

	"github.com/clavoie/jsonu"
)

func ExampleTypes() {
	type Struct struct {
		Integer    jsonu.Int
		Floating32 jsonu.Float32
		Floating64 jsonu.Float64
	}

	testStruct := new(Struct)

	expectedJson := `{"Integer":10,"Floating32":88.8,"Floating64":99.1}`
	err := jsonu.FromBytes([]byte(expectedJson), testStruct)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(testStruct.Integer, testStruct.Floating32, testStruct.Floating64)
	jsonBytes, err := jsonu.ToBytes(testStruct)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))

	// Output: 10 88.8 99.1
	//{"Integer":10,"Floating32":88.8,"Floating64":99.1}
}
