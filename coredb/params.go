package coredb

import "reflect"

type Params struct {
	params []any
}

func NewParams(params ...any) *Params {
	return &Params{params}
}

func (p *Params) Add(params ...any) {
	for i := 0; i < len(params); i++ {
		if v := reflect.ValueOf(params[i]); v.Kind() == reflect.Slice {
			// append all elements in params[0] to p.params
			requiredCapacity := len(p.params) + v.Len() + len(params)
			newSlice := make([]any, 0, requiredCapacity)
			newSlice = append(newSlice, p.params...)
			for j := 0; j < v.Len(); j++ {
				newSlice = append(newSlice, v.Index(j).Interface())
			}
			p.params = newSlice
		} else {
			p.params = append(p.params, params[i])
		}
	}
}

func (p *Params) Get() []any {
	return p.params
}
