package service

import (
	"context"
	"log"

	"github.com/aisalamdag23/MoneyMeExam/api/middleware"
	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/aisalamdag23/MoneyMeExam/app/database/models"
	"gorm.io/gorm"
)

type (
	Service interface {
		CreateLoanApplication(req structs.Request) error
	}

	server struct {
		ctx context.Context
		db  *gorm.DB
	}
)

func New(ctx context.Context) Service {
	db := ctx.Value(middleware.DBCtxKey).(*gorm.DB)

	return &server{
		db:  db,
		ctx: ctx,
	}
}

func (s *server) CreateLoanApplication(req structs.Request) error {
	rec := models.LoanApplications{
		AmountRequired: req.AmountRequired,
		Term:           req.Term,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DateOfBirth:    req.TDateOfBirth,
		Mobile:         req.Mobile,
		Email:          req.Email,
	}
	if err := s.db.Create(&rec).Error; err != nil {
		log.Printf("db creation failed: %v", err)
		return err
	}

	return nil
}
