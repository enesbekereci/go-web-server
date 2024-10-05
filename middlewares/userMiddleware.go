package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userRole", "user")
		ctx.Next()
		status := ctx.Writer.Status()
		log.Println(status)
	}
}
