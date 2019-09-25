package dbutils

import (
	"fmt"
	"os"
	"reflect"
)

type DataSource struct {
	Host     string `default:"localhost" env:"DB_HOST"`
	Port     string `default:"5432" env:"DB_PORT"`
	User     string `default:"postgres" env:"DB_USER"`
	Password string `default:"postgres" env:"DB_PASS"`
	DBName   string `default:"postgres" env:"DB_NAME"`
	SSLMode  string `default:"disable" env:"DB_SSLMODE"`
}

func (d DataSource) String() string {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(&d).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if field.String() == "" {
			envVar := t.Field(i).Tag.Get("env")
			if envVar != "" {
				val := os.Getenv(envVar)
				if val != "" {
					field.SetString(val)
					continue
				}
			}
			defaultVal := t.Field(i).Tag.Get("default")
			field.SetString(defaultVal)
		}
	}

	if d.DBName == "" {
		panic("DataSource.DBName must not be empty")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode)
}