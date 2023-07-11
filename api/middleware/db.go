package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CtxKey string

const (
	DBCtxKey CtxKey = "db"
)

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := c.MustGet("context").(context.Context)
		appCtx = context.WithValue(appCtx, DBCtxKey, db)
		c.Set("context", appCtx)
		c.Next()
	}
}
