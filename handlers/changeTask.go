package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo/models"
)

func (c *TaskRepository) ChangeTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cant read request", http.StatusInternalServerError)
		return
	}
	var inputData models.Todo

	err = json.Unmarshal(body, &inputData)
	if err != nil {
		http.Error(w, "error unmarshal request", http.StatusInternalServerError)
		return
	}

	query := `UPDATE tasks SET name = $1, content = $2, status = $3 WHERE id = $4`
	res, err := c.Db.Exec(query, inputData.Name, inputData.Content, inputData.Status, inputData.Id)
	if err != nil {
		http.Error(w, "error updating task", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "cant getting affceted rows", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "task not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("task updated successfully"))

}
