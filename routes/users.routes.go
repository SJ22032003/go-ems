package routes

import (
	cnt "github.com/SJ22032003/go-ems/controllers"
	gin "github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine, v *gin.RouterGroup) {
	v.POST("/sign-up", cnt.CreateUser)
	v.POST("/login", cnt.LoginUser)
}
