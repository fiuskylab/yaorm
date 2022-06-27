package pgsql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type migrator struct {
	db        *sql.DB
	model     any
	tableName string
}

func migrate(model any, db *sql.DB) error {
	m := &migrator{
		db:    db,
		model: model,
	}

	m.
		setTableName()

	return nil
}

func (m *migrator) setTableName() *migrator {
	splitted := strings.
		Split(
			fmt.Sprintf("%T", m.model),
			".",
		)

	tableName := splitted[len(splitted)-1]
	tableName = strings.ToLower(tableName)
	if tableName[len(tableName)-1] != 's' {
		tableName += "s"
	}

	m.tableName = tableName
	return m
}
