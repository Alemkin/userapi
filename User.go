package main

import (
	"net/http"
	"encoding/json"
	"github.com/labstack/echo"
)
/*
 * User Defines the user object
 */
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string
}

/*
 * Create User
 */
func CreateUser(c echo.Context) error {
	u := json.NewDecoder(c.Request().Body)
	var user = &User{}
	if err := u.Decode(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//Create user in mongodb

	return c.JSON(http.StatusOK, user)
}
