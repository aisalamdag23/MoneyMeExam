package models

import (
	"time"

	"gorm.io/gorm"
)

type LoanApplications struct {
	gorm.Model
	AmountRequired string    `gorm:"column:amount_required"`
	Term           string    `gorm:"column:term"`
	FirstName      string    `gorm:"column:first_name"`
	LastName       string    `gorm:"column:last_name"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth"`
	Mobile         string    `gorm:"column:mobile"`
	Email          string    `gorm:"column:email"`
}

func (LoanApplications) TableName() string {
	return "loan_applications"
}
