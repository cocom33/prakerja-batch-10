package day5

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	id int          `json:"id"`
	name string     `json:"name"`
	email string    `json:"email"`
}

type BaseResponse struct {
	Status bool      `json:"status"`
	Message string   `json:"message"`
	Data interface{} `json:"data"`
}

func day5() {
	e := echo.New()
	
	// routing
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUsersDetailController)
	e.Start(":8000")
}

func GetUsersController(c echo.Context) error {
	var users []User
	
	users = append(users, User{1, "A", "A"})
	users = append(users, User{2, "B", "B"})
	users = append(users, User{3, "C", "C"})
	
	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Success get data",
		Data: users,
	})
}

func GetUsersDetailController(c echo.Context) error {
	id := c.Param("id")
	var users User = User{1, id, "A"}
	
	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Success get data",
		Data: users,
	})
}