package handlers

import (
	"github.com/aisalamdag23/MoneyMeExam/api/structs"
	"github.com/gin-gonic/gin"
)

func PostLoanRequest(ctx *gin.Context) {
	var request structs.PostLoanRequest
	if err := ctx.Bind(&request); err != nil {

	}
}
