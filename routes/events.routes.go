package routes

import (
	cnt "github.com/SJ22032003/go-ems/controllers"
	middleware "github.com/SJ22032003/go-ems/middlewares"
	gin "github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine, v *gin.RouterGroup) {

	authorized := v.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/events", cnt.GetEvents)
		authorized.GET("/event/:id", cnt.GetEventById)
		authorized.PUT("/event/:id", cnt.UpdateEventById)
		authorized.POST("/events", cnt.CreateEvents)
		authorized.DELETE("/event/:id", cnt.DeleteEventById)
	}

}
