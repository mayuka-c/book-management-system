package models

type Book struct {
	Name        string `json:"name" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Publication string `json:"publication" validate:"required"`
}

type UpdateBook struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
