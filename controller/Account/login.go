package account

import (
	auth "DompetKilat-SimpleWeb/controller/Auth"
	"DompetKilat-SimpleWeb/database/mysql"
	modelsAccount "DompetKilat-SimpleWeb/models/Account"
	modelsAuth "DompetKilat-SimpleWeb/models/Auth"
	accountRepo "DompetKilat-SimpleWeb/repository/Account"
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(c echo.Context) error {
	db := mysql.GetConnection()
	defer db.Close()

	ctx := context.Background()
	loginRepo := accountRepo.AccountRepository(db)

	var login modelsAccount.Account
	var response modelsAccount.RegisterResponse

	err := json.NewDecoder(c.Request().Body).Decode(&login)

	if err != nil || login.Password == "" || login.Username == "" {
		response.Status = "error"
		response.Message = "All field is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	hashed, err := loginRepo.FindByUsername(ctx, login.Username)
	byteHashed := []byte(hashed)
	
	if bcrypt.CompareHashAndPassword(byteHashed, []byte(login.Password)) != nil {
		response.Status = "Login Failed"
		response.Message = "Username or password incorrect"
		return c.JSON(http.StatusUnauthorized, response)
	}

	token, err := auth.GenerateToken(login.Username)

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
