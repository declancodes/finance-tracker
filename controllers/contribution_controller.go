package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ContributionController .
type ContributionController struct{}

var contributionRepo = repositories.ContributionRepository{}

// CreateContribution .
func (c *ContributionController) CreateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Contribution

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			writeHeaderForBadRequestModel(w, "contribution", err)
			return
		}

		c.ContributionUUID, _ = uuid.NewUUID()
		cUUID := contributionRepo.CreateContribution(db, c)

		err = json.NewEncoder(w).Encode(cUUID)
		logError(err)
	}
}

// GetContribution .
func (c *ContributionController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		c := contributionRepo.GetContribution(db, cUUID)

		err = json.NewEncoder(w).Encode(c)
		logError(err)
	}
}

// GetContributions .
func (c *ContributionController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cs := contributionRepo.GetContributions(db)

		err := json.NewEncoder(w).Encode(cs)
		logError(err)
	}
}

// UpdateContribution .
func (c *ContributionController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var c models.Contribution

		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			writeHeaderForBadRequestModel(w, "contribution", err)
			return
		}

		c.ContributionUUID = cUUID
		contributionRepo.UpdateContribution(db, c)

		err = json.NewEncoder(w).Encode(c)
		logError(err)
	}
}

// DeleteContribution .
func (c *ContributionController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		contributionRepo.DeleteContribution(db, cUUID)
	}
}
