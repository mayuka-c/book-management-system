package routes

import (
	"github.com/labstack/echo"
	"github.com/mayuka-c/book-management-system/internal/app/controllers"
	"github.com/mayuka-c/book-management-system/internal/pkg/config"
)

func RegisterBookStoreRoutes(e *echo.Echo, dbConfig config.DBConfig) {
	controller := controllers.NewController(dbConfig)
	e.POST("/book", controller.CreateBook)
	e.GET("/books", controller.GetAllBooks)
	e.GET("/book/:id", controller.GetBookByID)
	e.PUT("/book/:id", controller.UpdateBook)
	e.DELETE("/book/:id", controller.DeleteBook)
}
