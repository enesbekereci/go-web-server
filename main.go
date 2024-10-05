package main

import (
	"fmt"
	"go-web-server/db"
	"go-web-server/middlewares"
	routes "go-web-server/routes"
	"go-web-server/system"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	sys := system.System{}
	sys.Initialize()
	go sys.RunCounter()
	db.Seed()
	r := gin.Default()
	r.Use(middlewares.CounterMiddleware(&sys))
	r.Use(middlewares.CounterEndMiddleware(&sys))
	r.GET("/alive", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message":         "Server is up and runnig.",
			"total_request":   sys.Counter,
			"current_request": sys.ReadCurrent(),
		})
	})
	r.GET("/wait", func(ctx *gin.Context) {
		time.Sleep(10 * time.Second)
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "Waited 10 sec.",
		})
	})
	routes.AddV1Routes(r)
	routes.AddV2Routes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
