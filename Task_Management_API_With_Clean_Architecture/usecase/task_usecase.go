package usecase

import (
	"errors"

	// "fmt"
	"net/http"

	// "log"
	"task_management_api_with_clean_architecture/domain"
	// "time"
	// "github.com/google/uuid"
)

type ITaskRepository interface {
	FetchAllTasks() ([]domain.Task, error)
	FetchTaskById(id string) (domain.Task, error)
	InsertTask(task domain.Task) error
	UpdateTask(id string, updatedTask domain.Task) error
	DeleteTask(id string) error
}

type TaskUsecase struct {
	TaskRepo ITaskRepository
}

func NewTaskUseCase(ur ITaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: ur}
}

var (
	ErrNoTask         = errors.New("task not found")
	ErrTaskExists     = errors.New("task already exists")
	ErrInternalServer = errors.New("internal server error")
)

// func getCollection(client *mongo.Client) *mongo.Collection {
// 	collection := client.Database("test").Collection("tasks")

// 	return collection
// }

// var client = connectToDB()
// var collection = getCollection(client)

func (tu *TaskUsecase) GetAllTasks() ([]domain.Task, int, error) {
	var Tasks []domain.Task

	Tasks, err := tu.TaskRepo.FetchAllTasks()

	if err != nil {
		return nil, http.StatusInternalServerError, ErrInternalServer
	}

	return Tasks, http.StatusOK, nil
}

func (tu *TaskUsecase) GetTaskById(id string) (domain.Task, int, error) {
	var task domain.Task

	task, err := tu.TaskRepo.FetchTaskById(id)

	if err != nil {
		return domain.Task{}, http.StatusNotFound, ErrNoTask
	}

	return task, http.StatusOK, nil

}

func (tu *TaskUsecase) CreateTask(newTask domain.Task) (int, error) {
	newId := newTask.Id

	// check whether id exists or not
	_, err := tu.TaskRepo.FetchTaskById(newId)

	if err == nil {
		return http.StatusConflict, ErrTaskExists
	}

	err = tu.TaskRepo.InsertTask(newTask)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (tu *TaskUsecase) UpdateTask(id string, updatedTask domain.Task) (int, error) {

	//check whether the task exists or not
	_, err := tu.TaskRepo.FetchTaskById(id)

	if err != nil {
		return http.StatusNotFound, ErrNoTask
	}

	err = tu.TaskRepo.UpdateTask(id, updatedTask)

	if err != nil {
		return http.StatusInternalServerError, ErrInternalServer
	}

	return http.StatusOK, nil
}

func (tu *TaskUsecase) DeleteTask(id string) (int, error) {
	//check whether the task exists or not
	_, err := tu.TaskRepo.FetchTaskById(id)

	if err != nil {
		return http.StatusNotFound, ErrNoTask
	}

	err = tu.TaskRepo.DeleteTask(id)
	if err != nil {
		return http.StatusInternalServerError, ErrInternalServer
	}

	return http.StatusInternalServerError, nil
}
