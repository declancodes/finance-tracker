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

func badRequestContribution(w http.ResponseWriter, err error) {
	badRequestModel(w, "contribution", err)
}

func errorExecutingContribution(w http.ResponseWriter, err error) {
	errorExecuting(w, "contribution", err)
}

// CreateContribution .
func (c *ContributionController) CreateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Contribution
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			badRequestContribution(w, err)
			return
		}

		c.ContributionUUID, _ = uuid.NewUUID()
		cUUID, err := contributionRepo.CreateContribution(db, c)
		if err != nil {
			errorCreating(w, "contribution", err)
			return
		}

		err = json.NewEncoder(w).Encode(cUUID)
		logError(err)
	}
}

// GetContribution .
func (c *ContributionController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		c, err := contributionRepo.GetContribution(db, cUUID)
		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(c)
		logError(err)
	}
}

// GetContributions .
func (c *ContributionController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := contributionRepo.GetContributions(db)
		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(cs)
		logError(err)
	}
}

// UpdateContribution .
func (c *ContributionController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var c models.Contribution
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			badRequestContribution(w, err)
			return
		}

		c.ContributionUUID = cUUID
		err = contributionRepo.UpdateContribution(db, c)
		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(c)
		logError(err)
	}
}

// DeleteContribution .
func (c *ContributionController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "contribution", contributionRepo.DeleteContribution)
	}
}
