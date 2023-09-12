package main

import (
	"fmt"
)

type User struct {
	name string
	age int
	email string
	address string
}

type UserInterface interface {
	gantiNama()
}

// * adalah pointer
func (user User) gantiNama() {
	user.name = "alta - " + user.name
}

func show(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println(data.(int) + 2)
	case string:
		fmt.Println(data.(string) + " hellow")
	}
}

func Tambah (a, b int) (int, error) {
	result := a + b
	return result, nil
}

func materi() {
	// handle error
	defer func () {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("a")
	// panic("ada custom")
	fmt.Println("b")

	// error 
	// result, err := Tambah(0, 20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)

	// type data interface
	// show(1)
	// show("1")

	// 
	// var userInterface UserInterface
	// var user User
	// user.name = "altera"
	// fmt.Println(user)
	
	// userInterface = user
	// userInterface.gantiNama()

	// pointer
	// var age int = 20
	// var ageAddress *int = &age
	// fmt.Println(age)
	// fmt.Println(ageAddress)
	// gantiAddress(ageAddress)
	// fmt.Println(*ageAddress)
}

func changeUser(user User) {
}


func gantiAddress(address *int) {
	*address = 30
}