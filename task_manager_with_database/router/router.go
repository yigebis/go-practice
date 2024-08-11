package router

import (
	"task_manager_with_database/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run()
}
