package db

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Publication string
}

type UpdateBook struct {
	ID          uint
	Name        string
	Author      string
	Publication string
}
