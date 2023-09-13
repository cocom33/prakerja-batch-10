package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// user model
type User struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

// slice of User - users data
var users []User

// -------------------- controller --------------------

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	// if length of slice users less then 1 id = 1, else id +1 from last id in data slice users
	if len(users) < 1 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	// append new data to slice users
	users = append(users, user)

	// render JSON response with success message and the new data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// get all users
func GetUsersController(c echo.Context) error {
	// render JSON response with success message and users data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "berhasil mendapatkan users",
		"user":     users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))

	// find data by id, then render data - JSON response
	for _, v := range users {
		if v.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success user found",
				"user":     v,
			})
		}
	}

	// when the data does not exist render JSON response with not found messaga, and nil data
	return c.JSON(echo.ErrBadRequest.Code, map[string]interface{}{
		"messages": "User Not Found",
		"user":     nil,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// catch path param from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "tidak user id",
		})
	}

	// binding data
	data := new(User)
  if err := c.Bind(data); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "messages": "input data tidak sesuai",
    })
  }

	// find data and change the value, then render JSON response
	for _, v := range users {
		if v.Id == id {
			v.Name = data.Name
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success updated user",
				"user":     v,
			})
		}
	}

	// when the data does not exist render JSON response with not found message
  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "messages": "user tidak ditemukan",
  })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// catch path param from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user id tidak ditemukan",
		})
	}

	// looking for data in index of slice users, then delete element in slice users and render JSON response
	for _, v := range users {
		if v.Id == id {
			users = append(users[:id], users[id+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "delete success",
			})
		}
	}

	// when the data does not exist render JSON response with not found message
	return c.JSON(http.StatusNotFound, map[string]interface{}{
    "messages": "user tidak ditemukan",
  })
}

// ---------------------------------------------------
func main() {
	// instance of Echo
	e := echo.New()

	// routing
	e.GET("/users", GetUsersController)          //get all users
	e.GET("/users/:id", GetUserController)       //get user by id
	e.POST("/users", CreateUserController)       //Create new user
	e.PUT("/users/:id", UpdateUserController)    //update user
	e.DELETE("/users/:id", DeleteUserController) //delete user

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
