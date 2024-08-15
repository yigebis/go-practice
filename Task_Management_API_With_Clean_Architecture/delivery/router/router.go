package router

import (
	"task_management_api_with_clean_architecture/delivery/controllers"
	"task_management_api_with_clean_architecture/infrastructure"
	"task_management_api_with_clean_architecture/usecase"

	"github.com/gin-gonic/gin"
)

func Setup(uuc *usecase.UserUsecase, tuc *usecase.TaskUsecase, secretKey string) {
	router := gin.Default()

	task_controller := controllers.NewTaskController(tuc)
	user_controller := controllers.NewUserController(uuc)

	router.GET("/tasks", infrastructure.AuthMiddleware(secretKey), task_controller.GetAllTasks)
	router.GET("/tasks/:id", infrastructure.AuthMiddleware(secretKey), task_controller.GetTask)
	router.PUT("/tasks/:id", infrastructure.AuthMiddleware(secretKey), infrastructure.AdminMiddleware(), task_controller.UpdateTask)
	router.POST("/tasks", infrastructure.AuthMiddleware(secretKey), infrastructure.AdminMiddleware(), task_controller.CreateTask)
	router.DELETE("/tasks/:id", infrastructure.AuthMiddleware(secretKey), infrastructure.AdminMiddleware(), task_controller.DeleteTask)

	router.POST("/register", user_controller.Register)
	router.POST("/login", user_controller.Login)
	router.PUT("/promote/:username", infrastructure.AuthMiddleware(secretKey), infrastructure.AdminMiddleware(), user_controller.Promote)

	router.Run()
}
