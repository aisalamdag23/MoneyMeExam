package structs

import "time"

type Request struct {
	AmountRequired string    `json:"AmountRequired" binding:"required,numeric"`
	Term           string    `json:"Term" binding:"required,numeric"`
	FirstName      string    `json:"FirstName" binding:"required,min=1,max=300"`
	LastName       string    `json:"LastName" binding:"required,min=1,max=300"`
	DateOfBirth    string    `json:"DateOfBirth" binding:"required,len=10"`
	TDateOfBirth   time.Time `json:"-"`
	Mobile         string    `json:"Mobile" binding:"required,numeric"`
	Email          string    `json:"Email" binding:"required,email"`
}
