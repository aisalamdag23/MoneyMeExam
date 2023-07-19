package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/aisalamdag23/MoneyMeExam/app/service"
	"github.com/gin-gonic/gin"
)

func CalculateLoanQuote(ctx *gin.Context) {
	loanIDStr := ctx.Param("id")
	loanID, err := strconv.Atoi(loanIDStr)
	if err != nil {
		log.Printf("failed loan id string to int conversion: %v", err)
		resp := structs.ErrorResponse{
			Message: "invalid id",
			Code:    fmt.Sprintf("%d_%s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	var request structs.UpdateLoanRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("failed request binding: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	// dob validations
	tDob, err := time.Parse("2006-01-02", request.DateOfBirth)
	if err != nil {
		log.Printf("failed dob validation: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	now := time.Now()
	diff, _ := time.ParseDuration(now.Sub(tDob).String())
	if !tDob.Before(now) || diff.Hours() < 18*24*365 {
		err = errors.New("must be 18 years old and above")
		log.Printf("failed dob validation: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	request.TDateOfBirth = tDob

	serv := service.New(ctx.MustGet("context").(context.Context))
	err = serv.UpdateLoanApplication(request, loanID)
	if err != nil {
		log.Printf("updateloanapplication failed: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		}
		ctx.JSON(http.StatusInternalServerError, resp)
	}

	ctx.JSON(http.StatusOK, structs.SuccessResponse{Message: "success"})
}
