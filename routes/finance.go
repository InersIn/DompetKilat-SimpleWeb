package routes

import (
	finance "DompetKilat-SimpleWeb/controller/Service/Finance"
	invoiceCO "DompetKilat-SimpleWeb/controller/Service/Invoice/ConventionalOsf"
	invoiceCI "DompetKilat-SimpleWeb/controller/Service/Invoice/ConventionalInv"

	"github.com/labstack/echo/v4"
)

func Service(e *echo.Group) {
	service := e.Group("/service")
	service.Add("POST", "/createFinances", finance.CreateFinance)
	service.Add("GET", "/createFinances", finance.GetAllFinance)

	service.Add("POST", "/createConventionalOsf", invoiceCO.CreateConventionalOsf)
	service.Add("GET", "/createConventionalOsf", invoiceCO.GetAllInvoiceOsf)

	service.Add("POST", "/createConventionalInvoice", invoiceCI.CreateConventionalInvoice)
	service.Add("GET", "/createConventionalInvoice", invoiceCI.GetAllInvoiceInv)
}
