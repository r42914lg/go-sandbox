package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/set_current", sendHandler)

	httpPort := "8091"
	e.Logger.Fatal(e.Start(":" + httpPort))
}

type RequestBody struct {
	Value int `json:"value"`
}

func sendHandler(c echo.Context) error {
	var requestBody RequestBody

	err := c.Bind(&requestBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Received int val in SERVICE #1 = %d", requestBody.Value))
}
