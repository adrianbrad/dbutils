package dbutils

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestDataSource_String(t *testing.T) {
	type table struct{
		Expected string
		Actual string
	}

	os.Clearenv()
	connInfosNoEnvVars := []table {
		{"host=localhost port=5432 user=postgres password=postgres dbname=nowaster sslmode=disable",  DataSource{DBName: "nowaster"}.String()},
		{"host=db port=123 user=a password=b dbname=c sslmode=enable",  DataSource{Host: "db", Port: "123", User: "a", Password: "b", DBName: "c", SSLMode: "enable"}.String()},
	}
	runTableTest(t, connInfosNoEnvVars)

	os.Setenv("DB_HOST", "ENV")
	connInfosHostVar := []table {
		{"host=ENV port=5432 user=postgres password=postgres dbname=nowaster sslmode=disable", DataSource{DBName: "nowaster"}.String()},
	}
	runTableTest(t, connInfosHostVar)

	os.Setenv("DB_PORT", "54320")
	connInfosHostPortVar := []table {
		{"host=ENV port=54320 user=postgres password=postgres dbname=nowaster sslmode=disable", DataSource{DBName: "nowaster"}.String()},
	}
	runTableTest(t, connInfosHostPortVar)
}

func runTableTest(t *testing.T, table interface{}) {
	slice := reflect.ValueOf(table)
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)
		expected := elem.Field(0).Interface()
		actual := elem.Field(1).Interface()
		t.Run(fmt.Sprintf("%s", expected), func(t *testing.T) {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("\nGOT: %s\n EXPECTED: %s", actual, expected)
			}
		})
	}
}
