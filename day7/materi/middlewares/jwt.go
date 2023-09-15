package middlewares

import (
	"prakerja/day7/materi/models/user/database"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	UserId int `json:"userId"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(user database.User) string {
	claims := &jwtCustomClaims{
		user.Id,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	
		// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte("123"))
	// if err != nil {
	// 	return err
	// }

	return t
}