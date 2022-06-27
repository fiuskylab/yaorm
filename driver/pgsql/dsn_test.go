package pgsql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	t.Run("Empty DSN", func(t *testing.T) {
		require := require.New(t)
		dsn := DSN{}

		actual := dsn.String()
		expected := "postgresql://:@:0/"

		require.Equal(expected, actual)
	})

	t.Run("Filled DSN", func(t *testing.T) {
		require := require.New(t)
		dsn := DSN{
			User:     "root",
			Password: "root",
		}

		actual := dsn.String()
		expected := "postgresql://root:root@:0/"

		require.Equal(expected, actual)
	})

	t.Run("Filled DSN", func(t *testing.T) {
		require := require.New(t)
		dsn := DSN{
			User:     "root",
			Password: "root",
			Host:     "localhost",
			Port:     5432,
		}

		actual := dsn.String()
		expected := "postgresql://root:root@localhost:5432/"

		require.Equal(expected, actual)
	})

	t.Run("Filled DSN", func(t *testing.T) {
		require := require.New(t)
		dsn := DSN{
			User:     "root",
			Password: "root",
			Host:     "localhost",
			Port:     5432,
			DBName:   "db",
		}

		actual := dsn.String()
		expected := "postgresql://root:root@localhost:5432/db"

		require.Equal(expected, actual)
	})

	t.Run("Filled DSN", func(t *testing.T) {
		require := require.New(t)
		dsn := DSN{
			User:     "root",
			Password: "root",
			Host:     "localhost",
			Port:     5432,
			DBName:   "db",
			Params: map[string]string{
				"sslmode": "true",
			},
		}

		actual := dsn.String()
		expected := "postgresql://root:root@localhost:5432/db?sslmode=true"

		require.Equal(expected, actual)
	})

	t.Run("Filled DSN", func(t *testing.T) {
		require := require.New(t)
		dsn := DSN{
			User:     "root",
			Password: "root",
			Host:     "localhost",
			Port:     5432,
			DBName:   "db",
			Params: map[string]string{
				"sslmode": "true",
				"param1":  "value1",
				"param2":  "value2",
			},
		}

		actual := dsn.String()
		expected := "postgresql://root:root@localhost:5432/db?sslmode=true&param1=value1&param2=value2"

		require.Equal(expected, actual)
	})
}
