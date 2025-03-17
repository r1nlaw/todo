package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/handlers"
	"todo/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func NewTaskRepository(db *sqlx.DB) *handlers.TaskRepository {
	return &handlers.TaskRepository{Db: db}
}

func Test_AddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании мока БД: %v", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "postgres")

	mock.ExpectQuery(`INSERT INTO tasks`).
		WithArgs("Test Task", "Test content", "Выполняется").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := NewTaskRepository(sqlxDB)

	task := models.Todo{Name: "Test Task", Content: "Test content", Status: "Выполняется"}
	body, _ := json.Marshal(task)
	request := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	repo.AddTask(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("ожидается статус 200, но получен %d", rr.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("ошибка при разборе JSON-ответа: %v", err)
	}

	if response["id"] != float64(1) {
		t.Errorf("ожидается ID = 1, но получен %v", response["id"])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("ожидание моков не были выполнены: %v", err)
	}

}
