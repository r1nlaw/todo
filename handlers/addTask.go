package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo/models"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (c *TaskRepository) AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	var inputData models.Todo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cant read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &inputData)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (name, content, status) VALUES ($1, $2, $3) RETURNING id`

	var newID int
	err = c.db.QueryRow(query, inputData.Name, inputData.Content, inputData.Status).Scan(&newID)
	if err != nil {
		http.Error(w, "failed to insert task", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Задача успешно добавлена",
		"id":      newID,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

}
