package routes

import (
	v1 "go-web-server/routes/v1"

	"github.com/gin-gonic/gin"
)

func AddV1Routes(r *gin.Engine) {
	// Simple group: v1
	v1Routes := r.Group("/v1")
	v1.AddUserRoutes(v1Routes)
	v1.AddKeyRoutes(v1Routes)
	v1.AddAuthRoutes(v1Routes)
}
