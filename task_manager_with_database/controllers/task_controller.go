package controllers

import (
	"net/http"
	"task_manager_with_database/data"
	"task_manager_with_database/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
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

	err = data.CreateTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
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

	err = data.UpdateTask(id, updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := data.DeleteTask(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
