package v1

import (
	"go-web-server/controllers"

	"github.com/gin-gonic/gin"
)

func AddKeyRoutes(r *gin.RouterGroup) {
	user := r.Group("/key")
	user.GET("/:key", controllers.GetUser())
	user.POST("/", controllers.AddUser())
	user.DELETE("/:key_id", controllers.DeleteUser())
}
