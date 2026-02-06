// echo-demo/handlers/profile.go
package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"echo-demo/utils"
)

func Profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*utils.JwtCustomClaims)

	return c.JSON(200, map[string]any{
		"username": claims.Username,
		"isAdmin":  claims.IsAdmin,
	})
}
