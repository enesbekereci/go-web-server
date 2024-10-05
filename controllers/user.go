package controllers

import (
	"go-web-server/models"
	"go-web-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, error := services.GetUsersService()
		if error == nil {
			ctx.JSON(http.StatusOK, result)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": error})
		}
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ctx.Param("name")
		desc := ctx.Query("with_description")
		ctx.JSON(http.StatusOK, gin.H{"name": userName, "description": desc})
	}
}

func AddUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user AddUserType
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		//add user service
		ctx.JSON(http.StatusOK, gin.H{"message": "User Added"})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userType := ctx.Param("type")
		ctx.JSON(http.StatusOK, gin.H{"message": "done" + userType})
	}
}

type AddUserType struct {
	Name        string          `json:"name"`
	Password    string          `json:"password"`
	UserRole    models.UserRole `json:"user_role"`
	Description *string         `json:"description"`
}
