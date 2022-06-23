package invoice

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsInvoice "DompetKilat-SimpleWeb/models/Finance"
	invoiceRepo "DompetKilat-SimpleWeb/repository/Service/Invoice"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var Types string = "CI"

func CreateConventionalInvoice(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	invoiceRepo := invoiceRepo.InvoiceRepository(db)

	var invoice modelsInvoice.Invoce
	var response models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&invoice)
	if err != nil {
		panic(err)
	}

	_, id := helper.GetToken(c)
	invoice.Type = Types
	err = invoiceRepo.CreateInvoice(ctx, invoice, id)

	if err != nil {
		response.Status = "error"
		response.Message = "Failed create invoice"
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = "success"
	response.Message = "Successfully create invoice"

	return c.JSON(http.StatusOK, response)
}
