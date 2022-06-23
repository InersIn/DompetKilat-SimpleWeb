package invoice

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsInvoice "DompetKilat-SimpleWeb/models/Finance"
	invoiceRepo "DompetKilat-SimpleWeb/repository/Service/Invoice"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllProductiveInv(c echo.Context) error {
	var responseData models.ResponseData
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	invoiceRepo := invoiceRepo.InvoiceRepository(db)
	var response models.Response

	_, id := helper.GetToken(c)

	var invoices []modelsInvoice.Invoce
	invoices, err := invoiceRepo.GetAllInvoice(ctx, id, Types)
	if err != nil {
		fmt.Println(err)
	}

	if len(invoices) == 0 {
		response.Status = "success"
		response.Message = "Data is empty"
		return c.JSON(http.StatusOK, response)
	}

	responseData.Status = "success"
	responseData.Data = invoices

	return c.JSON(http.StatusOK, responseData)
}
