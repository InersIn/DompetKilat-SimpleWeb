package sbn

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	modelsSbn "DompetKilat-SimpleWeb/models/Finance"
	sbn "DompetKilat-SimpleWeb/repository/Service/Sbn"
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateSbn(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	sbnRepo := sbn.SbnRepository(db)

	var sbn modelsSbn.Sbn
	var response models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&sbn)
	if err != nil {
		panic(err)
	}

	_, id := helper.GetToken(c)

	err = sbnRepo.CreateSbn(ctx, sbn, id)

	if err != nil {
		response.Status = "error"
		response.Message = "Failed create SBN"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = "success"
	response.Message = "Successfully create SBN"

	return c.JSON(http.StatusOK, response)
}
