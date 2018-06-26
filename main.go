package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct { 
	Status string `json:"status"`
}

func main() {
	e := echo.New()
	e.GET("/status", func (c echo.Context) error {
		return c.JSON(http.StatusOK, Response{Status: "OKAY"})
	})

	e.Logger.Fatal(e.Start(":8081"))
}