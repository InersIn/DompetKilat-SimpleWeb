package routes

import (
	account "DompetKilat-SimpleWeb/controller/Account"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/api/register", account.RegisterAccount)

	return e
}
