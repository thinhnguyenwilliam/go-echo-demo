// echo-demo/handlers/auth.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func Stream(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/plain")
	c.Response().WriteHeader(200)

	for i := 1; i <= 5; i++ {
		_, err := c.Response().Write([]byte(
			fmt.Sprintf("chunk %d\n", i),
		))
		if err != nil {
			return err
		}

		c.Response().Flush() // 🔥 IMPORTANT
		time.Sleep(1 * time.Second)
	}

	return nil
}

func StreamJSON(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "application/json")
	c.Response().WriteHeader(200)

	encoder := json.NewEncoder(c.Response())

	for i := 1; i <= 5; i++ {
		err := encoder.Encode(map[string]any{
			"step": i,
			"time": time.Now(),
		})
		if err != nil {
			return err
		}

		c.Response().Flush()
		time.Sleep(time.Second)
	}

	return nil
}

func StreamSSE(c echo.Context) error {
	res := c.Response()
	req := c.Request()

	res.Header().Set("Content-Type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")

	for i := 1; i <= 5; i++ {
		fmt.Fprintf(res, "data: message %d\n\n", i)
		res.Flush()

		time.Sleep(time.Second)

		// stop if client disconnects
		select {
		case <-req.Context().Done():
			return nil
		default:
		}
	}

	return nil
}
