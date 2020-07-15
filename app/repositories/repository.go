package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ErrNoRecord is returned by any operation that is performed for a nonexistent record.
var ErrNoRecord = errors.New("repositories: record does not exist")

func buildQueryClauses(baseQuery string, mValues map[string]interface{}, mFilters map[string]string) (string, []interface{}) {
	var values []interface{}
	var conditions []string

	for fKey, fVal := range mFilters {
		if val, ok := mValues[fKey]; ok {
			values = append(values, val)
			conditions = append(conditions, fVal+"(?)")
		}
	}

	var qb strings.Builder

	qb.WriteString(baseQuery)
	if len(conditions) > 0 {
		qb.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}
	qb.WriteString(";")

	return qb.String(), values
}

func getGetQueryAndValues(getQuery string, mValues map[string]interface{}, mFilters map[string]string) (string, []interface{}, error) {
	query, values := buildQueryClauses(getQuery, mValues, mFilters)

	q, args, err := sqlx.In(query, values...)
	if err != nil {
		return "", nil, err
	}
	return postgresRebind(q), args, nil
}

func createAndGetIDs(db *sqlx.DB, query string, es interface{}) ([]uuid.UUID, error) {
	q, args, err := sqlx.Named(query, es)
	if err != nil {
		return nil, err
	}

	q = postgresRebind(q)

	var IDs []uuid.UUID
	err = db.Select(&IDs, q, args...)
	if err != nil {
		return nil, err
	}
	return IDs, nil
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

func postgresRebind(query string) string {
	return sqlx.Rebind(sqlx.DOLLAR, query)
}
