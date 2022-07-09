package helpers

import (
	"strings"
)

// ParseColumnName will receive a "raw" name for a column,
// and return the parsed value of it, e.g:
// - ID -> id
// - PGsql -> pg_sql
// - CreatedAt -> created_at
// - UserAddress -> user_address
func ParseColumnName(name string) string {
	name = addUnderscores(name)
	return strings.ToLower(name)
}

func addUnderscores(str string) string {
	parsedStr := []byte(``)
	parsedStr = append(parsedStr, str[0])
	lenStr := len(str)
	for i := 1; i < lenStr; i++ {
		if (i + 1) == lenStr {
			parsedStr = append(parsedStr, str[i])
			continue
		}
		if !isUpper(str[i-1]) &&
			isUpper(str[i]) &&
			!isUpper(str[i+1]) {
			parsedStr = append(parsedStr, '_')
			parsedStr = append(parsedStr, str[i])
			continue
		}
		if isUpper(str[i]) && !isUpper(str[i+1]) {
			parsedStr = append(parsedStr, str[i])
			parsedStr = append(parsedStr, '_')
			continue
		}
		parsedStr = append(parsedStr, str[i])
	}

	return string(parsedStr)
}

func isUpper(b byte) bool {
	return b >= 65 && b <= 90
}
