package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	domain "r42914lg.com/domain"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/endpoint_2", sendHandler)

	httpPort := "8092"
	e.Logger.Fatal(e.Start(":" + httpPort))
}

func sendHandler(c echo.Context) error {
	var requestBody domain.RequestBody

	err := c.Bind(&requestBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	resp := &domain.ServiceTwoResponse{
		StrValue: fmt.Sprintf("in SERVICE_2 val incremented is %d !!!", requestBody.Value+1),
	}

	return c.JSON(http.StatusOK, resp)
}
