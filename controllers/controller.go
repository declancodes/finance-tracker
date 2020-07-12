package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func addJSONContentHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func created(w http.ResponseWriter, ID uuid.UUID) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(ID.String()))
}

func updated(w http.ResponseWriter, ID uuid.UUID) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ID.String()))
}

func delete(w http.ResponseWriter, r *http.Request, db *sqlx.DB, m string, fn func(*sqlx.DB, uuid.UUID) error) {
	ID, err := getID(r)
	if err != nil {
		badRequestID(w, err)
		return
	}

	err = fn(db, ID)
	if err != nil {
		errorExecuting(w, m, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func badRequestID(w http.ResponseWriter, err error) {
	badRequest(w, "invalid id", err)
}

func badRequestModel(w http.ResponseWriter, model string, err error) {
	badRequest(w, "invalid "+model, err)
}

func badRequest(w http.ResponseWriter, msg string, err error) {
	http.Error(w, msg, http.StatusBadRequest)
	log.Println(err)
}

func errorCreating(w http.ResponseWriter, m string, err error) {
	http.Error(w, "error creating "+m, http.StatusInternalServerError)
	log.Println(err)
}

func errorExecuting(w http.ResponseWriter, m string, err error) {
	switch err {
	case sql.ErrNoRows, repositories.ErrNoRecord:
		http.Error(w, m+" does not exist", http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(err)
}

func getFilters(r *http.Request) map[string]interface{} {
	q := r.URL.Query()

	accNames := getSlice(q, "account")
	catNames := getSlice(q, "category")
	fundSymbols := getSlice(q, "fund")

	start := getTime(q.Get("start"))
	end := getTime(q.Get("end"))

	mValues := make(map[string]interface{})
	if accNames != nil && len(accNames) > 0 {
		mValues["accounts"] = accNames
	}
	if catNames != nil && len(catNames) > 0 {
		mValues["categories"] = catNames
	}
	if fundSymbols != nil && len(fundSymbols) > 0 {
		mValues["funds"] = fundSymbols
	}
	if !start.IsZero() {
		mValues["start"] = start
	}
	if !end.IsZero() {
		mValues["end"] = end
	}

	return mValues
}

func getSlice(q url.Values, search string) []string {
	if vs, ok := q[search]; ok {
		return vs
	}
	return nil
}

func getTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}
	}
	return t
}

func getID(r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["id"])
	if err != nil {
		return uuid.Nil, err
	}
	return ID, nil
}
