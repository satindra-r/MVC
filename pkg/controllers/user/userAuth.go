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
		"timestamp": time.Now().Unix(),
	}).SignedString([]byte(config.JWTSecret))

	utils.LogIfErr(err, "JWT Signing Error")

	return JWT
}
