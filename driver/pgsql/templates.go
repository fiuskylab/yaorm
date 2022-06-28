package pgsql

import (
	"strings"
	"text/template"
)

const (
	checkTable = `SELECT EXISTS (
   SELECT FROM information_schema.tables 
   WHERE  table_schema = $1
   AND    table_name   = $2
   );`

	createTable = `CREATE TABLE IF NOT EXISTS %s.%s (
		{{buildRows .}}
	);`
)

var funcWriteRows = template.FuncMap{
	"buildRows": func(elems []string) string {
		return strings.Join(elems, ",\n")
	},
}
