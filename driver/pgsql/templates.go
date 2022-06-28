package pgsql

const (
	checkTable = `
	SELECT EXISTS (
   SELECT FROM information_schema.tables 
   WHERE  table_schema = $1
   AND    table_name   = $2
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
