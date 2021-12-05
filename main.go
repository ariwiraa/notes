package main

import (
	"fmt"
	"net/http"
	"notes/app/config"
	"notes/app/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "bisa nih")
	})

	database.Init(e)

	port := config.GetConfig(e).Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
