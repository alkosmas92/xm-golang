package models

import (
	"github.com/google/uuid"
)

type Company struct {
	CompanyID         uuid.UUID `json:"id" validate:"required"`
	Name              string    `json:"name" validate:"required,max=15,unique"`
	Description       string    `json:"description" validate:"max=3000"`
	AmountOfEmployees int       `json:"amount_of_employees" validate:"required"`
	Registered        bool      `json:"registered" validate:"required"`
	Type              string    `json:"type" validate:"required,oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}

func NewCompany(name string, description string, amountOfEmployees int, registered bool, companyType string) (*Company, error) {
	company := &Company{
		CompanyID:         uuid.New(),
		Name:              name,
		Description:       description,
		AmountOfEmployees: amountOfEmployees,
		Registered:        registered,
		Type:              companyType,
	}

	return company, nil
}
