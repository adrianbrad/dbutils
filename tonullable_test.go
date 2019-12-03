package dbutils

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToNullable(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(sql.NullInt64{
		Int64: 100,
		Valid: true,
	}, toNullable(100))
	assert.Equal(sql.NullInt64{
		Int64: 0,
		Valid: false,
	}, toNullable(0))

	assert.Equal([]interface{}{
		sql.NullString{
			String: "test",
			Valid:  true,
		},
		sql.NullString{
			String: "",
			Valid:  false,
		},
		sql.NullInt64{
			Int64: 100,
			Valid: true,
		},
		sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
	}, toNullableList("test", "", 100, 0))
}
