package main

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/database"
	"todo/handlers"

	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Сервер запущен"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8082"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("ошибка инициализации БД:", err)
	}
	defer db.Close()
	taskRepo := handlers.NewTaskRepository(db)

	http.HandleFunc("/", taskRepo.ShowTasks)

	handlerWithCors := c.Handler(http.DefaultServeMux)

	http.HandleFunc("/API/add-task", taskRepo.AddTask)
	http.HandleFunc("/API/delete-task", taskRepo.RemoveTask)
	http.HandleFunc("/API/update-task", taskRepo.ChangeTask)

	http.ListenAndServe(":8081", handlerWithCors)

}
