package driver

// Type is the type for defining which connection is going to be set.
type Type string

const (
	// PGSQL is the string representation for
	// setting up a Postgres connection.
	PGSQL Type = "postgres"
)
