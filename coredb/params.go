package coredb

import "reflect"

type Params struct {
	params []any
}

func NewParams(params ...any) *Params {
	return &Params{params}
}

func (p *Params) Add(params ...any) {
	totalLen := len(p.params)
	for i := 0; i < len(params); i++ {
		if v := reflect.ValueOf(params[i]); v.Kind() == reflect.Slice {
			totalLen += v.Len()
		} else {
			totalLen++
		}
	}

	newSlice := make([]any, 0, totalLen)
	newSlice = append(newSlice, p.params...)
	for i := 0; i < len(params); i++ {
		if v := reflect.ValueOf(params[i]); v.Kind() == reflect.Slice {
			// append all elements in params[0] to p.params
			for j := 0; j < v.Len(); j++ {
				newSlice = append(newSlice, v.Index(j).Interface())
			}
		} else {
			newSlice = append(newSlice, params[i])
		}
	}
	p.params = newSlice
}

func (p *Params) Get() []any {
	return p.params
}
