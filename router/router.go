package router

import (
	"parser/api"
	"parser/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
    e := echo.New()

	middlewares.SetLoggerMiddlewares(e)
	
	api.MainGroup(e)

    return e
}