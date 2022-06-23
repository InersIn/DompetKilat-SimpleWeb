package routes

import (
	finance "DompetKilat-SimpleWeb/controller/Service/Finance"
	invoiceCI "DompetKilat-SimpleWeb/controller/Service/Invoice/ConventionalInv"
	invoiceCO "DompetKilat-SimpleWeb/controller/Service/Invoice/ConventionalOsf"
	invoicePI "DompetKilat-SimpleWeb/controller/Service/Invoice/ProductiveInv"
	reksadana "DompetKilat-SimpleWeb/controller/Service/Reksadana"

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

	service.Add("POST", "/createProductiveInvoice", invoicePI.CreateProductiveInv)
	service.Add("GET", "/createProductiveInvoice", invoicePI.GetAllProductiveInv)

	service.Add("POST", "/createReksadana", reksadana.CreateReksadana)
	service.Add("GET", "/createReksadana", reksadana.GetAllReksadana)
}
