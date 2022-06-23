package finance

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsFinance "DompetKilat-SimpleWeb/models/Finance"
	financeRepo "DompetKilat-SimpleWeb/repository/Service/Finance"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllFinance(c echo.Context) error {
	var responseData models.ResponseData
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	financeRepo := financeRepo.FinanceRepository(db)
	var response models.Response

	_, id := helper.GetToken(c)

	var finances []modelsFinance.Finance
	finances, err := financeRepo.GetAllFinance(ctx, id)
	if err != nil {
		fmt.Println(err)
	}

	if len(finances) == 0 {
		response.Status = "success"
		response.Message = "Data is empty"
		return c.JSON(http.StatusOK, response)
	}

	responseData.Status = "success"
	responseData.Data = finances

	return c.JSON(http.StatusOK, responseData)
}
