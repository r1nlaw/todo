package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TodoList interface {
	AddTask(w http.ResponseWriter, r *http.Request)
	RemoveTask(w http.ResponseWriter, r *http.Request)
	ChangeTask(w http.ResponseWriter, r *http.Request)
	ShowTasks(w http.ResponseWriter, r *http.Request)
}

type TaskRepository struct {
	Db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{Db: db}
}
