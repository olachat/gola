package structs

import "strings"

func getGoName(sqlName string) string {
	sap := "_"
	if strings.Contains(sqlName, "-") {
		sap = "-"
	}

	parts := strings.Split(sqlName, sap)
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}

	joinString := strings.Join(parts, "")
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
