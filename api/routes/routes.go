package routes

import (
	controllers "infra-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	eventController := controllers.NewEventController()
	v1 := router.Group("/v1")
	{
		v1.GET("/events", eventController.FindAll)
		v1.POST("/event", eventController.Create)
	}

	return v1
}
