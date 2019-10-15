package dbutils

import (
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func TestAttemptConnect(t *testing.T) {
	db, ok, errs := AttemptConnect("postgres", DataSource{
		Host:     "test_db_1",
		DBName:   "dbutils",
		Password: "dbutils",
		User:     "dbutils",
	}, 10, time.Second)

	if ok != true {
		t.Errorf("connection failed")
		t.FailNow()
	}

	t.Logf("connection errors: %s", errs.Error())

	row := db.QueryRow("SELECT success FROM test")

	var success bool
	err := row.Scan(&success)
	if err != nil {
		t.Errorf("scanning failed")
		t.FailNow()
	}

	if success != true {
		t.Errorf("success is not true")
		t.FailNow()
	}

	db.Close()
	db, ok, errs = AttemptConnect("postgres", DataSource{
		Host:     "test_db_1",
		DBName:   "dbutils",
		Password: "dbutils",
		User:     "dbutils",
	}, 1, 1)

	if errs != nil {
		t.Errorf("got connection errors when db was up: %s", errs.Error())
		t.FailNow()
	}

	if ok != true {
		t.Errorf("connection failed when db was up")
		t.FailNow()
	}

	row = db.QueryRow("SELECT success FROM test")

	err = row.Scan(&success)
	if err != nil {
		t.Errorf("scanning failed with db up")
		t.FailNow()
	}

	if success != true {
		t.Errorf("success is not true with db up")
		t.FailNow()
	}
}
