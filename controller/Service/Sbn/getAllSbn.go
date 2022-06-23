package sbn

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsSbn "DompetKilat-SimpleWeb/models/Finance"
	sbn "DompetKilat-SimpleWeb/repository/Service/Sbn"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetALlSbn(c echo.Context) error {
	var responseData models.ResponseData
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	sbnRepo := sbn.SbnRepository(db)

	var response models.Response

	_, id := helper.GetToken(c)

	var sbn []modelsSbn.Sbn
	sbn, err := sbnRepo.GetAllSbn(ctx, id)
	if err != nil {
		fmt.Println(err)
	}

	if len(sbn) == 0 {
		response.Status = "success"
		response.Message = "Data is empty"
		return c.JSON(http.StatusOK, response)
	}

	responseData.Status = "success"
	responseData.Data = sbn

	return c.JSON(http.StatusOK, responseData)
}
