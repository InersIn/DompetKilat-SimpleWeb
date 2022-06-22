package routes

import (
	account "DompetKilat-SimpleWeb/controller/Account"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/api/register", account.RegisterAccount)
	e.POST("/api/login", account.LoginAccount)

	return e
}
