package handlers_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/alkosmas92/xm-golang/internal/handlers"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/repository"
	"github.com/alkosmas92/xm-golang/internal/services"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	query := `
	CREATE TABLE company (
		CompanyID TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		description TEXT,
		amountofemployees INTEGER NOT NULL,
		registered BOOLEAN NOT NULL,
		type TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestIntegrationCompanyHandler(t *testing.T) {
	// Set up in-memory database
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize repository, service, and handler
	repo := repository.NewCompanyRepository(db)
	service := services.NewCompanyService(repo)
	logger := logrus.New()
	handler := handlers.NewCompanyHandler(service, logger)

	// Test create company
	t.Run("CreateCompany", func(t *testing.T) {
		company := &models.Company{
			Name:              "TechCorp1",
			Description:       "A technology company",
			AmountOfEmployees: 150,
			Registered:        true,
			Type:              "Corporations",
		}
		body, _ := json.Marshal(company)
		req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handler.CreateCompany(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	// Test get company by ID
	t.Run("GetCompanyByID", func(t *testing.T) {
		// Create a company first
		companyID := uuid.New()
		company := &models.Company{
			CompanyID:         companyID,
			Name:              "TechCorp2",
			Description:       "A technology company",
			AmountOfEmployees: 150,
			Registered:        true,
			Type:              "Corporations",
		}
		err := repo.CreateCompany(context.Background(), company)
		assert.NoError(t, err)

		// Retrieve the company
		req, _ := http.NewRequest("GET", "/companies?company_id="+companyID.String(), nil)
		rr := httptest.NewRecorder()

		handler.GetCompanyByID(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.Company
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, company.CompanyID, response.CompanyID)
		assert.Equal(t, company.Name, response.Name)
	})

	// Test update company
	t.Run("UpdateCompany", func(t *testing.T) {
		companyID := uuid.New()
		company := &models.Company{
			CompanyID:         companyID,
			Name:              "TechCorp3",
			Description:       "A technology company",
			AmountOfEmployees: 150,
			Registered:        true,
			Type:              "Corporations",
		}
		err := repo.CreateCompany(context.Background(), company)
		assert.NoError(t, err)

		company.Description = "An updated description"
		body, _ := json.Marshal(company)
		req, _ := http.NewRequest("PUT", "/companies?company_id="+companyID.String(), bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handler.UpdateCompany(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		updatedCompany, err := repo.GetCompanyByCompanyID(context.Background(), companyID)
		assert.NoError(t, err)
		assert.Equal(t, "An updated description", updatedCompany.Description)
	})

	// Test delete company
	t.Run("DeleteCompany", func(t *testing.T) {
		companyID := uuid.New()
		company := &models.Company{
			CompanyID:         companyID,
			Name:              "TechCorp4",
			Description:       "A technology company",
			AmountOfEmployees: 150,
			Registered:        true,
			Type:              "Corporations",
		}
		err := repo.CreateCompany(context.Background(), company)
		assert.NoError(t, err)

		req, _ := http.NewRequest("DELETE", "/companies?company_id="+companyID.String(), nil)
		rr := httptest.NewRecorder()

		handler.DeleteCompany(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		deletedCompany, err := repo.GetCompanyByCompanyID(context.Background(), companyID)
		assert.Nil(t, deletedCompany)
	})
}
