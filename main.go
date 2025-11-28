package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	// Read raw body
	body, _ := io.ReadAll(c.Request().Body)
	log.Println("Raw body:", string(body))

	// Restore to allow Bind()
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	req := new(LoginRequest)

	// Bind JSON → struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	// Log request data
	log.Printf("LoginRequest is: %#v\n", req)

	// Dummy login logic
	if req.Username == "admin1" && req.Password == "123" {
		return c.JSON(http.StatusOK, LoginResponse{
			Token: "fake-jwt-token",
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"error": "invalid credentials",
	})
}

func main() {
	e := echo.New()

	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":8080"))
}
