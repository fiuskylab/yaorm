package yaorm

import (
	"fmt"

	"github.com/fiuskylab/yaorm/driver"
	"github.com/fiuskylab/yaorm/driver/pgsql"
)

// Open setup a connection to a given Database.
func Open(t driver.Type, dsn driver.DSN, schema string) (driver.Driver, error) {
	switch t {
	case driver.PGSQL:
		return pgsql.Open(dsn, schema)
	default:
		return nil, fmt.Errorf("driver %s not supported", t)
	}
}
