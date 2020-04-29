package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ContributionController is the means for interacting with Contribution entities from an http router.
type ContributionController struct{}

var contributionRepo = repositories.ContributionRepository{}

func badRequestContribution(w http.ResponseWriter, err error) {
	badRequestModel(w, "contribution", err)
}

func errorExecutingContribution(w http.ResponseWriter, err error) {
	errorExecuting(w, "contribution", err)
}

// CreateContribution creates a Contribution based on the r *http.Request Body.
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

// GetContribution gets a Contribution.
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

// GetContributions gets Contribution entities.
func (c *ContributionController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		accName := q.Get("account")
		catName := q.Get("category")
		start := getTime(q.Get("start"))
		end := getTime(q.Get("end"))

		m := make(map[string]interface{})
		if accName != "" {
			m["account"] = accName
		}
		if catName != "" {
			m["category"] = catName
		}
		if !start.IsZero() {
			m["start"] = start
		}
		if !end.IsZero() {
			m["end"] = end
		}

		cs, err := contributionRepo.GetContributions(db, m)

		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(cs)
		logError(err)
	}
}

// UpdateContribution updates a Contribution.
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

// DeleteContribution deletes a Contribution.
func (c *ContributionController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "contribution", contributionRepo.DeleteContribution)
	}
}
