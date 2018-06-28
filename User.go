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

var UserObject = User{
	Username: "ALEX",
	Email: "testuserapi@alemkin.com",
	Name: "Alexander Lemkin",
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

/*
 * Update User
 */
 func UpdateUser(c echo.Context) error {
	u := json.NewDecoder(c.Request().Body)
	var user = &User{}
	if err := u.Decode(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

  //Grab user from mongodb
  dbUser := User{
    Username: user.Username,
    Email: user.Email,
    Name: user.Name,
  }

	return c.JSON(http.StatusOK, dbUser)
}

/*
 * Get User
 */
func GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
    return c.JSON(http.StatusBadRequest, "id missing")
  }
	//TODO pull from mongoDB
	return c.JSON(http.StatusOK, UserObject)
}

/*
 * Delete User
 */
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
    return c.JSON(http.StatusBadRequest, "id missing")
  }
	//TODO delete from mongoDB
	return c.JSON(http.StatusOK, "")
}