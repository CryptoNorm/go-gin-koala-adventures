package controller

import (
	"github.com/gin-gonic/gin"
)

type gameEventController interface {
	FindByPlayer(ctx *gin.Context) []model.gameEvent
	FindAll(ctx *gin.Context) []model.gameEvent
	Save(ctx *gin.Context) model.gameEvent
}

type gameEventController struct {
	service service.gameEventService
}

func NewEvent(service service.gameEventService) gameEventController {
	return &gameEventController{
		service: service,
	}
}

func (c *gameEventController) FindAll(ctx *gin.Context) []model.gameEvent {
	return c.service.FindAll()
}

func (c *gameEventController) Save(ctx *gin.Context) model.gameEvent {
	var gameEvent model.gameEvent

	ctx.BindJSON(&gameEvent)

	c.service.Save(gameEvent)
	return gameEvent
}

func (c *gameEventController) FindByVin(ctx *gin.Context) []model.gameEvent {
	vin := ctx.Param("vin")

	return c.service.FindByVin(vin)
}
