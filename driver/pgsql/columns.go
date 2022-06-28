package pgsql

import (
	"fmt"
	"reflect"
	"strings"
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
		dataType, _ := dataTypes[field.Type.Kind().String()]
		cs = append(cs, column{
			name:     strings.ToLower(field.Name),
			dataType: dataType,
		})
	}

	return
}
