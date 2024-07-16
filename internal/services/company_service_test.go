package services

import (
	"context"
	"github.com/alkosmas92/xm-golang/internal/mocks"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompanyService_GetCompanysByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCompanyRepository(ctrl)
	service := NewCompanyService(mockRepo)

	company, _ := models.NewCompany("Test Company", "Description", 100, true, "Corporations")

	ctx := context.Background()
	mockRepo.EXPECT().GetCompanysByUserID(ctx, company.CompanyID).Return(company, nil)

	company, err := service.GetCompanysByUserID(ctx, company.CompanyID)

	assert.NoError(t, err)
	assert.Equal(t, company, company)
}

func TestCompanyService_CreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCompanyRepository(ctrl)
	service := NewCompanyService(mockRepo)

	company, _ := models.NewCompany("Test Company", "Description", 100, true, "Corporations")

	ctx := context.Background()
	mockRepo.EXPECT().CreateCompany(ctx, company).Return(nil)

	err := service.CreateCompany(ctx, company)

	assert.NoError(t, err)
}

func TestCompanyService_UpdateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCompanyRepository(ctrl)
	service := NewCompanyService(mockRepo)

	company, _ := models.NewCompany("Test Company", "Description", 100, true, "Corporations")

	ctx := context.Background()
	mockRepo.EXPECT().UpdateCompany(ctx, company.CompanyID, company).Return(nil)

	err := service.UpdateCompany(ctx, company.CompanyID, company)

	assert.NoError(t, err)
}

func TestCompanyService_DeleteCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCompanyRepository(ctrl)
	service := NewCompanyService(mockRepo)

	companyID := uuid.New()

	ctx := context.Background()
	mockRepo.EXPECT().DeleteCompany(ctx, companyID).Return(nil)

	err := service.DeleteCompany(ctx, companyID)

	assert.NoError(t, err)
}
