package tag

import (
	"fmt"
	"reflect"
)

// ParseTags receives a struct, and parse all yaORM tags.
func ParseTags(v any) (m map[string]string, err error) {
	typeOf := reflect.TypeOf(v)

	if typeOf.Kind() != reflect.Struct {
		err = fmt.Errorf(invalidTypeErr, reflect.Struct.String(), typeOf.Kind().String())
		return
	}

	return
}
