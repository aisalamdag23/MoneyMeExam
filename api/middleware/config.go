package middleware

import (
	"context"

	"github.com/aisalamdag23/MoneyMeExam/config"
	"github.com/gin-gonic/gin"
)

const (
	CfgCtxKey CtxKey = "config"
)

func ConfigMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := c.MustGet("context").(context.Context)
		appCtx = context.WithValue(appCtx, CfgCtxKey, cfg)
		c.Set("context", appCtx)
		c.Next()
	}
}
