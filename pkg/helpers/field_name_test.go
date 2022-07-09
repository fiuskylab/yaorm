package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseColumnName(t *testing.T) {
	t.Run("ID", func(t *testing.T) {
		require := require.New(t)

		expected := "id"
		actual := ParseColumnName("ID")

		require.Equal(expected, actual)
	})

	t.Run("PGsql", func(t *testing.T) {
		require := require.New(t)

		expected := "pg_sql"
		actual := ParseColumnName("PGsql")

		require.Equal(expected, actual)
	})

	t.Run("CreatedAt", func(t *testing.T) {
		require := require.New(t)

		expected := "created_at"
		actual := ParseColumnName("CreatedAt")

		require.Equal(expected, actual)
	})
	t.Run("UserAddress", func(t *testing.T) {
		require := require.New(t)

		expected := "user_address"
		actual := ParseColumnName("UserAddress")

		require.Equal(expected, actual)
	})
}
