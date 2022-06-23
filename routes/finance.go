package routes

import (
	finance "DompetKilat-SimpleWeb/controller/Service/Finance"
	invoice "DompetKilat-SimpleWeb/controller/Service/Invoice"

	"github.com/labstack/echo/v4"
)

func Service(e *echo.Group) {
	service := e.Group("/service")
	service.Add("POST", "/createFinances", finance.CreateFinance)
	service.Add("GET", "/createFinances", finance.GetAllFinance)

	service.Add("POST", "/createConventionalOsf", invoice.CreateConventionalOsf)
	service.Add("GET", "/createConventionalOsf", invoice.GetAllInvoice)
}
