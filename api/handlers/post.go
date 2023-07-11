package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/aisalamdag23/MoneyMeExam/app/service"
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
	if !tDob.Before(now) || now.Sub(tDob) < 18 {
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
	serv := service.New(ctx.MustGet("context").(context.Context))
	err = serv.CreateLoanApplication(request)
	if err != nil {
		log.Printf("createloadapplication failed: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		}
		ctx.JSON(http.StatusBadRequest, resp)
	}

	// 1.2 return url redirection

}
