package pgsql

import (
	"fmt"
	"reflect"

	"github.com/fiuskylab/yaorm/pkg/helpers"
)

type (
	columns []column

	column struct {
		name     string
		dataType string
	}
)

func (cs columns) toSliceStr() (slice []string) {
	for _, c := range cs {
		slice = append(slice, fmt.Sprintf("%s %s", c.name, c.dataType))
	}

	return
}

func getColums(t reflect.Type) (cs columns) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dataType := dataTypes[field.Type.String()]
		field.Name = helpers.ParseColumnName(field.Name)
		cs = append(cs, column{
			name:     field.Name,
			dataType: dataType,
		})
	}

	return
}
