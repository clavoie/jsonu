package jsonu

import "github.com/clavoie/di"

func NewDiDefs() []*di.Def {
	return []*di.Def{
		{newJson, di.Singleton},
	}
}
