# jsonu [![GoDoc](https://godoc.org/github.com/clavoie/jsonu?status.svg)](http://godoc.org/github.com/clavoie/jsonu) [![Build Status](https://travis-ci.org/clavoie/jsonu.svg?branch=master)](https://travis-ci.org/clavoie/jsonu) [![codecov](https://codecov.io/gh/clavoie/jsonu/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/jsonu) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/jsonu)](https://goreportcard.com/report/github.com/clavoie/jsonu)

JSON utilities for Go

## Quote Deserialization

Some client side JSON libraries quote numbers they serialize in AJAX requests. Go can't deserialize those values with encoding/json, so jsonu provides some numerical wrapper types that know how deserialize their values if they are quoted:

```go
  type Struct struct {
    Integer    jsonu.Int
    Floating32 jsonu.Float32
    Floating64 jsonu.Float64
  }
  quotedJson := `{"Integer":"10","Floating32":"88.8","Floating64":"99.1"}`
  decodeStruct = new(Struct)
  decoder = json.NewDecoder(strings.NewReader(quotedJson))
  err = decoder.Decode(decodeStruct)
  // err will be nil
```

## Bytes

Convenience functions are provided to serialize / deserialize JSON objects from a byte slice.

