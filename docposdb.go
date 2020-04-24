package dbutils

import (
	"database/sql"
	"fmt"
	"github.com/ory/dockertest"
	"time"
)

func NewDockerPostgresDB(user, password, dbName, schema string) (func() error, error) {
	dockerStartWait := 60 * time.Second

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, err
	}

	res, err := pool.Run("postgres", "11-alpine", []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		fmt.Sprintf("POSTGRES_USER=%s", user),
		fmt.Sprintf("POSTGRES_DB=%s", dbName),
	})
	if err != nil {
		return nil, err
	}

	purge := func() { pool.Purge(res) }

	errChan := make(chan error)
	done := make(chan struct{})

	var db *sql.DB

	go func() {
		if err := pool.Retry(func() error {
			db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=localhost port=%s sslmode=disable", user, password, dbName, res.GetPort("5432/tcp")))
			if err != nil {
				return err
			}
			return db.Ping()
		}); err != nil {
			errChan <- err
		}

		close(done)
	}()

	select {
	case err := <-errChan:
		purge()
		return nil, err
	case <-time.After(dockerStartWait):
		purge()
		return nil, fmt.Errorf("timeout on checking postgres connection")
	case <-done:
		close(errChan)
	}

	defer db.Close()
	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	teardown := func() error {
		return pool.Purge(res)
	}

	return teardown, nil
}
