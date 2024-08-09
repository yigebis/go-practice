package data

import (
	"task_manager/models"
	"time"
)

var TasksMap map[string]models.Task

var Tasks []models.Task

func init() {
	TasksMap = make(map[string]models.Task)

	Tasks = []models.Task{
		{
			Id:          "1",
			Title:       "Complete API Documentation",
			Description: "Write comprehensive documentation for the Task Management API.",
			DueDate:     time.Date(2024, time.August, 15, 23, 59, 0, 0, time.UTC),
			Status:      "Pending",
		},
		{
			Id:          "2",
			Title:       "Implement User Authentication",
			Description: "Develop and integrate user authentication for the API.",
			DueDate:     time.Date(2024, time.August, 10, 17, 0, 0, 0, time.UTC),
			Status:      "In Progress",
		},
		{
			Id:          "3",
			Title:       "Refactor Codebase",
			Description: "Refactor the codebase for better maintainability and performance.",
			DueDate:     time.Date(2024, time.August, 20, 12, 0, 0, 0, time.UTC),
			Status:      "Completed",
		},
		{
			Id:          "4",
			Title:       "Create Unit Tests",
			Description: "Write unit tests for the core functionalities of the API.",
			DueDate:     time.Date(2024, time.August, 12, 9, 0, 0, 0, time.UTC),
			Status:      "Pending",
		},
		{
			Id:          "5",
			Title:       "Deploy API to Production",
			Description: "Deploy the Task Management API to the production environment.",
			DueDate:     time.Date(2024, time.August, 18, 14, 30, 0, 0, time.UTC),
			Status:      "Pending",
		},
	}

	for _, task := range Tasks {
		TasksMap[task.Id] = task
	}

}

func GetAllTasks() []models.Task {
	return Tasks
}

func GetTaskById(id string) (bool, models.Task) {
	task, exists := TasksMap[id]

	if exists {
		return true, task
	}

	return false, task
}

func CreateTask(newTask models.Task) bool {
	newId := newTask.Id

	_, exists := TasksMap[newId]

	if exists {
		return false
	}

	TasksMap[newId] = newTask
	Tasks = append(Tasks, newTask)
	return true
}

func UpdateTask(id string, updatedTask models.Task) bool {
	_, exists := TasksMap[id]

	if !exists {
		return false
	}

	TasksMap[id] = updatedTask

	Tasks = Tasks[:0]
	for _, task := range TasksMap {
		Tasks = append(Tasks, task)
	}

	return true
}

func DeleteTask(id string) bool {
	_, exists := TasksMap[id]

	if !exists {
		return false
	}

	delete(TasksMap, id)

	Tasks = Tasks[:0]
	for _, task := range TasksMap {
		Tasks = append(Tasks, task)
	}
	return true
}
