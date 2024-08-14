package router

import (
	"task_manager_with_auth/controllers"
	"task_manager_with_auth/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware(), controllers.GetAllTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware(), controllers.GetTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.UpdateTask)
	router.POST("/tasks", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.CreateTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteTask)

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.PUT("/promote/:username", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.Promote)

	router.Run()
}
