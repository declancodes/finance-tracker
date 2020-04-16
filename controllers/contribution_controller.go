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
func (contributionController *ContributionController) CreateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contribution models.Contribution

		err := json.NewDecoder(r.Body).Decode(&contribution)
		if err != nil {
			writeHeaderForBadRequestModel(w, "contribution", err)
			return
		}

		contribution.ContributionUUID, _ = uuid.NewUUID()
		contributionUUID := contributionRepo.CreateContribution(db, contribution)

		err = json.NewEncoder(w).Encode(contributionUUID)
		logError(err)
	}
}

// GetContribution .
func (contributionController *ContributionController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributionUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		contribution := contributionRepo.GetContribution(db, contributionUUID)

		err = json.NewEncoder(w).Encode(contribution)
		logError(err)
	}
}

// GetContributions .
func (contributionController *ContributionController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributions := contributionRepo.GetContributions(db)

		err := json.NewEncoder(w).Encode(contributions)
		logError(err)
	}
}

// UpdateContribution .
func (contributionController *ContributionController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributionUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var contribution models.Contribution

		err = json.NewDecoder(r.Body).Decode(&contribution)
		if err != nil {
			writeHeaderForBadRequestModel(w, "contribution", err)
			return
		}

		contribution.ContributionUUID = contributionUUID
		contributionRepo.UpdateContribution(db, contribution)

		err = json.NewEncoder(w).Encode(contribution)
		logError(err)
	}
}

// DeleteContribution .
func (contributionController *ContributionController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributionUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		contributionRepo.DeleteContribution(db, contributionUUID)
	}
}
