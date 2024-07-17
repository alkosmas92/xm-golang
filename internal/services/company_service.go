package services

import (
	"context"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/repository"
	"github.com/google/uuid"
)

type CompanyService interface {
	GetCompanyByCompanyID(ctx context.Context, companyID uuid.UUID) (*models.Company, error)
	CreateCompany(ctx context.Context, company *models.Company) error
	UpdateCompany(ctx context.Context, companyID uuid.UUID, company *models.Company) error
	DeleteCompany(ctx context.Context, companyID uuid.UUID) error
}

type companyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) CompanyService {
	return &companyService{repo: repo}
}

func (s *companyService) GetCompanyByCompanyID(ctx context.Context, companyID uuid.UUID) (*models.Company, error) {
	return s.repo.GetCompanyByCompanyID(ctx, companyID)
}

func (s *companyService) CreateCompany(ctx context.Context, company *models.Company) error {
	return s.repo.CreateCompany(ctx, company)
}

func (s *companyService) UpdateCompany(ctx context.Context, companyID uuid.UUID, company *models.Company) error {
	return s.repo.UpdateCompany(ctx, companyID, company)
}

func (s *companyService) DeleteCompany(ctx context.Context, companyID uuid.UUID) error {
	return s.repo.DeleteCompany(ctx, companyID)
}
