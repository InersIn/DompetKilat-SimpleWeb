package finance

import (
	auth "DompetKilat-SimpleWeb/controller/Auth"
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/models"
	modelsFinance "DompetKilat-SimpleWeb/models/Finance"
	financeRepo "DompetKilat-SimpleWeb/repository/Finance"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetAllFinance(c echo.Context) error {
	var responseData models.ResponseData
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	financeRepo := financeRepo.FinanceRepository(db)
	var response models.Response

	token := c.Request().Header.Get("Authorization")
	token = strings.Split(token, "Bearer ")[1]

	_, id := auth.ExtractToken(token)

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
