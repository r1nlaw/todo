package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo/models"
)

func (c *TaskRepository) RemoveTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cant read request", http.StatusBadRequest)
		return
	}
	var inputData models.Todo

	err = json.Unmarshal(body, &inputData)
	if err != nil {
		http.Error(w, "cant unmarshal request", http.StatusInternalServerError)
		return
	}

	query := `DELETE FROM tasks WHERE id = $1`
	result, err := c.db.Exec(query, inputData.Id)
	if err != nil {
		http.Error(w, "cant deleting task", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "error getting affected rows", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("task deleted successfully"))

}
