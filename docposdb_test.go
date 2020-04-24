package dbutils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDockerPostgresDB(t *testing.T) {
	_, port, err := NewDockerPostgresDB("test", "test", "test", "")
	require.NoError(t, err)
	db, err := sql.Open("postgres", DataSource{
		Port:     port,
		User:     "test",
		Password: "test",
		DBName:   "test",
		SSLMode:  "disable",
	}.String())
	require.NoError(t, err)

	assert.NoError(t, db.Ping())
}
