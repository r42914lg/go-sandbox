package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/endpoint_1", endPointOneHandler)

	httpPort := "8091"
	e.Logger.Fatal(e.Start(":" + httpPort))
}

type ServiceOneRequestBody struct {
	IntValue int `json:"value"`
}

type ServiceTwoResponse struct {
	StrValue string `json:"str_val"`
}

func endPointOneHandler(c echo.Context) error {
	var requestBody ServiceOneRequestBody

	err := c.Bind(&requestBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	jsonStr := []byte(fmt.Sprintf(`{"value":%d}`, requestBody.IntValue))
	r := bytes.NewReader(jsonStr)

	resp, err := client.Post("http://host.docker.internal:8081/endpoint_2", "application/json", r)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var serviceTwoResponse ServiceTwoResponse
	if err = json.Unmarshal(body, &serviceTwoResponse); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Received INT val in SERVICE #1 = %d / Returning STR val from SERVICE #2 = %s", requestBody.IntValue, serviceTwoResponse.StrValue))
}
