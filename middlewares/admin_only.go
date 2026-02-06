// echo-demo/middlewares/admin_only.go
package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"echo-demo/utils"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// JWT token from context (set by echo-jwt)
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*utils.JwtCustomClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		if !claims.IsAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "admin access required",
			})
		}

		return next(c)
	}
}
