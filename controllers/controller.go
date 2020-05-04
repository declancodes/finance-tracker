package controllers

import (
	"database/sql"
	"log"
	"net/http"
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
	ID, err := getUUID(r)
	if err != nil {
		badRequestUUID(w, err)
		return
	}

	err = fn(db, ID)
	if err != nil {
		errorExecuting(w, m, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func badRequestUUID(w http.ResponseWriter, err error) {
	badRequest(w, "invalid uuid", err)
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

func getTime(s string) time.Time {
	const dateFormat = "2006-_1-_2"

	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return time.Time{}
	}
	return t
}

func getUUID(r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["uuid"])
	if err != nil {
		return uuid.Nil, err
	}
	return ID, nil
}
