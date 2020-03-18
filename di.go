package jsonu

import "github.com/clavoie/di/v2"

// NewDiDefs returns new dependency injection definitions
// for this package.
func NewDiDefs() []*di.Def {
	return []*di.Def{
		{NewSerializer, di.Singleton},
	}
}
