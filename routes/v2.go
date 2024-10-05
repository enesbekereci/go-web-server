package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddV2Routes(r *gin.Engine) {
	// Simple group: v1
	v2Routes := r.Group("/v2")

	v2Routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Not Ready",
		})
	})
}
