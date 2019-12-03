package dbutils

import (
	"database/sql"
	"reflect"
	"time"
)

func ToNullable(given interface{}) interface{} {
	switch given.(type) {
	case int, int8, int16, int32, int64:
		n := reflect.ValueOf(given).Int()
		return sql.NullInt64{
			Int64: n,
			Valid: n != 0,
		}

	case string:
		return sql.NullString{
			String: given.(string),
			Valid:  given.(string) != "",
		}

	case time.Time:
		defaultTime := time.Time{}
		return sql.NullTime{
			Time:  given.(time.Time),
			Valid: given.(time.Time) != defaultTime,
		}
	default:
		return nil
	}
}

func ToNullableList(args ...interface{}) []interface{} {
	res := make([]interface{}, len(args))
	for i, obj := range args {
		res[i] = ToNullable(obj)
	}
	return res
}
