package tag

import (
	"encoding/json"
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

	t.Run("Valid Struct", func(t *testing.T) {
		require := require.New(t)
		m := model.Model{}
		actual := ParseTags(m)
		expected := ParsedTags{
			ParsedTags: map[string][]ParsedTag{
				"ID": {
					{
						TagType: pk,
						Value:   "",
					},
					{
						TagType: autoincrement,
						Value:   "",
					},
				},
				"UpdatedAt": {
					{
						TagType: trigger,
						Value:   "update",
					},
				},
				"CreatedAt": {
					{
						TagType: trigger,
						Value:   "create",
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
			ValueOf: reflect.ValueOf(m),
			TypeOf:  reflect.TypeOf(m),
			Error:   nil,
		}

		bActual, _ := json.MarshalIndent(actual, "", "	")
		bExpected, _ := json.MarshalIndent(expected, "", "	")

		require.Equal(bActual, bExpected)
	})
}
