// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/company_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/alkosmas92/xm-golang/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockCompanyService is a mock of CompanyService interface.
type MockCompanyService struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyServiceMockRecorder
}

// MockCompanyServiceMockRecorder is the mock recorder for MockCompanyService.
type MockCompanyServiceMockRecorder struct {
	mock *MockCompanyService
}

// NewMockCompanyService creates a new mock instance.
func NewMockCompanyService(ctrl *gomock.Controller) *MockCompanyService {
	mock := &MockCompanyService{ctrl: ctrl}
	mock.recorder = &MockCompanyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyService) EXPECT() *MockCompanyServiceMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockCompanyService) CreateCompany(ctx context.Context, company *models.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockCompanyServiceMockRecorder) CreateCompany(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockCompanyService)(nil).CreateCompany), ctx, company)
}

// DeleteCompany mocks base method.
func (m *MockCompanyService) DeleteCompany(ctx context.Context, companyID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", ctx, companyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockCompanyServiceMockRecorder) DeleteCompany(ctx, companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockCompanyService)(nil).DeleteCompany), ctx, companyID)
}

// GetCompanysByUserID mocks base method.
func (m *MockCompanyService) GetCompanysByUserID(ctx context.Context, companyID uuid.UUID) (*models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanysByUserID", ctx, companyID)
	ret0, _ := ret[0].(*models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanysByUserID indicates an expected call of GetCompanysByUserID.
func (mr *MockCompanyServiceMockRecorder) GetCompanysByUserID(ctx, companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanysByUserID", reflect.TypeOf((*MockCompanyService)(nil).GetCompanysByUserID), ctx, companyID)
}

// UpdateCompany mocks base method.
func (m *MockCompanyService) UpdateCompany(ctx context.Context, companyID uuid.UUID, company *models.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", ctx, companyID, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockCompanyServiceMockRecorder) UpdateCompany(ctx, companyID, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockCompanyService)(nil).UpdateCompany), ctx, companyID, company)
}