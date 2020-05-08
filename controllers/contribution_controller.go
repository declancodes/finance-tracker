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

		c.ID, _ = uuid.NewUUID()
		cUUID, err := contributionRepo.CreateContribution(db, c)
		if err != nil {
			errorCreating(w, "contribution", err)
			return
		}

		created(w, cUUID)
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

		addJSONContentHeader(w)
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

		mValues := make(map[string]interface{})
		if accName != "" {
			mValues["account"] = accName
		}
		if catName != "" {
			mValues["category"] = catName
		}
		if !start.IsZero() {
			mValues["start"] = start
		}
		if !end.IsZero() {
			mValues["end"] = end
		}

		cs, err := contributionRepo.GetContributions(db, mValues)

		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		addJSONContentHeader(w)
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

		c.ID = cUUID
		err = contributionRepo.UpdateContribution(db, c)
		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		updated(w, c.ID)
	}
}

// DeleteContribution deletes a Contribution.
func (c *ContributionController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "contribution", contributionRepo.DeleteContribution)
	}
}
