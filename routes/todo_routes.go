package routes

import (
	"my-todo-api/controllers"
	"my-todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todos := r.Group("/todos")
	todos.Use(middleware.AuthMiddleware())
	{
		todos.GET("/", controllers.GetTodos)
		todos.POST("/", controllers.CreateTodo)
		todos.GET("/:id", controllers.GetTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}
}
