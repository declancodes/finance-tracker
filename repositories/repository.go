package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

// ErrNoRecord is returned by any operation that is performed for a nonexistent record.
var ErrNoRecord = errors.New("repositories: record does not exist")

// ErrFiltersToMap is returned by any operations expecting filters to map but having differences between those provided.
var ErrFiltersToMap = errors.New("repositories: mismatched filters to map provided")

func buildQueryClauses(mValues map[string]interface{}, mFilters map[string]string) (string, []interface{}, error) {
	var values []interface{}
	var conditions []string

	if len(mValues) != len(mFilters) {
		return "", []interface{}{}, ErrFiltersToMap
	}

	if len(mValues) > 0 {
		for k := range mFilters {
			if val, ok := mValues[k]; ok {
				values = append(values, val)
				conditions = append(conditions, fmt.Sprintf("%s$%d", mFilters[k], len(values)))
			} else {
				return "", []interface{}{}, ErrFiltersToMap
			}
		}
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE"
	}

	return fmt.Sprintf("%s %s;", where, strings.Join(conditions, " AND ")), values, nil
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
