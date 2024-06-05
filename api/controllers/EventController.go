package controllers

import (
	"infra-api/api/aws/s3"
	"infra-api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type eventConroller struct {
	events []entities.Event
}

func NewEventController() *eventConroller {
	return &eventConroller{}
}

func (e *eventConroller) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, e.events)
}

func (e *eventConroller) Create(ctx *gin.Context) {
	event := entities.NewEvent()

	if err := ctx.BindJSON(&event); err != nil {
		return
	}

	e.events = append(e.events, *event)

	formattedResults, err := s3.ListS3(event.Bucket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar objetos no bucket"})
		return
	}

	response := gin.H{
		"event":           event,
		"formattedResult": formattedResults,
	}
	ctx.JSON(http.StatusOK, response)
}
