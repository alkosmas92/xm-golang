package repository

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/google/uuid"
)

// CompanyRepository provides access to the company storage.
type CompanyRepository interface {
	GetCompanyByCompanyID(ctx context.Context, companyID uuid.UUID) (*models.Company, error)
	CreateCompany(ctx context.Context, company *models.Company) error
	UpdateCompany(ctx context.Context, companyID uuid.UUID, company *models.Company) error
	DeleteCompany(ctx context.Context, companyID uuid.UUID) error
}

// companyRepository provides access to the company database.
type companyRepository struct {
	db *sql.DB
}

// NewCompanyRepository creates a new CompanyRepository.
func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) GetCompanyByCompanyID(ctx context.Context, companyID uuid.UUID) (*models.Company, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		query := `
			SELECT companyID, name, description, amountOfEmployees, registered, type
			FROM company
			WHERE companyID = ?`

		row := r.db.QueryRowContext(ctx, query, companyID)

		var company models.Company
		if err := row.Scan(&company.CompanyID, &company.Name, &company.Description, &company.AmountOfEmployees, &company.Registered, &company.Type); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil // No company found
			}
			return nil, err
		}

		return &company, nil
	}
}

func (r *companyRepository) CreateCompany(ctx context.Context, company *models.Company) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := `
			INSERT INTO company (CompanyID, Name, Description, AmountOfEmployees, Registered, Type)
			VALUES (?, ?, ?, ?, ?, ?)`
		_, err := r.db.ExecContext(ctx, query, company.CompanyID, company.Name, company.Description, company.AmountOfEmployees, company.Registered, company.Type)
		return err
	}
}

func (r *companyRepository) UpdateCompany(ctx context.Context, companyID uuid.UUID, company *models.Company) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := `
			UPDATE company
			SET name = ?, description = ?, amountOfEmployees = ?, registered = ?, type = ?
			WHERE companyID = ?`
		_, err := r.db.ExecContext(ctx, query, company.Name, company.Description, company.AmountOfEmployees, company.Registered, company.Type, companyID)
		return err
	}
}

func (r *companyRepository) DeleteCompany(ctx context.Context, companyID uuid.UUID) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := "DELETE FROM company WHERE companyID = ?"
		_, err := r.db.ExecContext(ctx, query, companyID)
		return err
	}
}
