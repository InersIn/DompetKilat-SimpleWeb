package finance

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsFinance "DompetKilat-SimpleWeb/models/Finance"
	financeRepo "DompetKilat-SimpleWeb/repository/Service/Finance"
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateFinance(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	financeRepo := financeRepo.FinanceRepository(db)

	var finances modelsFinance.Finance
	var response models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&finances)
	if err != nil {
		panic(err)
	}

	_, id := helper.GetToken(c)

	err = financeRepo.CreateFinance(ctx, finances, id)
	if err != nil {
		response.Status = "error"
		response.Message = "Failed create finance"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = "success"
	response.Message = "Successfully create finance"

	return c.JSON(http.StatusOK, response)

}
