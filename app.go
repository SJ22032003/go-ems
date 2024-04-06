package main

import (
	db "github.com/SJ22032003/go-ems/db"
	routes "github.com/SJ22032003/go-ems/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	v1 := server.Group("/v1/api")

	routes.EventRoutes(server, v1)
	routes.UserRoutes(server, v1)

	server.Run(":8080")
}
