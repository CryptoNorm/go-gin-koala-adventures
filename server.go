package main

import (
	"github.com/CryptoNorm/go-gin-koala-adventures-api/controller"
	"github.com/CryptoNorm/go-gin-koala-adventures-api/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	gameEventService    service.gameEventService       = service.NewEvent()
	gameEventController controller.gameEventController = controller.NewEvent(gameEventService)
)

func main() {
	router := gin.Default()
	// Enable CORS for requests UI domain (port)
	router.Use(cors.Default())

	router.GET("/gameEvents/:player", func(ctx *gin.Context) {
		ctx.JSON(200, gameEventController.FindByPlayer(ctx))
	})

	router.GET("/gameEvents", func(ctx *gin.Context) {
		ctx.JSON(200, gameEventController.FindAll(ctx))
	})

	router.POST("/gameEvents", func(ctx *gin.Context) {
		ctx.JSON(200, gameEventController.Save(ctx))
	})

	router.Run(":8082")

}
