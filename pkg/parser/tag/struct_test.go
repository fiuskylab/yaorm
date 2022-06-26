package tag

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fiuskylab/yaorm/model"
	"github.com/stretchr/testify/require"
)

func TestParseTags(t *testing.T) {
	t.Run("Not Struct", func(t *testing.T) {
		require := require.New(t)

		expected := ParsedTags{
			Error: fmt.Errorf(invalidTypeErr, reflect.Struct.String(), "int"),
		}

		i := int(1)

		actualParsedTags := ParseTags(i)

		if actualParsedTags.Error == nil {
			require.NotNil(actualParsedTags.Error)
		}

		require.EqualError(actualParsedTags.Error, expected.Error.Error())
	})

	// TODO:
	// 	Find some way to validate the struct and ignore the pointers
	//t.Run("Valid Struct", func(t *testing.T) {
	//})

	m := model.Model{}
	_ = ParseTags(m)
	_ = ParsedTags{
		parsedTags: map[string][]parsedTag{
			"ID": {
				{
					tagType: pk,
					value:   "",
				},
				{
					tagType: autoincrement,
					value:   "",
				},
			},
			"UpdatedAt": {
				{
					tagType: trigger,
					value:   "update",
				},
			},
			"CreatedAt": {
				{
					tagType: trigger,
					value:   "create",
				},
			},
		},
		rawSlice: []rawTag{
			{
				rawStr: "pk,autoincrement",
				field:  reflect.TypeOf(m).Field(0),
				value:  reflect.ValueOf(m).Field(0),
			},
			{
				rawStr: "trigger:update",
				field:  reflect.TypeOf(m).Field(1),
				value:  reflect.ValueOf(m).Field(1),
			},
			{
				rawStr: "trigger:create",
				field:  reflect.TypeOf(m).Field(2),
				value:  reflect.ValueOf(m).Field(2),
			},
		},
		valueOf: reflect.ValueOf(m),
		typeOf:  reflect.TypeOf(m),
		Error:   nil,
	}
}
