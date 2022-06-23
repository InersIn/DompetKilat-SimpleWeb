package account

import (
	auth "DompetKilat-SimpleWeb/controller/Auth"
	"DompetKilat-SimpleWeb/database/mysql"
	models "DompetKilat-SimpleWeb/models"
	modelsAccount "DompetKilat-SimpleWeb/models/Account"
	modelsAuth "DompetKilat-SimpleWeb/models/Auth"
	accountRepo "DompetKilat-SimpleWeb/repository/Account"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	loginRepo := accountRepo.AccountRepository(db)

	var body modelsAccount.Account
	var login modelsAccount.Account
	var response models.Response

	err := json.NewDecoder(c.Request().Body).Decode(&body)

	if err != nil || body.Password == "" || body.Username == "" {
		response.Status = "error"
		response.Message = "All field is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	login, err = loginRepo.FindByUsername(ctx, body)
	if err != nil {
		panic(err)
	}

	byteHashed := []byte(login.Password)

	if bcrypt.CompareHashAndPassword(byteHashed, []byte(body.Password)) != nil {
		fmt.Println(login.Password, []byte(body.Password), login)
		response.Status = "Login Failed"
		response.Message = "Username or password incorrect"
		return c.JSON(http.StatusUnauthorized, response)
	}

	token, err := auth.GenerateToken(login.Username, login.User_id)

	if err != nil {
		response.Status = "error"
		response.Message = "Failed generated token"
		return c.JSON(http.StatusInternalServerError, response)
	}

	var tokenResponse modelsAuth.TokenResponse
	tokenResponse.Token = token
	tokenResponse.Type = "Bearer"

	return c.JSON(http.StatusOK, tokenResponse)
}
