package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ctx.Param("name")
		desc := ctx.Query("with_description")
		ctx.JSON(http.StatusOK, gin.H{"name": userName, "description": desc})
	}
}

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user AddUserType
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "done"})
	}
}
