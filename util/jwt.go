package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

/**
 * @Author wyf
 * @Date 2021/7/8 11:17
 **/
const Secret = "JWT_SECRET"

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func GenerateToken(username, password string) string {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
		"iat":      time.Now().Unix(),
		"username": username,
		"password": password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv(Secret)))
	return t
}
