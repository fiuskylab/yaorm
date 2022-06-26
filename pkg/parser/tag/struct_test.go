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

		var expectedMap map[string]string
		expectedError := fmt.Sprintf(invalidTypeErr, reflect.Struct.String(), "int")

		i := int(1)

		actualMap, actualErr := ParseTags(i)

		if actualErr == nil {
			require.NotNil(actualErr)
		}

		require.Equal(expectedMap, actualMap)
		require.EqualError(actualErr, expectedError)
	})
}
