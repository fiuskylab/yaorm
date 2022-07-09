package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTableName(t *testing.T) {
	t.Run("Pluralize Client", func(t *testing.T) {
		require := require.New(t)

		require.NotNil(pluClient)
	})

	t.Run("User", func(t *testing.T) {
		require := require.New(t)

		expected := "users"
		actual := ParseTableName("User")

		require.Equal(expected, actual)
	})
}
