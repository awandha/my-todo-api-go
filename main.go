package main

import (
	"my-todo-api/database"
	"my-todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	routes.AuthRoutes(r)
	routes.TodoRoutes(r) // <- Make this protected next

	r.Run(":8080")
}
