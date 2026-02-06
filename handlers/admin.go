// echo-demo/handlers/admin.go
package handlers

import "github.com/labstack/echo/v4"

func AdminDashboard(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"message": "welcome admin 🚀",
	})
}
