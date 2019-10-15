package dbutils

import (
	"database/sql"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"time"
)

// tries- default 0, runs 1000 attempts
// wait - default 0, wait 1 second
func AttemptConnect(driver string, ds DataSource, tries int, wait time.Duration) (*sql.DB, bool, error) {
	var db *sql.DB
	var err error

	if wait == 0 {
		wait = time.Second
	}
	if tries == 0 {
		tries = 1000
	}

	var result *multierror.Error
	for i := 1; i <= tries; i++ {
		db, err = open(driver, ds.String())
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("dbutils.AttemptConnect: attempt number: %d error: %w", i, err))
			time.Sleep(wait)
			continue
		}
		return db, true, result.ErrorOrNil()
	}

	return nil, false, result
}

func open(driver, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, fmt.Errorf("dbutils.open: error while opening db conn: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("dbutils.open: error while pinging db conn: %w", err)
	}
	return db, nil
}
