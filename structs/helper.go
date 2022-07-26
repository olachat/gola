package structs

import (
	"net/url"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getGoName(sqlName string) string {
	if sqlName == "" {
		return "Empty"
	}
	sap := "_"
	if strings.Contains(sqlName, "-") {
		sap = "-"
	}

	parts := strings.Split(sqlName, sap)
	for i, p := range parts {
		parts[i] = cases.Title(language.English).String(p)
	}

	joinString := strings.Join(parts, "")
	joinString = url.QueryEscape(joinString)
	joinString = strings.ReplaceAll(joinString, "%", "x")
	if sqlName[:1] == "_" {
		return "X" + joinString
	}

	return joinString
}

func getValue(fullDBType string) string {
	i := strings.Index(fullDBType, "(")
	j := strings.LastIndex(fullDBType, ")")
	if i < j && i > 0 {
		return fullDBType[i+1 : j]
	}

	return ""
}

func setInclude(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}

	return false
}
