package routes

import (
	account "DompetKilat-SimpleWeb/controller/Account"
	customeMiddleware "DompetKilat-SimpleWeb/middleware"
	models "DompetKilat-SimpleWeb/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func private(c echo.Context) error {
	var response models.Response
	response.Status = "Success"
	response.Message = "Welcome"
	return c.JSON(http.StatusOK, response)
}

func Api(e *echo.Echo) {
	api := e.Group("/api")
	api.Add("POST", "/login", account.LoginAccount)
	api.Add("POST", "/register", account.RegisterAccount)

	api.Use(customeMiddleware.ValidateToken)
	// Test
	Service(api)
}
