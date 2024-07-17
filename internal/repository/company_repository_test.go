package repository

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/xm-golang/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTestCompanyDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	createTable := `
	CREATE TABLE companies (
		companyID TEXT PRIMARY KEY,
		name TEXT,
		description TEXT,
		amountOfEmployees INTEGER,
		registered BOOLEAN,
		type TEXT
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}

	return db
}

func TestCompanyRepository(t *testing.T) {
	db := setupTestCompanyDB(t)
	defer db.Close()

	repo := NewCompanyRepository(db)
	ctx := context.Background()

	company, _ := models.NewCompany("Test Company", "Description", 100, true, "Corporations")

	err := repo.CreateCompany(ctx, company)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	retrievedCompany, err := repo.GetCompanyByCompanyID(ctx, company.CompanyID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, company.CompanyID, retrievedCompany.CompanyID)
	assert.Equal(t, company.Name, retrievedCompany.Name)
	assert.Equal(t, company.Description, retrievedCompany.Description)
	assert.Equal(t, company.AmountOfEmployees, retrievedCompany.AmountOfEmployees)
	assert.Equal(t, company.Registered, retrievedCompany.Registered)
	assert.Equal(t, company.Type, retrievedCompany.Type)

	updatedCompany := &models.Company{
		CompanyID:         company.CompanyID,
		Name:              "Updated Company",
		Description:       "Updated Description",
		AmountOfEmployees: 150,
		Registered:        false,
		Type:              "NonProfit",
	}

	err = repo.UpdateCompany(ctx, updatedCompany.CompanyID, updatedCompany)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	retrievedCompany, err = repo.GetCompanyByCompanyID(ctx, updatedCompany.CompanyID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, updatedCompany.CompanyID, retrievedCompany.CompanyID)
	assert.Equal(t, updatedCompany.Name, retrievedCompany.Name)
	assert.Equal(t, updatedCompany.Description, retrievedCompany.Description)
	assert.Equal(t, updatedCompany.AmountOfEmployees, retrievedCompany.AmountOfEmployees)
	assert.Equal(t, updatedCompany.Registered, retrievedCompany.Registered)
	assert.Equal(t, updatedCompany.Type, retrievedCompany.Type)

	err = repo.DeleteCompany(ctx, company.CompanyID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	retrievedCompany, err = repo.GetCompanyByCompanyID(ctx, company.CompanyID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Nil(t, retrievedCompany)
}
