package routes

import (
	handler "github.com/SJ22032003/go-ems/handler"
	middleware "github.com/SJ22032003/go-ems/middlewares"
	gin "github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine, v *gin.RouterGroup) {

	eventHandler := handler.EventHandler{}

	authorized := v.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/events", eventHandler.GetEvents)
		authorized.GET("/event/:id", eventHandler.GetEventById)
		authorized.PUT("/event/:id", eventHandler.UpdateEventById)
		authorized.POST("/events", eventHandler.CreateEvents)
		authorized.DELETE("/event/:id", eventHandler.DeleteEventById)
	}

}
