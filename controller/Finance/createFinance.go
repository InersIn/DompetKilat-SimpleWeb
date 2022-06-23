package finance

import (
	auth "DompetKilat-SimpleWeb/controller/Auth"
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/models"
	modelsFinance "DompetKilat-SimpleWeb/models/Finance"
	financeRepo "DompetKilat-SimpleWeb/repository/Finance"
	"context"
	"encoding/json"
	"net/http"
	"strings"

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

	token := c.Request().Header.Get("Authorization")
	token = strings.Split(token, "Bearer ")[1]

	_, id := auth.ExtractToken(token)

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
