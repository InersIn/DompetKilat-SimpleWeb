package reksadana

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsReksadana "DompetKilat-SimpleWeb/models/Finance"
	reksadanaRepo "DompetKilat-SimpleWeb/repository/Service/Reksadana"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllReksadana(c echo.Context) error {
	var responseData models.ResponseData
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	reksadanaRepo := reksadanaRepo.ReksadanaRepository(db)
	var response models.Response

	_, id := helper.GetToken(c)

	var reksadana []modelsReksadana.Reksadana
	reksadana, err := reksadanaRepo.GetAllReksadana(ctx, id)
	if err != nil {
		fmt.Println(err)
	}

	if len(reksadana) == 0 {
		response.Status = "success"
		response.Message = "Data is empty"
		return c.JSON(http.StatusOK, response)
	}

	responseData.Status = "success"
	responseData.Data = reksadana

	return c.JSON(http.StatusOK, responseData)
}
