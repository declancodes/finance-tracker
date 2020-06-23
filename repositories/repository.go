package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ErrNoRecord is returned by any operation that is performed for a nonexistent record.
var ErrNoRecord = errors.New("repositories: record does not exist")

// ErrFiltersToMap is returned by any operations expecting filters to map but having differences between those provided.
var ErrFiltersToMap = errors.New("repositories: mismatched filters to map provided")

func buildQueryClauses(mValues map[string]interface{}, mFilters map[string]string) (string, []interface{}, error) {
	var values []interface{}
	var conditions []string

	if len(mValues) > 0 {
		for k := range mFilters {
			if val, ok := mValues[k]; ok {
				values = append(values, val)
				conditions = append(conditions, fmt.Sprintf("%s(?)", mFilters[k]))
			}
		}
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE"
	}

	return fmt.Sprintf("%s %s;", where, strings.Join(conditions, " AND ")), values, nil
}

func getCreateQueryAndVals(query string, es interface{}) (string, []interface{}, error) {
	q, args, err := sqlx.Named(query, es)
	if err != nil {
		return "", nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	return q, args, nil
}

func createAndGetUUID(db *sqlx.DB, query string, e interface{}) (uuid.UUID, error) {
	rows, err := db.NamedQuery(query, e)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	var ID uuid.UUID
	for rows.Next() {
		err = rows.Scan(&ID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return ID, nil
}

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

func updateEntity(db *sqlx.DB, query string, e interface{}) error {
	res, err := db.NamedExec(query, e)
	_, err = getExecuted(res, err)
	return err
}

func deleteEntity(db *sqlx.DB, query string, ID uuid.UUID) error {
	res, err := db.Exec(query, ID.String())
	_, err = getExecuted(res, err)
	return err
}
