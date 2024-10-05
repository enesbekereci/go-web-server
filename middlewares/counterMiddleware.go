package middlewares

import (
	"go-web-server/system"
	"log"

	"github.com/gin-gonic/gin"
)

func CounterMiddleware(sys *system.System) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sys.CounterChan <- 1
		ctx.Set("Counter", "Done")
		ctx.Next()
	}
}

func CounterEndMiddleware(sys *system.System) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		sys.CounterChan <- -1
		ctx.Set("CounterEnd", "Done")

	}
}

func ErrorMiddleware(sys *system.System) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Counter", "Done")
		ctx.Next()
		status := ctx.Writer.Status()
		log.Println(status)
	}
}
