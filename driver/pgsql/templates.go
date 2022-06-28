package pgsql

const (
	checkTable = `
	SELECT EXISTS (
   SELECT FROM information_schema.tables 
   WHERE  table_schema = '?'
   AND    table_name   = '?'
   );
	`

	createTable = `
	CREATE TABLE IF NOT EXISTS ? (
		{{range .}}
			{{println . ","}}
		{{end}}
	)
 	`
)
