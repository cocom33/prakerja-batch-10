package database

import "prakerja/day7/materi/models/user/request"

type User struct {
	Id       int    `json:"id" gorm:"primaryKey autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) MapFromLogin(userLogin request.UserLogin) {
	user.Email = userLogin.Email
	user.Password = userLogin.Password
}

func (user *User) MapFromRegister(userRegister request.UserRegister) {
	user.Name = userRegister.Name
	user.Email = userRegister.Email
	user.Password = userRegister.Password
}