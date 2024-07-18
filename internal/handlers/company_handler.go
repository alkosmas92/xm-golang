package handlers

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/services"
	"github.com/google/uuid"
)

// CompanyHandler handles HTTP requests for company-related operations.
type CompanyHandler struct {
	Service services.CompanyService
	Logger  *logrus.Logger
}

// NewCompanyHandler create a new CompanyHandler.
func NewCompanyHandler(service services.CompanyService, logger *logrus.Logger) *CompanyHandler {
	return &CompanyHandler{Service: service, Logger: logger}
}

// CreateCompany handles creation of a new company.
func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	company.CompanyID = uuid.New() // Generate a new UUID for the company

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	if err := h.Service.CreateCompany(ctx, &company); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"company_id": company.CompanyID.String(),
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Error("failed to encode response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetCompanyByID handles fetching a company by its ID.
func (h *CompanyHandler) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	companyIDStr := r.URL.Query().Get("company_id")
	if companyIDStr == "" {
		http.Error(w, "Missing company_id query parameter", http.StatusBadRequest)
		return
	}

	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		http.Error(w, "Invalid company_id format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	company, err := h.Service.GetCompanyByCompanyID(ctx, companyID)
	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// UpdateCompany handles updating an existing company.
func (h *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	companyID, err := uuid.Parse(r.URL.Query().Get("company_id"))
	if err != nil {
		http.Error(w, "Invalid company_id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.Service.UpdateCompany(ctx, companyID, &company); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteCompany handles deleting a company.
func (h *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	companyID, err := uuid.Parse(r.URL.Query().Get("company_id"))
	if err != nil {
		http.Error(w, "Invalid company_id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.Service.DeleteCompany(ctx, companyID); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
