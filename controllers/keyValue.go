package controllers

import "github.com/gin-gonic/gin"

func GetKeyValue() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userType := ctx.Param("type")
		ctx.JSON(200, gin.H{"message": "GetUser" + userType})
	}
}

func AddKeyValue() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userType := ctx.Param("user")
		ctx.JSON(200, gin.H{"message": "done" + userType})
	}
}

func DeleteKeyValue() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userType := ctx.Param("type")
		ctx.JSON(200, gin.H{"message": "done" + userType})
	}
}
