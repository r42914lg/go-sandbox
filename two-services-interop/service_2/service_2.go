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

	e.POST("/endpoint_2", sendHandler)

	httpPort := "8092"
	e.Logger.Fatal(e.Start(":" + httpPort))
}

type RequestBody struct {
	Value int `json:"value"`
}

type ServiceTwoResponse struct {
	StrValue string `json:"str_val"`
}

func sendHandler(c echo.Context) error {
	var requestBody RequestBody

	err := c.Bind(&requestBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	resp := &ServiceTwoResponse{
		StrValue: fmt.Sprintf("in SERVICE_2 val incremented is %d !!!", requestBody.Value+1),
	}

	return c.JSON(http.StatusOK, resp)
}
