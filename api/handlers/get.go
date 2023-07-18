package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/aisalamdag23/MoneyMeExam/app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLoanQuote(ctx *gin.Context) {
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

	serv := service.New(ctx.MustGet("context").(context.Context))
	loanDetail, err := serv.GetLoanApplication(loanID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("getloanapplication failed: %v", err)
			resp := structs.ErrorResponse{
				Message: err.Error(),
				Code:    fmt.Sprintf("%d_%s", http.StatusNotFound, http.StatusText(http.StatusNotFound)),
			}
			ctx.JSON(http.StatusNotFound, resp)
			return
		}
		log.Printf("getloanapplication failed: %v", err)
		resp := structs.ErrorResponse{
			Message: err.Error(),
			Code:    fmt.Sprintf("%d_%s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx.JSON(http.StatusOK, *loanDetail)

}
