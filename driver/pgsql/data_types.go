package pgsql

var (
	// TODO:
	// 		Review if those make sense.
	// 		Following: https://www.postgresql.org/docs/13/datatype.html
	dataTypes = map[string]string{
		"int8":      "smallint",
		"int16":     "integer",
		"int32":     "bigint",
		"int64":     "bigint",
		"int":       "bigint",
		"string":    "text",
		"time.Time": "timestamp",
	}
)
