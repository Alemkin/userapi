package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
 * Response is used for status
 */
type Response struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{Status: "OKAY"})
	})
	e.GET("/user", CreateUser)
	
	e.Logger.Fatal(e.Start(":8081"))
}
