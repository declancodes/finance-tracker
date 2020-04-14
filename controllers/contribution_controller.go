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

		json.NewDecoder(r.Body).Decode(&contribution)
		contribution.ContributionUUID = contributionUUID

		contributionRepo := repositories.ContributionRepository{}
		contributionUUID = contributionRepo.CreateContribution(db, contribution)

		json.NewEncoder(w).Encode(contributionUUID)
	}
}

// GetContribution .
func (contributionController *ContributionController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contribution models.Contribution
		contributionUUID := getUUID(r)

		contributionRepo := repositories.ContributionRepository{}
		contribution = contributionRepo.GetContribution(db, contributionUUID)

		json.NewEncoder(w).Encode(contribution)
	}
}

// GetContributions .
func (contributionController *ContributionController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contributionRepo := repositories.ContributionRepository{}
		contributions := contributionRepo.GetContributions(db)

		json.NewEncoder(w).Encode(contributions)
	}
}

// UpdateContribution .
func (contributionController *ContributionController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contribution models.Contribution
		contributionUUID := getUUID(r)

		json.NewDecoder(r.Body).Decode(&contribution)
		contribution.ContributionUUID = contributionUUID

		contributionRepo := repositories.ContributionRepository{}
		contributionRepo.UpdateContribution(db, contribution)

		json.NewEncoder(w).Encode(contribution)
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
