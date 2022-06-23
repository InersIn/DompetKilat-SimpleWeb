package account

import (
	"DompetKilat-SimpleWeb/database/mysql"
	models "DompetKilat-SimpleWeb/models"
	modelsAccount "DompetKilat-SimpleWeb/models/Account"
	accountRepo "DompetKilat-SimpleWeb/repository/Account"
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterAccount(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	registerRepo := accountRepo.AccountRepository(db)

	var register modelsAccount.Account
	var response models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&register)

	if register.Email == "" || register.Password == "" || register.Username == "" {
		response.Status = "error"
		response.Message = "All field is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	register, err = registerRepo.Register(ctx, register)
	if err != nil {
		response.Status = "error"
		response.Message = "Email or Username already registered"
		return c.JSON(http.StatusConflict, response)
	}

	response.Status = "success"
	response.Message = "Account Success Register"

	return c.JSON(http.StatusOK, response)
}
