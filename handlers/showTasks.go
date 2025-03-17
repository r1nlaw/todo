package handlers

import (
	"encoding/json"
	"net/http"
	"todo/models"
)

func (c *TaskRepository) ShowTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	var tasks []models.Todo

	query := `SELECT id, name, content, status FROM tasks`
	err := c.Db.Select(&tasks, query)
	if err != nil {
		http.Error(w, "cant get tasks from database", http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		http.Error(w, "nothing tasks in database", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

}
