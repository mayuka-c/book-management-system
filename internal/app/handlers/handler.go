package handlers

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/mayuka-c/book-management-system/internal/pkg/config"
	"github.com/mayuka-c/book-management-system/internal/pkg/db"
	"github.com/mayuka-c/book-management-system/internal/pkg/models"
)

type Handler struct {
	db *db.DBClient
}

func NewHandler(dbConfig config.DBConfig) *Handler {
	return &Handler{
		db: db.Connect(dbConfig),
	}
}

func (h *Handler) CreateBook(ctx context.Context, input models.Book) error {

	dbObj := db.Book{}
	copier.Copy(&dbObj, &input)

	err := h.db.InsertBook(ctx, dbObj)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) GetAllBooks(ctx context.Context) ([]models.Book, error) {

	books := []models.Book{}

	dbBooks, err := h.db.GetAllBooks(ctx)
	if err != nil {
		return books, err
	}

	copier.Copy(&books, &dbBooks)

	return books, nil
}

func (h *Handler) GetBookByID(ctx context.Context, id int) (models.Book, error) {

	book := models.Book{}

	uint_id := uint(id)

	dbBook, err := h.db.GetBookByID(ctx, uint_id)
	if err != nil {
		return book, err
	}

	copier.Copy(&book, &dbBook)

	return book, nil
}

func (h *Handler) UpdateBook(ctx context.Context, id int, input models.UpdateBook) error {

	book, err := h.GetBookByID(ctx, id)
	if err != nil {
		return err
	}

	dbObj := db.UpdateBook{}
	copier.Copy(&dbObj, &book)

	if input.Name != "" {
		dbObj.Name = input.Name
	}
	if input.Author != "" {
		dbObj.Author = input.Author
	}
	if input.Publication != "" {
		dbObj.Publication = input.Publication
	}

	dbObj.ID = uint(id)

	err = h.db.UpdateBook(ctx, dbObj)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteBook(ctx context.Context, id int) error {

	uint_id := uint(id)

	err := h.db.DeleteBook(ctx, uint_id)
	if err != nil {
		return err
	}

	return nil
}
