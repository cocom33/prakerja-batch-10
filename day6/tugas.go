package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port string
	DB_Host string
	DB_Name string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Port: "3306",
		DB_Host: "localhost",
		DB_Name: "crud_go",
	}
	
	connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name string       `json:"name" form:"name"`
	Email string      `json:"email" form:"email"`
	Password string   `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users": users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var users User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	result := DB.First(&users, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,  result.Error.Error())
	}
	
	return c.JSON(http.StatusOK,  map[string]interface{}{
			"message": "berhasil mengambil data user",
			"user": users,
		})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user": user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	user := User{}	
	if result := DB.First(&user, id).Error; result != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error())
	}
	
	deleteUser := DB.Delete(&user, id)
	if deleteUser.Error != nil {
		return c.JSON(http.StatusInternalServerError, deleteUser.Error.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "berhasil menghapus data",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	inputUser := User{}
  if err := c.Bind(&inputUser); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "Data pengguna tidak valid",
    })
  }
  
  updateUser := User{}
  if err := DB.First(&updateUser, id).Error; err != nil {
    return c.JSON(http.StatusNotFound, map[string]interface{}{
      "message": "Pengguna tidak ditemukan",
    })
  }
  
  if inputUser.Name == "" || inputUser.Email == "" {
		return c.JSON(http.StatusBadRequest, "silahkan masukkan nama dan email")
	}

  updateUser.Name = inputUser.Name
  updateUser.Email = inputUser.Email
  updateUser.Password = inputUser.Password
  
  if err := DB.Save(&updateUser).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "message": "Gagal mengupdate pengguna",
    })
  }
	
	return c.JSON(http.StatusOK, map[string]interface{} {
		"message": "berhasil memperbarui data",
		"user": updateUser,
	})
}

func main() {
	// create a new echo instance
	e := echo.New()
	
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}