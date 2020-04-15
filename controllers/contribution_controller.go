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

// CreateContribution .
func (contributionController *ContributionController) CreateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contribution models.Contribution
		contributionUUID, _ := uuid.NewUUID()

		err := json.NewDecoder(r.Body).Decode(&contribution)
		logError(err)

		contribution.ContributionUUID = contributionUUID

		contributionRepo := repositories.ContributionRepository{}
		contributionUUID = contributionRepo.CreateContribution(db, contribution)

		err = json.NewEncoder(w).Encode(contributionUUID)
		logError(err)
	}
}

// GetContribution .
func (contributionController *ContributionController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contribution models.Contribution
		contributionUUID := getUUID(r)

		contributionRepo := repositories.ContributionRepository{}
		contribution = contributionRepo.GetContribution(db, contributionUUID)

		err := json.NewEncoder(w).Encode(contribution)
		logError(err)
	}
}

// GetContributions .
func (contributionController *ContributionController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributionRepo := repositories.ContributionRepository{}
		contributions := contributionRepo.GetContributions(db)

		err := json.NewEncoder(w).Encode(contributions)
		logError(err)
	}
}

// UpdateContribution .
func (contributionController *ContributionController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contribution models.Contribution
		contributionUUID := getUUID(r)

		err := json.NewDecoder(r.Body).Decode(&contribution)
		logError(err)

		contribution.ContributionUUID = contributionUUID

		contributionRepo := repositories.ContributionRepository{}
		contributionRepo.UpdateContribution(db, contribution)

		err = json.NewEncoder(w).Encode(contribution)
		logError(err)
	}
}

// DeleteContribution .
func (contributionController *ContributionController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributionUUID := getUUID(r)

		contributionRepo := repositories.ContributionRepository{}
		contributionRepo.DeleteContribution(db, contributionUUID)
	}
}
