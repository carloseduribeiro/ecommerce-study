package database

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnection(t *testing.T) {
	t.Run("should testing database connection", func(t *testing.T) {
		testDatabase := NewTestDatabase(t)
		defer testDatabase.Close(t)
		connection, err := NewConnection("postgres", testDatabase.ConnectionString(t))
		require.NoError(t, err)
		defer connection.Close()
		err = connection.Exec("CREATE TABLE IF NOT EXISTS teste (id INTEGER NOT NULL PRIMARY KEY)")
		assert.NoError(t, err)
	})
}
