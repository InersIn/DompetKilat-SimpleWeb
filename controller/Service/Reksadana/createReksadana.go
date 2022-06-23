package reksadana

import (
	"DompetKilat-SimpleWeb/database/mysql"
	"DompetKilat-SimpleWeb/helper"
	"DompetKilat-SimpleWeb/models"
	reksadana "DompetKilat-SimpleWeb/repository/Service/Reksadana"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	modelsReksadana "DompetKilat-SimpleWeb/models/Finance"

	"github.com/labstack/echo/v4"
)

func CreateReksadana(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	reksadanaRepo := reksadana.ReksadanaRepository(db)

	var reksadana modelsReksadana.Reksadana
	var response models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&reksadana)
	if err != nil {
		panic(err)
	}

	_, id := helper.GetToken(c)

	err = reksadanaRepo.CreateReksadana(ctx, reksadana, id)

	if err != nil {
		response.Status = "error"
		response.Message = "Failed create reksadana"
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = "success"
	response.Message = "Successfully create reksadana"

	return c.JSON(http.StatusOK, response)
}
