package yaorm

import (
	"fmt"

	"github.com/fiuskylab/yaorm/driver"
	"github.com/fiuskylab/yaorm/driver/pgsql"
)

// Open setup a connection to a given Database.
func Open(t driver.Type, dsn driver.DSN) (driver.Driver, error) {
	switch t {
	case driver.PGSQL:
		return pgsql.Open(dsn)
	default:
		return nil, fmt.Errorf("driver %s not supported", t)
	}
}
