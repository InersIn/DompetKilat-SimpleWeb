package middleware

import (
	auth "DompetKilat-SimpleWeb/controller/Auth"
	"DompetKilat-SimpleWeb/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	var response models.Response
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			response.Status = "Request Unauthorized"
			response.Message = "Token missing"
			return c.JSON(http.StatusUnauthorized, response)
		}
		err := auth.ValidateToken(token)
		if err != nil {
			response.Status = "Request Unauthorized"
			response.Message = err.Error()
			return c.JSON(http.StatusUnauthorized, response)
		}
		return next(c)
	}
}
