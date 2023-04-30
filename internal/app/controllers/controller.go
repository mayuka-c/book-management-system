package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/mayuka-c/book-management-system/internal/app/handlers"
	"github.com/mayuka-c/book-management-system/internal/pkg/config"
	"github.com/mayuka-c/book-management-system/internal/pkg/models"
	"github.com/mayuka-c/book-management-system/pkg/log"
)

type Controller struct {
	handler *handlers.Handler
}

func NewController(dbConfig config.DBConfig) *Controller {
	return &Controller{
		handler: handlers.NewHandler(dbConfig),
	}
}

func (con *Controller) CreateBook(c echo.Context) (err error) {

	ctx := c.Request().Context()

	reqBody := new(models.Book)
	if err = c.Bind(reqBody); err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(reqBody); err != nil {
		log.Errorf(ctx, err.Error())
		return err
	}

	err = con.handler.CreateBook(ctx, *reqBody)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "msg: Server is unable to process request")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"msg": "Book has been successfully created"})
}

func (con *Controller) GetAllBooks(c echo.Context) (err error) {

	ctx := c.Request().Context()

	books, err := con.handler.GetAllBooks(ctx)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "msg: Server is unable to process request")
	}

	return c.JSON(http.StatusOK, books)
}

func (con *Controller) GetBookByID(c echo.Context) (err error) {

	ctx := c.Request().Context()

	id := c.Param("id")

	int_id, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "msg: ID provided is not numeric only")
	}

	book, err := con.handler.GetBookByID(ctx, int_id)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "msg: Server is unable to process request")
	}

	return c.JSON(http.StatusFound, book)
}

func (con *Controller) UpdateBook(c echo.Context) (err error) {

	ctx := c.Request().Context()

	id := c.Param("id")

	int_id, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "msg: ID provided is not numeric only")
	}

	reqBody := new(models.UpdateBook)
	if err = c.Bind(reqBody); err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = con.handler.UpdateBook(ctx, int_id, *reqBody)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "msg: Server is unable to process request")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"msg": "Book has been successfully updated"})
}

func (con *Controller) DeleteBook(c echo.Context) (err error) {

	ctx := c.Request().Context()

	id := c.Param("id")

	int_id, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "msg: ID provided is not numeric only")
	}

	err = con.handler.DeleteBook(ctx, int_id)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "msg: Server is unable to process request")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"msg": "Book has been successfully deleted"})
}
