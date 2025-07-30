package routes

import (
	"my-todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todos := router.Group("/todos")
	{
		todos.GET("/", controllers.GetTodos)
		todos.POST("/", controllers.CreateTodo)
		todos.GET("/:id", controllers.GetTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}
}
