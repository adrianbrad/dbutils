package dbutils

import (
	"database/sql"
	"time"
)

// tries- default 0, runs 1000 attempts
// wait - default 0, wait 1 second
func AttemptConnectPostgres(ds DataSource, tries int, wait time.Duration) (*sql.DB, error) {
	var db *sql.DB
	var err error

	if wait == 0 {
		wait = 1
	}
	if tries == 0 {
		tries = 1000
	}

	for i := 1; i <= tries; i++ {
		db, err = sql.Open("postgres", ds.String())
		if err != nil {
			time.Sleep(time.Second * wait)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}