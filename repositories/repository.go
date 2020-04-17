package repositories

import (
	"database/sql"
	"errors"
	"log"
)

// ErrNoRecord is returned by any operation that is performed for a nonexistent record.
var ErrNoRecord = errors.New("repositories: record does not exist")

func getExecuted(r sql.Result, err error) (int64, error) {
	if err != nil {
		log.Println(err)
		return 0, err
	}

	updated, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	if updated == 0 {
		return 0, ErrNoRecord
	}
	return updated, nil
}
