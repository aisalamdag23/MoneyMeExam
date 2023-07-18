package service

import (
	"context"
	"log"
	"strconv"

	"github.com/aisalamdag23/MoneyMeExam/api/middleware"
	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/aisalamdag23/MoneyMeExam/app/database/models"
	gofinancial "github.com/razorpay/go-financial"
	"github.com/razorpay/go-financial/enums/paymentperiod"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type (
	Service interface {
		CreateLoanApplication(req structs.Request) (*uint, error)
		GetLoanApplication(id int) (*structs.LoanApplicationDetails, error)
		UpdateLoanApplication(req structs.UpdateLoanRequest, id int) error
		pmt(loanAmount decimal.Decimal, term int) decimal.Decimal
		ipmt(loanAmount decimal.Decimal, term int) decimal.Decimal
	}

	server struct {
		ctx context.Context
		db  *gorm.DB
	}
)

const (
	LOAN_PERCENTAGE = 0.092
	DEFAULT_FEE     = 300.00
)

func New(ctx context.Context) Service {
	db := ctx.Value(middleware.DBCtxKey).(*gorm.DB)

	return &server{
		db:  db,
		ctx: ctx,
	}
}

func (s *server) CreateLoanApplication(req structs.Request) (*uint, error) {
	rec := models.LoanApplications{
		AmountRequired: req.AmountRequired,
		Term:           req.Term,
		Title:          req.Title,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DateOfBirth:    req.TDateOfBirth,
		Mobile:         req.Mobile,
		Email:          req.Email,
	}
	if err := s.db.Create(&rec).Error; err != nil {
		log.Printf("db creation failed: %v", err)
		return nil, err
	}

	return &rec.ID, nil
}

func (s *server) GetLoanApplication(id int) (*structs.LoanApplicationDetails, error) {
	loanApp := models.LoanApplications{}

	if err := s.db.First(&loanApp, id).Error; err != nil {
		log.Printf("db retrieval failed: %v", err)
		return nil, err
	}

	resp := structs.LoanApplicationDetails{
		AmountRequired:   loanApp.AmountRequired,
		Term:             loanApp.Term,
		Title:            loanApp.Title,
		FirstName:        loanApp.FirstName,
		LastName:         loanApp.LastName,
		DateOfBirth:      loanApp.DateOfBirth.Format("2006-01-02"),
		Mobile:           loanApp.Mobile,
		Email:            loanApp.Email,
		Repayment:        loanApp.Repayment,
		TotalInterest:    loanApp.TotalInterest,
		EstablishmentFee: loanApp.EstablishmentFee,
	}

	return &resp, nil
}

func (s *server) UpdateLoanApplication(req structs.UpdateLoanRequest, id int) error {
	decLoanAmt, err := decimal.NewFromString(req.AmountRequired)
	if err != nil {
		log.Printf("string to decimal failed: %v", err)
		return err
	}

	intTerm, err := strconv.Atoi(req.Term)
	if err != nil {
		log.Printf("string to int failed: %v", err)
		return err
	}

	rec := models.LoanApplications{
		AmountRequired:   req.AmountRequired,
		Term:             req.Term,
		Title:            req.Title,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		DateOfBirth:      req.TDateOfBirth,
		Mobile:           req.Mobile,
		Email:            req.Email,
		Repayment:        s.pmt(decLoanAmt, intTerm).String(),
		TotalInterest:    s.ipmt(decLoanAmt, intTerm).String(),
		EstablishmentFee: strconv.Itoa(DEFAULT_FEE),
	}

	if err := s.db.Model(models.LoanApplications{}).Where("id = ?", id).Updates(rec).Error; err != nil {
		log.Printf("db update failed: %v", err)
		return err
	}

	return nil
}

func (s *server) pmt(loanAmount decimal.Decimal, term int) decimal.Decimal {
	rate := decimal.NewFromFloat(LOAN_PERCENTAGE / 12)
	nper := int64(term)
	pv := loanAmount
	fv := decimal.NewFromInt(0)
	when := paymentperiod.ENDING

	return gofinancial.Pmt(rate, nper, pv, fv, when).Round(2).Abs()

}

func (s *server) ipmt(loanAmount decimal.Decimal, term int) decimal.Decimal {
	rate := decimal.NewFromFloat(LOAN_PERCENTAGE / 12)
	nper := int64(term)
	pv := loanAmount
	fv := decimal.NewFromInt(0)
	when := paymentperiod.ENDING
	var total decimal.Decimal

	for i := int64(0); i < nper; i++ {
		pmt := gofinancial.IPmt(rate, i+1, nper, pv, fv, when)
		total = total.Add(pmt.Round(0).Abs())
	}

	return total.Round(2)
}
