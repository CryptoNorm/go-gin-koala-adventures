package controller

import (
	"github.com/CryptoNorm/go-gin-koala-adventures-api/model"
	"github.com/CryptoNorm/go-gin-koala-adventures-api/service"
	"github.com/gin-gonic/gin"
)

type GameEventController interface {
	FindByPlayer(ctx *gin.Context) []model.GameEvent
	FindAll(ctx *gin.Context) []model.GameEvent
	Save(ctx *gin.Context) model.GameEvent
}

type gameEventController struct {
	service service.GameEventService
}

func NewEvent(service service.GameEventService) GameEventController {
	return &gameEventController{
		service: service,
	}
}

func (c *gameEventController) FindAll(ctx *gin.Context) []model.GameEvent {
	return c.service.FindAll()
}

func (c *gameEventController) Save(ctx *gin.Context) model.GameEvent {
	var gameEvent model.GameEvent

	ctx.BindJSON(&gameEvent)

	c.service.Save(gameEvent)
	return gameEvent
}

func (c *gameEventController) FindByPlayer(ctx *gin.Context) []model.GameEvent {
	player := ctx.Param("player")

	return c.service.FindByPlayer(player)
}
