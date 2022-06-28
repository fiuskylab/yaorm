package pgsql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type migrator struct {
	db        *sql.DB
	err       error
	model     any
	schema    string
	stmt      *sql.Stmt
	tableName string
	tx        *sql.Tx
}

func migrate(model any, db *sql.DB, schema string) error {
	m := &migrator{
		db:     db,
		model:  model,
		schema: schema,
	}

	m.
		begin().
		setTableName().
		buildCheckStmt().
		runCheckStmt()

	return m.err
}

// begin will setup the sql.Tx
func (m *migrator) begin() *migrator {
	m.tx, m.err = m.db.Begin()
	return m
}

func (m *migrator) setTableName() *migrator {
	if m.err != nil {
		return m
	}
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

func (m *migrator) buildCheckStmt() *migrator {
	if m.err != nil {
		return m
	}

	query := fmt.Sprintf(
		checkTable,
		m.schema,
		m.tableName,
	)

	m.stmt, m.err = m.tx.Prepare(query)

	if m.err != nil {
		return m
	}

	return m
}

func (m *migrator) runCheckStmt() *migrator {
	if m.err != nil {
		return m
	}
}
