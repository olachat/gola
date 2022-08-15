package ormtpl

import (
	"embed"
	"text/template"
)

//go:embed *.go*
var templates embed.FS

// GetTpl return Template of giving path
func GetTpl(path string) *template.Template {
	var err error
	tpl := template.New("tpl")
	tplStr, _ := templates.ReadFile(path)
	result, err := tpl.Parse(string(tplStr))
	if err != nil {
		panic(err)
	}
	return result
}

// TplStruct define object used for tpl code generation
type TplStruct interface {
	SetVersion(version string)
	GetVersion() string
	GetName() string
}
