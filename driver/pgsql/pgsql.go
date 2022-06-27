package pgsql

import "github.com/fiuskylab/yaorm/driver"

// PGSQL is PostgreSQL driver
type PGSQL struct {
	dsn string
}

// Open sets up a Connection with PostgreSQL DB.
func Open(dsn driver.DSN) (driver.Driver, error) {
	p := &PGSQL{
		dsn: dsn.String(),
	}

	return p, nil
}

// Ping will check the connection with Postgre DB.
func (p *PGSQL) Ping() error {

	return nil
}
