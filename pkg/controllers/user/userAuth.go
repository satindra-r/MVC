package user

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"mvc/pkg/config"
	"mvc/pkg/utils"
	"time"
)

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckUserPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return (err != nil)
}

func GenerateJWT(userId int) string {
	var JWT, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    userId,
		"Timestamp": time.Now().Unix() / 60,
	}).SignedString([]byte(config.JWTSecret))

	utils.LogIfErr(err, "JWT Signing Error")

	return JWT
}

func JWTGetUserId(JWT string) int {
	var claims = jwt.MapClaims{}
	var userID int
	var timestamp int
	var err error
	_, err = jwt.ParseWithClaims(JWT, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return -1
	}

	userID = (int)(claims["UserId"].(float64))
	timestamp = (int)(claims["Timestamp"].(float64))

	if timestamp < (int)(time.Now().Unix()/60)-60*24 {
		return -1
	}

	return userID

}
