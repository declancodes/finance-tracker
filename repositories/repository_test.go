package repositories

import (
	"testing"
)

func TestBuildQueryClauses(t *testing.T) {
	tables := []struct {
		mValues         map[string]interface{}
		mFilters        map[string]string
		expectedClauses string
		expectedValues  []interface{}
	}{
		{
			map[string]interface{}{
				"category": "foo",
			},
			map[string]string{
				"category": "account_category.name = ",
			},
			"WHERE account_category.name = $1;",
			[]interface{}{"foo"},
		},
	}

	for _, table := range tables {
		actualClauses, actualValues := buildQueryClauses(table.mValues, table.mFilters)

		if actualClauses != table.expectedClauses {
			t.Errorf("Query clauses built were incorrect, got: %s, want: %s", actualClauses, table.expectedClauses)
		}

		for i, v := range actualValues {
			if v != table.expectedValues[i] {
				t.Errorf("Query values built were incorrect, got: %v, want: %v", v, table.expectedValues[i])
			}
		}
	}
}
