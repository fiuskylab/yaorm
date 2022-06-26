package tag

import (
	"fmt"
	"reflect"
	"testing"

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

}
