// echo-demo/middlewares/basic_auth.go
package middlewares

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Password string
	IsAdmin  bool
}

var users = map[string]User{
	"admin": {
		Password: "123",
		IsAdmin:  true,
	},
	"ThinhNguyen": {
		Password: "123",
		IsAdmin:  false,
	},
}

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	user, exists := users[username]
	if !exists {
		return false, nil
	}

	if user.Password != password {
		return false, nil
	}

	c.Set("username", username)
	c.Set("admin", user.IsAdmin)

	return true, nil
}
