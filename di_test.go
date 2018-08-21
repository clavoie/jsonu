package jsonu_test

import (
	"testing"

	"github.com/clavoie/jsonu"
)

func TestNewDiDefs(t *testing.T) {
	defs := jsonu.NewDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := jsonu.NewDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
