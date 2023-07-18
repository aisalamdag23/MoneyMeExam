package structs

import "time"

type Request struct {
	AmountRequired string    `json:"AmountRequired" binding:"required,numeric"`
	Term           string    `json:"Term" binding:"required,numeric"`
	Title          string    `json:"Title" binding:"required,oneof=mr. mrs. ms. MR. MRS. MS. Mr. Ms. Mrs. mR. mS. mRs. mRS. MrS."`
	FirstName      string    `json:"FirstName" binding:"required,min=1,max=300"`
	LastName       string    `json:"LastName" binding:"required,min=1,max=300"`
	DateOfBirth    string    `json:"DateOfBirth" binding:"required,len=10"`
	TDateOfBirth   time.Time `json:"-"`
	Mobile         string    `json:"Mobile" binding:"required,numeric"`
	Email          string    `json:"Email" binding:"required,email"`
}

type UpdateLoanRequest struct {
	AmountRequired string    `json:"AmountRequired" binding:"numeric"`
	Term           string    `json:"Term" binding:"numeric"`
	Title          string    `json:"Title" binding:"oneof=mr. mrs. ms. MR. MRS. MS. Mr. Ms. Mrs. mR. mS. mRs. mRS. MrS."`
	FirstName      string    `json:"FirstName" binding:"min=1,max=300"`
	LastName       string    `json:"LastName" binding:"min=1,max=300"`
	DateOfBirth    string    `json:"DateOfBirth" binding:"len=10"`
	TDateOfBirth   time.Time `json:"-"`
	Mobile         string    `json:"Mobile" binding:"numeric"`
	Email          string    `json:"Email" binding:"email"`
}
