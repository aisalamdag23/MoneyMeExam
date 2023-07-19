package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/aisalamdag23/MoneyMeExam/api/middleware"
	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/aisalamdag23/MoneyMeExam/app/service"
	"github.com/aisalamdag23/MoneyMeExam/config"
	"github.com/gin-gonic/gin"
)

func PostLoanRequest(ctx *gin.Context) {
	// bind request data - return failed when error
	var request structs.Request

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

	// 1.1 save data into db
	appCtx := ctx.MustGet("context").(context.Context)
	serv := service.New(appCtx)
	loanAppID, err := serv.CreateLoanApplication(request)
	if err != nil {
		log.Printf("createloanapplication failed: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		}
		ctx.JSON(http.StatusInternalServerError, resp)
	}

	// 1.2 return url redirection
	cfg := appCtx.Value(middleware.CfgCtxKey).(*config.Config)
	if cfg == nil {
		errStr := "config nil"
		log.Println(errStr)
		resp := structs.ErrorResponse{
			Message: errStr,
			Code:    fmt.Sprintf("%d_%s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		}
		ctx.JSON(http.StatusInternalServerError, resp)
	}

	if cfg != nil && (cfg.Portal.BaseURL == "" || cfg.Portal.QuoteCalcURL == "") {
		errStr := "config base url or quote calc url is empty"
		log.Println(errStr)
		resp := structs.ErrorResponse{
			Message: errStr,
			Code:    fmt.Sprintf("%d_%s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		}
		ctx.JSON(http.StatusInternalServerError, resp)
	}

	quoteCalcURL := fmt.Sprintf("%s/%s?id=%d", cfg.Portal.BaseURL, cfg.Portal.QuoteCalcURL, *loanAppID)
	resp := structs.LoanApplicationRedirection{URL: quoteCalcURL}

	ctx.JSON(http.StatusOK, resp)
}
