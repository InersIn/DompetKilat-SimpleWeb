package routes

import (
	finance "DompetKilat-SimpleWeb/controller/Finance"

	"github.com/labstack/echo/v4"
)

func Service(e *echo.Group) {
	service := e.Group("/service")
	service.Add("POST", "/createFinances", finance.CreateFinance)
	service.Add("GET", "/createFinances", finance.GetAllFinance)
}