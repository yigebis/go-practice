package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks := data.GetAllTasks()

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	found, task := data.GetTaskById(id)

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	err := c.BindJSON(&newTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Payload"})
		return
	}

	status := data.CreateTask(newTask)

	if !status {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Id already exists"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	err := c.BindJSON(&updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	updated := data.UpdateTask(id, updatedTask)

	if !updated {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	deleted := data.DeleteTask(id)

	if !deleted {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
