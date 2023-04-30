package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"

	"github.com/mayuka-c/book-management-system/internal/app/routes"
	"github.com/mayuka-c/book-management-system/internal/pkg/config"
	"github.com/mayuka-c/book-management-system/pkg/log"
)

var (
	ctx = context.Background()
)

var serviceConfig config.ServiceConfig
var dbConfig config.DBConfig

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func init() {
	serviceConfig = config.GetServiceConfig(ctx)
	dbConfig = config.GetDBConfig(ctx)
}

func startAPIServer(e *echo.Echo) {
	log.Infof(ctx, "Starting service on port: %v", serviceConfig.APIPort)
	err := e.Start(":" + strconv.Itoa(serviceConfig.APIPort))
	if err != nil {
		log.Errorf(ctx, "Failed running the router. Please restart the service and try again.")
	}
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	routes.RegisterBookStoreRoutes(e, dbConfig)
	startAPIServer(e)
}
