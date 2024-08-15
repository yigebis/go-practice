package controllers

import (
	"net/http"
	"task_management_api_with_clean_architecture/domain"
	"task_management_api_with_clean_architecture/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase *usecase.TaskUsecase
}

func NewTaskController(uuc *usecase.TaskUsecase) *TaskController {
	return &TaskController{TaskUsecase: uuc}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, code, err := tc.TaskUsecase.GetAllTasks()

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, code, err := tc.TaskUsecase.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task
	err := c.BindJSON(&newTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Payload"})
		return
	}

	code, err := tc.TaskUsecase.CreateTask(newTask)

	if err != nil {
		c.IndentedJSON(code, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask domain.Task

	err := c.BindJSON(&updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	code, err := tc.TaskUsecase.UpdateTask(id, updatedTask)

	if err != nil {
		c.IndentedJSON(code, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	code, err := tc.TaskUsecase.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(code, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
