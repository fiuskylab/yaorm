package pgsql

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
)

type migrator struct {
	buf        *bytes.Buffer
	db         *sql.DB
	err        error
	model      any
	resultRows []any
	rows       *sql.Rows
	schema     string
	stmt       *sql.Stmt
	tableName  string
	template   *template.Template
	tx         *sql.Tx
}

func migrate(model any, db *sql.DB, schema string) error {
	// TODO:
	// 		Check if receive 'model' is a struct pointer
	m := &migrator{
		db:     db,
		model:  model,
		schema: schema,
	}

	m.
		setTableName()

	if m.tableExists() {
		return m.err
	}

	m.
		begin().
		setCreateTableTemplate().
		fillCreateTable().
		exec(
			fmt.Sprintf(
				m.buf.String(),
				m.schema,
				m.tableName,
			),
		)

	return m.err
}

func (m *migrator) tableExists() bool {
	m.
		begin().
		prepare(checkTable).
		query(m.schema, m.tableName)

	b := new(bool)

	for m.rows.Next() {
		if m.err = m.rows.Scan(b); m.err != nil {
			return false
		}
	}

	return *b
}

func (m *migrator) setCreateTableTemplate() *migrator {
	if m.err != nil {
		return m
	}

	m.template, m.err = template.
		New("create_table_query").
		Funcs(funcWriteRows).
		Parse(createTable)

	return m
}

func (m *migrator) fillCreateTable() *migrator {
	if m.err != nil {
		return m
	}

	cs := getColums(reflect.TypeOf(m.model)).
		toSliceStr()

	m.buf = new(bytes.Buffer)

	m.err = m.template.Execute(m.buf, cs)

	return m
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

func (m *migrator) prepare(query string) *migrator {
	if m.err != nil {
		return m
	}

	m.stmt, m.err = m.tx.Prepare(query)

	return m
}

func (m *migrator) exec(query string, params ...any) *migrator {
	_, m.err = m.db.Exec(query, params...)

	return m
}

func (m *migrator) query(params ...any) *migrator {
	if m.err != nil {
		return m
	}

	m.rows, m.err = m.stmt.Query(params...)

	return m
}
