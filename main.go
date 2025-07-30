package main

import (
	"my-todo-api/database"
	"my-todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	routes.RegisterTodoRoutes(router)
	router.Run(":8080")
}
