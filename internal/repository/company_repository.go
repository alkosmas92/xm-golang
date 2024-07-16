package repository

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/google/uuid"
)

type CompanyRepository interface {
	GetCompanysByUserID(ctx context.Context, companyID uuid.UUID) (*models.Company, error)
	CreateCompany(ctx context.Context, company *models.Company) error
	UpdateCompany(ctx context.Context, companyID uuid.UUID, company *models.Company) error
	DeleteCompany(ctx context.Context, companyID uuid.UUID) error
}

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) GetCompanysByUserID(ctx context.Context, companyID uuid.UUID) (*models.Company, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		query := `
			SELECT id, name, description, amount_of_employees, registered, type
			FROM companies
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
			INSERT INTO companies (Company_id, Name, Description, AmountOfEmployees, Registered, Type)
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
			UPDATE companies
			SET name = ?, description = ?, amount_of_employees = ?, registered = ?, type = ?
			WHERE company_id = ?`
		_, err := r.db.ExecContext(ctx, query, company.Name, company.Description, company.AmountOfEmployees, company.Registered, company.Type, companyID)
		return err
	}
}

func (r *companyRepository) DeleteCompany(ctx context.Context, companyID uuid.UUID) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		query := "DELETE FROM companies WHERE company_id = ?"
		_, err := r.db.ExecContext(ctx, query, companyID)
		return err
	}
}
