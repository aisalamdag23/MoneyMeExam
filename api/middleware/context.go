package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := context.TODO()
		c.Set("context", appCtx)
	}
}
