package routes

import (
	handler "github.com/SJ22032003/go-ems/handler"
	gin "github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine, v *gin.RouterGroup) {

	userHandler := handler.UserHandler{}

	v.POST("/sign-up", userHandler.CreateUser)
	v.POST("/login", userHandler.LoginUser)
	v.GET("/login-page", userHandler.RenderLoginPage)
}
