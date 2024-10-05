package v1

import (
	"go-web-server/controllers"

	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.GET("/login", controllers.Login())
	auth.POST("/logout", controllers.Logout())
}
