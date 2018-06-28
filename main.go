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

	e.GET("/status", status)

	userRoutes(e)
	
	e.Logger.Fatal(e.Start(":8081"))
}

func userRoutes(e *echo.Echo) {
	e.POST("/user", CreateUser)
	e.PUT("/user/:id", UpdateUser)
	e.GET("/user/:id", GetUser)
	e.DELETE("/user/:id", DeleteUser)
}

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Status: "OKAY"})
}