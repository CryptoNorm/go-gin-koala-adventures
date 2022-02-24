package controller

import (
	"github.com/CryptoNorm/go-gin-koala-adventures-api/model"
	"github.com/CryptoNorm/go-gin-koala-adventures-api/service"
	"github.com/gin-gonic/gin"
)

type gameEventController interface {
	FindByPlayer(ctx *gin.Context) []model.GameEvent
	FindAll(ctx *gin.Context) []model.GameEvent
	Save(ctx *gin.Context) model.GameEvent
}

type gameEventController struct {
	service service.GameEventService
}

func NewEvent(service service.gameEventService) gameEventController {
	return &gameEventController{
		service: service,
	}
}

func (c *gameEventController) FindAll(ctx *gin.Context) []model.gGameEvent {
	return c.service.FindAll()
}

func (c *gameEventController) Save(ctx *gin.Context) model.GameEvent {
	var gameEvent model.gameEvent

	ctx.BindJSON(&gameEvent)

	c.service.Save(gameEvent)
	return gameEvent
}

func (c *gameEventController) FindByVin(ctx *gin.Context) []model.gameEvent {
	vin := ctx.Param("vin")

	return c.service.FindByVin(vin)
}
