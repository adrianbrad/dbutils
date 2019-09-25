package dbutils

import (
	"database/sql"
	"fmt"
	"time"
)

type errors []error

func (e errors) Error() string {
	return ""
}

// tries- default 0, runs 1000 attempts
// wait - default 0, wait 1 second
func AttemptConnect(driver string, ds DataSource, tries int, wait time.Duration) (*sql.DB, bool, error) {
	var db *sql.DB
	var err error

	if wait == 0 {
		wait = 1
	}
	if tries == 0 {
		tries = 1000
	}

	var errs errors
	var ok bool
	for i := 1; i <= tries; i++ {
		db, err = sql.Open(driver, ds.String())
		if err != nil {
			errs = append(errs, fmt.Errorf("dbutils.AttemptConnect: attempt number: %d error :%w", i, err))
			time.Sleep(time.Second * wait)
			continue
		}
		ok = true
		break
	}
	if !ok {
		return nil, false, errs
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, false, fmt.Errorf("dbutils.AttemptConnect: while pinging database: %w", err)
	}

	return db, true, errs
}
