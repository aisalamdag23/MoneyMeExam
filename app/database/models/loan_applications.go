package models

import (
	"time"

	"gorm.io/gorm"
)

type LoanApplications struct {
	gorm.Model
	AmountRequired   string    `gorm:"column:amount_required"`
	Term             string    `gorm:"column:term"`
	Title            string    `gorm:"column:title"`
	FirstName        string    `gorm:"column:first_name"`
	LastName         string    `gorm:"column:last_name"`
	DateOfBirth      time.Time `gorm:"column:date_of_birth"`
	Mobile           string    `gorm:"column:mobile"`
	Email            string    `gorm:"column:email"`
	Repayment        string    `gorm:"column:repayment"`
	EstablishmentFee string    `gorm:"column:establishment_fee"`
	TotalInterest    string    `gorm:"column:total_interest"`
}

func (LoanApplications) TableName() string {
	return "loan_applications"
}
