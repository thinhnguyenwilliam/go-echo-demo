// echo-demo/handlers/auth.go
package handlers

import (
	"net/http"

	"echo-demo/models"
	"echo-demo/utils"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	req := new(models.LoginRequest)

	if err := c.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	isAdmin := false

	token, err := utils.GenerateToken(req.Username, isAdmin)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
	})
}
