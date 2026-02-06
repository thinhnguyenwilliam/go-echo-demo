package main

import (
	"echo-demo/handlers"
	"echo-demo/utils"

	"github.com/golang-jwt/jwt/v5"
	jwtmiddleware "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // 👈 VERY important

	// Public route
	e.POST("/login", handlers.Login)

	// JWT protected routes
	api := e.Group("/api")
	api.Use(jwtmiddleware.WithConfig(jwtmiddleware.Config{
		SigningKey: utils.JwtSecret,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.JwtCustomClaims)
		},
	}))

	api.GET("/profile", handlers.Profile)

	e.Logger.Fatal(e.Start(":8080"))
}
