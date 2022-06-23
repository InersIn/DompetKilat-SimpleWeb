package helper

import (
	auth "DompetKilat-SimpleWeb/controller/Auth"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) (string, int) {
	token := c.Request().Header.Get("Authorization")
	token = strings.Split(token, "Bearer ")[1]

	username, id := auth.ExtractToken(token)
	return username, id
}
