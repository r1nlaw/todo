package handlers

import "net/http"

type TodoList interface {
	AddTask(w http.ResponseWriter, r *http.Request)
	RemoveTask(w http.ResponseWriter, r *http.Request)
	ChangeTask(w http.ResponseWriter, r *http.Request)
	ShowTasks(w http.ResponseWriter, r *http.Request)
}
