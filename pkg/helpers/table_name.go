package helpers

import (
	"strings"

	"github.com/gertd/go-pluralize"
)

var (
	pluClient *pluralize.Client
)

func init() {
	pluClient = pluralize.NewClient()
}

// ParseTableName will just add an 's' at the of
// a table name:
// "User" -> "users"
// "Car" -> "cars"
func ParseTableName(tableName string) string {
	tableName = pluClient.Plural(tableName)
	tableName = strings.ToLower(tableName)
	return tableName
}
