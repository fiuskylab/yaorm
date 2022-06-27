package pgsql

import (
	"database/sql"

	"github.com/fiuskylab/yaorm/driver"
	// Importing as blank to apply PQ interface.
	_ "github.com/lib/pq"
)

// PGSQL is PostgreSQL driver
type PGSQL struct {
	dsn string
	db  *sql.DB
}

// Open sets up a Connection with PostgreSQL DB.
func Open(dsn driver.DSN) (driver.Driver, error) {
	p := &PGSQL{
		dsn: dsn.String(),
	}

	if err := p.setConnection(); err != nil {
		return p, err
	}

	return p, nil
}

// Migrate will receive a slice of models to migrate
func (p *PGSQL) Migrate(models ...any) error {
	return nil
}

func (p *PGSQL) setConnection() error {
	db, err := sql.Open(string(driver.PGSQL), p.dsn)
	if err != nil {
		return err
	}
	p.db = db

	return err
}

// Ping will check the connection with Postgre DB.
func (p *PGSQL) Ping() error {
	return p.db.Ping()
}
