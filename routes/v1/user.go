package v1

import (
	"go-web-server/controllers"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.GET("/", controllers.GetUsers())
	user.GET("/:user", controllers.GetUser())
	user.POST("/", controllers.AddUser())
	user.DELETE("/:user_id", controllers.DeleteUser())
}
