package models

type Todo struct {
	Id      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Content string `json:"content" db:"content"`
	Status  string `json:"status" db:"status"`
}

type DeleteRequest struct {
	Id int64 `json:"id" db:"id"`
}
